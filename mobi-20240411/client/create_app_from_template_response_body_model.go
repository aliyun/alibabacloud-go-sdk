// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppFromTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateAppFromTemplateResponseBodyData) *CreateAppFromTemplateResponseBody
	GetData() *CreateAppFromTemplateResponseBodyData
	SetRequestId(v string) *CreateAppFromTemplateResponseBody
	GetRequestId() *string
}

type CreateAppFromTemplateResponseBody struct {
	Data *CreateAppFromTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 36F552F7-E61E-556A-9957-8284318D1B9C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppFromTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppFromTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppFromTemplateResponseBody) GetData() *CreateAppFromTemplateResponseBodyData {
	return s.Data
}

func (s *CreateAppFromTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppFromTemplateResponseBody) SetData(v *CreateAppFromTemplateResponseBodyData) *CreateAppFromTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppFromTemplateResponseBody) SetRequestId(v string) *CreateAppFromTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppFromTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAppFromTemplateResponseBodyData struct {
	// example:
	//
	// 172050620*****
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2024-03-26T10:22Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-03-26T10:22Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// -1
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1731664463*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateAppFromTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAppFromTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppFromTemplateResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppFromTemplateResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateAppFromTemplateResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateAppFromTemplateResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateAppFromTemplateResponseBodyData) GetIcon() *string {
	return s.Icon
}

func (s *CreateAppFromTemplateResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateAppFromTemplateResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateAppFromTemplateResponseBodyData) SetAppId(v string) *CreateAppFromTemplateResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetDescription(v string) *CreateAppFromTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetGmtCreate(v string) *CreateAppFromTemplateResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetGmtModified(v string) *CreateAppFromTemplateResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetIcon(v string) *CreateAppFromTemplateResponseBodyData {
	s.Icon = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetName(v string) *CreateAppFromTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetWorkspaceId(v string) *CreateAppFromTemplateResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
