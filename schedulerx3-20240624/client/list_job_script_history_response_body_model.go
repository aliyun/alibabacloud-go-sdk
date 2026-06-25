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
	SetMaxResults(v int32) *ListJobScriptHistoryResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListJobScriptHistoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListJobScriptHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobScriptHistoryResponseBody
	GetSuccess() *bool
}

type ListJobScriptHistoryResponseBody struct {
	// The response code. A value of `200` indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// - The response data.
	Data *ListJobScriptHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The response message.
	//
	// This parameter is required.
	//
	// example:
	//
	// not support query script history, please upgrade engine version to 2.2.2+
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A unique ID for the request. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
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

func (s *ListJobScriptHistoryResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
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

func (s *ListJobScriptHistoryResponseBody) SetMaxResults(v int32) *ListJobScriptHistoryResponseBody {
	s.MaxResults = &v
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
	// The token to retrieve the next page of results. If this parameter is empty, no more data is available.
	//
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// - A list of script history records.
	Records []*ListJobScriptHistoryResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total count of entries.
	//
	// example:
	//
	// 21
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListJobScriptHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListJobScriptHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobScriptHistoryResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListJobScriptHistoryResponseBodyData) GetRecords() []*ListJobScriptHistoryResponseBodyDataRecords {
	return s.Records
}

func (s *ListJobScriptHistoryResponseBodyData) GetTotal() *string {
	return s.Total
}

func (s *ListJobScriptHistoryResponseBodyData) SetNextToken(v string) *ListJobScriptHistoryResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListJobScriptHistoryResponseBodyData) SetRecords(v []*ListJobScriptHistoryResponseBodyDataRecords) *ListJobScriptHistoryResponseBodyData {
	s.Records = v
	return s
}

func (s *ListJobScriptHistoryResponseBodyData) SetTotal(v string) *ListJobScriptHistoryResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListJobScriptHistoryResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobScriptHistoryResponseBodyDataRecords struct {
	// The timestamp when the script version was created.
	//
	// example:
	//
	// 2025-06-29 15:56:36
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the script version.
	//
	// example:
	//
	// 1963096506470832
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The content of the script.
	//
	// example:
	//
	// #!/bin/bash
	//
	// echo "xxl-job: hello shell"
	//
	// echo "脚本位置: $0"
	//
	// echo "任务参数: $1"
	//
	// echo "分片序号 = $2"
	//
	// echo "分片总数 = $3"
	//
	// echo "Good bye!"
	//
	// exit 0
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	// The script version description.
	//
	// example:
	//
	// init version
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
}

func (s ListJobScriptHistoryResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListJobScriptHistoryResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListJobScriptHistoryResponseBodyDataRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListJobScriptHistoryResponseBodyDataRecords) GetCreator() *string {
	return s.Creator
}

func (s *ListJobScriptHistoryResponseBodyDataRecords) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *ListJobScriptHistoryResponseBodyDataRecords) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *ListJobScriptHistoryResponseBodyDataRecords) SetCreateTime(v string) *ListJobScriptHistoryResponseBodyDataRecords {
	s.CreateTime = &v
	return s
}

func (s *ListJobScriptHistoryResponseBodyDataRecords) SetCreator(v string) *ListJobScriptHistoryResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *ListJobScriptHistoryResponseBodyDataRecords) SetScriptContent(v string) *ListJobScriptHistoryResponseBodyDataRecords {
	s.ScriptContent = &v
	return s
}

func (s *ListJobScriptHistoryResponseBodyDataRecords) SetVersionDescription(v string) *ListJobScriptHistoryResponseBodyDataRecords {
	s.VersionDescription = &v
	return s
}

func (s *ListJobScriptHistoryResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
