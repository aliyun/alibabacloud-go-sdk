// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckPolarFsQuotaConsistencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableRepair(v bool) *CheckPolarFsQuotaConsistencyRequest
	GetEnableRepair() *bool
	SetEnableStrictCalculate(v bool) *CheckPolarFsQuotaConsistencyRequest
	GetEnableStrictCalculate() *bool
	SetPath(v string) *CheckPolarFsQuotaConsistencyRequest
	GetPath() *string
	SetPolarFsInstanceId(v string) *CheckPolarFsQuotaConsistencyRequest
	GetPolarFsInstanceId() *string
}

type CheckPolarFsQuotaConsistencyRequest struct {
	// example:
	//
	// false
	EnableRepair *bool `json:"EnableRepair,omitempty" xml:"EnableRepair,omitempty"`
	// example:
	//
	// false
	EnableStrictCalculate *bool `json:"EnableStrictCalculate,omitempty" xml:"EnableStrictCalculate,omitempty"`
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-test****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s CheckPolarFsQuotaConsistencyRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckPolarFsQuotaConsistencyRequest) GoString() string {
	return s.String()
}

func (s *CheckPolarFsQuotaConsistencyRequest) GetEnableRepair() *bool {
	return s.EnableRepair
}

func (s *CheckPolarFsQuotaConsistencyRequest) GetEnableStrictCalculate() *bool {
	return s.EnableStrictCalculate
}

func (s *CheckPolarFsQuotaConsistencyRequest) GetPath() *string {
	return s.Path
}

func (s *CheckPolarFsQuotaConsistencyRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *CheckPolarFsQuotaConsistencyRequest) SetEnableRepair(v bool) *CheckPolarFsQuotaConsistencyRequest {
	s.EnableRepair = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyRequest) SetEnableStrictCalculate(v bool) *CheckPolarFsQuotaConsistencyRequest {
	s.EnableStrictCalculate = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyRequest) SetPath(v string) *CheckPolarFsQuotaConsistencyRequest {
	s.Path = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyRequest) SetPolarFsInstanceId(v string) *CheckPolarFsQuotaConsistencyRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *CheckPolarFsQuotaConsistencyRequest) Validate() error {
	return dara.Validate(s)
}
