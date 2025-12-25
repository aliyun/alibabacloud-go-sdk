// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishSceneResponse
	GetStatusCode() *int32
	SetBody(v *PublishSceneResponseBody) *PublishSceneResponse
	GetBody() *PublishSceneResponseBody
}

type PublishSceneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishSceneResponse) GoString() string {
	return s.String()
}

func (s *PublishSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishSceneResponse) GetBody() *PublishSceneResponseBody {
	return s.Body
}

func (s *PublishSceneResponse) SetHeaders(v map[string]*string) *PublishSceneResponse {
	s.Headers = v
	return s
}

func (s *PublishSceneResponse) SetStatusCode(v int32) *PublishSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishSceneResponse) SetBody(v *PublishSceneResponseBody) *PublishSceneResponse {
	s.Body = v
	return s
}

func (s *PublishSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
