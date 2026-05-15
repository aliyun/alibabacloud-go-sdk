// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeKnowledgeBaseStatsResponseBodyData) *DescribeKnowledgeBaseStatsResponseBody
	GetData() *DescribeKnowledgeBaseStatsResponseBodyData
	SetErrorCode(v string) *DescribeKnowledgeBaseStatsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeKnowledgeBaseStatsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeKnowledgeBaseStatsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeKnowledgeBaseStatsResponseBody
	GetSuccess() *bool
}

type DescribeKnowledgeBaseStatsResponseBody struct {
	Data *DescribeKnowledgeBaseStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Permission denied.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Knowledge base limit exceeded. Current: xxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeKnowledgeBaseStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseStatsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseStatsResponseBody) GetData() *DescribeKnowledgeBaseStatsResponseBodyData {
	return s.Data
}

func (s *DescribeKnowledgeBaseStatsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeKnowledgeBaseStatsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeKnowledgeBaseStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKnowledgeBaseStatsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeKnowledgeBaseStatsResponseBody) SetData(v *DescribeKnowledgeBaseStatsResponseBodyData) *DescribeKnowledgeBaseStatsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponseBody) SetErrorCode(v string) *DescribeKnowledgeBaseStatsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponseBody) SetErrorMessage(v string) *DescribeKnowledgeBaseStatsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponseBody) SetRequestId(v string) *DescribeKnowledgeBaseStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponseBody) SetSuccess(v bool) *DescribeKnowledgeBaseStatsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKnowledgeBaseStatsResponseBodyData struct {
	// example:
	//
	// 100
	DocumentCount *int32 `json:"DocumentCount,omitempty" xml:"DocumentCount,omitempty"`
	// example:
	//
	// 18
	KbHits *int64 `json:"KbHits,omitempty" xml:"KbHits,omitempty"`
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// example:
	//
	// 10
	TotalChunkCount *int32 `json:"TotalChunkCount,omitempty" xml:"TotalChunkCount,omitempty"`
	// example:
	//
	// 4194588751
	TotalFileSize *int64 `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
}

func (s DescribeKnowledgeBaseStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseStatsResponseBodyData) GetDocumentCount() *int32 {
	return s.DocumentCount
}

func (s *DescribeKnowledgeBaseStatsResponseBodyData) GetKbHits() *int64 {
	return s.KbHits
}

func (s *DescribeKnowledgeBaseStatsResponseBodyData) GetKbUuid() *string {
	return s.KbUuid
}

func (s *DescribeKnowledgeBaseStatsResponseBodyData) GetTotalChunkCount() *int32 {
	return s.TotalChunkCount
}

func (s *DescribeKnowledgeBaseStatsResponseBodyData) GetTotalFileSize() *int64 {
	return s.TotalFileSize
}

func (s *DescribeKnowledgeBaseStatsResponseBodyData) SetDocumentCount(v int32) *DescribeKnowledgeBaseStatsResponseBodyData {
	s.DocumentCount = &v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponseBodyData) SetKbHits(v int64) *DescribeKnowledgeBaseStatsResponseBodyData {
	s.KbHits = &v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponseBodyData) SetKbUuid(v string) *DescribeKnowledgeBaseStatsResponseBodyData {
	s.KbUuid = &v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponseBodyData) SetTotalChunkCount(v int32) *DescribeKnowledgeBaseStatsResponseBodyData {
	s.TotalChunkCount = &v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponseBodyData) SetTotalFileSize(v int64) *DescribeKnowledgeBaseStatsResponseBodyData {
	s.TotalFileSize = &v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
