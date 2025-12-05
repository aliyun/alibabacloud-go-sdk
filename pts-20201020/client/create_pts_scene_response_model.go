// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePtsSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePtsSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePtsSceneResponse
	GetStatusCode() *int32
	SetBody(v *CreatePtsSceneResponseBody) *CreatePtsSceneResponse
	GetBody() *CreatePtsSceneResponseBody
}

type CreatePtsSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePtsSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePtsSceneResponse) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePtsSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePtsSceneResponse) GetBody() *CreatePtsSceneResponseBody {
	return s.Body
}

func (s *CreatePtsSceneResponse) SetHeaders(v map[string]*string) *CreatePtsSceneResponse {
	s.Headers = v
	return s
}

func (s *CreatePtsSceneResponse) SetStatusCode(v int32) *CreatePtsSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePtsSceneResponse) SetBody(v *CreatePtsSceneResponseBody) *CreatePtsSceneResponse {
	s.Body = v
	return s
}

func (s *CreatePtsSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
