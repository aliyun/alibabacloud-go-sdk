// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFirstRanksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListFirstRanksResponseBody
	GetRequestId() *string
	SetResult(v []*ListFirstRanksResponseBodyResult) *ListFirstRanksResponseBody
	GetResult() []*ListFirstRanksResponseBodyResult
}

type ListFirstRanksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort expression.
	//
	// For more information, see [FirstRank](https://help.aliyun.com/document_detail/170007.html).
	//
	// example:
	//
	// []
	Result []*ListFirstRanksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListFirstRanksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFirstRanksResponseBody) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFirstRanksResponseBody) GetResult() []*ListFirstRanksResponseBodyResult {
	return s.Result
}

func (s *ListFirstRanksResponseBody) SetRequestId(v string) *ListFirstRanksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFirstRanksResponseBody) SetResult(v []*ListFirstRanksResponseBodyResult) *ListFirstRanksResponseBody {
	s.Result = v
	return s
}

func (s *ListFirstRanksResponseBody) Validate() error {
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

type ListFirstRanksResponseBodyResult struct {
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
	// 0
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// Description
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The information about the expression.
	//
	// example:
	//
	// []
	Meta []*ListFirstRanksResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name.
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The time when the expression was updated.
	//
	// example:
	//
	// 0
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListFirstRanksResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListFirstRanksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *ListFirstRanksResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListFirstRanksResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListFirstRanksResponseBodyResult) GetMeta() []*ListFirstRanksResponseBodyResultMeta {
	return s.Meta
}

func (s *ListFirstRanksResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListFirstRanksResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListFirstRanksResponseBodyResult) SetActive(v bool) *ListFirstRanksResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetCreated(v int32) *ListFirstRanksResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetDescription(v string) *ListFirstRanksResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetMeta(v []*ListFirstRanksResponseBodyResultMeta) *ListFirstRanksResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetName(v string) *ListFirstRanksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetUpdated(v int32) *ListFirstRanksResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) Validate() error {
	if s.Meta != nil {
		for _, item := range s.Meta {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFirstRanksResponseBodyResultMeta struct {
	// The parameters that are used by a function in the expression.
	//
	// For more information, see [Rough sort functions](https://help.aliyun.com/document_detail/180765.html).
	//
	// example:
	//
	// ""
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	//
	// For more information about supported feature functions, see [Rough sort functions](https://help.aliyun.com/document_detail/180765.html).
	//
	// example:
	//
	// static_bm25()
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight. Valid values: -100000 to 100000. The value cannot be 0.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ListFirstRanksResponseBodyResultMeta) String() string {
	return dara.Prettify(s)
}

func (s ListFirstRanksResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponseBodyResultMeta) GetArg() *string {
	return s.Arg
}

func (s *ListFirstRanksResponseBodyResultMeta) GetAttribute() *string {
	return s.Attribute
}

func (s *ListFirstRanksResponseBodyResultMeta) GetWeight() *int32 {
	return s.Weight
}

func (s *ListFirstRanksResponseBodyResultMeta) SetArg(v string) *ListFirstRanksResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *ListFirstRanksResponseBodyResultMeta) SetAttribute(v string) *ListFirstRanksResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *ListFirstRanksResponseBodyResultMeta) SetWeight(v int32) *ListFirstRanksResponseBodyResultMeta {
	s.Weight = &v
	return s
}

func (s *ListFirstRanksResponseBodyResultMeta) Validate() error {
	return dara.Validate(s)
}
