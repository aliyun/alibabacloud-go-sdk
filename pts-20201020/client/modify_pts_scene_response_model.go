// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPtsSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPtsSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPtsSceneResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPtsSceneResponseBody) *ModifyPtsSceneResponse
	GetBody() *ModifyPtsSceneResponseBody
}

type ModifyPtsSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPtsSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *ModifyPtsSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPtsSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPtsSceneResponse) GetBody() *ModifyPtsSceneResponseBody {
	return s.Body
}

func (s *ModifyPtsSceneResponse) SetHeaders(v map[string]*string) *ModifyPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *ModifyPtsSceneResponse) SetStatusCode(v int32) *ModifyPtsSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPtsSceneResponse) SetBody(v *ModifyPtsSceneResponseBody) *ModifyPtsSceneResponse {
	s.Body = v
	return s
}

func (s *ModifyPtsSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
