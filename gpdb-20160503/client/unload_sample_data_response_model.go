// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnloadSampleDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnloadSampleDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnloadSampleDataResponse
	GetStatusCode() *int32
	SetBody(v *UnloadSampleDataResponseBody) *UnloadSampleDataResponse
	GetBody() *UnloadSampleDataResponseBody
}

type UnloadSampleDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnloadSampleDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnloadSampleDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UnloadSampleDataResponse) GoString() string {
	return s.String()
}

func (s *UnloadSampleDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnloadSampleDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnloadSampleDataResponse) GetBody() *UnloadSampleDataResponseBody {
	return s.Body
}

func (s *UnloadSampleDataResponse) SetHeaders(v map[string]*string) *UnloadSampleDataResponse {
	s.Headers = v
	return s
}

func (s *UnloadSampleDataResponse) SetStatusCode(v int32) *UnloadSampleDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UnloadSampleDataResponse) SetBody(v *UnloadSampleDataResponseBody) *UnloadSampleDataResponse {
	s.Body = v
	return s
}

func (s *UnloadSampleDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
