// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyExternalNodeStatusUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyExternalNodeStatusUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyExternalNodeStatusUpdateResponse
	GetStatusCode() *int32
	SetBody(v *ApplyExternalNodeStatusUpdateResponseBody) *ApplyExternalNodeStatusUpdateResponse
	GetBody() *ApplyExternalNodeStatusUpdateResponseBody
}

type ApplyExternalNodeStatusUpdateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyExternalNodeStatusUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyExternalNodeStatusUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyExternalNodeStatusUpdateResponse) GoString() string {
	return s.String()
}

func (s *ApplyExternalNodeStatusUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyExternalNodeStatusUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyExternalNodeStatusUpdateResponse) GetBody() *ApplyExternalNodeStatusUpdateResponseBody {
	return s.Body
}

func (s *ApplyExternalNodeStatusUpdateResponse) SetHeaders(v map[string]*string) *ApplyExternalNodeStatusUpdateResponse {
	s.Headers = v
	return s
}

func (s *ApplyExternalNodeStatusUpdateResponse) SetStatusCode(v int32) *ApplyExternalNodeStatusUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyExternalNodeStatusUpdateResponse) SetBody(v *ApplyExternalNodeStatusUpdateResponseBody) *ApplyExternalNodeStatusUpdateResponse {
	s.Body = v
	return s
}

func (s *ApplyExternalNodeStatusUpdateResponse) Validate() error {
	return dara.Validate(s)
}
