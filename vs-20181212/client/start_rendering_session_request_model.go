// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRenderingSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartRenderingSessionRequest
	GetAppId() *string
	SetClientId(v string) *StartRenderingSessionRequest
	GetClientId() *string
	SetClientParams(v *StartRenderingSessionRequestClientParams) *StartRenderingSessionRequest
	GetClientParams() *StartRenderingSessionRequestClientParams
	SetPatchId(v string) *StartRenderingSessionRequest
	GetPatchId() *string
	SetProjectId(v string) *StartRenderingSessionRequest
	GetProjectId() *string
}

type StartRenderingSessionRequest struct {
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 04c30850-1d91-4da1-b811-66d0ee94af7d
	ClientId     *string                                   `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientParams *StartRenderingSessionRequestClientParams `json:"ClientParams,omitempty" xml:"ClientParams,omitempty" type:"Struct"`
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

func (s StartRenderingSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRenderingSessionRequest) GoString() string {
	return s.String()
}

func (s *StartRenderingSessionRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartRenderingSessionRequest) GetClientId() *string {
	return s.ClientId
}

func (s *StartRenderingSessionRequest) GetClientParams() *StartRenderingSessionRequestClientParams {
	return s.ClientParams
}

func (s *StartRenderingSessionRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *StartRenderingSessionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *StartRenderingSessionRequest) SetAppId(v string) *StartRenderingSessionRequest {
	s.AppId = &v
	return s
}

func (s *StartRenderingSessionRequest) SetClientId(v string) *StartRenderingSessionRequest {
	s.ClientId = &v
	return s
}

func (s *StartRenderingSessionRequest) SetClientParams(v *StartRenderingSessionRequestClientParams) *StartRenderingSessionRequest {
	s.ClientParams = v
	return s
}

func (s *StartRenderingSessionRequest) SetPatchId(v string) *StartRenderingSessionRequest {
	s.PatchId = &v
	return s
}

func (s *StartRenderingSessionRequest) SetProjectId(v string) *StartRenderingSessionRequest {
	s.ProjectId = &v
	return s
}

func (s *StartRenderingSessionRequest) Validate() error {
	return dara.Validate(s)
}

type StartRenderingSessionRequestClientParams struct {
	// example:
	//
	// 106.11.43.1
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
}

func (s StartRenderingSessionRequestClientParams) String() string {
	return dara.Prettify(s)
}

func (s StartRenderingSessionRequestClientParams) GoString() string {
	return s.String()
}

func (s *StartRenderingSessionRequestClientParams) GetClientIp() *string {
	return s.ClientIp
}

func (s *StartRenderingSessionRequestClientParams) SetClientIp(v string) *StartRenderingSessionRequestClientParams {
	s.ClientIp = &v
	return s
}

func (s *StartRenderingSessionRequestClientParams) Validate() error {
	return dara.Validate(s)
}
