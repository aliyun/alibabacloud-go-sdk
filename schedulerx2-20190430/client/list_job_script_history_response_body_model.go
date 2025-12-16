// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobScriptHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListJobScriptHistoryResponseBody
	GetCode() *int32
	SetData(v *ListJobScriptHistoryResponseBodyData) *ListJobScriptHistoryResponseBody
	GetData() *ListJobScriptHistoryResponseBodyData
	SetMessage(v string) *ListJobScriptHistoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListJobScriptHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobScriptHistoryResponseBody
	GetSuccess() *bool
}

type ListJobScriptHistoryResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the jobs.
	Data *ListJobScriptHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information returned only if an error occurs.
	//
	// example:
	//
	// job is not existed, jobId=302
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true
	//
	// false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListJobScriptHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobScriptHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobScriptHistoryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListJobScriptHistoryResponseBody) GetData() *ListJobScriptHistoryResponseBodyData {
	return s.Data
}

func (s *ListJobScriptHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListJobScriptHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobScriptHistoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobScriptHistoryResponseBody) SetCode(v int32) *ListJobScriptHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobScriptHistoryResponseBody) SetData(v *ListJobScriptHistoryResponseBodyData) *ListJobScriptHistoryResponseBody {
	s.Data = v
	return s
}

func (s *ListJobScriptHistoryResponseBody) SetMessage(v string) *ListJobScriptHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListJobScriptHistoryResponseBody) SetRequestId(v string) *ListJobScriptHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobScriptHistoryResponseBody) SetSuccess(v bool) *ListJobScriptHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobScriptHistoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobScriptHistoryResponseBodyData struct {
	// The information about the job\\"s historical scripts.
	JobScriptHistoryInfos []*ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos `json:"JobScriptHistoryInfos,omitempty" xml:"JobScriptHistoryInfos,omitempty" type:"Repeated"`
}

func (s ListJobScriptHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListJobScriptHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobScriptHistoryResponseBodyData) GetJobScriptHistoryInfos() []*ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos {
	return s.JobScriptHistoryInfos
}

func (s *ListJobScriptHistoryResponseBodyData) SetJobScriptHistoryInfos(v []*ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) *ListJobScriptHistoryResponseBodyData {
	s.JobScriptHistoryInfos = v
	return s
}

func (s *ListJobScriptHistoryResponseBodyData) Validate() error {
	if s.JobScriptHistoryInfos != nil {
		for _, item := range s.JobScriptHistoryInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos struct {
	// The creation time.
	//
	// example:
	//
	// 2025-03-12 14:52:42
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	//
	// example:
	//
	// 1272118248844842
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The script content.
	//
	// example:
	//
	// #!/bin/bash
	//
	// # The following are predefined variables provided by the system. You can use them to obtain information about the job run.
	//
	// echo "Job parameters: #{schedulerx.jobParameters}"
	//
	// echo "Shard index: #{schedulerx.shardingId}"
	//
	// echo "Shard parameters: #{schedulerx.shardingParameters}"
	//
	// echo "Total number of shards: #{schedulerx.shardingNum}"
	//
	// echo "Current retry count: #{schedulerx.attempt}"
	//
	// echo "Trigger type: #{schedulerx.triggerType}"
	//
	// echo "Scheduled timestamp: #{schedulerx.scheduleTime}"
	//
	// echo "Data timestamp: #{schedulerx.dataTime}"
	//
	// # The output of the last line will be returned as the result
	//
	// echo "hello world"
	//
	// # exit 1 indicates failure
	//
	// exit 0
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	// The description of the script version.
	//
	// example:
	//
	// init version
	VersionesDescription *string `json:"VersionesDescription,omitempty" xml:"VersionesDescription,omitempty"`
}

func (s ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) String() string {
	return dara.Prettify(s)
}

func (s ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) GoString() string {
	return s.String()
}

func (s *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) GetCreator() *string {
	return s.Creator
}

func (s *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) GetVersionesDescription() *string {
	return s.VersionesDescription
}

func (s *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) SetCreateTime(v string) *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos {
	s.CreateTime = &v
	return s
}

func (s *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) SetCreator(v string) *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos {
	s.Creator = &v
	return s
}

func (s *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) SetScriptContent(v string) *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos {
	s.ScriptContent = &v
	return s
}

func (s *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) SetVersionesDescription(v string) *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos {
	s.VersionesDescription = &v
	return s
}

func (s *ListJobScriptHistoryResponseBodyDataJobScriptHistoryInfos) Validate() error {
	return dara.Validate(s)
}
