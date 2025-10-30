// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSampleApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSampleApiResponse
	GetStatusCode() *int32
	SetBody(v *CreateSampleApiResponseBody) *CreateSampleApiResponse
	GetBody() *CreateSampleApiResponseBody
}

type CreateSampleApiResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSampleApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSampleApiResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleApiResponse) GoString() string {
	return s.String()
}

func (s *CreateSampleApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSampleApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSampleApiResponse) GetBody() *CreateSampleApiResponseBody {
	return s.Body
}

func (s *CreateSampleApiResponse) SetHeaders(v map[string]*string) *CreateSampleApiResponse {
	s.Headers = v
	return s
}

func (s *CreateSampleApiResponse) SetStatusCode(v int32) *CreateSampleApiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSampleApiResponse) SetBody(v *CreateSampleApiResponseBody) *CreateSampleApiResponse {
	s.Body = v
	return s
}

func (s *CreateSampleApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
