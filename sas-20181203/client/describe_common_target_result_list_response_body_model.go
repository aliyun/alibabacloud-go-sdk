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
	// The request ID. Alibaba Cloud generates a unique identifier for each request. You can use the request ID to troubleshoot issues.
	//
	// example:
	//
	// 6673D49C-A9AB-40DD-B4A2-B92306701AE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configuration information.
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
	if s.TargetConfig != nil {
		if err := s.TargetConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCommonTargetResultListResponseBodyTargetConfig struct {
	// The asset configuration flag. Valid values:
	//
	// - **add**: The configuration takes effect on the asset.
	//
	// - **del**: The configuration does not take effect on the asset.
	//
	// example:
	//
	// del
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The default flag for asset configuration.
	//
	// example:
	//
	// add
	TargetDefault *string `json:"TargetDefault,omitempty" xml:"TargetDefault,omitempty"`
	// The group ID or asset UUID on which the configuration takes effect.
	//
	// > If **TargetType*	- returns **uuid**, this field indicates the **UUID*	- of the asset. If **TargetType*	- returns **groupId**, this field indicates the group ID.
	TargetList []*string `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	// The selection mode for the assets on which the configuration takes effect. Valid values:
	//
	// - **uuid**: Added by individual asset.
	//
	// - **groupId**: Added by server group.
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
	// The configuration type. Valid values:
	//
	// - **webshell_timescan**: web shell scan.
	//
	// - **aliscriptengine**: deep detection engine.
	//
	// - **alidetect**: installation scope of the local file detection engine.
	//
	// - **alidetect-scan-enable**: detection scope of the local file detection engine.
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
