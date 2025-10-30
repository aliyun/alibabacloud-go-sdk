// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamingDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStreamingDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStreamingDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStreamingDataSourceResponseBody) *DeleteStreamingDataSourceResponse
	GetBody() *DeleteStreamingDataSourceResponseBody
}

type DeleteStreamingDataSourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStreamingDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStreamingDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamingDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteStreamingDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStreamingDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStreamingDataSourceResponse) GetBody() *DeleteStreamingDataSourceResponseBody {
	return s.Body
}

func (s *DeleteStreamingDataSourceResponse) SetHeaders(v map[string]*string) *DeleteStreamingDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteStreamingDataSourceResponse) SetStatusCode(v int32) *DeleteStreamingDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStreamingDataSourceResponse) SetBody(v *DeleteStreamingDataSourceResponseBody) *DeleteStreamingDataSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteStreamingDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
