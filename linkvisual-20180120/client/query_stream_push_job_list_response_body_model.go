// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStreamPushJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryStreamPushJobListResponseBody
	GetCode() *string
	SetData(v *QueryStreamPushJobListResponseBodyData) *QueryStreamPushJobListResponseBody
	GetData() *QueryStreamPushJobListResponseBodyData
	SetErrorMessage(v string) *QueryStreamPushJobListResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryStreamPushJobListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryStreamPushJobListResponseBody
	GetSuccess() *bool
}

type QueryStreamPushJobListResponseBody struct {
	// The response code. Valid values:
	//
	// - **200**: The call was successful.
	//
	// - Other values: error codes. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 400
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned if the call is successful.
	Data *QueryStreamPushJobListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// Request parameter error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E55***B7-4***-4***-8***-D3******F565
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryStreamPushJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryStreamPushJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryStreamPushJobListResponseBody) GetData() *QueryStreamPushJobListResponseBodyData {
	return s.Data
}

func (s *QueryStreamPushJobListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryStreamPushJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryStreamPushJobListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryStreamPushJobListResponseBody) SetCode(v string) *QueryStreamPushJobListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStreamPushJobListResponseBody) SetData(v *QueryStreamPushJobListResponseBodyData) *QueryStreamPushJobListResponseBody {
	s.Data = v
	return s
}

func (s *QueryStreamPushJobListResponseBody) SetErrorMessage(v string) *QueryStreamPushJobListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryStreamPushJobListResponseBody) SetRequestId(v string) *QueryStreamPushJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStreamPushJobListResponseBody) SetSuccess(v bool) *QueryStreamPushJobListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryStreamPushJobListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryStreamPushJobListResponseBodyData struct {
	// The job list information.
	JobList []*QueryStreamPushJobListResponseBodyDataJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	// The total number of jobs.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryStreamPushJobListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryStreamPushJobListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobListResponseBodyData) GetJobList() []*QueryStreamPushJobListResponseBodyDataJobList {
	return s.JobList
}

func (s *QueryStreamPushJobListResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryStreamPushJobListResponseBodyData) SetJobList(v []*QueryStreamPushJobListResponseBodyDataJobList) *QueryStreamPushJobListResponseBodyData {
	s.JobList = v
	return s
}

func (s *QueryStreamPushJobListResponseBodyData) SetTotal(v int32) *QueryStreamPushJobListResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyData) Validate() error {
	if s.JobList != nil {
		for _, item := range s.JobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryStreamPushJobListResponseBodyDataJobList struct {
	// The time when the job was created. Unit: seconds.
	//
	// example:
	//
	// 1581595942
	CreateTime *int32 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the job.
	//
	// example:
	//
	// 23********
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The type of the stream push job. Valid values:
	//
	// - **1**: continuous job.
	//
	// - **2**: triggered job. The job ends after it is interrupted.
	//
	// example:
	//
	// 1
	JobType *int32 `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The current stream push status. Valid values:
	//
	// - **1**: pushing.
	//
	// - **2**: waiting to push.
	//
	// - **3**: push ended.
	//
	// example:
	//
	// 1
	PushStatus *int32 `json:"PushStatus,omitempty" xml:"PushStatus,omitempty"`
	// The stream push URL.
	//
	// example:
	//
	// rtmp://****
	PushUrl *string `json:"PushUrl,omitempty" xml:"PushUrl,omitempty"`
	// The stream type. Valid values:
	//
	// - **0**: main stream.
	//
	// - **1**: sub stream.
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s QueryStreamPushJobListResponseBodyDataJobList) String() string {
	return dara.Prettify(s)
}

func (s QueryStreamPushJobListResponseBodyDataJobList) GoString() string {
	return s.String()
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) GetJobId() *string {
	return s.JobId
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) GetJobType() *int32 {
	return s.JobType
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) GetPushStatus() *int32 {
	return s.PushStatus
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) GetPushUrl() *string {
	return s.PushUrl
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) GetStreamType() *int32 {
	return s.StreamType
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetCreateTime(v int32) *QueryStreamPushJobListResponseBodyDataJobList {
	s.CreateTime = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetJobId(v string) *QueryStreamPushJobListResponseBodyDataJobList {
	s.JobId = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetJobType(v int32) *QueryStreamPushJobListResponseBodyDataJobList {
	s.JobType = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetPushStatus(v int32) *QueryStreamPushJobListResponseBodyDataJobList {
	s.PushStatus = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetPushUrl(v string) *QueryStreamPushJobListResponseBodyDataJobList {
	s.PushUrl = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) SetStreamType(v int32) *QueryStreamPushJobListResponseBodyDataJobList {
	s.StreamType = &v
	return s
}

func (s *QueryStreamPushJobListResponseBodyDataJobList) Validate() error {
	return dara.Validate(s)
}
