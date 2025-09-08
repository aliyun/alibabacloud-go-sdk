// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectLogStoresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListProjectLogStoresResponseBodyData) *ListProjectLogStoresResponseBody
	GetData() []*ListProjectLogStoresResponseBodyData
	SetRequestId(v string) *ListProjectLogStoresResponseBody
	GetRequestId() *string
}

type ListProjectLogStoresResponseBody struct {
	// The data returned.
	Data []*ListProjectLogStoresResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectLogStoresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectLogStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectLogStoresResponseBody) GetData() []*ListProjectLogStoresResponseBodyData {
	return s.Data
}

func (s *ListProjectLogStoresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectLogStoresResponseBody) SetData(v []*ListProjectLogStoresResponseBodyData) *ListProjectLogStoresResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectLogStoresResponseBody) SetRequestId(v string) *ListProjectLogStoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectLogStoresResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProjectLogStoresResponseBodyData struct {
	// The endpoint of the Simple Log Service project.
	//
	// example:
	//
	// cn-hangzhou.log.aliyuncs.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// The name of the region in which the Simple Log Service project resides.
	//
	// example:
	//
	// hangzhou
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// cloud-siem-logstore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// 123XXXXXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// cloud-siem-project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region in which the Simple Log Service project resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account that can be used to perform operations supported by the threat analysis feature.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The username of the Alibaba Cloud account that can be used to perform operations supported by the threat analysis feature.
	//
	// example:
	//
	// sas_account_xxxx
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
}

func (s ListProjectLogStoresResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProjectLogStoresResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectLogStoresResponseBodyData) GetEndPoint() *string {
	return s.EndPoint
}

func (s *ListProjectLogStoresResponseBodyData) GetLocalName() *string {
	return s.LocalName
}

func (s *ListProjectLogStoresResponseBodyData) GetLogStore() *string {
	return s.LogStore
}

func (s *ListProjectLogStoresResponseBodyData) GetMainUserId() *int64 {
	return s.MainUserId
}

func (s *ListProjectLogStoresResponseBodyData) GetProject() *string {
	return s.Project
}

func (s *ListProjectLogStoresResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProjectLogStoresResponseBodyData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *ListProjectLogStoresResponseBodyData) GetSubUserName() *string {
	return s.SubUserName
}

func (s *ListProjectLogStoresResponseBodyData) SetEndPoint(v string) *ListProjectLogStoresResponseBodyData {
	s.EndPoint = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetLocalName(v string) *ListProjectLogStoresResponseBodyData {
	s.LocalName = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetLogStore(v string) *ListProjectLogStoresResponseBodyData {
	s.LogStore = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetMainUserId(v int64) *ListProjectLogStoresResponseBodyData {
	s.MainUserId = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetProject(v string) *ListProjectLogStoresResponseBodyData {
	s.Project = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetRegionId(v string) *ListProjectLogStoresResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetSubUserId(v int64) *ListProjectLogStoresResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) SetSubUserName(v string) *ListProjectLogStoresResponseBodyData {
	s.SubUserName = &v
	return s
}

func (s *ListProjectLogStoresResponseBodyData) Validate() error {
	return dara.Validate(s)
}
