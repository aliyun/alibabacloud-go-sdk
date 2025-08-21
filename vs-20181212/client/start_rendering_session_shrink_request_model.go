// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRenderingSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartRenderingSessionShrinkRequest
	GetAppId() *string
	SetClientId(v string) *StartRenderingSessionShrinkRequest
	GetClientId() *string
	SetClientParamsShrink(v string) *StartRenderingSessionShrinkRequest
	GetClientParamsShrink() *string
	SetPatchId(v string) *StartRenderingSessionShrinkRequest
	GetPatchId() *string
	SetProjectId(v string) *StartRenderingSessionShrinkRequest
	GetProjectId() *string
}

type StartRenderingSessionShrinkRequest struct {
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 04c30850-1d91-4da1-b811-66d0ee94af7d
	ClientId           *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientParamsShrink *string `json:"ClientParams,omitempty" xml:"ClientParams,omitempty"`
	// example:
	//
	// patch-03fa76e8e13a49b63456b063dgh309b4
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s StartRenderingSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRenderingSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartRenderingSessionShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartRenderingSessionShrinkRequest) GetClientId() *string {
	return s.ClientId
}

func (s *StartRenderingSessionShrinkRequest) GetClientParamsShrink() *string {
	return s.ClientParamsShrink
}

func (s *StartRenderingSessionShrinkRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *StartRenderingSessionShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *StartRenderingSessionShrinkRequest) SetAppId(v string) *StartRenderingSessionShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StartRenderingSessionShrinkRequest) SetClientId(v string) *StartRenderingSessionShrinkRequest {
	s.ClientId = &v
	return s
}

func (s *StartRenderingSessionShrinkRequest) SetClientParamsShrink(v string) *StartRenderingSessionShrinkRequest {
	s.ClientParamsShrink = &v
	return s
}

func (s *StartRenderingSessionShrinkRequest) SetPatchId(v string) *StartRenderingSessionShrinkRequest {
	s.PatchId = &v
	return s
}

func (s *StartRenderingSessionShrinkRequest) SetProjectId(v string) *StartRenderingSessionShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *StartRenderingSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
