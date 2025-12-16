// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirstRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeFirstRankResponseBody
	GetRequestId() *string
	SetResult(v *DescribeFirstRankResponseBodyResult) *DescribeFirstRankResponseBody
	GetResult() *DescribeFirstRankResponseBodyResult
}

type DescribeFirstRankResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort expression.
	Result *DescribeFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeFirstRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFirstRankResponseBody) GetResult() *DescribeFirstRankResponseBodyResult {
	return s.Result
}

func (s *DescribeFirstRankResponseBody) SetRequestId(v string) *DescribeFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirstRankResponseBody) SetResult(v *DescribeFirstRankResponseBodyResult) *DescribeFirstRankResponseBody {
	s.Result = v
	return s
}

func (s *DescribeFirstRankResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFirstRankResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	//
	// example:
	//
	// false
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// Description
	//
	// example:
	//
	// -
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The information about the expression.
	Meta []*DescribeFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// Parameter
	//
	// example:
	//
	// ar_wear_edit_time
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeFirstRankResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *DescribeFirstRankResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeFirstRankResponseBodyResult) GetMeta() []*DescribeFirstRankResponseBodyResultMeta {
	return s.Meta
}

func (s *DescribeFirstRankResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeFirstRankResponseBodyResult) SetActive(v bool) *DescribeFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResult) SetDescription(v string) *DescribeFirstRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResult) SetMeta(v []*DescribeFirstRankResponseBodyResultMeta) *DescribeFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *DescribeFirstRankResponseBodyResult) SetName(v string) *DescribeFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResult) Validate() error {
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

type DescribeFirstRankResponseBodyResultMeta struct {
	// The parameters that are used by a function in the expression.
	//
	// example:
	//
	// ar_edit_time
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	//
	// example:
	//
	// timeliness_ms()
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight. Valid values: -100000 to 100000. The value cannot be 0.
	//
	// example:
	//
	// 1
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s DescribeFirstRankResponseBodyResultMeta) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponseBodyResultMeta) GetArg() *string {
	return s.Arg
}

func (s *DescribeFirstRankResponseBodyResultMeta) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeFirstRankResponseBodyResultMeta) GetWeight() *float32 {
	return s.Weight
}

func (s *DescribeFirstRankResponseBodyResultMeta) SetArg(v string) *DescribeFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResultMeta) SetAttribute(v string) *DescribeFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResultMeta) SetWeight(v float32) *DescribeFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResultMeta) Validate() error {
	return dara.Validate(s)
}
