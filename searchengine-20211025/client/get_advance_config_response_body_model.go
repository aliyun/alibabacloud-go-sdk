// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAdvanceConfigResponseBody
	GetRequestId() *string
	SetResult(v *GetAdvanceConfigResponseBodyResult) *GetAdvanceConfigResponseBody
	GetResult() *GetAdvanceConfigResponseBodyResult
}

type GetAdvanceConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	Result *GetAdvanceConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAdvanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAdvanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAdvanceConfigResponseBody) GetResult() *GetAdvanceConfigResponseBodyResult {
	return s.Result
}

func (s *GetAdvanceConfigResponseBody) SetRequestId(v string) *GetAdvanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdvanceConfigResponseBody) SetResult(v *GetAdvanceConfigResponseBodyResult) *GetAdvanceConfigResponseBody {
	s.Result = v
	return s
}

func (s *GetAdvanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAdvanceConfigResponseBodyResult struct {
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
	// The description of the advanced configuration.
	//
	// example:
	//
	// close alarm, chiji id 37080
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The files.
	Files []*GetAdvanceConfigResponseBodyResultFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The name of the advanced configuration.
	//
	// example:
	//
	// ha-cn-0ju2s170b03_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the advanced configuration. Valid values: drafting: The advanced configuration is in the draft state. used: The advanced configuration is being used. unused: The advanced configuration is not used. trash: The advanced configuration is being deleted.
	//
	// example:
	//
	// 0,1,3,6,8
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the advanced configuration was updated.
	//
	// example:
	//
	// ""
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetAdvanceConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetAdvanceConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *GetAdvanceConfigResponseBodyResult) GetContentType() *string {
	return s.ContentType
}

func (s *GetAdvanceConfigResponseBodyResult) GetDesc() *string {
	return s.Desc
}

func (s *GetAdvanceConfigResponseBodyResult) GetFiles() []*GetAdvanceConfigResponseBodyResultFiles {
	return s.Files
}

func (s *GetAdvanceConfigResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetAdvanceConfigResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetAdvanceConfigResponseBodyResult) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetAdvanceConfigResponseBodyResult) SetContent(v string) *GetAdvanceConfigResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetContentType(v string) *GetAdvanceConfigResponseBodyResult {
	s.ContentType = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetDesc(v string) *GetAdvanceConfigResponseBodyResult {
	s.Desc = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetFiles(v []*GetAdvanceConfigResponseBodyResultFiles) *GetAdvanceConfigResponseBodyResult {
	s.Files = v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetName(v string) *GetAdvanceConfigResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetStatus(v string) *GetAdvanceConfigResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) SetUpdateTime(v int64) *GetAdvanceConfigResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetAdvanceConfigResponseBodyResultFiles struct {
	// The file path.
	//
	// example:
	//
	// ""
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether the file is a directory.
	//
	// example:
	//
	// True
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether the file is a container.
	//
	// example:
	//
	// True
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// The file name.
	//
	// example:
	//
	// ha-cn-2r42ostoc01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAdvanceConfigResponseBodyResultFiles) String() string {
	return dara.Prettify(s)
}

func (s GetAdvanceConfigResponseBodyResultFiles) GoString() string {
	return s.String()
}

func (s *GetAdvanceConfigResponseBodyResultFiles) GetFullPathName() *string {
	return s.FullPathName
}

func (s *GetAdvanceConfigResponseBodyResultFiles) GetIsDir() *bool {
	return s.IsDir
}

func (s *GetAdvanceConfigResponseBodyResultFiles) GetIsTemplate() *bool {
	return s.IsTemplate
}

func (s *GetAdvanceConfigResponseBodyResultFiles) GetName() *string {
	return s.Name
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetFullPathName(v string) *GetAdvanceConfigResponseBodyResultFiles {
	s.FullPathName = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetIsDir(v bool) *GetAdvanceConfigResponseBodyResultFiles {
	s.IsDir = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetIsTemplate(v bool) *GetAdvanceConfigResponseBodyResultFiles {
	s.IsTemplate = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResultFiles) SetName(v string) *GetAdvanceConfigResponseBodyResultFiles {
	s.Name = &v
	return s
}

func (s *GetAdvanceConfigResponseBodyResultFiles) Validate() error {
	return dara.Validate(s)
}
