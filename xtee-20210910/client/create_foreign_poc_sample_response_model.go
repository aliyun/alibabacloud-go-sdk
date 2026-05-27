// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForeignPocSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateForeignPocSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateForeignPocSampleResponse
	GetStatusCode() *int32
	SetBody(v *CreateForeignPocSampleResponseBody) *CreateForeignPocSampleResponse
	GetBody() *CreateForeignPocSampleResponseBody
}

type CreateForeignPocSampleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateForeignPocSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateForeignPocSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateForeignPocSampleResponse) GoString() string {
	return s.String()
}

func (s *CreateForeignPocSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateForeignPocSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateForeignPocSampleResponse) GetBody() *CreateForeignPocSampleResponseBody {
	return s.Body
}

func (s *CreateForeignPocSampleResponse) SetHeaders(v map[string]*string) *CreateForeignPocSampleResponse {
	s.Headers = v
	return s
}

func (s *CreateForeignPocSampleResponse) SetStatusCode(v int32) *CreateForeignPocSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateForeignPocSampleResponse) SetBody(v *CreateForeignPocSampleResponseBody) *CreateForeignPocSampleResponse {
	s.Body = v
	return s
}

func (s *CreateForeignPocSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
