// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonTargetResultListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCommonTargetResultListResponseBody
	GetRequestId() *string
	SetTargetConfig(v *DescribeCommonTargetResultListResponseBodyTargetConfig) *DescribeCommonTargetResultListResponseBody
	GetTargetConfig() *DescribeCommonTargetResultListResponseBodyTargetConfig
}

type DescribeCommonTargetResultListResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6673D49C-A9AB-40DD-B4A2-B92306701AE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the configuration item.
	TargetConfig *DescribeCommonTargetResultListResponseBodyTargetConfig `json:"TargetConfig,omitempty" xml:"TargetConfig,omitempty" type:"Struct"`
}

func (s DescribeCommonTargetResultListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonTargetResultListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommonTargetResultListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCommonTargetResultListResponseBody) GetTargetConfig() *DescribeCommonTargetResultListResponseBodyTargetConfig {
	return s.TargetConfig
}

func (s *DescribeCommonTargetResultListResponseBody) SetRequestId(v string) *DescribeCommonTargetResultListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommonTargetResultListResponseBody) SetTargetConfig(v *DescribeCommonTargetResultListResponseBodyTargetConfig) *DescribeCommonTargetResultListResponseBody {
	s.TargetConfig = v
	return s
}

func (s *DescribeCommonTargetResultListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCommonTargetResultListResponseBodyTargetConfig struct {
	// The identifier that indicates whether the configuration item is applied to the server. Valid values:
	//
	// 	- **add**: applied
	//
	// 	- **del**: not applied
	//
	// example:
	//
	// del
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The default identifier.
	//
	// example:
	//
	// add
	TargetDefault *string `json:"TargetDefault,omitempty" xml:"TargetDefault,omitempty"`
	// An array that consists of the IDs of the server groups or the UUIDs of the servers.
	//
	// >  If **uuid*	- is returned for the **TargetType*	- parameter, **UUIDs*	- of the servers are returned. If **groupId*	- is returned for the **TargetType*	- parameter, IDs of the server groups are returned.
	TargetList []*string `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	// The type of the server to which the configuration item is applied. Valid values:
	//
	// 	- **uuid**: a server
	//
	// 	- **groupId**: a server group
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 22
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The type of the configuration item. Valid values:
	//
	// 	- **webshell_timescan**: webshell detection and removal
	//
	// 	- **aliscriptengine**: in-depth detection engine
	//
	// 	- **alidetect**: installation scope of local file detection
	//
	// 	- **alidetect-scan-enable**: detection scope of local file detection
	//
	// example:
	//
	// webshell_timescan
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCommonTargetResultListResponseBodyTargetConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonTargetResultListResponseBodyTargetConfig) GoString() string {
	return s.String()
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) GetFlag() *string {
	return s.Flag
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) GetTargetDefault() *string {
	return s.TargetDefault
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) GetTargetList() []*string {
	return s.TargetList
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) GetType() *string {
	return s.Type
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) SetFlag(v string) *DescribeCommonTargetResultListResponseBodyTargetConfig {
	s.Flag = &v
	return s
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) SetTargetDefault(v string) *DescribeCommonTargetResultListResponseBodyTargetConfig {
	s.TargetDefault = &v
	return s
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) SetTargetList(v []*string) *DescribeCommonTargetResultListResponseBodyTargetConfig {
	s.TargetList = v
	return s
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) SetTargetType(v string) *DescribeCommonTargetResultListResponseBodyTargetConfig {
	s.TargetType = &v
	return s
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) SetTotalCount(v string) *DescribeCommonTargetResultListResponseBodyTargetConfig {
	s.TotalCount = &v
	return s
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) SetType(v string) *DescribeCommonTargetResultListResponseBodyTargetConfig {
	s.Type = &v
	return s
}

func (s *DescribeCommonTargetResultListResponseBodyTargetConfig) Validate() error {
	return dara.Validate(s)
}
