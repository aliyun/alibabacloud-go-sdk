// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdvanceConfigDirResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAdvanceConfigDirResponseBody
	GetRequestId() *string
	SetResult(v []*ListAdvanceConfigDirResponseBodyResult) *ListAdvanceConfigDirResponseBody
	GetResult() []*ListAdvanceConfigDirResponseBodyResult
}

type ListAdvanceConfigDirResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The advanced configuration files.
	Result []*ListAdvanceConfigDirResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAdvanceConfigDirResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAdvanceConfigDirResponseBody) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAdvanceConfigDirResponseBody) GetResult() []*ListAdvanceConfigDirResponseBodyResult {
	return s.Result
}

func (s *ListAdvanceConfigDirResponseBody) SetRequestId(v string) *ListAdvanceConfigDirResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBody) SetResult(v []*ListAdvanceConfigDirResponseBodyResult) *ListAdvanceConfigDirResponseBody {
	s.Result = v
	return s
}

func (s *ListAdvanceConfigDirResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAdvanceConfigDirResponseBodyResult struct {
	// The absolute path in which the file is stored.
	//
	// example:
	//
	// "/path/wpd/nae"
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether the file is a directory. Valid values: true and false.
	//
	// example:
	//
	// true
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether the file is a template. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// true
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// file_name_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAdvanceConfigDirResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAdvanceConfigDirResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirResponseBodyResult) GetFullPathName() *string {
	return s.FullPathName
}

func (s *ListAdvanceConfigDirResponseBodyResult) GetIsDir() *bool {
	return s.IsDir
}

func (s *ListAdvanceConfigDirResponseBodyResult) GetIsTemplate() *bool {
	return s.IsTemplate
}

func (s *ListAdvanceConfigDirResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetFullPathName(v string) *ListAdvanceConfigDirResponseBodyResult {
	s.FullPathName = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetIsDir(v bool) *ListAdvanceConfigDirResponseBodyResult {
	s.IsDir = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetIsTemplate(v bool) *ListAdvanceConfigDirResponseBodyResult {
	s.IsTemplate = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBodyResult) SetName(v string) *ListAdvanceConfigDirResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListAdvanceConfigDirResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
