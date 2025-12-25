// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopySceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopySceneResponse
	GetStatusCode() *int32
	SetBody(v *CopySceneResponseBody) *CopySceneResponse
	GetBody() *CopySceneResponseBody
}

type CopySceneResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopySceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopySceneResponse) String() string {
	return dara.Prettify(s)
}

func (s CopySceneResponse) GoString() string {
	return s.String()
}

func (s *CopySceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopySceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopySceneResponse) GetBody() *CopySceneResponseBody {
	return s.Body
}

func (s *CopySceneResponse) SetHeaders(v map[string]*string) *CopySceneResponse {
	s.Headers = v
	return s
}

func (s *CopySceneResponse) SetStatusCode(v int32) *CopySceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CopySceneResponse) SetBody(v *CopySceneResponseBody) *CopySceneResponse {
	s.Body = v
	return s
}

func (s *CopySceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
