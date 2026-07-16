// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatas(v []*DescribeRuleMetadataResponseBodyDatas) *DescribeRuleMetadataResponseBody
	GetDatas() []*DescribeRuleMetadataResponseBodyDatas
	SetRequestId(v string) *DescribeRuleMetadataResponseBody
	GetRequestId() *string
}

type DescribeRuleMetadataResponseBody struct {
	// The list of metadata.
	Datas []*DescribeRuleMetadataResponseBodyDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B6947DF8-7AC0-50D0-BADA-177646ABB85A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRuleMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleMetadataResponseBody) GetDatas() []*DescribeRuleMetadataResponseBodyDatas {
	return s.Datas
}

func (s *DescribeRuleMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleMetadataResponseBody) SetDatas(v []*DescribeRuleMetadataResponseBodyDatas) *DescribeRuleMetadataResponseBody {
	s.Datas = v
	return s
}

func (s *DescribeRuleMetadataResponseBody) SetRequestId(v string) *DescribeRuleMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleMetadataResponseBody) Validate() error {
	if s.Datas != nil {
		for _, item := range s.Datas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRuleMetadataResponseBodyDatas struct {
	// The subset of metadata.
	Children []*DescribeRuleMetadataResponseBodyDatasChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The human-readable content.
	//
	// example:
	//
	// 中国
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The value.
	//
	// example:
	//
	// CN
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRuleMetadataResponseBodyDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleMetadataResponseBodyDatas) GoString() string {
	return s.String()
}

func (s *DescribeRuleMetadataResponseBodyDatas) GetChildren() []*DescribeRuleMetadataResponseBodyDatasChildren {
	return s.Children
}

func (s *DescribeRuleMetadataResponseBodyDatas) GetText() *string {
	return s.Text
}

func (s *DescribeRuleMetadataResponseBodyDatas) GetValue() *string {
	return s.Value
}

func (s *DescribeRuleMetadataResponseBodyDatas) SetChildren(v []*DescribeRuleMetadataResponseBodyDatasChildren) *DescribeRuleMetadataResponseBodyDatas {
	s.Children = v
	return s
}

func (s *DescribeRuleMetadataResponseBodyDatas) SetText(v string) *DescribeRuleMetadataResponseBodyDatas {
	s.Text = &v
	return s
}

func (s *DescribeRuleMetadataResponseBodyDatas) SetValue(v string) *DescribeRuleMetadataResponseBodyDatas {
	s.Value = &v
	return s
}

func (s *DescribeRuleMetadataResponseBodyDatas) Validate() error {
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

type DescribeRuleMetadataResponseBodyDatasChildren struct {
	// The human-readable content.
	//
	// example:
	//
	// 北京市
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The value.
	//
	// example:
	//
	// CN-BJ
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRuleMetadataResponseBodyDatasChildren) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleMetadataResponseBodyDatasChildren) GoString() string {
	return s.String()
}

func (s *DescribeRuleMetadataResponseBodyDatasChildren) GetText() *string {
	return s.Text
}

func (s *DescribeRuleMetadataResponseBodyDatasChildren) GetValue() *string {
	return s.Value
}

func (s *DescribeRuleMetadataResponseBodyDatasChildren) SetText(v string) *DescribeRuleMetadataResponseBodyDatasChildren {
	s.Text = &v
	return s
}

func (s *DescribeRuleMetadataResponseBodyDatasChildren) SetValue(v string) *DescribeRuleMetadataResponseBodyDatasChildren {
	s.Value = &v
	return s
}

func (s *DescribeRuleMetadataResponseBodyDatasChildren) Validate() error {
	return dara.Validate(s)
}
