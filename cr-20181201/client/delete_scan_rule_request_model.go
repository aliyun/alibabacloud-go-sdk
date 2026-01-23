// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScanRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteScanRuleRequest
	GetInstanceId() *string
	SetScanRuleId(v string) *DeleteScanRuleRequest
	GetScanRuleId() *string
}

type DeleteScanRuleRequest struct {
	// The instance ID
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rule ID
	//
	// example:
	//
	// crscnr-aemytkwad2h7fyhb
	ScanRuleId *string `json:"ScanRuleId,omitempty" xml:"ScanRuleId,omitempty"`
}

func (s DeleteScanRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScanRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteScanRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteScanRuleRequest) GetScanRuleId() *string {
	return s.ScanRuleId
}

func (s *DeleteScanRuleRequest) SetInstanceId(v string) *DeleteScanRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteScanRuleRequest) SetScanRuleId(v string) *DeleteScanRuleRequest {
	s.ScanRuleId = &v
	return s
}

func (s *DeleteScanRuleRequest) Validate() error {
	return dara.Validate(s)
}
