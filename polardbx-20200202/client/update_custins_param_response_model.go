// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustinsParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustinsParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustinsParamResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustinsParamResponseBody) *UpdateCustinsParamResponse
	GetBody() *UpdateCustinsParamResponseBody
}

type UpdateCustinsParamResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustinsParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustinsParamResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustinsParamResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustinsParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustinsParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustinsParamResponse) GetBody() *UpdateCustinsParamResponseBody {
	return s.Body
}

func (s *UpdateCustinsParamResponse) SetHeaders(v map[string]*string) *UpdateCustinsParamResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustinsParamResponse) SetStatusCode(v int32) *UpdateCustinsParamResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustinsParamResponse) SetBody(v *UpdateCustinsParamResponseBody) *UpdateCustinsParamResponse {
	s.Body = v
	return s
}

func (s *UpdateCustinsParamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
