// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorRequestSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetErrorRequestSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetErrorRequestSampleResponse
	GetStatusCode() *int32
	SetBody(v *GetErrorRequestSampleResponseBody) *GetErrorRequestSampleResponse
	GetBody() *GetErrorRequestSampleResponseBody
}

type GetErrorRequestSampleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErrorRequestSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErrorRequestSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetErrorRequestSampleResponse) GoString() string {
	return s.String()
}

func (s *GetErrorRequestSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetErrorRequestSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetErrorRequestSampleResponse) GetBody() *GetErrorRequestSampleResponseBody {
	return s.Body
}

func (s *GetErrorRequestSampleResponse) SetHeaders(v map[string]*string) *GetErrorRequestSampleResponse {
	s.Headers = v
	return s
}

func (s *GetErrorRequestSampleResponse) SetStatusCode(v int32) *GetErrorRequestSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErrorRequestSampleResponse) SetBody(v *GetErrorRequestSampleResponseBody) *GetErrorRequestSampleResponse {
	s.Body = v
	return s
}

func (s *GetErrorRequestSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
