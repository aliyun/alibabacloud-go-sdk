// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWarningThresholdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUid(v int64) *SetWarningThresholdRequest
	GetUid() *int64
	SetWarningValue(v string) *SetWarningThresholdRequest
	GetWarningValue() *string
}

type SetWarningThresholdRequest struct {
	// The UID of the reseller customer account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1792155717328010
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The percentage value. Valid values: 1 to 100. When the available credit drops below this percentage of the total credit limit, an alert email is sent to the reseller partner account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	WarningValue *string `json:"WarningValue,omitempty" xml:"WarningValue,omitempty"`
}

func (s SetWarningThresholdRequest) String() string {
	return dara.Prettify(s)
}

func (s SetWarningThresholdRequest) GoString() string {
	return s.String()
}

func (s *SetWarningThresholdRequest) GetUid() *int64 {
	return s.Uid
}

func (s *SetWarningThresholdRequest) GetWarningValue() *string {
	return s.WarningValue
}

func (s *SetWarningThresholdRequest) SetUid(v int64) *SetWarningThresholdRequest {
	s.Uid = &v
	return s
}

func (s *SetWarningThresholdRequest) SetWarningValue(v string) *SetWarningThresholdRequest {
	s.WarningValue = &v
	return s
}

func (s *SetWarningThresholdRequest) Validate() error {
	return dara.Validate(s)
}
