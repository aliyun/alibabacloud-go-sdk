// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsInstanceId(v string) *DescribeInvadeEventDetailResponseBody
	GetAssetsInstanceId() *string
	SetAssetsInstanceName(v string) *DescribeInvadeEventDetailResponseBody
	GetAssetsInstanceName() *string
	SetEventDesc(v string) *DescribeInvadeEventDetailResponseBody
	GetEventDesc() *string
	SetEventDetail(v string) *DescribeInvadeEventDetailResponseBody
	GetEventDetail() *string
	SetEventKey(v string) *DescribeInvadeEventDetailResponseBody
	GetEventKey() *string
	SetEventName(v string) *DescribeInvadeEventDetailResponseBody
	GetEventName() *string
	SetEventUuid(v string) *DescribeInvadeEventDetailResponseBody
	GetEventUuid() *string
	SetFirstTime(v int32) *DescribeInvadeEventDetailResponseBody
	GetFirstTime() *int32
	SetIsIgnore(v bool) *DescribeInvadeEventDetailResponseBody
	GetIsIgnore() *bool
	SetLastTime(v int32) *DescribeInvadeEventDetailResponseBody
	GetLastTime() *int32
	SetOperationList(v []*DescribeInvadeEventDetailResponseBodyOperationList) *DescribeInvadeEventDetailResponseBody
	GetOperationList() []*DescribeInvadeEventDetailResponseBodyOperationList
	SetPrivateIP(v string) *DescribeInvadeEventDetailResponseBody
	GetPrivateIP() *string
	SetProcessStatus(v int32) *DescribeInvadeEventDetailResponseBody
	GetProcessStatus() *int32
	SetPublicIP(v string) *DescribeInvadeEventDetailResponseBody
	GetPublicIP() *string
	SetReference(v string) *DescribeInvadeEventDetailResponseBody
	GetReference() *string
	SetRegionNo(v string) *DescribeInvadeEventDetailResponseBody
	GetRegionNo() *string
	SetRequestId(v string) *DescribeInvadeEventDetailResponseBody
	GetRequestId() *string
	SetRiskLevel(v int32) *DescribeInvadeEventDetailResponseBody
	GetRiskLevel() *int32
	SetUnhandleOperationList(v []*DescribeInvadeEventDetailResponseBodyUnhandleOperationList) *DescribeInvadeEventDetailResponseBody
	GetUnhandleOperationList() []*DescribeInvadeEventDetailResponseBodyUnhandleOperationList
}

type DescribeInvadeEventDetailResponseBody struct {
	// example:
	//
	// i-8vb2nmm070m****
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// example:
	//
	// ECS_test
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// example:
	//
	// test
	EventDesc *string `json:"EventDesc,omitempty" xml:"EventDesc,omitempty"`
	// example:
	//
	// test
	EventDetail *string `json:"EventDetail,omitempty" xml:"EventDetail,omitempty"`
	// example:
	//
	// C&CActivity
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// example:
	//
	// event_test
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// aa6e786c-5034-457a-8e05-1c63fab****
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// example:
	//
	// 1735006706
	FirstTime *int32 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// example:
	//
	// false
	IsIgnore *bool `json:"IsIgnore,omitempty" xml:"IsIgnore,omitempty"`
	// example:
	//
	// 1740104108
	LastTime      *int32                                                `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	OperationList []*DescribeInvadeEventDetailResponseBodyOperationList `json:"OperationList,omitempty" xml:"OperationList,omitempty" type:"Repeated"`
	// example:
	//
	// 10.21.186.XXX
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// example:
	//
	// 1
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// example:
	//
	// 106.15.185.XXX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// example:
	//
	// test
	Reference *string `json:"Reference,omitempty" xml:"Reference,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// 8022D695-4A35-50BC-8697-EA9C233A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	RiskLevel             *int32                                                        `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	UnhandleOperationList []*DescribeInvadeEventDetailResponseBodyUnhandleOperationList `json:"UnhandleOperationList,omitempty" xml:"UnhandleOperationList,omitempty" type:"Repeated"`
}

func (s DescribeInvadeEventDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventDetailResponseBody) GetAssetsInstanceId() *string {
	return s.AssetsInstanceId
}

func (s *DescribeInvadeEventDetailResponseBody) GetAssetsInstanceName() *string {
	return s.AssetsInstanceName
}

func (s *DescribeInvadeEventDetailResponseBody) GetEventDesc() *string {
	return s.EventDesc
}

func (s *DescribeInvadeEventDetailResponseBody) GetEventDetail() *string {
	return s.EventDetail
}

func (s *DescribeInvadeEventDetailResponseBody) GetEventKey() *string {
	return s.EventKey
}

func (s *DescribeInvadeEventDetailResponseBody) GetEventName() *string {
	return s.EventName
}

func (s *DescribeInvadeEventDetailResponseBody) GetEventUuid() *string {
	return s.EventUuid
}

func (s *DescribeInvadeEventDetailResponseBody) GetFirstTime() *int32 {
	return s.FirstTime
}

func (s *DescribeInvadeEventDetailResponseBody) GetIsIgnore() *bool {
	return s.IsIgnore
}

func (s *DescribeInvadeEventDetailResponseBody) GetLastTime() *int32 {
	return s.LastTime
}

func (s *DescribeInvadeEventDetailResponseBody) GetOperationList() []*DescribeInvadeEventDetailResponseBodyOperationList {
	return s.OperationList
}

func (s *DescribeInvadeEventDetailResponseBody) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeInvadeEventDetailResponseBody) GetProcessStatus() *int32 {
	return s.ProcessStatus
}

func (s *DescribeInvadeEventDetailResponseBody) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeInvadeEventDetailResponseBody) GetReference() *string {
	return s.Reference
}

func (s *DescribeInvadeEventDetailResponseBody) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeInvadeEventDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvadeEventDetailResponseBody) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeInvadeEventDetailResponseBody) GetUnhandleOperationList() []*DescribeInvadeEventDetailResponseBodyUnhandleOperationList {
	return s.UnhandleOperationList
}

func (s *DescribeInvadeEventDetailResponseBody) SetAssetsInstanceId(v string) *DescribeInvadeEventDetailResponseBody {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetAssetsInstanceName(v string) *DescribeInvadeEventDetailResponseBody {
	s.AssetsInstanceName = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetEventDesc(v string) *DescribeInvadeEventDetailResponseBody {
	s.EventDesc = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetEventDetail(v string) *DescribeInvadeEventDetailResponseBody {
	s.EventDetail = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetEventKey(v string) *DescribeInvadeEventDetailResponseBody {
	s.EventKey = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetEventName(v string) *DescribeInvadeEventDetailResponseBody {
	s.EventName = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetEventUuid(v string) *DescribeInvadeEventDetailResponseBody {
	s.EventUuid = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetFirstTime(v int32) *DescribeInvadeEventDetailResponseBody {
	s.FirstTime = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetIsIgnore(v bool) *DescribeInvadeEventDetailResponseBody {
	s.IsIgnore = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetLastTime(v int32) *DescribeInvadeEventDetailResponseBody {
	s.LastTime = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetOperationList(v []*DescribeInvadeEventDetailResponseBodyOperationList) *DescribeInvadeEventDetailResponseBody {
	s.OperationList = v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetPrivateIP(v string) *DescribeInvadeEventDetailResponseBody {
	s.PrivateIP = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetProcessStatus(v int32) *DescribeInvadeEventDetailResponseBody {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetPublicIP(v string) *DescribeInvadeEventDetailResponseBody {
	s.PublicIP = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetReference(v string) *DescribeInvadeEventDetailResponseBody {
	s.Reference = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetRegionNo(v string) *DescribeInvadeEventDetailResponseBody {
	s.RegionNo = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetRequestId(v string) *DescribeInvadeEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetRiskLevel(v int32) *DescribeInvadeEventDetailResponseBody {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) SetUnhandleOperationList(v []*DescribeInvadeEventDetailResponseBodyUnhandleOperationList) *DescribeInvadeEventDetailResponseBody {
	s.UnhandleOperationList = v
	return s
}

func (s *DescribeInvadeEventDetailResponseBody) Validate() error {
	if s.OperationList != nil {
		for _, item := range s.OperationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UnhandleOperationList != nil {
		for _, item := range s.UnhandleOperationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInvadeEventDetailResponseBodyOperationList struct {
	// example:
	//
	// test
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// RunMode
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
}

func (s DescribeInvadeEventDetailResponseBodyOperationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventDetailResponseBodyOperationList) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventDetailResponseBodyOperationList) GetArgs() *string {
	return s.Args
}

func (s *DescribeInvadeEventDetailResponseBodyOperationList) GetOperate() *string {
	return s.Operate
}

func (s *DescribeInvadeEventDetailResponseBodyOperationList) SetArgs(v string) *DescribeInvadeEventDetailResponseBodyOperationList {
	s.Args = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBodyOperationList) SetOperate(v string) *DescribeInvadeEventDetailResponseBodyOperationList {
	s.Operate = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBodyOperationList) Validate() error {
	return dara.Validate(s)
}

type DescribeInvadeEventDetailResponseBodyUnhandleOperationList struct {
	// example:
	//
	// test
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// example:
	//
	// RunMode
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
}

func (s DescribeInvadeEventDetailResponseBodyUnhandleOperationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventDetailResponseBodyUnhandleOperationList) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventDetailResponseBodyUnhandleOperationList) GetArgs() *string {
	return s.Args
}

func (s *DescribeInvadeEventDetailResponseBodyUnhandleOperationList) GetOperate() *string {
	return s.Operate
}

func (s *DescribeInvadeEventDetailResponseBodyUnhandleOperationList) SetArgs(v string) *DescribeInvadeEventDetailResponseBodyUnhandleOperationList {
	s.Args = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBodyUnhandleOperationList) SetOperate(v string) *DescribeInvadeEventDetailResponseBodyUnhandleOperationList {
	s.Operate = &v
	return s
}

func (s *DescribeInvadeEventDetailResponseBodyUnhandleOperationList) Validate() error {
	return dara.Validate(s)
}
