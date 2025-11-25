// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchSlsDispatchStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItemList(v []*DescribeBatchSlsDispatchStatusResponseBodyItemList) *DescribeBatchSlsDispatchStatusResponseBody
	GetItemList() []*DescribeBatchSlsDispatchStatusResponseBodyItemList
	SetLogstoreName(v string) *DescribeBatchSlsDispatchStatusResponseBody
	GetLogstoreName() *string
	SetProjectName(v string) *DescribeBatchSlsDispatchStatusResponseBody
	GetProjectName() *string
	SetRequestId(v string) *DescribeBatchSlsDispatchStatusResponseBody
	GetRequestId() *string
}

type DescribeBatchSlsDispatchStatusResponseBody struct {
	ItemList []*DescribeBatchSlsDispatchStatusResponseBodyItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	// example:
	//
	// rs-stats
	LogstoreName *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	// example:
	//
	// cloudfirewallnew-project-199053910542****-cn-hangzhou
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 7036EBAB-F85F-5AAE-976F-C75AEE59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) GetItemList() []*DescribeBatchSlsDispatchStatusResponseBodyItemList {
	return s.ItemList
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetItemList(v []*DescribeBatchSlsDispatchStatusResponseBodyItemList) *DescribeBatchSlsDispatchStatusResponseBody {
	s.ItemList = v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetLogstoreName(v string) *DescribeBatchSlsDispatchStatusResponseBody {
	s.LogstoreName = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetProjectName(v string) *DescribeBatchSlsDispatchStatusResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetRequestId(v string) *DescribeBatchSlsDispatchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) Validate() error {
	if s.ItemList != nil {
		for _, item := range s.ItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBatchSlsDispatchStatusResponseBodyItemList struct {
	// example:
	//
	// success_finished
	ConfigStatus *string `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
	DispatchName *string `json:"DispatchName,omitempty" xml:"DispatchName,omitempty"`
	// example:
	//
	// internet_log
	DispatchValue *string `json:"DispatchValue,omitempty" xml:"DispatchValue,omitempty"`
	// example:
	//
	// true
	Enable     *bool     `json:"Enable,omitempty" xml:"Enable,omitempty"`
	FilterKeys []*string `json:"FilterKeys,omitempty" xml:"FilterKeys,omitempty" type:"Repeated"`
	// example:
	//
	// log_type
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusResponseBodyItemList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponseBodyItemList) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) GetConfigStatus() *string {
	return s.ConfigStatus
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) GetDispatchName() *string {
	return s.DispatchName
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) GetDispatchValue() *string {
	return s.DispatchValue
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) GetFilterKeys() []*string {
	return s.FilterKeys
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) GetSearchName() *string {
	return s.SearchName
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) SetConfigStatus(v string) *DescribeBatchSlsDispatchStatusResponseBodyItemList {
	s.ConfigStatus = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) SetDispatchName(v string) *DescribeBatchSlsDispatchStatusResponseBodyItemList {
	s.DispatchName = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) SetDispatchValue(v string) *DescribeBatchSlsDispatchStatusResponseBodyItemList {
	s.DispatchValue = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) SetEnable(v bool) *DescribeBatchSlsDispatchStatusResponseBodyItemList {
	s.Enable = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) SetFilterKeys(v []*string) *DescribeBatchSlsDispatchStatusResponseBodyItemList {
	s.FilterKeys = v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) SetSearchName(v string) *DescribeBatchSlsDispatchStatusResponseBodyItemList {
	s.SearchName = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyItemList) Validate() error {
	return dara.Validate(s)
}
