// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeRecordFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkDescribeRecordFileUrlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkDescribeRecordFileUrlResponseBody
	GetCode() *string
	SetData(v *ClinkDescribeRecordFileUrlResponseBodyData) *ClinkDescribeRecordFileUrlResponseBody
	GetData() *ClinkDescribeRecordFileUrlResponseBodyData
	SetMessage(v string) *ClinkDescribeRecordFileUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkDescribeRecordFileUrlResponseBody
	GetRequestId() *string
}

type ClinkDescribeRecordFileUrlResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkDescribeRecordFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkDescribeRecordFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeRecordFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkDescribeRecordFileUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkDescribeRecordFileUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkDescribeRecordFileUrlResponseBody) GetData() *ClinkDescribeRecordFileUrlResponseBodyData {
	return s.Data
}

func (s *ClinkDescribeRecordFileUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkDescribeRecordFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkDescribeRecordFileUrlResponseBody) SetAccessDeniedDetail(v string) *ClinkDescribeRecordFileUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponseBody) SetCode(v string) *ClinkDescribeRecordFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponseBody) SetData(v *ClinkDescribeRecordFileUrlResponseBodyData) *ClinkDescribeRecordFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponseBody) SetMessage(v string) *ClinkDescribeRecordFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponseBody) SetRequestId(v string) *ClinkDescribeRecordFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDescribeRecordFileUrlResponseBodyData struct {
	// 录音流程列表。仅当获取 mp3 格式通话录音时返回（不传 recordSide 且 recordType 为 "record" 或未传）。
	AudioFlows []*ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows `json:"AudioFlows,omitempty" xml:"AudioFlows,omitempty" type:"Repeated"`
	// 请求 id
	//
	// example:
	//
	// 示例值示例值
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 通话录音地址
	//
	// example:
	//
	// xxx
	RecordFileUrl *string `json:"RecordFileUrl,omitempty" xml:"RecordFileUrl,omitempty"`
}

func (s ClinkDescribeRecordFileUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeRecordFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkDescribeRecordFileUrlResponseBodyData) GetAudioFlows() []*ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows {
	return s.AudioFlows
}

func (s *ClinkDescribeRecordFileUrlResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkDescribeRecordFileUrlResponseBodyData) GetRecordFileUrl() *string {
	return s.RecordFileUrl
}

func (s *ClinkDescribeRecordFileUrlResponseBodyData) SetAudioFlows(v []*ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows) *ClinkDescribeRecordFileUrlResponseBodyData {
	s.AudioFlows = v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponseBodyData) SetClinkRequestId(v string) *ClinkDescribeRecordFileUrlResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponseBodyData) SetRecordFileUrl(v string) *ClinkDescribeRecordFileUrlResponseBodyData {
	s.RecordFileUrl = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponseBodyData) Validate() error {
	if s.AudioFlows != nil {
		for _, item := range s.AudioFlows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows struct {
	// 节点，1：录音，2：保持，3：静音，4：咨询
	//
	// example:
	//
	// 1
	Node *int64 `json:"Node,omitempty" xml:"Node,omitempty"`
	// 时间戳，单位：秒
	//
	// example:
	//
	// 19
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// 类型，0：开始，1：结束
	//
	// example:
	//
	// 0
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows) GoString() string {
	return s.String()
}

func (s *ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows) GetNode() *int64 {
	return s.Node
}

func (s *ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows) GetType() *int64 {
	return s.Type
}

func (s *ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows) SetNode(v int64) *ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows {
	s.Node = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows) SetTimestamp(v int64) *ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows {
	s.Timestamp = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows) SetType(v int64) *ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows {
	s.Type = &v
	return s
}

func (s *ClinkDescribeRecordFileUrlResponseBodyDataAudioFlows) Validate() error {
	return dara.Validate(s)
}
