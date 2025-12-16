// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSortScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSortScriptResponseBody
	GetRequestId() *string
	SetResult(v *GetSortScriptResponseBodyResult) *GetSortScriptResponseBody
	GetResult() *GetSortScriptResponseBodyResult
}

type GetSortScriptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the script.
	Result *GetSortScriptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSortScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *GetSortScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSortScriptResponseBody) GetResult() *GetSortScriptResponseBodyResult {
	return s.Result
}

func (s *GetSortScriptResponseBody) SetRequestId(v string) *GetSortScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSortScriptResponseBody) SetResult(v *GetSortScriptResponseBodyResult) *GetSortScriptResponseBody {
	s.Result = v
	return s
}

func (s *GetSortScriptResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSortScriptResponseBodyResult struct {
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
	// rank_cava_20230606_v7
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// The status of the script. For more information, see the description of the status response parameter in the ListSortScripts topic.
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

func (s GetSortScriptResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetSortScriptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSortScriptResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSortScriptResponseBodyResult) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetSortScriptResponseBodyResult) GetScope() *string {
	return s.Scope
}

func (s *GetSortScriptResponseBodyResult) GetScriptName() *string {
	return s.ScriptName
}

func (s *GetSortScriptResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetSortScriptResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetSortScriptResponseBodyResult) SetCreateTime(v string) *GetSortScriptResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetModifyTime(v string) *GetSortScriptResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetScope(v string) *GetSortScriptResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetScriptName(v string) *GetSortScriptResponseBodyResult {
	s.ScriptName = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetStatus(v string) *GetSortScriptResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetType(v string) *GetSortScriptResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
