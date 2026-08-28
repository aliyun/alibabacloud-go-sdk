// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPluginWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPluginWorkspaceResponseBody
	GetCode() *string
	SetData(v *GetPluginWorkspaceResponseBodyData) *GetPluginWorkspaceResponseBody
	GetData() *GetPluginWorkspaceResponseBodyData
	SetMessage(v string) *GetPluginWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPluginWorkspaceResponseBody
	GetRequestId() *string
}

type GetPluginWorkspaceResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetPluginWorkspaceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4AFA893B-A75B-5002-AACF-84CABE06197A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPluginWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPluginWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPluginWorkspaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPluginWorkspaceResponseBody) GetData() *GetPluginWorkspaceResponseBodyData {
	return s.Data
}

func (s *GetPluginWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPluginWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPluginWorkspaceResponseBody) SetCode(v string) *GetPluginWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetPluginWorkspaceResponseBody) SetData(v *GetPluginWorkspaceResponseBodyData) *GetPluginWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *GetPluginWorkspaceResponseBody) SetMessage(v string) *GetPluginWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetPluginWorkspaceResponseBody) SetRequestId(v string) *GetPluginWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPluginWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPluginWorkspaceResponseBodyData struct {
	// example:
	//
	// 664f1e2xxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 987654
	PipelineRunId *string `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	// example:
	//
	// 12345678
	RepoId *string `json:"repoId,omitempty" xml:"repoId,omitempty"`
	// example:
	//
	// my-custom-plugin
	RepoName *string `json:"repoName,omitempty" xml:"repoName,omitempty"`
	// example:
	//
	// https://apigw-console-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/{uid}/plugin/plugin_1756262400.wasm
	WasmUrl *string `json:"wasmUrl,omitempty" xml:"wasmUrl,omitempty"`
	// example:
	//
	// plw-xxxxxxxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetPluginWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPluginWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPluginWorkspaceResponseBodyData) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetPluginWorkspaceResponseBodyData) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *GetPluginWorkspaceResponseBodyData) GetRepoId() *string {
	return s.RepoId
}

func (s *GetPluginWorkspaceResponseBodyData) GetRepoName() *string {
	return s.RepoName
}

func (s *GetPluginWorkspaceResponseBodyData) GetWasmUrl() *string {
	return s.WasmUrl
}

func (s *GetPluginWorkspaceResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetPluginWorkspaceResponseBodyData) SetOrganizationId(v string) *GetPluginWorkspaceResponseBodyData {
	s.OrganizationId = &v
	return s
}

func (s *GetPluginWorkspaceResponseBodyData) SetPipelineRunId(v string) *GetPluginWorkspaceResponseBodyData {
	s.PipelineRunId = &v
	return s
}

func (s *GetPluginWorkspaceResponseBodyData) SetRepoId(v string) *GetPluginWorkspaceResponseBodyData {
	s.RepoId = &v
	return s
}

func (s *GetPluginWorkspaceResponseBodyData) SetRepoName(v string) *GetPluginWorkspaceResponseBodyData {
	s.RepoName = &v
	return s
}

func (s *GetPluginWorkspaceResponseBodyData) SetWasmUrl(v string) *GetPluginWorkspaceResponseBodyData {
	s.WasmUrl = &v
	return s
}

func (s *GetPluginWorkspaceResponseBodyData) SetWorkspaceId(v string) *GetPluginWorkspaceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetPluginWorkspaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
