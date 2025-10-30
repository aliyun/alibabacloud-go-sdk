// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSampleDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSampleDataResponse
	GetStatusCode() *int32
	SetBody(v *CreateSampleDataResponseBody) *CreateSampleDataResponse
	GetBody() *CreateSampleDataResponseBody
}

type CreateSampleDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSampleDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSampleDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleDataResponse) GoString() string {
	return s.String()
}

func (s *CreateSampleDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSampleDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSampleDataResponse) GetBody() *CreateSampleDataResponseBody {
	return s.Body
}

func (s *CreateSampleDataResponse) SetHeaders(v map[string]*string) *CreateSampleDataResponse {
	s.Headers = v
	return s
}

func (s *CreateSampleDataResponse) SetStatusCode(v int32) *CreateSampleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSampleDataResponse) SetBody(v *CreateSampleDataResponseBody) *CreateSampleDataResponse {
	s.Body = v
	return s
}

func (s *CreateSampleDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
