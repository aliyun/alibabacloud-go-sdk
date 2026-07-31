// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySelectOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QuerySelectOptionsResponseBody
	GetRequestId() *string
	SetSelectOptions(v []*QuerySelectOptionsResponseBodySelectOptions) *QuerySelectOptionsResponseBody
	GetSelectOptions() []*QuerySelectOptionsResponseBodySelectOptions
}

type QuerySelectOptionsResponseBody struct {
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SelectOptions []*QuerySelectOptionsResponseBodySelectOptions `json:"SelectOptions,omitempty" xml:"SelectOptions,omitempty" type:"Repeated"`
}

func (s QuerySelectOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySelectOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySelectOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySelectOptionsResponseBody) GetSelectOptions() []*QuerySelectOptionsResponseBodySelectOptions {
	return s.SelectOptions
}

func (s *QuerySelectOptionsResponseBody) SetRequestId(v string) *QuerySelectOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySelectOptionsResponseBody) SetSelectOptions(v []*QuerySelectOptionsResponseBodySelectOptions) *QuerySelectOptionsResponseBody {
	s.SelectOptions = v
	return s
}

func (s *QuerySelectOptionsResponseBody) Validate() error {
	if s.SelectOptions != nil {
		for _, item := range s.SelectOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySelectOptionsResponseBodySelectOptions struct {
	BizConfig   map[string]interface{}                                 `json:"BizConfig,omitempty" xml:"BizConfig,omitempty"`
	Category    *string                                                `json:"Category,omitempty" xml:"Category,omitempty"`
	Children    []*QuerySelectOptionsResponseBodySelectOptionsChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	Description *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string                                                `json:"Label,omitempty" xml:"Label,omitempty"`
	Tags        []*string                                              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Value       *string                                                `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QuerySelectOptionsResponseBodySelectOptions) String() string {
	return dara.Prettify(s)
}

func (s QuerySelectOptionsResponseBodySelectOptions) GoString() string {
	return s.String()
}

func (s *QuerySelectOptionsResponseBodySelectOptions) GetBizConfig() map[string]interface{} {
	return s.BizConfig
}

func (s *QuerySelectOptionsResponseBodySelectOptions) GetCategory() *string {
	return s.Category
}

func (s *QuerySelectOptionsResponseBodySelectOptions) GetChildren() []*QuerySelectOptionsResponseBodySelectOptionsChildren {
	return s.Children
}

func (s *QuerySelectOptionsResponseBodySelectOptions) GetDescription() *string {
	return s.Description
}

func (s *QuerySelectOptionsResponseBodySelectOptions) GetLabel() *string {
	return s.Label
}

func (s *QuerySelectOptionsResponseBodySelectOptions) GetTags() []*string {
	return s.Tags
}

func (s *QuerySelectOptionsResponseBodySelectOptions) GetValue() *string {
	return s.Value
}

func (s *QuerySelectOptionsResponseBodySelectOptions) SetBizConfig(v map[string]interface{}) *QuerySelectOptionsResponseBodySelectOptions {
	s.BizConfig = v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptions) SetCategory(v string) *QuerySelectOptionsResponseBodySelectOptions {
	s.Category = &v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptions) SetChildren(v []*QuerySelectOptionsResponseBodySelectOptionsChildren) *QuerySelectOptionsResponseBodySelectOptions {
	s.Children = v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptions) SetDescription(v string) *QuerySelectOptionsResponseBodySelectOptions {
	s.Description = &v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptions) SetLabel(v string) *QuerySelectOptionsResponseBodySelectOptions {
	s.Label = &v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptions) SetTags(v []*string) *QuerySelectOptionsResponseBodySelectOptions {
	s.Tags = v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptions) SetValue(v string) *QuerySelectOptionsResponseBodySelectOptions {
	s.Value = &v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptions) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySelectOptionsResponseBodySelectOptionsChildren struct {
	BizConfig   map[string]interface{} `json:"BizConfig,omitempty" xml:"BizConfig,omitempty"`
	Category    *string                `json:"Category,omitempty" xml:"Category,omitempty"`
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string                `json:"Label,omitempty" xml:"Label,omitempty"`
	Tags        []*string              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Value       *string                `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QuerySelectOptionsResponseBodySelectOptionsChildren) String() string {
	return dara.Prettify(s)
}

func (s QuerySelectOptionsResponseBodySelectOptionsChildren) GoString() string {
	return s.String()
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) GetBizConfig() map[string]interface{} {
	return s.BizConfig
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) GetCategory() *string {
	return s.Category
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) GetDescription() *string {
	return s.Description
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) GetLabel() *string {
	return s.Label
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) GetTags() []*string {
	return s.Tags
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) GetValue() *string {
	return s.Value
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) SetBizConfig(v map[string]interface{}) *QuerySelectOptionsResponseBodySelectOptionsChildren {
	s.BizConfig = v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) SetCategory(v string) *QuerySelectOptionsResponseBodySelectOptionsChildren {
	s.Category = &v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) SetDescription(v string) *QuerySelectOptionsResponseBodySelectOptionsChildren {
	s.Description = &v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) SetLabel(v string) *QuerySelectOptionsResponseBodySelectOptionsChildren {
	s.Label = &v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) SetTags(v []*string) *QuerySelectOptionsResponseBodySelectOptionsChildren {
	s.Tags = v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) SetValue(v string) *QuerySelectOptionsResponseBodySelectOptionsChildren {
	s.Value = &v
	return s
}

func (s *QuerySelectOptionsResponseBodySelectOptionsChildren) Validate() error {
	return dara.Validate(s)
}
