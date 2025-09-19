// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonTargetResultListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DescribeCommonTargetResultListRequest
	GetSourceIp() *string
	SetType(v string) *DescribeCommonTargetResultListRequest
	GetType() *string
}

type DescribeCommonTargetResultListRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 113.57.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	// >  You can call the [ListClientUserDefineRules](~~ListClientUserDefineRules~~) and [ListSystemClientRules](~~ListSystemClientRules~~) operations to obtain more types of custom and system configuration items.
	//
	// This parameter is required.
	//
	// example:
	//
	// webshell_timescan
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCommonTargetResultListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonTargetResultListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommonTargetResultListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeCommonTargetResultListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeCommonTargetResultListRequest) SetSourceIp(v string) *DescribeCommonTargetResultListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCommonTargetResultListRequest) SetType(v string) *DescribeCommonTargetResultListRequest {
	s.Type = &v
	return s
}

func (s *DescribeCommonTargetResultListRequest) Validate() error {
	return dara.Validate(s)
}
