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
	// The source IP address to query.
	//
	// example:
	//
	// 113.57.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the configuration rule. Valid values:
	//
	// - **webshell_timescan**: Web shell detection and removal.
	//
	// - **aliscriptengine**: Deep inspection engine.
	//
	// - **alidetect**: Installation scope of the local file detection engine.
	//
	// - **alidetect-scan-enable**: Scan scope of the local file detection engine.
	//
	// > You can call the [ListClientUserDefineRules](~~ListClientUserDefineRules~~) and [ListSystemClientRules](~~ListSystemClientRules~~) operations to obtain more custom and system-defined rule types.
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
