// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsItemStatistic(v bool) *GetCheckSummaryRequest
	GetIsItemStatistic() *bool
	SetLang(v string) *GetCheckSummaryRequest
	GetLang() *string
	SetResourceDirectoryAccountId(v string) *GetCheckSummaryRequest
	GetResourceDirectoryAccountId() *string
	SetTaskSources(v []*string) *GetCheckSummaryRequest
	GetTaskSources() []*string
	SetVendors(v []*string) *GetCheckSummaryRequest
	GetVendors() []*string
}

type GetCheckSummaryRequest struct {
	// Specifies whether to return the statistics of the check items, including the number of check items supported by the system and the number of check items available to you. Default value: **false**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsItemStatistic *bool `json:"IsItemStatistic,omitempty" xml:"IsItemStatistic,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to query the IDs of Alibaba Cloud accounts.
	//
	// example:
	//
	// 000
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// List of task sources.
	TaskSources []*string `json:"TaskSources,omitempty" xml:"TaskSources,omitempty" type:"Repeated"`
	// The cloud service providers.
	Vendors []*string `json:"Vendors,omitempty" xml:"Vendors,omitempty" type:"Repeated"`
}

func (s GetCheckSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetCheckSummaryRequest) GetIsItemStatistic() *bool {
	return s.IsItemStatistic
}

func (s *GetCheckSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *GetCheckSummaryRequest) GetResourceDirectoryAccountId() *string {
	return s.ResourceDirectoryAccountId
}

func (s *GetCheckSummaryRequest) GetTaskSources() []*string {
	return s.TaskSources
}

func (s *GetCheckSummaryRequest) GetVendors() []*string {
	return s.Vendors
}

func (s *GetCheckSummaryRequest) SetIsItemStatistic(v bool) *GetCheckSummaryRequest {
	s.IsItemStatistic = &v
	return s
}

func (s *GetCheckSummaryRequest) SetLang(v string) *GetCheckSummaryRequest {
	s.Lang = &v
	return s
}

func (s *GetCheckSummaryRequest) SetResourceDirectoryAccountId(v string) *GetCheckSummaryRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *GetCheckSummaryRequest) SetTaskSources(v []*string) *GetCheckSummaryRequest {
	s.TaskSources = v
	return s
}

func (s *GetCheckSummaryRequest) SetVendors(v []*string) *GetCheckSummaryRequest {
	s.Vendors = v
	return s
}

func (s *GetCheckSummaryRequest) Validate() error {
	return dara.Validate(s)
}
