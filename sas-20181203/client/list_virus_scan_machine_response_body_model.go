// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListVirusScanMachineResponseBodyData) *ListVirusScanMachineResponseBody
	GetData() []*ListVirusScanMachineResponseBodyData
	SetPageInfo(v *ListVirusScanMachineResponseBodyPageInfo) *ListVirusScanMachineResponseBody
	GetPageInfo() *ListVirusScanMachineResponseBodyPageInfo
	SetRequestId(v string) *ListVirusScanMachineResponseBody
	GetRequestId() *string
}

type ListVirusScanMachineResponseBody struct {
	// The data details.
	Data []*ListVirusScanMachineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListVirusScanMachineResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request and can be used to troubleshoot issues.
	//
	// example:
	//
	// A27C1C09-828B-5CB8-9203-F55423BE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVirusScanMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineResponseBody) GetData() []*ListVirusScanMachineResponseBodyData {
	return s.Data
}

func (s *ListVirusScanMachineResponseBody) GetPageInfo() *ListVirusScanMachineResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListVirusScanMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirusScanMachineResponseBody) SetData(v []*ListVirusScanMachineResponseBodyData) *ListVirusScanMachineResponseBody {
	s.Data = v
	return s
}

func (s *ListVirusScanMachineResponseBody) SetPageInfo(v *ListVirusScanMachineResponseBodyPageInfo) *ListVirusScanMachineResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListVirusScanMachineResponseBody) SetRequestId(v string) *ListVirusScanMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirusScanMachineResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVirusScanMachineResponseBodyData struct {
	// The number of security alert occurrences.
	//
	// example:
	//
	// 28
	EventCount *int32 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The instance ID of the server.
	//
	// example:
	//
	// i-bp180bogui4fc0z4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// centos****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the associated instance.
	//
	// example:
	//
	// 172.16.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the associated instance.
	//
	// example:
	//
	// 10.42.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The UUID of the asset instance.
	//
	// example:
	//
	// 6690a46c-0edb-4663-a641-3629d1a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListVirusScanMachineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineResponseBodyData) GetEventCount() *int32 {
	return s.EventCount
}

func (s *ListVirusScanMachineResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVirusScanMachineResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListVirusScanMachineResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListVirusScanMachineResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListVirusScanMachineResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *ListVirusScanMachineResponseBodyData) SetEventCount(v int32) *ListVirusScanMachineResponseBodyData {
	s.EventCount = &v
	return s
}

func (s *ListVirusScanMachineResponseBodyData) SetInstanceId(v string) *ListVirusScanMachineResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListVirusScanMachineResponseBodyData) SetInstanceName(v string) *ListVirusScanMachineResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListVirusScanMachineResponseBodyData) SetInternetIp(v string) *ListVirusScanMachineResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *ListVirusScanMachineResponseBodyData) SetIntranetIp(v string) *ListVirusScanMachineResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *ListVirusScanMachineResponseBodyData) SetUuid(v string) *ListVirusScanMachineResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListVirusScanMachineResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListVirusScanMachineResponseBodyPageInfo struct {
	// The page number of the current page in a paging query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries returned per page in a paging query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 168
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVirusScanMachineResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListVirusScanMachineResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVirusScanMachineResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVirusScanMachineResponseBodyPageInfo) SetCurrentPage(v int32) *ListVirusScanMachineResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListVirusScanMachineResponseBodyPageInfo) SetPageSize(v int32) *ListVirusScanMachineResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListVirusScanMachineResponseBodyPageInfo) SetTotalCount(v int32) *ListVirusScanMachineResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListVirusScanMachineResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
