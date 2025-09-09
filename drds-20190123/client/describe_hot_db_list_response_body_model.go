// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHotDbListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeHotDbListResponseBodyData) *DescribeHotDbListResponseBody
	GetData() *DescribeHotDbListResponseBodyData
	SetMsg(v string) *DescribeHotDbListResponseBody
	GetMsg() *string
	SetRequestId(v string) *DescribeHotDbListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeHotDbListResponseBody
	GetSuccess() *bool
}

type DescribeHotDbListResponseBody struct {
	// The result that is returned.
	Data *DescribeHotDbListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned.
	//
	// example:
	//
	// msg
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0B6B7BDC-575D-4A77-A4F8-24B7EF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHotDbListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotDbListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBody) GetData() *DescribeHotDbListResponseBodyData {
	return s.Data
}

func (s *DescribeHotDbListResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DescribeHotDbListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHotDbListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHotDbListResponseBody) SetData(v *DescribeHotDbListResponseBodyData) *DescribeHotDbListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeHotDbListResponseBody) SetMsg(v string) *DescribeHotDbListResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeHotDbListResponseBody) SetRequestId(v string) *DescribeHotDbListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHotDbListResponseBody) SetSuccess(v bool) *DescribeHotDbListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHotDbListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHotDbListResponseBodyData struct {
	// The information about the databases on which hot-spot scale-out is performed.
	List *DescribeHotDbListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The random number.
	//
	// example:
	//
	// jzhz
	RandomCode *string `json:"RandomCode,omitempty" xml:"RandomCode,omitempty"`
}

func (s DescribeHotDbListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotDbListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyData) GetList() *DescribeHotDbListResponseBodyDataList {
	return s.List
}

func (s *DescribeHotDbListResponseBodyData) GetRandomCode() *string {
	return s.RandomCode
}

func (s *DescribeHotDbListResponseBodyData) SetList(v *DescribeHotDbListResponseBodyDataList) *DescribeHotDbListResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeHotDbListResponseBodyData) SetRandomCode(v string) *DescribeHotDbListResponseBodyData {
	s.RandomCode = &v
	return s
}

func (s *DescribeHotDbListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeHotDbListResponseBodyDataList struct {
	InstanceDb []*DescribeHotDbListResponseBodyDataListInstanceDb `json:"InstanceDb,omitempty" xml:"InstanceDb,omitempty" type:"Repeated"`
}

func (s DescribeHotDbListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotDbListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyDataList) GetInstanceDb() []*DescribeHotDbListResponseBodyDataListInstanceDb {
	return s.InstanceDb
}

func (s *DescribeHotDbListResponseBodyDataList) SetInstanceDb(v []*DescribeHotDbListResponseBodyDataListInstanceDb) *DescribeHotDbListResponseBodyDataList {
	s.InstanceDb = v
	return s
}

func (s *DescribeHotDbListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeHotDbListResponseBodyDataListInstanceDb struct {
	HotDbList *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList `json:"HotDbList,omitempty" xml:"HotDbList,omitempty" type:"Struct"`
	// The name of the instance.
	//
	// example:
	//
	// instanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DescribeHotDbListResponseBodyDataListInstanceDb) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotDbListResponseBodyDataListInstanceDb) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDb) GetHotDbList() *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList {
	return s.HotDbList
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDb) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDb) SetHotDbList(v *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) *DescribeHotDbListResponseBodyDataListInstanceDb {
	s.HotDbList = v
	return s
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDb) SetInstanceName(v string) *DescribeHotDbListResponseBodyDataListInstanceDb {
	s.InstanceName = &v
	return s
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDb) Validate() error {
	return dara.Validate(s)
}

type DescribeHotDbListResponseBodyDataListInstanceDbHotDbList struct {
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) GetData() []*string {
	return s.Data
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) SetData(v []*string) *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList {
	s.Data = v
	return s
}

func (s *DescribeHotDbListResponseBodyDataListInstanceDbHotDbList) Validate() error {
	return dara.Validate(s)
}
