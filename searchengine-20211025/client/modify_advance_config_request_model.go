// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ModifyAdvanceConfigRequest
	GetContent() *string
	SetContentType(v string) *ModifyAdvanceConfigRequest
	GetContentType() *string
	SetDesc(v string) *ModifyAdvanceConfigRequest
	GetDesc() *string
	SetFiles(v []*ModifyAdvanceConfigRequestFiles) *ModifyAdvanceConfigRequest
	GetFiles() []*ModifyAdvanceConfigRequestFiles
	SetName(v string) *ModifyAdvanceConfigRequest
	GetName() *string
	SetStatus(v string) *ModifyAdvanceConfigRequest
	GetStatus() *string
	SetUpdateTime(v int64) *ModifyAdvanceConfigRequest
	GetUpdateTime() *int64
}

type ModifyAdvanceConfigRequest struct {
	// The content of the advanced configuration that is returned.
	//
	// example:
	//
	// ""
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The type of the configuration content. Valid values: FILE, GIT, HTTP, and ODPS.
	//
	// example:
	//
	// FILE
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The description of the advanced configuration.
	//
	// example:
	//
	// test
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The files.
	Files []*ModifyAdvanceConfigRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The name of the advanced configuration.
	//
	// example:
	//
	// ha-cn-zvp2qr1sk01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the advanced configuration. Valid values: drafting: The advanced configuration is in the draft state. used: The advanced configuration is being used. unused: The advanced configuration is not used. trash: The advanced configuration is being deleted.
	//
	// example:
	//
	// used
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the advanced configuration was updated.
	//
	// example:
	//
	// 2024-02-27T07:50:55Z
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ModifyAdvanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigRequest) GetContent() *string {
	return s.Content
}

func (s *ModifyAdvanceConfigRequest) GetContentType() *string {
	return s.ContentType
}

func (s *ModifyAdvanceConfigRequest) GetDesc() *string {
	return s.Desc
}

func (s *ModifyAdvanceConfigRequest) GetFiles() []*ModifyAdvanceConfigRequestFiles {
	return s.Files
}

func (s *ModifyAdvanceConfigRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAdvanceConfigRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyAdvanceConfigRequest) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ModifyAdvanceConfigRequest) SetContent(v string) *ModifyAdvanceConfigRequest {
	s.Content = &v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetContentType(v string) *ModifyAdvanceConfigRequest {
	s.ContentType = &v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetDesc(v string) *ModifyAdvanceConfigRequest {
	s.Desc = &v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetFiles(v []*ModifyAdvanceConfigRequestFiles) *ModifyAdvanceConfigRequest {
	s.Files = v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetName(v string) *ModifyAdvanceConfigRequest {
	s.Name = &v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetStatus(v string) *ModifyAdvanceConfigRequest {
	s.Status = &v
	return s
}

func (s *ModifyAdvanceConfigRequest) SetUpdateTime(v int64) *ModifyAdvanceConfigRequest {
	s.UpdateTime = &v
	return s
}

func (s *ModifyAdvanceConfigRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyAdvanceConfigRequestFiles struct {
	// The full path of the file.
	//
	// example:
	//
	// /cluster.json
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Specifies whether the file is a directory.
	//
	// example:
	//
	// true
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Specifies whether the file is a template.
	//
	// example:
	//
	// true
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// The node name.
	//
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModifyAdvanceConfigRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvanceConfigRequestFiles) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigRequestFiles) GetFullPathName() *string {
	return s.FullPathName
}

func (s *ModifyAdvanceConfigRequestFiles) GetIsDir() *bool {
	return s.IsDir
}

func (s *ModifyAdvanceConfigRequestFiles) GetIsTemplate() *bool {
	return s.IsTemplate
}

func (s *ModifyAdvanceConfigRequestFiles) GetName() *string {
	return s.Name
}

func (s *ModifyAdvanceConfigRequestFiles) SetFullPathName(v string) *ModifyAdvanceConfigRequestFiles {
	s.FullPathName = &v
	return s
}

func (s *ModifyAdvanceConfigRequestFiles) SetIsDir(v bool) *ModifyAdvanceConfigRequestFiles {
	s.IsDir = &v
	return s
}

func (s *ModifyAdvanceConfigRequestFiles) SetIsTemplate(v bool) *ModifyAdvanceConfigRequestFiles {
	s.IsTemplate = &v
	return s
}

func (s *ModifyAdvanceConfigRequestFiles) SetName(v string) *ModifyAdvanceConfigRequestFiles {
	s.Name = &v
	return s
}

func (s *ModifyAdvanceConfigRequestFiles) Validate() error {
	return dara.Validate(s)
}
