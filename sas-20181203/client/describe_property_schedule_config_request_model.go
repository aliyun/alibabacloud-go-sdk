// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScheduleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *DescribePropertyScheduleConfigRequest
	GetType() *string
}

type DescribePropertyScheduleConfigRequest struct {
	// The type of the asset fingerprints. Valid values:
	//
	// 	- **scheduler_port_period**: listening port
	//
	// 	- **scheduler_process_period**: running process
	//
	// 	- **scheduler_account_period**: account
	//
	// 	- **scheduler_software_period**: software
	//
	// 	- **scheduler_cron_period**: scheduled task
	//
	// 	- **scheduler_sca_period**: middleware
	//
	// 	- **scheduler_autorun_period**: startup item
	//
	// 	- **scheduler_lkm_period**: kernel module
	//
	// 	- **scheduler_sca_proxy_period**: website
	//
	// This parameter is required.
	//
	// example:
	//
	// scheduler_autorun_period
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePropertyScheduleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScheduleConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyScheduleConfigRequest) GetType() *string {
	return s.Type
}

func (s *DescribePropertyScheduleConfigRequest) SetType(v string) *DescribePropertyScheduleConfigRequest {
	s.Type = &v
	return s
}

func (s *DescribePropertyScheduleConfigRequest) Validate() error {
	return dara.Validate(s)
}
