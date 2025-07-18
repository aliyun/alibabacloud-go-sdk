// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamingDataServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStreamingDataServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStreamingDataServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStreamingDataServiceResponseBody) *DeleteStreamingDataServiceResponse
	GetBody() *DeleteStreamingDataServiceResponseBody
}

type DeleteStreamingDataServiceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStreamingDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStreamingDataServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamingDataServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteStreamingDataServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStreamingDataServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStreamingDataServiceResponse) GetBody() *DeleteStreamingDataServiceResponseBody {
	return s.Body
}

func (s *DeleteStreamingDataServiceResponse) SetHeaders(v map[string]*string) *DeleteStreamingDataServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteStreamingDataServiceResponse) SetStatusCode(v int32) *DeleteStreamingDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStreamingDataServiceResponse) SetBody(v *DeleteStreamingDataServiceResponseBody) *DeleteStreamingDataServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteStreamingDataServiceResponse) Validate() error {
	return dara.Validate(s)
}
