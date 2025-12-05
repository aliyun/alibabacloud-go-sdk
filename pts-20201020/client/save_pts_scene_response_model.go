// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSavePtsSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SavePtsSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SavePtsSceneResponse
	GetStatusCode() *int32
	SetBody(v *SavePtsSceneResponseBody) *SavePtsSceneResponse
	GetBody() *SavePtsSceneResponseBody
}

type SavePtsSceneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SavePtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SavePtsSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneResponse) GoString() string {
	return s.String()
}

func (s *SavePtsSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SavePtsSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SavePtsSceneResponse) GetBody() *SavePtsSceneResponseBody {
	return s.Body
}

func (s *SavePtsSceneResponse) SetHeaders(v map[string]*string) *SavePtsSceneResponse {
	s.Headers = v
	return s
}

func (s *SavePtsSceneResponse) SetStatusCode(v int32) *SavePtsSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *SavePtsSceneResponse) SetBody(v *SavePtsSceneResponseBody) *SavePtsSceneResponse {
	s.Body = v
	return s
}

func (s *SavePtsSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
