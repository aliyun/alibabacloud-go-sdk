// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySceneConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySceneConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySceneConfigsResponse
	GetStatusCode() *int32
	SetBody(v *QuerySceneConfigsResponseBody) *QuerySceneConfigsResponse
	GetBody() *QuerySceneConfigsResponseBody
}

type QuerySceneConfigsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySceneConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySceneConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneConfigsResponse) GoString() string {
	return s.String()
}

func (s *QuerySceneConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySceneConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySceneConfigsResponse) GetBody() *QuerySceneConfigsResponseBody {
	return s.Body
}

func (s *QuerySceneConfigsResponse) SetHeaders(v map[string]*string) *QuerySceneConfigsResponse {
	s.Headers = v
	return s
}

func (s *QuerySceneConfigsResponse) SetStatusCode(v int32) *QuerySceneConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySceneConfigsResponse) SetBody(v *QuerySceneConfigsResponseBody) *QuerySceneConfigsResponse {
	s.Body = v
	return s
}

func (s *QuerySceneConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
