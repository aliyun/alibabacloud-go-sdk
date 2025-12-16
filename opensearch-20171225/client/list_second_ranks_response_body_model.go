// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecondRanksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSecondRanksResponseBody
	GetRequestId() *string
	SetResult(v []*ListSecondRanksResponseBodyResult) *ListSecondRanksResponseBody
	GetResult() []*ListSecondRanksResponseBodyResult
	SetTotalCount(v int32) *ListSecondRanksResponseBody
	GetTotalCount() *int32
}

type ListSecondRanksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the fine sort expression.
	//
	// For more information, see [SecondRank](https://help.aliyun.com/document_detail/170008.html).
	Result []*ListSecondRanksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSecondRanksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecondRanksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecondRanksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecondRanksResponseBody) GetResult() []*ListSecondRanksResponseBodyResult {
	return s.Result
}

func (s *ListSecondRanksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSecondRanksResponseBody) SetRequestId(v string) *ListSecondRanksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecondRanksResponseBody) SetResult(v []*ListSecondRanksResponseBodyResult) *ListSecondRanksResponseBody {
	s.Result = v
	return s
}

func (s *ListSecondRanksResponseBody) SetTotalCount(v int32) *ListSecondRanksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecondRanksResponseBody) Validate() error {
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

type ListSecondRanksResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	//
	// example:
	//
	// false
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the expression was created.
	//
	// example:
	//
	// 0
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// Description
	//
	// example:
	//
	// ""
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
	// false
	IsDefault *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// Indicates whether the expression is a system expression. This parameter is displayed only in the response. Valid values:
	//
	// 	- true: The expression is a system expression.
	//
	// 	- false:The expression is not a system expression
	//
	// example:
	//
	// true
	IsSys *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	// The content of the fine sort expression. You can define an expression that consists of fields, feature functions, and mathematical functions to implement complex sort logic.
	//
	// example:
	//
	// random()+now()
	Meta *string `json:"meta,omitempty" xml:"meta,omitempty"`
	// Parameter
	//
	// example:
	//
	// tests
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The time when the expression was updated.
	//
	// example:
	//
	// 1587052801
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListSecondRanksResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSecondRanksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSecondRanksResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *ListSecondRanksResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListSecondRanksResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListSecondRanksResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListSecondRanksResponseBodyResult) GetIsDefault() *string {
	return s.IsDefault
}

func (s *ListSecondRanksResponseBodyResult) GetIsSys() *string {
	return s.IsSys
}

func (s *ListSecondRanksResponseBodyResult) GetMeta() *string {
	return s.Meta
}

func (s *ListSecondRanksResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListSecondRanksResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListSecondRanksResponseBodyResult) SetActive(v bool) *ListSecondRanksResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetCreated(v int32) *ListSecondRanksResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetDescription(v string) *ListSecondRanksResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetId(v string) *ListSecondRanksResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetIsDefault(v string) *ListSecondRanksResponseBodyResult {
	s.IsDefault = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetIsSys(v string) *ListSecondRanksResponseBodyResult {
	s.IsSys = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetMeta(v string) *ListSecondRanksResponseBodyResult {
	s.Meta = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetName(v string) *ListSecondRanksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetUpdated(v int32) *ListSecondRanksResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
