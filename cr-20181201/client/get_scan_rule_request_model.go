// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScanRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetScanRuleRequest
	GetInstanceId() *string
	SetScanRuleId(v string) *GetScanRuleRequest
	GetScanRuleId() *string
}

type GetScanRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crscnr-2sdveqjhpzdcafyq
	ScanRuleId *string `json:"ScanRuleId,omitempty" xml:"ScanRuleId,omitempty"`
}

func (s GetScanRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScanRuleRequest) GoString() string {
	return s.String()
}

func (s *GetScanRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetScanRuleRequest) GetScanRuleId() *string {
	return s.ScanRuleId
}

func (s *GetScanRuleRequest) SetInstanceId(v string) *GetScanRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetScanRuleRequest) SetScanRuleId(v string) *GetScanRuleRequest {
	s.ScanRuleId = &v
	return s
}

func (s *GetScanRuleRequest) Validate() error {
	return dara.Validate(s)
}
