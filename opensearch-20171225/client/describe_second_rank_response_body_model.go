// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecondRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSecondRankResponseBody
	GetRequestId() *string
	SetResult(v *DescribeSecondRankResponseBodyResult) *DescribeSecondRankResponseBody
	GetResult() *DescribeSecondRankResponseBodyResult
}

type DescribeSecondRankResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the fine sort expression.
	Result *DescribeSecondRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeSecondRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecondRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecondRankResponseBody) GetResult() *DescribeSecondRankResponseBodyResult {
	return s.Result
}

func (s *DescribeSecondRankResponseBody) SetRequestId(v string) *DescribeSecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecondRankResponseBody) SetResult(v *DescribeSecondRankResponseBodyResult) *DescribeSecondRankResponseBody {
	s.Result = v
	return s
}

func (s *DescribeSecondRankResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecondRankResponseBodyResult struct {
	// Specifies whether to set the fine sort expression as the default sort expression.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the expression was created.
	//
	// example:
	//
	// 1587052801
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// Description
	//
	// example:
	//
	// -
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the expression. This parameter appears only in the response.
	//
	// example:
	//
	// 89047
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the expression is the default one. This parameter appears only in the response. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsDefault *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// Indicates whether the expression is a system expression. This parameter appears only in the response. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsSys *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	// The content of the fine sort expression. You can define an expression that consists of fields, feature functions, and mathematical functions to implement complex sort logic.
	//
	// example:
	//
	// random()+now()
	Meta *string `json:"meta,omitempty" xml:"meta,omitempty"`
	// The name.
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

func (s DescribeSecondRankResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecondRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSecondRankResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *DescribeSecondRankResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *DescribeSecondRankResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecondRankResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *DescribeSecondRankResponseBodyResult) GetIsDefault() *string {
	return s.IsDefault
}

func (s *DescribeSecondRankResponseBodyResult) GetIsSys() *string {
	return s.IsSys
}

func (s *DescribeSecondRankResponseBodyResult) GetMeta() *string {
	return s.Meta
}

func (s *DescribeSecondRankResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeSecondRankResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *DescribeSecondRankResponseBodyResult) SetActive(v bool) *DescribeSecondRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetCreated(v int32) *DescribeSecondRankResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetDescription(v string) *DescribeSecondRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetId(v string) *DescribeSecondRankResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetIsDefault(v string) *DescribeSecondRankResponseBodyResult {
	s.IsDefault = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetIsSys(v string) *DescribeSecondRankResponseBodyResult {
	s.IsSys = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetMeta(v string) *DescribeSecondRankResponseBodyResult {
	s.Meta = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetName(v string) *DescribeSecondRankResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetUpdated(v int32) *DescribeSecondRankResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
