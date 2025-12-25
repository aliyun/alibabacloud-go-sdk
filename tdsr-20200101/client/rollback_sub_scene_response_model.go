// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackSubSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackSubSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackSubSceneResponse
	GetStatusCode() *int32
	SetBody(v *RollbackSubSceneResponseBody) *RollbackSubSceneResponse
	GetBody() *RollbackSubSceneResponseBody
}

type RollbackSubSceneResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackSubSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackSubSceneResponse) GoString() string {
	return s.String()
}

func (s *RollbackSubSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackSubSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackSubSceneResponse) GetBody() *RollbackSubSceneResponseBody {
	return s.Body
}

func (s *RollbackSubSceneResponse) SetHeaders(v map[string]*string) *RollbackSubSceneResponse {
	s.Headers = v
	return s
}

func (s *RollbackSubSceneResponse) SetStatusCode(v int32) *RollbackSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackSubSceneResponse) SetBody(v *RollbackSubSceneResponseBody) *RollbackSubSceneResponse {
	s.Body = v
	return s
}

func (s *RollbackSubSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
