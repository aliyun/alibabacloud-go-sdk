// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamingDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStreamingDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStreamingDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateStreamingDataSourceResponseBody) *CreateStreamingDataSourceResponse
	GetBody() *CreateStreamingDataSourceResponseBody
}

type CreateStreamingDataSourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStreamingDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStreamingDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamingDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateStreamingDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStreamingDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStreamingDataSourceResponse) GetBody() *CreateStreamingDataSourceResponseBody {
	return s.Body
}

func (s *CreateStreamingDataSourceResponse) SetHeaders(v map[string]*string) *CreateStreamingDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateStreamingDataSourceResponse) SetStatusCode(v int32) *CreateStreamingDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStreamingDataSourceResponse) SetBody(v *CreateStreamingDataSourceResponseBody) *CreateStreamingDataSourceResponse {
	s.Body = v
	return s
}

func (s *CreateStreamingDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
