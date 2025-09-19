// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckCountStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatisticType(v string) *GetCheckCountStatisticRequest
	GetStatisticType() *string
	SetTaskSources(v []*string) *GetCheckCountStatisticRequest
	GetTaskSources() []*string
	SetVendors(v []*string) *GetCheckCountStatisticRequest
	GetVendors() []*string
}

type GetCheckCountStatisticRequest struct {
	// The type of the statistics. Valid values:
	//
	// 	- **user**: the top five users that are granted excessive permissions.
	//
	// 	- **role**: the top five roles that are granted excessive permissions.
	//
	// 	- **instance**: the top five cloud services on which risks are detected.
	//
	// 	- **host**: the top five servers on which baseline risks are detected.
	//
	// example:
	//
	// instance
	StatisticType *string   `json:"StatisticType,omitempty" xml:"StatisticType,omitempty"`
	TaskSources   []*string `json:"TaskSources,omitempty" xml:"TaskSources,omitempty" type:"Repeated"`
	// The cloud service providers.
	Vendors []*string `json:"Vendors,omitempty" xml:"Vendors,omitempty" type:"Repeated"`
}

func (s GetCheckCountStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckCountStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetCheckCountStatisticRequest) GetStatisticType() *string {
	return s.StatisticType
}

func (s *GetCheckCountStatisticRequest) GetTaskSources() []*string {
	return s.TaskSources
}

func (s *GetCheckCountStatisticRequest) GetVendors() []*string {
	return s.Vendors
}

func (s *GetCheckCountStatisticRequest) SetStatisticType(v string) *GetCheckCountStatisticRequest {
	s.StatisticType = &v
	return s
}

func (s *GetCheckCountStatisticRequest) SetTaskSources(v []*string) *GetCheckCountStatisticRequest {
	s.TaskSources = v
	return s
}

func (s *GetCheckCountStatisticRequest) SetVendors(v []*string) *GetCheckCountStatisticRequest {
	s.Vendors = v
	return s
}

func (s *GetCheckCountStatisticRequest) Validate() error {
	return dara.Validate(s)
}
