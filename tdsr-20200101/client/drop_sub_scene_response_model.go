// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropSubSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DropSubSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DropSubSceneResponse
	GetStatusCode() *int32
	SetBody(v *DropSubSceneResponseBody) *DropSubSceneResponse
	GetBody() *DropSubSceneResponseBody
}

type DropSubSceneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DropSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DropSubSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DropSubSceneResponse) GoString() string {
	return s.String()
}

func (s *DropSubSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DropSubSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DropSubSceneResponse) GetBody() *DropSubSceneResponseBody {
	return s.Body
}

func (s *DropSubSceneResponse) SetHeaders(v map[string]*string) *DropSubSceneResponse {
	s.Headers = v
	return s
}

func (s *DropSubSceneResponse) SetStatusCode(v int32) *DropSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DropSubSceneResponse) SetBody(v *DropSubSceneResponseBody) *DropSubSceneResponse {
	s.Body = v
	return s
}

func (s *DropSubSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
