// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSortScriptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSortScriptsResponseBody
	GetRequestId() *string
	SetResult(v []*ListSortScriptsResponseBodyResult) *ListSortScriptsResponseBody
	GetResult() []*ListSortScriptsResponseBodyResult
}

type ListSortScriptsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The scripts.
	Result []*ListSortScriptsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSortScriptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSortScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSortScriptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSortScriptsResponseBody) GetResult() []*ListSortScriptsResponseBodyResult {
	return s.Result
}

func (s *ListSortScriptsResponseBody) SetRequestId(v string) *ListSortScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSortScriptsResponseBody) SetResult(v []*ListSortScriptsResponseBodyResult) *ListSortScriptsResponseBody {
	s.Result = v
	return s
}

func (s *ListSortScriptsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSortScriptsResponseBodyResult struct {
	// The time when the script was created.
	//
	// example:
	//
	// 2020-04-02 20:21:14
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the script was last modified.
	//
	// example:
	//
	// 2020-04-02 21:21:14
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// The sort phase to which the script applies.
	//
	// example:
	//
	// second_rank
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The name of the script.
	//
	// example:
	//
	// test
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// The status of the script. Valid values:
	//
	// 	- configurable: The script is created, but no script files are uploaded.
	//
	// 	- not compiled: The script is not compiled.
	//
	// 	- compile failed: The compilation of the script failed.
	//
	// 	- compile successful: The script is compiled.
	//
	// 	- released: The script is published.
	//
	// example:
	//
	// released
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the script.
	//
	// example:
	//
	// cava_script
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSortScriptsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSortScriptsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSortScriptsResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSortScriptsResponseBodyResult) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListSortScriptsResponseBodyResult) GetScope() *string {
	return s.Scope
}

func (s *ListSortScriptsResponseBodyResult) GetScriptName() *string {
	return s.ScriptName
}

func (s *ListSortScriptsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListSortScriptsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListSortScriptsResponseBodyResult) SetCreateTime(v string) *ListSortScriptsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetModifyTime(v string) *ListSortScriptsResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetScope(v string) *ListSortScriptsResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetScriptName(v string) *ListSortScriptsResponseBodyResult {
	s.ScriptName = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetStatus(v string) *ListSortScriptsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetType(v string) *ListSortScriptsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
