// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllCustomTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAllCustomTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAllCustomTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *GetAllCustomTemplatesResponseBody) *GetAllCustomTemplatesResponse
	GetBody() *GetAllCustomTemplatesResponseBody
}

type GetAllCustomTemplatesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllCustomTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllCustomTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAllCustomTemplatesResponse) GoString() string {
	return s.String()
}

func (s *GetAllCustomTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAllCustomTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAllCustomTemplatesResponse) GetBody() *GetAllCustomTemplatesResponseBody {
	return s.Body
}

func (s *GetAllCustomTemplatesResponse) SetHeaders(v map[string]*string) *GetAllCustomTemplatesResponse {
	s.Headers = v
	return s
}

func (s *GetAllCustomTemplatesResponse) SetStatusCode(v int32) *GetAllCustomTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllCustomTemplatesResponse) SetBody(v *GetAllCustomTemplatesResponseBody) *GetAllCustomTemplatesResponse {
	s.Body = v
	return s
}

func (s *GetAllCustomTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
