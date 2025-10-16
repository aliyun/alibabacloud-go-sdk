// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomizedListHeadersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v []*ModifyCustomizedListHeadersRequestHeaders) *ModifyCustomizedListHeadersRequest
	GetHeaders() []*ModifyCustomizedListHeadersRequestHeaders
	SetListType(v string) *ModifyCustomizedListHeadersRequest
	GetListType() *string
	SetRegionId(v string) *ModifyCustomizedListHeadersRequest
	GetRegionId() *string
}

type ModifyCustomizedListHeadersRequest struct {
	// The headers.
	Headers []*ModifyCustomizedListHeadersRequestHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The type of the list.
	//
	// Valid values:
	//
	// 	- desktop: cloud computer
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// desktop
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyCustomizedListHeadersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomizedListHeadersRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomizedListHeadersRequest) GetHeaders() []*ModifyCustomizedListHeadersRequestHeaders {
	return s.Headers
}

func (s *ModifyCustomizedListHeadersRequest) GetListType() *string {
	return s.ListType
}

func (s *ModifyCustomizedListHeadersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCustomizedListHeadersRequest) SetHeaders(v []*ModifyCustomizedListHeadersRequestHeaders) *ModifyCustomizedListHeadersRequest {
	s.Headers = v
	return s
}

func (s *ModifyCustomizedListHeadersRequest) SetListType(v string) *ModifyCustomizedListHeadersRequest {
	s.ListType = &v
	return s
}

func (s *ModifyCustomizedListHeadersRequest) SetRegionId(v string) *ModifyCustomizedListHeadersRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCustomizedListHeadersRequest) Validate() error {
	if s.Headers != nil {
		for _, item := range s.Headers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyCustomizedListHeadersRequestHeaders struct {
	// The display type of the header.
	//
	// > For the desktop_id_name and office_site_id_name head keys, set the value of this parameter to required. For other header keys, set the value of this parameter to display or hide based on your requirements.
	//
	// example:
	//
	// display
	DisplayType *string `json:"DisplayType,omitempty" xml:"DisplayType,omitempty"`
	// The key of the header.
	//
	// > All header keys of the list must be specified.
	//
	// Valid values:
	//
	// 	- desktop_id_name: the IDs and names of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- system_data_disk: the system disks and data disks of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- office_site_type: the office network types of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- create_time: the time when the cloud computers are created.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ip: the IP addresses of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- spec_system_protocol: the instance types, OSs, and protocol types of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- monitor: the monitoring information of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- assigned_users: the number of end users that are assigned to the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- encryption: indicates whether the cloud computers are encrypted.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- office_site_id_name: the IDs and names of the office networks.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- pay_type: the billing methods of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- tag: the tags that are attached to the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- hostname: the hostnames of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- status: the statuses of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- current_user: the current end users of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// pay_type
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
}

func (s ModifyCustomizedListHeadersRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomizedListHeadersRequestHeaders) GoString() string {
	return s.String()
}

func (s *ModifyCustomizedListHeadersRequestHeaders) GetDisplayType() *string {
	return s.DisplayType
}

func (s *ModifyCustomizedListHeadersRequestHeaders) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *ModifyCustomizedListHeadersRequestHeaders) SetDisplayType(v string) *ModifyCustomizedListHeadersRequestHeaders {
	s.DisplayType = &v
	return s
}

func (s *ModifyCustomizedListHeadersRequestHeaders) SetHeaderKey(v string) *ModifyCustomizedListHeadersRequestHeaders {
	s.HeaderKey = &v
	return s
}

func (s *ModifyCustomizedListHeadersRequestHeaders) Validate() error {
	return dara.Validate(s)
}
