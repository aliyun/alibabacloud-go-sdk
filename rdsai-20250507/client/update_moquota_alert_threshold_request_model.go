// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMOQuotaAlertThresholdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApikey(v []*UpdateMOQuotaAlertThresholdRequestApikey) *UpdateMOQuotaAlertThresholdRequest
	GetApikey() []*UpdateMOQuotaAlertThresholdRequestApikey
	SetInstanceId(v string) *UpdateMOQuotaAlertThresholdRequest
	GetInstanceId() *string
}

type UpdateMOQuotaAlertThresholdRequest struct {
	// This parameter is required.
	Apikey []*UpdateMOQuotaAlertThresholdRequestApikey `json:"Apikey,omitempty" xml:"Apikey,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateMOQuotaAlertThresholdRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMOQuotaAlertThresholdRequest) GoString() string {
	return s.String()
}

func (s *UpdateMOQuotaAlertThresholdRequest) GetApikey() []*UpdateMOQuotaAlertThresholdRequestApikey {
	return s.Apikey
}

func (s *UpdateMOQuotaAlertThresholdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateMOQuotaAlertThresholdRequest) SetApikey(v []*UpdateMOQuotaAlertThresholdRequestApikey) *UpdateMOQuotaAlertThresholdRequest {
	s.Apikey = v
	return s
}

func (s *UpdateMOQuotaAlertThresholdRequest) SetInstanceId(v string) *UpdateMOQuotaAlertThresholdRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdRequest) Validate() error {
	if s.Apikey != nil {
		for _, item := range s.Apikey {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMOQuotaAlertThresholdRequestApikey struct {
	// ApiKey
	//
	// example:
	//
	// sk-rds-*****
	Apikey           *string `json:"Apikey,omitempty" xml:"Apikey,omitempty"`
	ThresholdPercent *int32  `json:"ThresholdPercent,omitempty" xml:"ThresholdPercent,omitempty"`
}

func (s UpdateMOQuotaAlertThresholdRequestApikey) String() string {
	return dara.Prettify(s)
}

func (s UpdateMOQuotaAlertThresholdRequestApikey) GoString() string {
	return s.String()
}

func (s *UpdateMOQuotaAlertThresholdRequestApikey) GetApikey() *string {
	return s.Apikey
}

func (s *UpdateMOQuotaAlertThresholdRequestApikey) GetThresholdPercent() *int32 {
	return s.ThresholdPercent
}

func (s *UpdateMOQuotaAlertThresholdRequestApikey) SetApikey(v string) *UpdateMOQuotaAlertThresholdRequestApikey {
	s.Apikey = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdRequestApikey) SetThresholdPercent(v int32) *UpdateMOQuotaAlertThresholdRequestApikey {
	s.ThresholdPercent = &v
	return s
}

func (s *UpdateMOQuotaAlertThresholdRequestApikey) Validate() error {
	return dara.Validate(s)
}
