// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFirstRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveFirstRankResponseBody
	GetRequestId() *string
	SetResult(v *RemoveFirstRankResponseBodyResult) *RemoveFirstRankResponseBody
	GetResult() *RemoveFirstRankResponseBodyResult
}

type RemoveFirstRankResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E676FAB6-A0AC-64D9-F9D7-D0D33C930CFF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort expression.
	Result *RemoveFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s RemoveFirstRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveFirstRankResponseBody) GetResult() *RemoveFirstRankResponseBodyResult {
	return s.Result
}

func (s *RemoveFirstRankResponseBody) SetRequestId(v string) *RemoveFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFirstRankResponseBody) SetResult(v *RemoveFirstRankResponseBodyResult) *RemoveFirstRankResponseBody {
	s.Result = v
	return s
}

func (s *RemoveFirstRankResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveFirstRankResponseBodyResult struct {
	// Specifies whether to set the fine sort expression as the default sort expression.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// Description
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The information about the expression.
	Meta []*RemoveFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// Parameter
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RemoveFirstRankResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s RemoveFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *RemoveFirstRankResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *RemoveFirstRankResponseBodyResult) GetMeta() []*RemoveFirstRankResponseBodyResultMeta {
	return s.Meta
}

func (s *RemoveFirstRankResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *RemoveFirstRankResponseBodyResult) SetActive(v bool) *RemoveFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResult) SetDescription(v string) *RemoveFirstRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResult) SetMeta(v []*RemoveFirstRankResponseBodyResultMeta) *RemoveFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *RemoveFirstRankResponseBodyResult) SetName(v string) *RemoveFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResult) Validate() error {
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

type RemoveFirstRankResponseBodyResultMeta struct {
	// The parameters that are used by a function in the expression. For more information, see Rough sort functions.[](~~170007~~)
	//
	// example:
	//
	// ""
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, which refers to the scoring feature or search field, For more information about supported feature functions, see Rough sort functions.[](~~170007~~)
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

func (s RemoveFirstRankResponseBodyResultMeta) String() string {
	return dara.Prettify(s)
}

func (s RemoveFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponseBodyResultMeta) GetArg() *string {
	return s.Arg
}

func (s *RemoveFirstRankResponseBodyResultMeta) GetAttribute() *string {
	return s.Attribute
}

func (s *RemoveFirstRankResponseBodyResultMeta) GetWeight() *float32 {
	return s.Weight
}

func (s *RemoveFirstRankResponseBodyResultMeta) SetArg(v string) *RemoveFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResultMeta) SetAttribute(v string) *RemoveFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResultMeta) SetWeight(v float32) *RemoveFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResultMeta) Validate() error {
	return dara.Validate(s)
}
