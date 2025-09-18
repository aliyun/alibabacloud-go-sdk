// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSceneListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSceneListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSceneListResponse
	GetStatusCode() *int32
	SetBody(v *GetSceneListResponseBody) *GetSceneListResponse
	GetBody() *GetSceneListResponseBody
}

type GetSceneListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSceneListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSceneListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSceneListResponse) GoString() string {
	return s.String()
}

func (s *GetSceneListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSceneListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSceneListResponse) GetBody() *GetSceneListResponseBody {
	return s.Body
}

func (s *GetSceneListResponse) SetHeaders(v map[string]*string) *GetSceneListResponse {
	s.Headers = v
	return s
}

func (s *GetSceneListResponse) SetStatusCode(v int32) *GetSceneListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSceneListResponse) SetBody(v *GetSceneListResponseBody) *GetSceneListResponse {
	s.Body = v
	return s
}

func (s *GetSceneListResponse) Validate() error {
	return dara.Validate(s)
}
