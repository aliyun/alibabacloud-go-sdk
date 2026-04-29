// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnknownThreatDetectMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperateType(v string) *OperateUnknownThreatDetectMachineRequest
	GetOperateType() *string
	SetStatus(v string) *OperateUnknownThreatDetectMachineRequest
	GetStatus() *string
	SetUuidList(v []*string) *OperateUnknownThreatDetectMachineRequest
	GetUuidList() []*string
}

type OperateUnknownThreatDetectMachineRequest struct {
	// example:
	//
	// restart_study
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// monitoring
	Status   *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s OperateUnknownThreatDetectMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateUnknownThreatDetectMachineRequest) GoString() string {
	return s.String()
}

func (s *OperateUnknownThreatDetectMachineRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *OperateUnknownThreatDetectMachineRequest) GetStatus() *string {
	return s.Status
}

func (s *OperateUnknownThreatDetectMachineRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *OperateUnknownThreatDetectMachineRequest) SetOperateType(v string) *OperateUnknownThreatDetectMachineRequest {
	s.OperateType = &v
	return s
}

func (s *OperateUnknownThreatDetectMachineRequest) SetStatus(v string) *OperateUnknownThreatDetectMachineRequest {
	s.Status = &v
	return s
}

func (s *OperateUnknownThreatDetectMachineRequest) SetUuidList(v []*string) *OperateUnknownThreatDetectMachineRequest {
	s.UuidList = v
	return s
}

func (s *OperateUnknownThreatDetectMachineRequest) Validate() error {
	return dara.Validate(s)
}
