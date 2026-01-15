// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFeatureTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeFeatureTemplateListResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *DescribeFeatureTemplateListResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *DescribeFeatureTemplateListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeFeatureTemplateListResponseBodyResultObject) *DescribeFeatureTemplateListResponseBody
	GetResultObject() []*DescribeFeatureTemplateListResponseBodyResultObject
	SetSuccess(v bool) *DescribeFeatureTemplateListResponseBody
	GetSuccess() *bool
}

type DescribeFeatureTemplateListResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId    *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject []*DescribeFeatureTemplateListResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFeatureTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFeatureTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFeatureTemplateListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeFeatureTemplateListResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeFeatureTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFeatureTemplateListResponseBody) GetResultObject() []*DescribeFeatureTemplateListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeFeatureTemplateListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeFeatureTemplateListResponseBody) SetCode(v string) *DescribeFeatureTemplateListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFeatureTemplateListResponseBody) SetHttpStatusCode(v int64) *DescribeFeatureTemplateListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeFeatureTemplateListResponseBody) SetRequestId(v string) *DescribeFeatureTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFeatureTemplateListResponseBody) SetResultObject(v []*DescribeFeatureTemplateListResponseBodyResultObject) *DescribeFeatureTemplateListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeFeatureTemplateListResponseBody) SetSuccess(v bool) *DescribeFeatureTemplateListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFeatureTemplateListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFeatureTemplateListResponseBodyResultObject struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// FINANCE_51
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFeatureTemplateListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeFeatureTemplateListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeFeatureTemplateListResponseBodyResultObject) GetLabel() *string {
	return s.Label
}

func (s *DescribeFeatureTemplateListResponseBodyResultObject) GetValue() *string {
	return s.Value
}

func (s *DescribeFeatureTemplateListResponseBodyResultObject) SetLabel(v string) *DescribeFeatureTemplateListResponseBodyResultObject {
	s.Label = &v
	return s
}

func (s *DescribeFeatureTemplateListResponseBodyResultObject) SetValue(v string) *DescribeFeatureTemplateListResponseBodyResultObject {
	s.Value = &v
	return s
}

func (s *DescribeFeatureTemplateListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
