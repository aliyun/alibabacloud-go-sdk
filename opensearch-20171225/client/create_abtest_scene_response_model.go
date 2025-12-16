// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABTestSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateABTestSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateABTestSceneResponse
	GetStatusCode() *int32
	SetBody(v *CreateABTestSceneResponseBody) *CreateABTestSceneResponse
	GetBody() *CreateABTestSceneResponseBody
}

type CreateABTestSceneResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateABTestSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateABTestSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateABTestSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateABTestSceneResponse) GetBody() *CreateABTestSceneResponseBody {
	return s.Body
}

func (s *CreateABTestSceneResponse) SetHeaders(v map[string]*string) *CreateABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *CreateABTestSceneResponse) SetStatusCode(v int32) *CreateABTestSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateABTestSceneResponse) SetBody(v *CreateABTestSceneResponseBody) *CreateABTestSceneResponse {
	s.Body = v
	return s
}

func (s *CreateABTestSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
