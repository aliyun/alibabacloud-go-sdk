// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileTypeStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgramTypeAndCounts(v []*GetFileTypeStatisticResponseBodyProgramTypeAndCounts) *GetFileTypeStatisticResponseBody
	GetProgramTypeAndCounts() []*GetFileTypeStatisticResponseBodyProgramTypeAndCounts
	SetRequestId(v string) *GetFileTypeStatisticResponseBody
	GetRequestId() *string
}

type GetFileTypeStatisticResponseBody struct {
	// An array of node types and quantity.
	ProgramTypeAndCounts []*GetFileTypeStatisticResponseBodyProgramTypeAndCounts `json:"ProgramTypeAndCounts,omitempty" xml:"ProgramTypeAndCounts,omitempty" type:"Repeated"`
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// SDFSDFSDF-SDFSDF-SDFDSF-SDFSDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileTypeStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileTypeStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileTypeStatisticResponseBody) GetProgramTypeAndCounts() []*GetFileTypeStatisticResponseBodyProgramTypeAndCounts {
	return s.ProgramTypeAndCounts
}

func (s *GetFileTypeStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileTypeStatisticResponseBody) SetProgramTypeAndCounts(v []*GetFileTypeStatisticResponseBodyProgramTypeAndCounts) *GetFileTypeStatisticResponseBody {
	s.ProgramTypeAndCounts = v
	return s
}

func (s *GetFileTypeStatisticResponseBody) SetRequestId(v string) *GetFileTypeStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileTypeStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFileTypeStatisticResponseBodyProgramTypeAndCounts struct {
	// The number of nodes.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The type of the node.
	//
	// Valid values:
	//
	// 6 (Shell node), 10 (ODPS SQL node), 11 (ODPS MR node), 23 (Data Integration node), 24 (ODPS Script node), 99 (zero load node), 221 (PyODPS 2 node), 225 (ODPS Spark node), 227 (EMR Hive node), 228 (EMR Spark node), 229 (EMR Spark SQL node), 230 (EMR MR node), 239 (OSS object inspection node), 257 (EMR Shell node), 258 (EMR Spark Shell node), 259 (EMR Presto node), 260 (EMR Impala node), 900 (real-time data synchronization node), 1089 (cross-tenant collaboration node), 1091 (Hologres development node), 1093 (Hologres SQL node), 1100 (assignment node), and 1221 (PyODPS 3 node).
	//
	// example:
	//
	// ODPS_SQL
	ProgramType *string `json:"ProgramType,omitempty" xml:"ProgramType,omitempty"`
}

func (s GetFileTypeStatisticResponseBodyProgramTypeAndCounts) String() string {
	return dara.Prettify(s)
}

func (s GetFileTypeStatisticResponseBodyProgramTypeAndCounts) GoString() string {
	return s.String()
}

func (s *GetFileTypeStatisticResponseBodyProgramTypeAndCounts) GetCount() *int32 {
	return s.Count
}

func (s *GetFileTypeStatisticResponseBodyProgramTypeAndCounts) GetProgramType() *string {
	return s.ProgramType
}

func (s *GetFileTypeStatisticResponseBodyProgramTypeAndCounts) SetCount(v int32) *GetFileTypeStatisticResponseBodyProgramTypeAndCounts {
	s.Count = &v
	return s
}

func (s *GetFileTypeStatisticResponseBodyProgramTypeAndCounts) SetProgramType(v string) *GetFileTypeStatisticResponseBodyProgramTypeAndCounts {
	s.ProgramType = &v
	return s
}

func (s *GetFileTypeStatisticResponseBodyProgramTypeAndCounts) Validate() error {
	return dara.Validate(s)
}
