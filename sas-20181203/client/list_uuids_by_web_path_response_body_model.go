// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUuidsByWebPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListUuidsByWebPathResponseBodyList) *ListUuidsByWebPathResponseBody
	GetList() []*ListUuidsByWebPathResponseBodyList
	SetPageInfo(v *ListUuidsByWebPathResponseBodyPageInfo) *ListUuidsByWebPathResponseBody
	GetPageInfo() *ListUuidsByWebPathResponseBodyPageInfo
	SetRequestId(v string) *ListUuidsByWebPathResponseBody
	GetRequestId() *string
}

type ListUuidsByWebPathResponseBody struct {
	// An array that consists of the protected assets.
	List []*ListUuidsByWebPathResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListUuidsByWebPathResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A3C1240F-9DAC-5EE8-ADF5-2F930A95****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUuidsByWebPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUuidsByWebPathResponseBody) GoString() string {
	return s.String()
}

func (s *ListUuidsByWebPathResponseBody) GetList() []*ListUuidsByWebPathResponseBodyList {
	return s.List
}

func (s *ListUuidsByWebPathResponseBody) GetPageInfo() *ListUuidsByWebPathResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListUuidsByWebPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUuidsByWebPathResponseBody) SetList(v []*ListUuidsByWebPathResponseBodyList) *ListUuidsByWebPathResponseBody {
	s.List = v
	return s
}

func (s *ListUuidsByWebPathResponseBody) SetPageInfo(v *ListUuidsByWebPathResponseBodyPageInfo) *ListUuidsByWebPathResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListUuidsByWebPathResponseBody) SetRequestId(v string) *ListUuidsByWebPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUuidsByWebPathResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
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

type ListUuidsByWebPathResponseBodyList struct {
	// The public IP address of the server.
	//
	// example:
	//
	// 8.210.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.25.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// test****
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListUuidsByWebPathResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListUuidsByWebPathResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListUuidsByWebPathResponseBodyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListUuidsByWebPathResponseBodyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListUuidsByWebPathResponseBodyList) GetMachineName() *string {
	return s.MachineName
}

func (s *ListUuidsByWebPathResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *ListUuidsByWebPathResponseBodyList) SetInternetIp(v string) *ListUuidsByWebPathResponseBodyList {
	s.InternetIp = &v
	return s
}

func (s *ListUuidsByWebPathResponseBodyList) SetIntranetIp(v string) *ListUuidsByWebPathResponseBodyList {
	s.IntranetIp = &v
	return s
}

func (s *ListUuidsByWebPathResponseBodyList) SetMachineName(v string) *ListUuidsByWebPathResponseBodyList {
	s.MachineName = &v
	return s
}

func (s *ListUuidsByWebPathResponseBodyList) SetUuid(v string) *ListUuidsByWebPathResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *ListUuidsByWebPathResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListUuidsByWebPathResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUuidsByWebPathResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUuidsByWebPathResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListUuidsByWebPathResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListUuidsByWebPathResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUuidsByWebPathResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUuidsByWebPathResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUuidsByWebPathResponseBodyPageInfo) SetCount(v int32) *ListUuidsByWebPathResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListUuidsByWebPathResponseBodyPageInfo) SetCurrentPage(v int32) *ListUuidsByWebPathResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListUuidsByWebPathResponseBodyPageInfo) SetPageSize(v int32) *ListUuidsByWebPathResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListUuidsByWebPathResponseBodyPageInfo) SetTotalCount(v int32) *ListUuidsByWebPathResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListUuidsByWebPathResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
