// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserBusinessFormResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserBusinessFormResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserBusinessFormResponse
	GetStatusCode() *int32
	SetBody(v *AddUserBusinessFormResponseBody) *AddUserBusinessFormResponse
	GetBody() *AddUserBusinessFormResponseBody
}

type AddUserBusinessFormResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserBusinessFormResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserBusinessFormResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserBusinessFormResponse) GoString() string {
	return s.String()
}

func (s *AddUserBusinessFormResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserBusinessFormResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserBusinessFormResponse) GetBody() *AddUserBusinessFormResponseBody {
	return s.Body
}

func (s *AddUserBusinessFormResponse) SetHeaders(v map[string]*string) *AddUserBusinessFormResponse {
	s.Headers = v
	return s
}

func (s *AddUserBusinessFormResponse) SetStatusCode(v int32) *AddUserBusinessFormResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserBusinessFormResponse) SetBody(v *AddUserBusinessFormResponseBody) *AddUserBusinessFormResponse {
	s.Body = v
	return s
}

func (s *AddUserBusinessFormResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
