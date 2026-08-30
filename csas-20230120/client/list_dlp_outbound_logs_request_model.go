// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDlpOutboundLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDlpOutboundLogsRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *ListDlpOutboundLogsRequest
	GetEndTime() *int64
	SetLogId(v string) *ListDlpOutboundLogsRequest
	GetLogId() *string
	SetPageSize(v int32) *ListDlpOutboundLogsRequest
	GetPageSize() *int32
	SetPolicyAction(v string) *ListDlpOutboundLogsRequest
	GetPolicyAction() *string
	SetSrcFileName(v string) *ListDlpOutboundLogsRequest
	GetSrcFileName() *string
	SetStartTime(v int64) *ListDlpOutboundLogsRequest
	GetStartTime() *int64
	SetSubChannelType(v string) *ListDlpOutboundLogsRequest
	GetSubChannelType() *string
	SetUserName(v string) *ListDlpOutboundLogsRequest
	GetUserName() *string
}

type ListDlpOutboundLogsRequest struct {
	// The current page number, starting from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the query. UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1754956800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The log ID.
	//
	// example:
	//
	// da817m4mfrcs6xxxx3hg
	LogId *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The policy action. Single-value exact match.
	//
	// example:
	//
	// block_and_hint
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// The original file name. Fuzzy match.
	//
	// example:
	//
	// TestFile
	SrcFileName *string `json:"SrcFileName,omitempty" xml:"SrcFileName,omitempty"`
	// The start time of the query. UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1754870400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The secondary channel ID in the format of `PrimaryChannelID.SubChannelID`. Separate multiple values with commas.
	//
	// example:
	//
	// 3.1,3.2
	SubChannelType *string `json:"SubChannelType,omitempty" xml:"SubChannelType,omitempty"`
	// The username. Exact match.
	//
	// example:
	//
	// zhangsan
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListDlpOutboundLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDlpOutboundLogsRequest) GoString() string {
	return s.String()
}

func (s *ListDlpOutboundLogsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDlpOutboundLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDlpOutboundLogsRequest) GetLogId() *string {
	return s.LogId
}

func (s *ListDlpOutboundLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDlpOutboundLogsRequest) GetPolicyAction() *string {
	return s.PolicyAction
}

func (s *ListDlpOutboundLogsRequest) GetSrcFileName() *string {
	return s.SrcFileName
}

func (s *ListDlpOutboundLogsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDlpOutboundLogsRequest) GetSubChannelType() *string {
	return s.SubChannelType
}

func (s *ListDlpOutboundLogsRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListDlpOutboundLogsRequest) SetCurrentPage(v int32) *ListDlpOutboundLogsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDlpOutboundLogsRequest) SetEndTime(v int64) *ListDlpOutboundLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDlpOutboundLogsRequest) SetLogId(v string) *ListDlpOutboundLogsRequest {
	s.LogId = &v
	return s
}

func (s *ListDlpOutboundLogsRequest) SetPageSize(v int32) *ListDlpOutboundLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDlpOutboundLogsRequest) SetPolicyAction(v string) *ListDlpOutboundLogsRequest {
	s.PolicyAction = &v
	return s
}

func (s *ListDlpOutboundLogsRequest) SetSrcFileName(v string) *ListDlpOutboundLogsRequest {
	s.SrcFileName = &v
	return s
}

func (s *ListDlpOutboundLogsRequest) SetStartTime(v int64) *ListDlpOutboundLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDlpOutboundLogsRequest) SetSubChannelType(v string) *ListDlpOutboundLogsRequest {
	s.SubChannelType = &v
	return s
}

func (s *ListDlpOutboundLogsRequest) SetUserName(v string) *ListDlpOutboundLogsRequest {
	s.UserName = &v
	return s
}

func (s *ListDlpOutboundLogsRequest) Validate() error {
	return dara.Validate(s)
}
