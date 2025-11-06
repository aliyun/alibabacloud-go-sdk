// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneNacosConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneNacosConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneNacosConfigResponse
	GetStatusCode() *int32
	SetBody(v *CloneNacosConfigResponseBody) *CloneNacosConfigResponse
	GetBody() *CloneNacosConfigResponseBody
}

type CloneNacosConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneNacosConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneNacosConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneNacosConfigResponse) GetBody() *CloneNacosConfigResponseBody {
	return s.Body
}

func (s *CloneNacosConfigResponse) SetHeaders(v map[string]*string) *CloneNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *CloneNacosConfigResponse) SetStatusCode(v int32) *CloneNacosConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneNacosConfigResponse) SetBody(v *CloneNacosConfigResponseBody) *CloneNacosConfigResponse {
	s.Body = v
	return s
}

func (s *CloneNacosConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
