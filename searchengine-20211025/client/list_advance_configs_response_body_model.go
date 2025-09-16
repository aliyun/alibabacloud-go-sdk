// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdvanceConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAdvanceConfigsResponseBody
	GetRequestId() *string
	SetResult(v []*ListAdvanceConfigsResponseBodyResult) *ListAdvanceConfigsResponseBody
	GetResult() []*ListAdvanceConfigsResponseBodyResult
}

type ListAdvanceConfigsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4FB0325E-8C37-5525-96AC-0333523170A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The advanced configurations.
	Result []*ListAdvanceConfigsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAdvanceConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAdvanceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAdvanceConfigsResponseBody) GetResult() []*ListAdvanceConfigsResponseBodyResult {
	return s.Result
}

func (s *ListAdvanceConfigsResponseBody) SetRequestId(v string) *ListAdvanceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAdvanceConfigsResponseBody) SetResult(v []*ListAdvanceConfigsResponseBodyResult) *ListAdvanceConfigsResponseBody {
	s.Result = v
	return s
}

func (s *ListAdvanceConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAdvanceConfigsResponseBodyResult struct {
	// 	- The type of the advanced configuration. Valid values: -ONLINE: online configuration
	//
	// 	- \\-ONLINE_CAVA: online Cava configuration
	//
	// 	- \\-ONLINE_PLUGIN: online plug-in configuration
	//
	// 	- \\-ONLINE_QUERY: query configuration
	//
	// 	- \\-OFFLINE_DICT: offline dictionary configuration
	//
	// 	- \\-OFFLINE_TABLE: offline table configuration
	//
	// 	- \\-OFFLINE_COMMON: offline configuration
	//
	// 	- \\-OFFLINE_PLUGIN: offline plug-in configuration
	//
	// 	- \\-OFFLINE_INDEX: index configuration
	//
	// example:
	//
	// ONLINE
	AdvanceConfigType *string `json:"advanceConfigType,omitempty" xml:"advanceConfigType,omitempty"`
	// The content of the advanced configuration that is returned.
	//
	// example:
	//
	// {\\"url\\":\\"http://xxxxxx.aliyuncs.com/outnet_hz/packages/xxxxx/opensearch_offline_plugins_xxxxx.tar\\"}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The type of the configuration content. Valid values: FILE, GIT, HTTP, and ODPS.
	//
	// example:
	//
	// FILE
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The Alibaba Cloud account ID of the user who created the advanced configuration.
	//
	// example:
	//
	// 123456
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The description of the advanced configuration.
	//
	// example:
	//
	// test
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The files.
	Files []*ListAdvanceConfigsResponseBodyResultFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The name of the advanced configuration.
	//
	// example:
	//
	// my_index
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the advanced configuration. Valid values: drafting: The advanced configuration is in the draft state. used: The advanced configuration is being used. unused: The advanced configuration is not used. trash: The advanced configuration is being deleted.
	//
	// example:
	//
	// drafting
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the advanced configuration was updated.
	//
	// example:
	//
	// 1631070464000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListAdvanceConfigsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAdvanceConfigsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponseBodyResult) GetAdvanceConfigType() *string {
	return s.AdvanceConfigType
}

func (s *ListAdvanceConfigsResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *ListAdvanceConfigsResponseBodyResult) GetContentType() *string {
	return s.ContentType
}

func (s *ListAdvanceConfigsResponseBodyResult) GetCreator() *string {
	return s.Creator
}

func (s *ListAdvanceConfigsResponseBodyResult) GetDesc() *string {
	return s.Desc
}

func (s *ListAdvanceConfigsResponseBodyResult) GetFiles() []*ListAdvanceConfigsResponseBodyResultFiles {
	return s.Files
}

func (s *ListAdvanceConfigsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListAdvanceConfigsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListAdvanceConfigsResponseBodyResult) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListAdvanceConfigsResponseBodyResult) SetAdvanceConfigType(v string) *ListAdvanceConfigsResponseBodyResult {
	s.AdvanceConfigType = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetContent(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetContentType(v string) *ListAdvanceConfigsResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetCreator(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Creator = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetDesc(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Desc = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetFiles(v []*ListAdvanceConfigsResponseBodyResultFiles) *ListAdvanceConfigsResponseBodyResult {
	s.Files = v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetName(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetStatus(v string) *ListAdvanceConfigsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) SetUpdateTime(v int64) *ListAdvanceConfigsResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListAdvanceConfigsResponseBodyResultFiles struct {
	// The absolute path in which the file is stored.
	//
	// example:
	//
	// /path/wpd/nae
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether the file is a directory. Valid values: true and false.
	//
	// example:
	//
	// true
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether the file is a template. Valid values: true and false.
	//
	// example:
	//
	// true
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// The file name.
	//
	// example:
	//
	// file_name_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAdvanceConfigsResponseBodyResultFiles) String() string {
	return dara.Prettify(s)
}

func (s ListAdvanceConfigsResponseBodyResultFiles) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) GetFullPathName() *string {
	return s.FullPathName
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) GetIsDir() *bool {
	return s.IsDir
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) GetIsTemplate() *bool {
	return s.IsTemplate
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) GetName() *string {
	return s.Name
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetFullPathName(v string) *ListAdvanceConfigsResponseBodyResultFiles {
	s.FullPathName = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetIsDir(v bool) *ListAdvanceConfigsResponseBodyResultFiles {
	s.IsDir = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetIsTemplate(v bool) *ListAdvanceConfigsResponseBodyResultFiles {
	s.IsTemplate = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) SetName(v string) *ListAdvanceConfigsResponseBodyResultFiles {
	s.Name = &v
	return s
}

func (s *ListAdvanceConfigsResponseBodyResultFiles) Validate() error {
	return dara.Validate(s)
}
