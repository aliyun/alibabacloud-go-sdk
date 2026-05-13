// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkspaceCodePublishSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetWorkspaceCodePublishSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetWorkspaceCodePublishSettingResponse
	GetStatusCode() *int32
	SetBody(v *SetWorkspaceCodePublishSettingResponseBody) *SetWorkspaceCodePublishSettingResponse
	GetBody() *SetWorkspaceCodePublishSettingResponseBody
}

type SetWorkspaceCodePublishSettingResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWorkspaceCodePublishSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWorkspaceCodePublishSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s SetWorkspaceCodePublishSettingResponse) GoString() string {
	return s.String()
}

func (s *SetWorkspaceCodePublishSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetWorkspaceCodePublishSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetWorkspaceCodePublishSettingResponse) GetBody() *SetWorkspaceCodePublishSettingResponseBody {
	return s.Body
}

func (s *SetWorkspaceCodePublishSettingResponse) SetHeaders(v map[string]*string) *SetWorkspaceCodePublishSettingResponse {
	s.Headers = v
	return s
}

func (s *SetWorkspaceCodePublishSettingResponse) SetStatusCode(v int32) *SetWorkspaceCodePublishSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWorkspaceCodePublishSettingResponse) SetBody(v *SetWorkspaceCodePublishSettingResponseBody) *SetWorkspaceCodePublishSettingResponse {
	s.Body = v
	return s
}

func (s *SetWorkspaceCodePublishSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
