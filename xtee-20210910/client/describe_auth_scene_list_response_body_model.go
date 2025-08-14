// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthSceneListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAuthSceneListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeAuthSceneListResponseBodyResultObject) *DescribeAuthSceneListResponseBody
	GetResultObject() []*DescribeAuthSceneListResponseBodyResultObject
}

type DescribeAuthSceneListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject []*DescribeAuthSceneListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeAuthSceneListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthSceneListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuthSceneListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuthSceneListResponseBody) GetResultObject() []*DescribeAuthSceneListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAuthSceneListResponseBody) SetRequestId(v string) *DescribeAuthSceneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuthSceneListResponseBody) SetResultObject(v []*DescribeAuthSceneListResponseBodyResultObject) *DescribeAuthSceneListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAuthSceneListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAuthSceneListResponseBodyResultObject struct {
	// Service code
	//
	// example:
	//
	// account_abuse
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
}

func (s DescribeAuthSceneListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthSceneListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAuthSceneListResponseBodyResultObject) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeAuthSceneListResponseBodyResultObject) SetServiceCode(v string) *DescribeAuthSceneListResponseBodyResultObject {
	s.ServiceCode = &v
	return s
}

func (s *DescribeAuthSceneListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
