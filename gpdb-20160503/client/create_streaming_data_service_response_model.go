// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamingDataServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStreamingDataServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStreamingDataServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateStreamingDataServiceResponseBody) *CreateStreamingDataServiceResponse
	GetBody() *CreateStreamingDataServiceResponseBody
}

type CreateStreamingDataServiceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStreamingDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStreamingDataServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamingDataServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamingDataServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStreamingDataServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStreamingDataServiceResponse) GetBody() *CreateStreamingDataServiceResponseBody {
	return s.Body
}

func (s *CreateStreamingDataServiceResponse) SetHeaders(v map[string]*string) *CreateStreamingDataServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamingDataServiceResponse) SetStatusCode(v int32) *CreateStreamingDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStreamingDataServiceResponse) SetBody(v *CreateStreamingDataServiceResponseBody) *CreateStreamingDataServiceResponse {
	s.Body = v
	return s
}

func (s *CreateStreamingDataServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
