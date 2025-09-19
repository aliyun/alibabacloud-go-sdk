// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientConfSetupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClientConfSetupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClientConfSetupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClientConfSetupResponseBody) *ModifyClientConfSetupResponse
	GetBody() *ModifyClientConfSetupResponseBody
}

type ModifyClientConfSetupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClientConfSetupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClientConfSetupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientConfSetupResponse) GoString() string {
	return s.String()
}

func (s *ModifyClientConfSetupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClientConfSetupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClientConfSetupResponse) GetBody() *ModifyClientConfSetupResponseBody {
	return s.Body
}

func (s *ModifyClientConfSetupResponse) SetHeaders(v map[string]*string) *ModifyClientConfSetupResponse {
	s.Headers = v
	return s
}

func (s *ModifyClientConfSetupResponse) SetStatusCode(v int32) *ModifyClientConfSetupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClientConfSetupResponse) SetBody(v *ModifyClientConfSetupResponseBody) *ModifyClientConfSetupResponse {
	s.Body = v
	return s
}

func (s *ModifyClientConfSetupResponse) Validate() error {
	return dara.Validate(s)
}
