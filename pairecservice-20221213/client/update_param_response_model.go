// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateParamResponse
	GetStatusCode() *int32
	SetBody(v *UpdateParamResponseBody) *UpdateParamResponse
	GetBody() *UpdateParamResponseBody
}

type UpdateParamResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateParamResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateParamResponse) GoString() string {
	return s.String()
}

func (s *UpdateParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateParamResponse) GetBody() *UpdateParamResponseBody {
	return s.Body
}

func (s *UpdateParamResponse) SetHeaders(v map[string]*string) *UpdateParamResponse {
	s.Headers = v
	return s
}

func (s *UpdateParamResponse) SetStatusCode(v int32) *UpdateParamResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateParamResponse) SetBody(v *UpdateParamResponseBody) *UpdateParamResponse {
	s.Body = v
	return s
}

func (s *UpdateParamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
