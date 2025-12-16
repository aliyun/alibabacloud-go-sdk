// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirstRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateFirstRankResponseBody
	GetRequestId() *string
	SetResult(v *CreateFirstRankResponseBodyResult) *CreateFirstRankResponseBody
	GetResult() *CreateFirstRankResponseBodyResult
}

type CreateFirstRankResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort expression.
	Result *CreateFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateFirstRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFirstRankResponseBody) GetResult() *CreateFirstRankResponseBodyResult {
	return s.Result
}

func (s *CreateFirstRankResponseBody) SetRequestId(v string) *CreateFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFirstRankResponseBody) SetResult(v *CreateFirstRankResponseBodyResult) *CreateFirstRankResponseBody {
	s.Result = v
	return s
}

func (s *CreateFirstRankResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFirstRankResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The content of the expression.
	Meta []*CreateFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateFirstRankResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *CreateFirstRankResponseBodyResult) GetMeta() []*CreateFirstRankResponseBodyResultMeta {
	return s.Meta
}

func (s *CreateFirstRankResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateFirstRankResponseBodyResult) SetActive(v bool) *CreateFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *CreateFirstRankResponseBodyResult) SetMeta(v []*CreateFirstRankResponseBodyResultMeta) *CreateFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *CreateFirstRankResponseBodyResult) SetName(v string) *CreateFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateFirstRankResponseBodyResult) Validate() error {
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

type CreateFirstRankResponseBodyResultMeta struct {
	// The parameters that are used by a function in the expression.
	//
	// example:
	//
	// 1
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature functions, or field to be searched for.
	//
	// example:
	//
	// static_bm25()
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight. Valid values: [-100000,100000]. The value cannot be 0.
	//
	// example:
	//
	// 10
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s CreateFirstRankResponseBodyResultMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponseBodyResultMeta) GetArg() *string {
	return s.Arg
}

func (s *CreateFirstRankResponseBodyResultMeta) GetAttribute() *string {
	return s.Attribute
}

func (s *CreateFirstRankResponseBodyResultMeta) GetWeight() *float32 {
	return s.Weight
}

func (s *CreateFirstRankResponseBodyResultMeta) SetArg(v string) *CreateFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *CreateFirstRankResponseBodyResultMeta) SetAttribute(v string) *CreateFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *CreateFirstRankResponseBodyResultMeta) SetWeight(v float32) *CreateFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

func (s *CreateFirstRankResponseBodyResultMeta) Validate() error {
	return dara.Validate(s)
}
