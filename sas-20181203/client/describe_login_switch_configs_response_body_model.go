// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoginSwitchConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigList(v []*DescribeLoginSwitchConfigsResponseBodyConfigList) *DescribeLoginSwitchConfigsResponseBody
	GetConfigList() []*DescribeLoginSwitchConfigsResponseBodyConfigList
	SetCount(v int32) *DescribeLoginSwitchConfigsResponseBody
	GetCount() *int32
	SetRequestId(v string) *DescribeLoginSwitchConfigsResponseBody
	GetRequestId() *string
}

type DescribeLoginSwitchConfigsResponseBody struct {
	// The configuration item returned.
	ConfigList []*DescribeLoginSwitchConfigsResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The number of returned configuration items.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLoginSwitchConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginSwitchConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoginSwitchConfigsResponseBody) GetConfigList() []*DescribeLoginSwitchConfigsResponseBodyConfigList {
	return s.ConfigList
}

func (s *DescribeLoginSwitchConfigsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeLoginSwitchConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoginSwitchConfigsResponseBody) SetConfigList(v []*DescribeLoginSwitchConfigsResponseBodyConfigList) *DescribeLoginSwitchConfigsResponseBody {
	s.ConfigList = v
	return s
}

func (s *DescribeLoginSwitchConfigsResponseBody) SetCount(v int32) *DescribeLoginSwitchConfigsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeLoginSwitchConfigsResponseBody) SetRequestId(v string) *DescribeLoginSwitchConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoginSwitchConfigsResponseBody) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoginSwitchConfigsResponseBodyConfigList struct {
	// The type of the alert that you enabled or disabled. Valid values:
	//
	// 	- **login_common_ip**: alerts for unapproved logon IP addresses
	//
	// 	- **login_common_time**: alerts for unapproved logon time ranges
	//
	// 	- **login_common_account**: alerts for unapproved logon accounts
	//
	// example:
	//
	// login_common_ip
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The status of the Log Service feature. Valid values:
	//
	// 	- **0**: The feature is disabled.
	//
	// 	- **1**: The feature is enabled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeLoginSwitchConfigsResponseBodyConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginSwitchConfigsResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLoginSwitchConfigsResponseBodyConfigList) GetItem() *string {
	return s.Item
}

func (s *DescribeLoginSwitchConfigsResponseBodyConfigList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeLoginSwitchConfigsResponseBodyConfigList) SetItem(v string) *DescribeLoginSwitchConfigsResponseBodyConfigList {
	s.Item = &v
	return s
}

func (s *DescribeLoginSwitchConfigsResponseBodyConfigList) SetStatus(v int32) *DescribeLoginSwitchConfigsResponseBodyConfigList {
	s.Status = &v
	return s
}

func (s *DescribeLoginSwitchConfigsResponseBodyConfigList) Validate() error {
	return dara.Validate(s)
}
