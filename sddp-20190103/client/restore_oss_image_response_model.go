// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreOssImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreOssImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreOssImageResponse
	GetStatusCode() *int32
	SetBody(v *RestoreOssImageResponseBody) *RestoreOssImageResponse
	GetBody() *RestoreOssImageResponseBody
}

type RestoreOssImageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreOssImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreOssImageResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreOssImageResponse) GoString() string {
	return s.String()
}

func (s *RestoreOssImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreOssImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreOssImageResponse) GetBody() *RestoreOssImageResponseBody {
	return s.Body
}

func (s *RestoreOssImageResponse) SetHeaders(v map[string]*string) *RestoreOssImageResponse {
	s.Headers = v
	return s
}

func (s *RestoreOssImageResponse) SetStatusCode(v int32) *RestoreOssImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreOssImageResponse) SetBody(v *RestoreOssImageResponseBody) *RestoreOssImageResponse {
	s.Body = v
	return s
}

func (s *RestoreOssImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
