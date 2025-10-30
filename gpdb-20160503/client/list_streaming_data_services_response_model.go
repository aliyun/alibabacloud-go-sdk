// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamingDataServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStreamingDataServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStreamingDataServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListStreamingDataServicesResponseBody) *ListStreamingDataServicesResponse
	GetBody() *ListStreamingDataServicesResponseBody
}

type ListStreamingDataServicesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStreamingDataServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStreamingDataServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingDataServicesResponse) GoString() string {
	return s.String()
}

func (s *ListStreamingDataServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStreamingDataServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStreamingDataServicesResponse) GetBody() *ListStreamingDataServicesResponseBody {
	return s.Body
}

func (s *ListStreamingDataServicesResponse) SetHeaders(v map[string]*string) *ListStreamingDataServicesResponse {
	s.Headers = v
	return s
}

func (s *ListStreamingDataServicesResponse) SetStatusCode(v int32) *ListStreamingDataServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStreamingDataServicesResponse) SetBody(v *ListStreamingDataServicesResponseBody) *ListStreamingDataServicesResponse {
	s.Body = v
	return s
}

func (s *ListStreamingDataServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
