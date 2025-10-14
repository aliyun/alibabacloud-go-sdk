// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetErServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetErServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetErServiceResponseBody) *GetErServiceResponse
	GetBody() *GetErServiceResponseBody
}

type GetErServiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetErServiceResponse) GoString() string {
	return s.String()
}

func (s *GetErServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetErServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetErServiceResponse) GetBody() *GetErServiceResponseBody {
	return s.Body
}

func (s *GetErServiceResponse) SetHeaders(v map[string]*string) *GetErServiceResponse {
	s.Headers = v
	return s
}

func (s *GetErServiceResponse) SetStatusCode(v int32) *GetErServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErServiceResponse) SetBody(v *GetErServiceResponseBody) *GetErServiceResponse {
	s.Body = v
	return s
}

func (s *GetErServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
