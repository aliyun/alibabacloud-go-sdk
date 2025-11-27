// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudAppResponseBody) *DeleteCloudAppResponse
	GetBody() *DeleteCloudAppResponseBody
}

type DeleteCloudAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudAppResponse) GetBody() *DeleteCloudAppResponseBody {
	return s.Body
}

func (s *DeleteCloudAppResponse) SetHeaders(v map[string]*string) *DeleteCloudAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudAppResponse) SetStatusCode(v int32) *DeleteCloudAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudAppResponse) SetBody(v *DeleteCloudAppResponseBody) *DeleteCloudAppResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
