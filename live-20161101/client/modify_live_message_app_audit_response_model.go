// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageAppAuditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLiveMessageAppAuditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLiveMessageAppAuditResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLiveMessageAppAuditResponseBody) *ModifyLiveMessageAppAuditResponse
	GetBody() *ModifyLiveMessageAppAuditResponseBody
}

type ModifyLiveMessageAppAuditResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLiveMessageAppAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLiveMessageAppAuditResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageAppAuditResponse) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageAppAuditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLiveMessageAppAuditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLiveMessageAppAuditResponse) GetBody() *ModifyLiveMessageAppAuditResponseBody {
	return s.Body
}

func (s *ModifyLiveMessageAppAuditResponse) SetHeaders(v map[string]*string) *ModifyLiveMessageAppAuditResponse {
	s.Headers = v
	return s
}

func (s *ModifyLiveMessageAppAuditResponse) SetStatusCode(v int32) *ModifyLiveMessageAppAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLiveMessageAppAuditResponse) SetBody(v *ModifyLiveMessageAppAuditResponseBody) *ModifyLiveMessageAppAuditResponse {
	s.Body = v
	return s
}

func (s *ModifyLiveMessageAppAuditResponse) Validate() error {
	return dara.Validate(s)
}
