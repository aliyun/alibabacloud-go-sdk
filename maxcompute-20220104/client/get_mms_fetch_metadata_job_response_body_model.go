// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsFetchMetadataJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMmsFetchMetadataJobResponseBodyData) *GetMmsFetchMetadataJobResponseBody
	GetData() *GetMmsFetchMetadataJobResponseBodyData
	SetRequestId(v string) *GetMmsFetchMetadataJobResponseBody
	GetRequestId() *string
}

type GetMmsFetchMetadataJobResponseBody struct {
	Data *GetMmsFetchMetadataJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 5CA6292A-E301-5CD8-B4E2-AF060F99147B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsFetchMetadataJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMmsFetchMetadataJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsFetchMetadataJobResponseBody) GetData() *GetMmsFetchMetadataJobResponseBodyData {
	return s.Data
}

func (s *GetMmsFetchMetadataJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMmsFetchMetadataJobResponseBody) SetData(v *GetMmsFetchMetadataJobResponseBodyData) *GetMmsFetchMetadataJobResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBody) SetRequestId(v string) *GetMmsFetchMetadataJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsFetchMetadataJobResponseBodyData struct {
	// example:
	//
	// 2024-12-16 19:10:07
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// unexpected exception
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 1000002
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 5000
	Progress *float32 `json:"progress,omitempty" xml:"progress,omitempty"`
	// example:
	//
	// {"databases":5,"tables":75,"partitions":215}
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 2000015
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// 2024-12-16 19:09:37
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SCAN_DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetMmsFetchMetadataJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMmsFetchMetadataJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsFetchMetadataJobResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMmsFetchMetadataJobResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetMmsFetchMetadataJobResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMmsFetchMetadataJobResponseBodyData) GetProgress() *float32 {
	return s.Progress
}

func (s *GetMmsFetchMetadataJobResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *GetMmsFetchMetadataJobResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *GetMmsFetchMetadataJobResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMmsFetchMetadataJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetEndTime(v string) *GetMmsFetchMetadataJobResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetErrorMsg(v string) *GetMmsFetchMetadataJobResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetId(v int64) *GetMmsFetchMetadataJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetProgress(v float32) *GetMmsFetchMetadataJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetResult(v string) *GetMmsFetchMetadataJobResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetSourceId(v int64) *GetMmsFetchMetadataJobResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetStartTime(v string) *GetMmsFetchMetadataJobResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) SetStatus(v string) *GetMmsFetchMetadataJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
