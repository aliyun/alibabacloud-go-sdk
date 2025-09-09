// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDbRdsNameListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceNameList(v *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) *DescribeDrdsDbRdsNameListResponseBody
	GetInstanceNameList() *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList
	SetRequestId(v string) *DescribeDrdsDbRdsNameListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsDbRdsNameListResponseBody
	GetSuccess() *bool
}

type DescribeDrdsDbRdsNameListResponseBody struct {
	// Indicates the instances that are used to store the data of a database.
	InstanceNameList *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList `json:"InstanceNameList,omitempty" xml:"InstanceNameList,omitempty" type:"Struct"`
	// Indicates the ID of the request.
	//
	// example:
	//
	// 7E6FA2BF-05F2-44DD-95C0-D1B5B8xxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDbRdsNameListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListResponseBody) GetInstanceNameList() *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList {
	return s.InstanceNameList
}

func (s *DescribeDrdsDbRdsNameListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsDbRdsNameListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetInstanceNameList(v *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) *DescribeDrdsDbRdsNameListResponseBody {
	s.InstanceNameList = v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetRequestId(v string) *DescribeDrdsDbRdsNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponseBody) SetSuccess(v bool) *DescribeDrdsDbRdsNameListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDbRdsNameListResponseBodyInstanceNameList struct {
	InstanceName []*string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) GetInstanceName() []*string {
	return s.InstanceName
}

func (s *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) SetInstanceName(v []*string) *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList {
	s.InstanceName = v
	return s
}

func (s *DescribeDrdsDbRdsNameListResponseBodyInstanceNameList) Validate() error {
	return dara.Validate(s)
}
