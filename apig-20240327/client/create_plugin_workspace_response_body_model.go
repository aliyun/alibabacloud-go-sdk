// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePluginWorkspaceResponseBody
	GetCode() *string
	SetData(v *CreatePluginWorkspaceResponseBodyData) *CreatePluginWorkspaceResponseBody
	GetData() *CreatePluginWorkspaceResponseBodyData
	SetMessage(v string) *CreatePluginWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePluginWorkspaceResponseBody
	GetRequestId() *string
}

type CreatePluginWorkspaceResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreatePluginWorkspaceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 01A02219-8028-57D8-9D60-2D167FF9118E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePluginWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePluginWorkspaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePluginWorkspaceResponseBody) GetData() *CreatePluginWorkspaceResponseBodyData {
	return s.Data
}

func (s *CreatePluginWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePluginWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePluginWorkspaceResponseBody) SetCode(v string) *CreatePluginWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePluginWorkspaceResponseBody) SetData(v *CreatePluginWorkspaceResponseBodyData) *CreatePluginWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *CreatePluginWorkspaceResponseBody) SetMessage(v string) *CreatePluginWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePluginWorkspaceResponseBody) SetRequestId(v string) *CreatePluginWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePluginWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePluginWorkspaceResponseBodyData struct {
	// example:
	//
	// 12345678
	RepoId *string `json:"repoId,omitempty" xml:"repoId,omitempty"`
	// example:
	//
	// plw-xxxxxxxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreatePluginWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePluginWorkspaceResponseBodyData) GetRepoId() *string {
	return s.RepoId
}

func (s *CreatePluginWorkspaceResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreatePluginWorkspaceResponseBodyData) SetRepoId(v string) *CreatePluginWorkspaceResponseBodyData {
	s.RepoId = &v
	return s
}

func (s *CreatePluginWorkspaceResponseBodyData) SetWorkspaceId(v string) *CreatePluginWorkspaceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreatePluginWorkspaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
