// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPluginWorkspaceResponseBody
	GetCode() *string
	SetData(v []*ListPluginWorkspaceResponseBodyData) *ListPluginWorkspaceResponseBody
	GetData() []*ListPluginWorkspaceResponseBodyData
	SetMessage(v string) *ListPluginWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPluginWorkspaceResponseBody
	GetRequestId() *string
}

type ListPluginWorkspaceResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListPluginWorkspaceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B690F39C-1BDA-55E0-9E94-5358E758C772
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPluginWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPluginWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListPluginWorkspaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPluginWorkspaceResponseBody) GetData() []*ListPluginWorkspaceResponseBodyData {
	return s.Data
}

func (s *ListPluginWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPluginWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPluginWorkspaceResponseBody) SetCode(v string) *ListPluginWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *ListPluginWorkspaceResponseBody) SetData(v []*ListPluginWorkspaceResponseBodyData) *ListPluginWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *ListPluginWorkspaceResponseBody) SetMessage(v string) *ListPluginWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *ListPluginWorkspaceResponseBody) SetRequestId(v string) *ListPluginWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPluginWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPluginWorkspaceResponseBodyData struct {
	// example:
	//
	// 664f1e2xxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
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
	// plw-xxxxxxxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListPluginWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPluginWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPluginWorkspaceResponseBodyData) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListPluginWorkspaceResponseBodyData) GetRepoId() *string {
	return s.RepoId
}

func (s *ListPluginWorkspaceResponseBodyData) GetRepoName() *string {
	return s.RepoName
}

func (s *ListPluginWorkspaceResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListPluginWorkspaceResponseBodyData) SetOrganizationId(v string) *ListPluginWorkspaceResponseBodyData {
	s.OrganizationId = &v
	return s
}

func (s *ListPluginWorkspaceResponseBodyData) SetRepoId(v string) *ListPluginWorkspaceResponseBodyData {
	s.RepoId = &v
	return s
}

func (s *ListPluginWorkspaceResponseBodyData) SetRepoName(v string) *ListPluginWorkspaceResponseBodyData {
	s.RepoName = &v
	return s
}

func (s *ListPluginWorkspaceResponseBodyData) SetWorkspaceId(v string) *ListPluginWorkspaceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *ListPluginWorkspaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
