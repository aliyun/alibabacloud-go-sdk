// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegisterCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegisterCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegisterCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetRegisterCodeResponseBody) *GetRegisterCodeResponse
	GetBody() *GetRegisterCodeResponseBody
}

type GetRegisterCodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegisterCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegisterCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegisterCodeResponse) GoString() string {
	return s.String()
}

func (s *GetRegisterCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegisterCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegisterCodeResponse) GetBody() *GetRegisterCodeResponseBody {
	return s.Body
}

func (s *GetRegisterCodeResponse) SetHeaders(v map[string]*string) *GetRegisterCodeResponse {
	s.Headers = v
	return s
}

func (s *GetRegisterCodeResponse) SetStatusCode(v int32) *GetRegisterCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegisterCodeResponse) SetBody(v *GetRegisterCodeResponseBody) *GetRegisterCodeResponse {
	s.Body = v
	return s
}

func (s *GetRegisterCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
