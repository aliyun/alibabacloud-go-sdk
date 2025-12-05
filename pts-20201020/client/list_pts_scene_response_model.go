// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPtsSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPtsSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPtsSceneResponse
	GetStatusCode() *int32
	SetBody(v *ListPtsSceneResponseBody) *ListPtsSceneResponse
	GetBody() *ListPtsSceneResponseBody
}

type ListPtsSceneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPtsSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *ListPtsSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPtsSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPtsSceneResponse) GetBody() *ListPtsSceneResponseBody {
	return s.Body
}

func (s *ListPtsSceneResponse) SetHeaders(v map[string]*string) *ListPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *ListPtsSceneResponse) SetStatusCode(v int32) *ListPtsSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPtsSceneResponse) SetBody(v *ListPtsSceneResponseBody) *ListPtsSceneResponse {
	s.Body = v
	return s
}

func (s *ListPtsSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
