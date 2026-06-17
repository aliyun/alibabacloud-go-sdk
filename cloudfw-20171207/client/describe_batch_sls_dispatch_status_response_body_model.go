// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchSlsDispatchStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInfoList(v []*DescribeBatchSlsDispatchStatusResponseBodyInfoList) *DescribeBatchSlsDispatchStatusResponseBody
	GetInfoList() []*DescribeBatchSlsDispatchStatusResponseBodyInfoList
	SetItemList(v []*DescribeBatchSlsDispatchStatusResponseBodyItemList) *DescribeBatchSlsDispatchStatusResponseBody
	GetItemList() []*DescribeBatchSlsDispatchStatusResponseBodyItemList
	SetLogVersion(v string) *DescribeBatchSlsDispatchStatusResponseBody
	GetLogVersion() *string
	SetLogstoreName(v string) *DescribeBatchSlsDispatchStatusResponseBody
	GetLogstoreName() *string
	SetProjectName(v string) *DescribeBatchSlsDispatchStatusResponseBody
	GetProjectName() *string
	SetRequestId(v string) *DescribeBatchSlsDispatchStatusResponseBody
	GetRequestId() *string
}

type DescribeBatchSlsDispatchStatusResponseBody struct {
	// A list of detailed information about the Logstores.
	InfoList []*DescribeBatchSlsDispatchStatusResponseBodyInfoList `json:"InfoList,omitempty" xml:"InfoList,omitempty" type:"Repeated"`
	// A list of Simple Log Service projects.
	ItemList []*DescribeBatchSlsDispatchStatusResponseBodyItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	// The log version. A value of 1 indicates that there is one Logstore. A value of 2 indicates that there are two Logstores.
	//
	// example:
	//
	// 1
	LogVersion *string `json:"LogVersion,omitempty" xml:"LogVersion,omitempty"`
	// The name of the Logstore in Simple Log Service.
	//
	// example:
	//
	// rs-stats
	LogstoreName *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	// The name of the project in Simple Log Service.
	//
	// example:
	//
	// cloudfirewallnew-project-199053910542****-cn-hangzhou
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the request.
	//
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

func (s *DescribeBatchSlsDispatchStatusResponseBody) GetInfoList() []*DescribeBatchSlsDispatchStatusResponseBodyInfoList {
	return s.InfoList
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) GetItemList() []*DescribeBatchSlsDispatchStatusResponseBodyItemList {
	return s.ItemList
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) GetLogVersion() *string {
	return s.LogVersion
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

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetInfoList(v []*DescribeBatchSlsDispatchStatusResponseBodyInfoList) *DescribeBatchSlsDispatchStatusResponseBody {
	s.InfoList = v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetItemList(v []*DescribeBatchSlsDispatchStatusResponseBodyItemList) *DescribeBatchSlsDispatchStatusResponseBody {
	s.ItemList = v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetLogVersion(v string) *DescribeBatchSlsDispatchStatusResponseBody {
	s.LogVersion = &v
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
	if s.InfoList != nil {
		for _, item := range s.InfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeBatchSlsDispatchStatusResponseBodyInfoList struct {
	// The details of the log delivery configuration.
	ItemList []*DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	// The name of the Logstore in Simple Log Service.
	//
	// example:
	//
	// cloudfirewall-logstore
	LogstoreName *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	// The name of the project in Simple Log Service.
	//
	// example:
	//
	// cloudfirewall-project-1204872307283650-cn-hangzhou
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// \\`cn\\` indicates the Chinese mainland. \\`intl\\` indicates regions outside the Chinese mainland. \\`global\\` indicates global.
	//
	// example:
	//
	// cn
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusResponseBodyInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponseBodyInfoList) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoList) GetItemList() []*DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList {
	return s.ItemList
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoList) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoList) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoList) GetSite() *string {
	return s.Site
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoList) SetItemList(v []*DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) *DescribeBatchSlsDispatchStatusResponseBodyInfoList {
	s.ItemList = v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoList) SetLogstoreName(v string) *DescribeBatchSlsDispatchStatusResponseBodyInfoList {
	s.LogstoreName = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoList) SetProjectName(v string) *DescribeBatchSlsDispatchStatusResponseBodyInfoList {
	s.ProjectName = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoList) SetSite(v string) *DescribeBatchSlsDispatchStatusResponseBodyInfoList {
	s.Site = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoList) Validate() error {
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

type DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList struct {
	// The configuration status.
	//
	// example:
	//
	// success_finished
	ConfigStatus *string `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
	// The name of the delivery type.
	//
	// example:
	//
	// internet-traffic-log
	DispatchName *string `json:"DispatchName,omitempty" xml:"DispatchName,omitempty"`
	// The key for the log categorization. Valid values:
	//
	// **internet_log**
	//
	// **vpc_firewall_log**
	//
	// **nat_firewall_log**
	//
	// **ipv6_firewall_log**
	//
	// **dns_firewall_log**
	//
	// example:
	//
	// ipv6_firewall_log
	DispatchValue *string `json:"DispatchValue,omitempty" xml:"DispatchValue,omitempty"`
	// Indicates whether this delivery type is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The detailed delivery configurations for the Internet and VPCs.
	FilterKeys []*string `json:"FilterKeys,omitempty" xml:"FilterKeys,omitempty" type:"Repeated"`
	// The value is fixed to log_type. You can ignore this parameter.
	//
	// example:
	//
	// log_type
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) GetConfigStatus() *string {
	return s.ConfigStatus
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) GetDispatchName() *string {
	return s.DispatchName
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) GetDispatchValue() *string {
	return s.DispatchValue
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) GetFilterKeys() []*string {
	return s.FilterKeys
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) GetSearchName() *string {
	return s.SearchName
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) SetConfigStatus(v string) *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList {
	s.ConfigStatus = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) SetDispatchName(v string) *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList {
	s.DispatchName = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) SetDispatchValue(v string) *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList {
	s.DispatchValue = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) SetEnable(v bool) *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList {
	s.Enable = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) SetFilterKeys(v []*string) *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList {
	s.FilterKeys = v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) SetSearchName(v string) *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList {
	s.SearchName = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodyInfoListItemList) Validate() error {
	return dara.Validate(s)
}

type DescribeBatchSlsDispatchStatusResponseBodyItemList struct {
	// The status of the delivery configuration.
	//
	// example:
	//
	// success_finished
	ConfigStatus *string `json:"ConfigStatus,omitempty" xml:"ConfigStatus,omitempty"`
	// The name of the log delivery.
	//
	// example:
	//
	// internet-traffic-log
	DispatchName *string `json:"DispatchName,omitempty" xml:"DispatchName,omitempty"`
	// The value of the log to be delivered.
	//
	// example:
	//
	// internet_log
	DispatchValue *string `json:"DispatchValue,omitempty" xml:"DispatchValue,omitempty"`
	// The delivery status.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The supported filter conditions.
	FilterKeys []*string `json:"FilterKeys,omitempty" xml:"FilterKeys,omitempty" type:"Repeated"`
	// The name of the search type.
	//
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
