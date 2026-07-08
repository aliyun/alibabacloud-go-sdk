// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudAccessResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudAccessResponseBody) *DeleteCloudAccessResponse
	GetBody() *DeleteCloudAccessResponseBody
}

type DeleteCloudAccessResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAccessResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudAccessResponse) GetBody() *DeleteCloudAccessResponseBody {
	return s.Body
}

func (s *DeleteCloudAccessResponse) SetHeaders(v map[string]*string) *DeleteCloudAccessResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudAccessResponse) SetStatusCode(v int32) *DeleteCloudAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudAccessResponse) SetBody(v *DeleteCloudAccessResponseBody) *DeleteCloudAccessResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
