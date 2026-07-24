// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDebugDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetJobDebugDataResponseBody
	GetCode() *int64
	SetData(v *GetJobDebugDataResponseBodyData) *GetJobDebugDataResponseBody
	GetData() *GetJobDebugDataResponseBodyData
	SetRequestId(v string) *GetJobDebugDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJobDebugDataResponseBody
	GetSuccess() *bool
}

type GetJobDebugDataResponseBody struct {
	Code      *int64                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetJobDebugDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobDebugDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobDebugDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobDebugDataResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetJobDebugDataResponseBody) GetData() *GetJobDebugDataResponseBodyData {
	return s.Data
}

func (s *GetJobDebugDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobDebugDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJobDebugDataResponseBody) SetCode(v int64) *GetJobDebugDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobDebugDataResponseBody) SetData(v *GetJobDebugDataResponseBodyData) *GetJobDebugDataResponseBody {
	s.Data = v
	return s
}

func (s *GetJobDebugDataResponseBody) SetRequestId(v string) *GetJobDebugDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobDebugDataResponseBody) SetSuccess(v bool) *GetJobDebugDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobDebugDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobDebugDataResponseBodyData struct {
	DataRows   []*GetJobDebugDataResponseBodyDataDataRows `json:"DataRows,omitempty" xml:"DataRows,omitempty" type:"Repeated"`
	DebugField *string                                    `json:"DebugField,omitempty" xml:"DebugField,omitempty"`
	HasMore    *bool                                      `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	Limit      *string                                    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextCursor *string                                    `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
}

func (s GetJobDebugDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetJobDebugDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobDebugDataResponseBodyData) GetDataRows() []*GetJobDebugDataResponseBodyDataDataRows {
	return s.DataRows
}

func (s *GetJobDebugDataResponseBodyData) GetDebugField() *string {
	return s.DebugField
}

func (s *GetJobDebugDataResponseBodyData) GetHasMore() *bool {
	return s.HasMore
}

func (s *GetJobDebugDataResponseBodyData) GetLimit() *string {
	return s.Limit
}

func (s *GetJobDebugDataResponseBodyData) GetNextCursor() *string {
	return s.NextCursor
}

func (s *GetJobDebugDataResponseBodyData) SetDataRows(v []*GetJobDebugDataResponseBodyDataDataRows) *GetJobDebugDataResponseBodyData {
	s.DataRows = v
	return s
}

func (s *GetJobDebugDataResponseBodyData) SetDebugField(v string) *GetJobDebugDataResponseBodyData {
	s.DebugField = &v
	return s
}

func (s *GetJobDebugDataResponseBodyData) SetHasMore(v bool) *GetJobDebugDataResponseBodyData {
	s.HasMore = &v
	return s
}

func (s *GetJobDebugDataResponseBodyData) SetLimit(v string) *GetJobDebugDataResponseBodyData {
	s.Limit = &v
	return s
}

func (s *GetJobDebugDataResponseBodyData) SetNextCursor(v string) *GetJobDebugDataResponseBodyData {
	s.NextCursor = &v
	return s
}

func (s *GetJobDebugDataResponseBodyData) Validate() error {
	if s.DataRows != nil {
		for _, item := range s.DataRows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetJobDebugDataResponseBodyDataDataRows struct {
	FlinkInstanceId *string `json:"FlinkInstanceId,omitempty" xml:"FlinkInstanceId,omitempty"`
	JobName         *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Offset          *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Partition       *int32  `json:"Partition,omitempty" xml:"Partition,omitempty"`
	ProcessedValue  *string `json:"ProcessedValue,omitempty" xml:"ProcessedValue,omitempty"`
	Timestamp       *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetJobDebugDataResponseBodyDataDataRows) String() string {
	return dara.Prettify(s)
}

func (s GetJobDebugDataResponseBodyDataDataRows) GoString() string {
	return s.String()
}

func (s *GetJobDebugDataResponseBodyDataDataRows) GetFlinkInstanceId() *string {
	return s.FlinkInstanceId
}

func (s *GetJobDebugDataResponseBodyDataDataRows) GetJobName() *string {
	return s.JobName
}

func (s *GetJobDebugDataResponseBodyDataDataRows) GetOffset() *int64 {
	return s.Offset
}

func (s *GetJobDebugDataResponseBodyDataDataRows) GetPartition() *int32 {
	return s.Partition
}

func (s *GetJobDebugDataResponseBodyDataDataRows) GetProcessedValue() *string {
	return s.ProcessedValue
}

func (s *GetJobDebugDataResponseBodyDataDataRows) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetJobDebugDataResponseBodyDataDataRows) GetUuid() *string {
	return s.Uuid
}

func (s *GetJobDebugDataResponseBodyDataDataRows) SetFlinkInstanceId(v string) *GetJobDebugDataResponseBodyDataDataRows {
	s.FlinkInstanceId = &v
	return s
}

func (s *GetJobDebugDataResponseBodyDataDataRows) SetJobName(v string) *GetJobDebugDataResponseBodyDataDataRows {
	s.JobName = &v
	return s
}

func (s *GetJobDebugDataResponseBodyDataDataRows) SetOffset(v int64) *GetJobDebugDataResponseBodyDataDataRows {
	s.Offset = &v
	return s
}

func (s *GetJobDebugDataResponseBodyDataDataRows) SetPartition(v int32) *GetJobDebugDataResponseBodyDataDataRows {
	s.Partition = &v
	return s
}

func (s *GetJobDebugDataResponseBodyDataDataRows) SetProcessedValue(v string) *GetJobDebugDataResponseBodyDataDataRows {
	s.ProcessedValue = &v
	return s
}

func (s *GetJobDebugDataResponseBodyDataDataRows) SetTimestamp(v int64) *GetJobDebugDataResponseBodyDataDataRows {
	s.Timestamp = &v
	return s
}

func (s *GetJobDebugDataResponseBodyDataDataRows) SetUuid(v string) *GetJobDebugDataResponseBodyDataDataRows {
	s.Uuid = &v
	return s
}

func (s *GetJobDebugDataResponseBodyDataDataRows) Validate() error {
	return dara.Validate(s)
}
