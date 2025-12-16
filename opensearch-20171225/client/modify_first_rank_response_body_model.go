// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirstRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyFirstRankResponseBody
	GetRequestId() *string
	SetResult(v *ModifyFirstRankResponseBodyResult) *ModifyFirstRankResponseBody
	GetResult() *ModifyFirstRankResponseBodyResult
}

type ModifyFirstRankResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort expression.
	Result *ModifyFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyFirstRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFirstRankResponseBody) GetResult() *ModifyFirstRankResponseBodyResult {
	return s.Result
}

func (s *ModifyFirstRankResponseBody) SetRequestId(v string) *ModifyFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFirstRankResponseBody) SetResult(v *ModifyFirstRankResponseBodyResult) *ModifyFirstRankResponseBody {
	s.Result = v
	return s
}

func (s *ModifyFirstRankResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyFirstRankResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The description of the rough sort expression.
	//
	// example:
	//
	// 1
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The information about the expression.
	Meta []*ModifyFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModifyFirstRankResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *ModifyFirstRankResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ModifyFirstRankResponseBodyResult) GetMeta() []*ModifyFirstRankResponseBodyResultMeta {
	return s.Meta
}

func (s *ModifyFirstRankResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ModifyFirstRankResponseBodyResult) SetActive(v bool) *ModifyFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResult) SetDescription(v string) *ModifyFirstRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResult) SetMeta(v []*ModifyFirstRankResponseBodyResultMeta) *ModifyFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ModifyFirstRankResponseBodyResult) SetName(v string) *ModifyFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResult) Validate() error {
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

type ModifyFirstRankResponseBodyResultMeta struct {
	// The parameters that are used by a function in the expression.
	//
	// example:
	//
	// “1 ”
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	//
	// example:
	//
	// static_bm25()
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight. Valid values: -100000 to 100000. The value cannot be 0.
	//
	// example:
	//
	// 10
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ModifyFirstRankResponseBodyResultMeta) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponseBodyResultMeta) GetArg() *string {
	return s.Arg
}

func (s *ModifyFirstRankResponseBodyResultMeta) GetAttribute() *string {
	return s.Attribute
}

func (s *ModifyFirstRankResponseBodyResultMeta) GetWeight() *float32 {
	return s.Weight
}

func (s *ModifyFirstRankResponseBodyResultMeta) SetArg(v string) *ModifyFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResultMeta) SetAttribute(v string) *ModifyFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResultMeta) SetWeight(v float32) *ModifyFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResultMeta) Validate() error {
	return dara.Validate(s)
}
