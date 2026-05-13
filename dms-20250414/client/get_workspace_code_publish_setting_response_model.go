// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceCodePublishSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkspaceCodePublishSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkspaceCodePublishSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkspaceCodePublishSettingResponseBody) *GetWorkspaceCodePublishSettingResponse
	GetBody() *GetWorkspaceCodePublishSettingResponseBody
}

type GetWorkspaceCodePublishSettingResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspaceCodePublishSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspaceCodePublishSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceCodePublishSettingResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceCodePublishSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkspaceCodePublishSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkspaceCodePublishSettingResponse) GetBody() *GetWorkspaceCodePublishSettingResponseBody {
	return s.Body
}

func (s *GetWorkspaceCodePublishSettingResponse) SetHeaders(v map[string]*string) *GetWorkspaceCodePublishSettingResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponse) SetStatusCode(v int32) *GetWorkspaceCodePublishSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponse) SetBody(v *GetWorkspaceCodePublishSettingResponseBody) *GetWorkspaceCodePublishSettingResponse {
	s.Body = v
	return s
}

func (s *GetWorkspaceCodePublishSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
