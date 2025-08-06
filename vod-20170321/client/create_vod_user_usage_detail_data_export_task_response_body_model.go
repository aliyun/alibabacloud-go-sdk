// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodUserUsageDetailDataExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CreateVodUserUsageDetailDataExportTaskResponseBody
	GetEndTime() *string
	SetRequestId(v string) *CreateVodUserUsageDetailDataExportTaskResponseBody
	GetRequestId() *string
	SetStartTime(v string) *CreateVodUserUsageDetailDataExportTaskResponseBody
	GetStartTime() *string
	SetTaskId(v string) *CreateVodUserUsageDetailDataExportTaskResponseBody
	GetTaskId() *string
}

type CreateVodUserUsageDetailDataExportTaskResponseBody struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVodUserUsageDetailDataExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVodUserUsageDetailDataExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVodUserUsageDetailDataExportTaskResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateVodUserUsageDetailDataExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVodUserUsageDetailDataExportTaskResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateVodUserUsageDetailDataExportTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVodUserUsageDetailDataExportTaskResponseBody) SetEndTime(v string) *CreateVodUserUsageDetailDataExportTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskResponseBody) SetRequestId(v string) *CreateVodUserUsageDetailDataExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskResponseBody) SetStartTime(v string) *CreateVodUserUsageDetailDataExportTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskResponseBody) SetTaskId(v string) *CreateVodUserUsageDetailDataExportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVodUserUsageDetailDataExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
