// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateShardTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ValidateShardTaskResponseBodyList) *ValidateShardTaskResponseBody
	GetList() []*ValidateShardTaskResponseBodyList
	SetRequestId(v string) *ValidateShardTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ValidateShardTaskResponseBody
	GetSuccess() *bool
}

type ValidateShardTaskResponseBody struct {
	// Indicates the check results.
	List []*ValidateShardTaskResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Indicates the ID of the request.
	//
	// example:
	//
	// 0B6B7BDC-575D-4A77-A4F8-24B7EFAS45FG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ValidateShardTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateShardTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskResponseBody) GetList() []*ValidateShardTaskResponseBodyList {
	return s.List
}

func (s *ValidateShardTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateShardTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidateShardTaskResponseBody) SetList(v []*ValidateShardTaskResponseBodyList) *ValidateShardTaskResponseBody {
	s.List = v
	return s
}

func (s *ValidateShardTaskResponseBody) SetRequestId(v string) *ValidateShardTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateShardTaskResponseBody) SetSuccess(v bool) *ValidateShardTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateShardTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ValidateShardTaskResponseBodyList struct {
	// Indicates the name of a check item.
	//
	// example:
	//
	// same_schema
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// Indicates the result of the check item. Valid values:
	//
	// 	- **0**: indicates the task is valid.
	//
	// 	- **1**: indicates the task is invalid.
	//
	// example:
	//
	// 0
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateShardTaskResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ValidateShardTaskResponseBodyList) GoString() string {
	return s.String()
}

func (s *ValidateShardTaskResponseBodyList) GetItem() *string {
	return s.Item
}

func (s *ValidateShardTaskResponseBodyList) GetResult() *int32 {
	return s.Result
}

func (s *ValidateShardTaskResponseBodyList) SetItem(v string) *ValidateShardTaskResponseBodyList {
	s.Item = &v
	return s
}

func (s *ValidateShardTaskResponseBodyList) SetResult(v int32) *ValidateShardTaskResponseBodyList {
	s.Result = &v
	return s
}

func (s *ValidateShardTaskResponseBodyList) Validate() error {
	return dara.Validate(s)
}
