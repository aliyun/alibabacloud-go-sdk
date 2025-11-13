// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterScannerListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstalledCount(v int32) *DescribeClusterScannerListResponseBody
	GetInstalledCount() *int32
	SetList(v []*DescribeClusterScannerListResponseBodyList) *DescribeClusterScannerListResponseBody
	GetList() []*DescribeClusterScannerListResponseBodyList
	SetRequestId(v string) *DescribeClusterScannerListResponseBody
	GetRequestId() *string
}

type DescribeClusterScannerListResponseBody struct {
	// example:
	//
	// 1
	InstalledCount *int32                                        `json:"InstalledCount,omitempty" xml:"InstalledCount,omitempty"`
	List           []*DescribeClusterScannerListResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 11C96623-E106-59C9-866D-A6C82911459F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterScannerListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterScannerListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterScannerListResponseBody) GetInstalledCount() *int32 {
	return s.InstalledCount
}

func (s *DescribeClusterScannerListResponseBody) GetList() []*DescribeClusterScannerListResponseBodyList {
	return s.List
}

func (s *DescribeClusterScannerListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterScannerListResponseBody) SetInstalledCount(v int32) *DescribeClusterScannerListResponseBody {
	s.InstalledCount = &v
	return s
}

func (s *DescribeClusterScannerListResponseBody) SetList(v []*DescribeClusterScannerListResponseBodyList) *DescribeClusterScannerListResponseBody {
	s.List = v
	return s
}

func (s *DescribeClusterScannerListResponseBody) SetRequestId(v string) *DescribeClusterScannerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterScannerListResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterScannerListResponseBodyList struct {
	// example:
	//
	// 1693446913000
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// example:
	//
	// true
	NeedUpdate *bool `json:"NeedUpdate,omitempty" xml:"NeedUpdate,omitempty"`
	// example:
	//
	// offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 20241111
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 1
	WebhookOpen *int32 `json:"WebhookOpen,omitempty" xml:"WebhookOpen,omitempty"`
	// example:
	//
	// n
	WebhookStatus *string `json:"WebhookStatus,omitempty" xml:"WebhookStatus,omitempty"`
}

func (s DescribeClusterScannerListResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterScannerListResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeClusterScannerListResponseBodyList) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeClusterScannerListResponseBodyList) GetNeedUpdate() *bool {
	return s.NeedUpdate
}

func (s *DescribeClusterScannerListResponseBodyList) GetStatus() *string {
	return s.Status
}

func (s *DescribeClusterScannerListResponseBodyList) GetVersion() *string {
	return s.Version
}

func (s *DescribeClusterScannerListResponseBodyList) GetWebhookOpen() *int32 {
	return s.WebhookOpen
}

func (s *DescribeClusterScannerListResponseBodyList) GetWebhookStatus() *string {
	return s.WebhookStatus
}

func (s *DescribeClusterScannerListResponseBodyList) SetLastTime(v int64) *DescribeClusterScannerListResponseBodyList {
	s.LastTime = &v
	return s
}

func (s *DescribeClusterScannerListResponseBodyList) SetNeedUpdate(v bool) *DescribeClusterScannerListResponseBodyList {
	s.NeedUpdate = &v
	return s
}

func (s *DescribeClusterScannerListResponseBodyList) SetStatus(v string) *DescribeClusterScannerListResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeClusterScannerListResponseBodyList) SetVersion(v string) *DescribeClusterScannerListResponseBodyList {
	s.Version = &v
	return s
}

func (s *DescribeClusterScannerListResponseBodyList) SetWebhookOpen(v int32) *DescribeClusterScannerListResponseBodyList {
	s.WebhookOpen = &v
	return s
}

func (s *DescribeClusterScannerListResponseBodyList) SetWebhookStatus(v string) *DescribeClusterScannerListResponseBodyList {
	s.WebhookStatus = &v
	return s
}

func (s *DescribeClusterScannerListResponseBodyList) Validate() error {
	return dara.Validate(s)
}
