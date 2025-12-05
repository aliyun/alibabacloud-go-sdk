// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPtsSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPtsSceneResponse
	GetStatusCode() *int32
	SetBody(v *GetPtsSceneResponseBody) *GetPtsSceneResponse
	GetBody() *GetPtsSceneResponseBody
}

type GetPtsSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPtsSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPtsSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPtsSceneResponse) GetBody() *GetPtsSceneResponseBody {
	return s.Body
}

func (s *GetPtsSceneResponse) SetHeaders(v map[string]*string) *GetPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *GetPtsSceneResponse) SetStatusCode(v int32) *GetPtsSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPtsSceneResponse) SetBody(v *GetPtsSceneResponseBody) *GetPtsSceneResponse {
	s.Body = v
	return s
}

func (s *GetPtsSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
