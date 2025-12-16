// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecondRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySecondRankResponseBody
	GetRequestId() *string
	SetResult(v *ModifySecondRankResponseBodyResult) *ModifySecondRankResponseBody
	GetResult() *ModifySecondRankResponseBodyResult
}

type ModifySecondRankResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C5E2F73C-8241-81F8-3A62-65478C5A3111
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the fine sort expression.
	//
	// example:
	//
	// {}
	Result *ModifySecondRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifySecondRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecondRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecondRankResponseBody) GetResult() *ModifySecondRankResponseBodyResult {
	return s.Result
}

func (s *ModifySecondRankResponseBody) SetRequestId(v string) *ModifySecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecondRankResponseBody) SetResult(v *ModifySecondRankResponseBodyResult) *ModifySecondRankResponseBody {
	s.Result = v
	return s
}

func (s *ModifySecondRankResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifySecondRankResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the expression was created.
	//
	// example:
	//
	// 1
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the expression.
	//
	// example:
	//
	// "11"
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The expression ID. This parameter is displayed only in the response.
	//
	// example:
	//
	// 890473
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the expression is the default one. This parameter is displayed only in the response. Valid values:
	//
	// 	- true: the expression is the default one.
	//
	// 	- false: the expression is not the default one.
	//
	// example:
	//
	// true
	IsDefault *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// Indicates whether the expression is a system expression. This parameter is displayed only in the response. Valid values:
	//
	// 	- true: The expression is a system expression.
	//
	// 	- false:The expression is not a system expression
	//
	// example:
	//
	// false
	IsSys *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	// The content of the fine sort expression. You can define an expression that consists of fields, feature functions, and mathematical functions to implement complex sort logic.
	//
	// example:
	//
	// cate_id > 0 and cate_id < 1000
	Meta *string `json:"meta,omitempty" xml:"meta,omitempty"`
	// The expression name.
	//
	// example:
	//
	// lsh_second_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The time when the expression was updated.
	//
	// example:
	//
	// 1
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifySecondRankResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ModifySecondRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifySecondRankResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *ModifySecondRankResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ModifySecondRankResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ModifySecondRankResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ModifySecondRankResponseBodyResult) GetIsDefault() *string {
	return s.IsDefault
}

func (s *ModifySecondRankResponseBodyResult) GetIsSys() *string {
	return s.IsSys
}

func (s *ModifySecondRankResponseBodyResult) GetMeta() *string {
	return s.Meta
}

func (s *ModifySecondRankResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ModifySecondRankResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ModifySecondRankResponseBodyResult) SetActive(v bool) *ModifySecondRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetCreated(v int32) *ModifySecondRankResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetDescription(v string) *ModifySecondRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetId(v string) *ModifySecondRankResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetIsDefault(v string) *ModifySecondRankResponseBodyResult {
	s.IsDefault = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetIsSys(v string) *ModifySecondRankResponseBodyResult {
	s.IsSys = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetMeta(v string) *ModifySecondRankResponseBodyResult {
	s.Meta = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetName(v string) *ModifySecondRankResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetUpdated(v int32) *ModifySecondRankResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
