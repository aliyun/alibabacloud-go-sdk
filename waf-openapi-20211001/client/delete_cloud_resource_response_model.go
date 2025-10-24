// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudResourceResponseBody) *DeleteCloudResourceResponse
	GetBody() *DeleteCloudResourceResponseBody
}

type DeleteCloudResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudResourceResponse) GetBody() *DeleteCloudResourceResponseBody {
	return s.Body
}

func (s *DeleteCloudResourceResponse) SetHeaders(v map[string]*string) *DeleteCloudResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudResourceResponse) SetStatusCode(v int32) *DeleteCloudResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudResourceResponse) SetBody(v *DeleteCloudResourceResponseBody) *DeleteCloudResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
