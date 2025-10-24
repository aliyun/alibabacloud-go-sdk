// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseSceneConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefenseSceneConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefenseSceneConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefenseSceneConfigResponseBody) *ModifyDefenseSceneConfigResponse
	GetBody() *ModifyDefenseSceneConfigResponseBody
}

type ModifyDefenseSceneConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseSceneConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseSceneConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseSceneConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefenseSceneConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefenseSceneConfigResponse) GetBody() *ModifyDefenseSceneConfigResponseBody {
	return s.Body
}

func (s *ModifyDefenseSceneConfigResponse) SetHeaders(v map[string]*string) *ModifyDefenseSceneConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseSceneConfigResponse) SetStatusCode(v int32) *ModifyDefenseSceneConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseSceneConfigResponse) SetBody(v *ModifyDefenseSceneConfigResponseBody) *ModifyDefenseSceneConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyDefenseSceneConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
