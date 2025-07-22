// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoSqlOptimizeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v string) *UpdateAutoSqlOptimizeStatusRequest
	GetInstances() *string
	SetStatus(v int32) *UpdateAutoSqlOptimizeStatusRequest
	GetStatus() *int32
}

type UpdateAutoSqlOptimizeStatusRequest struct {
	// The database instance IDs. Separate multiple IDs with commas (,).
	//
	// >  You can specify up to 50 instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp10usoc1erj7****,rm-bp10usoc1erj7****
	Instances *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// The status of the automatic SQL optimization feature. Valid values:
	//
	// 	- **0**: The automatic SQL optimization feature is disabled.
	//
	// 	- **1**: **SQL diagnosis and automatic index creation*	- is specified.
	//
	// 	- **3**: **SQL diagnosis only*	- is specified.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAutoSqlOptimizeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoSqlOptimizeStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoSqlOptimizeStatusRequest) GetInstances() *string {
	return s.Instances
}

func (s *UpdateAutoSqlOptimizeStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateAutoSqlOptimizeStatusRequest) SetInstances(v string) *UpdateAutoSqlOptimizeStatusRequest {
	s.Instances = &v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusRequest) SetStatus(v int32) *UpdateAutoSqlOptimizeStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusRequest) Validate() error {
	return dara.Validate(s)
}
