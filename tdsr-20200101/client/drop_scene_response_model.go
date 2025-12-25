// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DropSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DropSceneResponse
	GetStatusCode() *int32
	SetBody(v *DropSceneResponseBody) *DropSceneResponse
	GetBody() *DropSceneResponseBody
}

type DropSceneResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DropSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DropSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DropSceneResponse) GoString() string {
	return s.String()
}

func (s *DropSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DropSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DropSceneResponse) GetBody() *DropSceneResponseBody {
	return s.Body
}

func (s *DropSceneResponse) SetHeaders(v map[string]*string) *DropSceneResponse {
	s.Headers = v
	return s
}

func (s *DropSceneResponse) SetStatusCode(v int32) *DropSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DropSceneResponse) SetBody(v *DropSceneResponseBody) *DropSceneResponse {
	s.Body = v
	return s
}

func (s *DropSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
