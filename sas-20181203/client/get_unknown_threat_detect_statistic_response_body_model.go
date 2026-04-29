// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnknownThreatDetectStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetUnknownThreatDetectStatisticResponseBodyData) *GetUnknownThreatDetectStatisticResponseBody
	GetData() *GetUnknownThreatDetectStatisticResponseBodyData
	SetRequestId(v string) *GetUnknownThreatDetectStatisticResponseBody
	GetRequestId() *string
}

type GetUnknownThreatDetectStatisticResponseBody struct {
	Data *GetUnknownThreatDetectStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// FD394AF6-591E-5168-8C8C-4C7847******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUnknownThreatDetectStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUnknownThreatDetectStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnknownThreatDetectStatisticResponseBody) GetData() *GetUnknownThreatDetectStatisticResponseBodyData {
	return s.Data
}

func (s *GetUnknownThreatDetectStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUnknownThreatDetectStatisticResponseBody) SetData(v *GetUnknownThreatDetectStatisticResponseBodyData) *GetUnknownThreatDetectStatisticResponseBody {
	s.Data = v
	return s
}

func (s *GetUnknownThreatDetectStatisticResponseBody) SetRequestId(v string) *GetUnknownThreatDetectStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUnknownThreatDetectStatisticResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUnknownThreatDetectStatisticResponseBodyData struct {
	// example:
	//
	// 1
	BlockMachineCount *int32 `json:"BlockMachineCount,omitempty" xml:"BlockMachineCount,omitempty"`
	// example:
	//
	// 1
	MachineCount *int32 `json:"MachineCount,omitempty" xml:"MachineCount,omitempty"`
	// example:
	//
	// 1
	MonitorMachineCount *int32 `json:"MonitorMachineCount,omitempty" xml:"MonitorMachineCount,omitempty"`
	// example:
	//
	// 1
	OpenMachineCount *int32 `json:"OpenMachineCount,omitempty" xml:"OpenMachineCount,omitempty"`
	// example:
	//
	// 1
	StudyingMachineCount *int32 `json:"StudyingMachineCount,omitempty" xml:"StudyingMachineCount,omitempty"`
}

func (s GetUnknownThreatDetectStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUnknownThreatDetectStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUnknownThreatDetectStatisticResponseBodyData) GetBlockMachineCount() *int32 {
	return s.BlockMachineCount
}

func (s *GetUnknownThreatDetectStatisticResponseBodyData) GetMachineCount() *int32 {
	return s.MachineCount
}

func (s *GetUnknownThreatDetectStatisticResponseBodyData) GetMonitorMachineCount() *int32 {
	return s.MonitorMachineCount
}

func (s *GetUnknownThreatDetectStatisticResponseBodyData) GetOpenMachineCount() *int32 {
	return s.OpenMachineCount
}

func (s *GetUnknownThreatDetectStatisticResponseBodyData) GetStudyingMachineCount() *int32 {
	return s.StudyingMachineCount
}

func (s *GetUnknownThreatDetectStatisticResponseBodyData) SetBlockMachineCount(v int32) *GetUnknownThreatDetectStatisticResponseBodyData {
	s.BlockMachineCount = &v
	return s
}

func (s *GetUnknownThreatDetectStatisticResponseBodyData) SetMachineCount(v int32) *GetUnknownThreatDetectStatisticResponseBodyData {
	s.MachineCount = &v
	return s
}

func (s *GetUnknownThreatDetectStatisticResponseBodyData) SetMonitorMachineCount(v int32) *GetUnknownThreatDetectStatisticResponseBodyData {
	s.MonitorMachineCount = &v
	return s
}

func (s *GetUnknownThreatDetectStatisticResponseBodyData) SetOpenMachineCount(v int32) *GetUnknownThreatDetectStatisticResponseBodyData {
	s.OpenMachineCount = &v
	return s
}

func (s *GetUnknownThreatDetectStatisticResponseBodyData) SetStudyingMachineCount(v int32) *GetUnknownThreatDetectStatisticResponseBodyData {
	s.StudyingMachineCount = &v
	return s
}

func (s *GetUnknownThreatDetectStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}
