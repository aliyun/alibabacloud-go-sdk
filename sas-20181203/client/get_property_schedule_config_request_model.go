// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPropertyScheduleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *GetPropertyScheduleConfigRequest
	GetType() *string
	SetUuid(v string) *GetPropertyScheduleConfigRequest
	GetUuid() *string
}

type GetPropertyScheduleConfigRequest struct {
	// The type of Asset Fingerprints for which the automatic collection frequency is configured. Valid values:
	//
	// - **scheduler_port_period**: listening port
	//
	// - **scheduler_process_period**: running process
	//
	// - **scheduler_account_period**: account asset
	//
	// - **scheduler_software_period**: software asset
	//
	// - **scheduler_cron_period**: scheduled node
	//
	// - **scheduler_sca_period**: middleware
	//
	// - **scheduler_autorun_period**: startup item
	//
	// - **scheduler_lkm_period**: kernel module
	//
	// - **scheduler_sca_proxy_period**: website
	//
	// example:
	//
	// scheduler_sca_period
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server to query.
	//
	// >You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// 50d213b4-3a35-427a-b8a5-04b0c7e1****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetPropertyScheduleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPropertyScheduleConfigRequest) GoString() string {
	return s.String()
}

func (s *GetPropertyScheduleConfigRequest) GetType() *string {
	return s.Type
}

func (s *GetPropertyScheduleConfigRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetPropertyScheduleConfigRequest) SetType(v string) *GetPropertyScheduleConfigRequest {
	s.Type = &v
	return s
}

func (s *GetPropertyScheduleConfigRequest) SetUuid(v string) *GetPropertyScheduleConfigRequest {
	s.Uuid = &v
	return s
}

func (s *GetPropertyScheduleConfigRequest) Validate() error {
	return dara.Validate(s)
}
