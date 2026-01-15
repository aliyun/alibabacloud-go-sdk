// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFeatureOptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeFeatureOptionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *DescribeFeatureOptionResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *DescribeFeatureOptionResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeFeatureOptionResponseBodyResultObject) *DescribeFeatureOptionResponseBody
	GetResultObject() []*DescribeFeatureOptionResponseBodyResultObject
	SetSuccess(v bool) *DescribeFeatureOptionResponseBody
	GetSuccess() *bool
}

type DescribeFeatureOptionResponseBody struct {
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
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject []*DescribeFeatureOptionResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFeatureOptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFeatureOptionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFeatureOptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeFeatureOptionResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeFeatureOptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFeatureOptionResponseBody) GetResultObject() []*DescribeFeatureOptionResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeFeatureOptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeFeatureOptionResponseBody) SetCode(v string) *DescribeFeatureOptionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFeatureOptionResponseBody) SetHttpStatusCode(v int64) *DescribeFeatureOptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeFeatureOptionResponseBody) SetRequestId(v string) *DescribeFeatureOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFeatureOptionResponseBody) SetResultObject(v []*DescribeFeatureOptionResponseBodyResultObject) *DescribeFeatureOptionResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeFeatureOptionResponseBody) SetSuccess(v bool) *DescribeFeatureOptionResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFeatureOptionResponseBody) Validate() error {
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

type DescribeFeatureOptionResponseBodyResultObject struct {
	// example:
	//
	// xxx
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
}

func (s DescribeFeatureOptionResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeFeatureOptionResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeFeatureOptionResponseBodyResultObject) GetFeatureName() *string {
	return s.FeatureName
}

func (s *DescribeFeatureOptionResponseBodyResultObject) SetFeatureName(v string) *DescribeFeatureOptionResponseBodyResultObject {
	s.FeatureName = &v
	return s
}

func (s *DescribeFeatureOptionResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
