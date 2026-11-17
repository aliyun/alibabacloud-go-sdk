// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelApiResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelApiResponseBody) *CreateModelApiResponse
	GetBody() *CreateModelApiResponseBody
}

type CreateModelApiResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelApiResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelApiResponse) GoString() string {
	return s.String()
}

func (s *CreateModelApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelApiResponse) GetBody() *CreateModelApiResponseBody {
	return s.Body
}

func (s *CreateModelApiResponse) SetHeaders(v map[string]*string) *CreateModelApiResponse {
	s.Headers = v
	return s
}

func (s *CreateModelApiResponse) SetStatusCode(v int32) *CreateModelApiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelApiResponse) SetBody(v *CreateModelApiResponseBody) *CreateModelApiResponse {
	s.Body = v
	return s
}

func (s *CreateModelApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
