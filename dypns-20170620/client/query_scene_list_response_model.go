// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySceneListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySceneListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySceneListResponse
	GetStatusCode() *int32
	SetBody(v *QuerySceneListResponseBody) *QuerySceneListResponse
	GetBody() *QuerySceneListResponseBody
}

type QuerySceneListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySceneListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySceneListResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneListResponse) GoString() string {
	return s.String()
}

func (s *QuerySceneListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySceneListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySceneListResponse) GetBody() *QuerySceneListResponseBody {
	return s.Body
}

func (s *QuerySceneListResponse) SetHeaders(v map[string]*string) *QuerySceneListResponse {
	s.Headers = v
	return s
}

func (s *QuerySceneListResponse) SetStatusCode(v int32) *QuerySceneListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySceneListResponse) SetBody(v *QuerySceneListResponseBody) *QuerySceneListResponse {
	s.Body = v
	return s
}

func (s *QuerySceneListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
