// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSceneResponse
	GetStatusCode() *int32
	SetBody(v *GetSceneResponseBody) *GetSceneResponse
	GetBody() *GetSceneResponseBody
}

type GetSceneResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSceneResponse) GoString() string {
	return s.String()
}

func (s *GetSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSceneResponse) GetBody() *GetSceneResponseBody {
	return s.Body
}

func (s *GetSceneResponse) SetHeaders(v map[string]*string) *GetSceneResponse {
	s.Headers = v
	return s
}

func (s *GetSceneResponse) SetStatusCode(v int32) *GetSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSceneResponse) SetBody(v *GetSceneResponseBody) *GetSceneResponse {
	s.Body = v
	return s
}

func (s *GetSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
