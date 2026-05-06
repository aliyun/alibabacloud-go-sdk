// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePocSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePocSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePocSampleResponse
	GetStatusCode() *int32
	SetBody(v *CreatePocSampleResponseBody) *CreatePocSampleResponse
	GetBody() *CreatePocSampleResponseBody
}

type CreatePocSampleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePocSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePocSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePocSampleResponse) GoString() string {
	return s.String()
}

func (s *CreatePocSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePocSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePocSampleResponse) GetBody() *CreatePocSampleResponseBody {
	return s.Body
}

func (s *CreatePocSampleResponse) SetHeaders(v map[string]*string) *CreatePocSampleResponse {
	s.Headers = v
	return s
}

func (s *CreatePocSampleResponse) SetStatusCode(v int32) *CreatePocSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePocSampleResponse) SetBody(v *CreatePocSampleResponseBody) *CreatePocSampleResponse {
	s.Body = v
	return s
}

func (s *CreatePocSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
