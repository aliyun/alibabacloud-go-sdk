// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBlacklistCallTaggingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBlacklistCallTaggingsResponseBody
	GetCode() *string
	SetData(v []*ListBlacklistCallTaggingsResponseBodyData) *ListBlacklistCallTaggingsResponseBody
	GetData() []*ListBlacklistCallTaggingsResponseBodyData
	SetHttpStatusCode(v int32) *ListBlacklistCallTaggingsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBlacklistCallTaggingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBlacklistCallTaggingsResponseBody
	GetRequestId() *string
}

type ListBlacklistCallTaggingsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListBlacklistCallTaggingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 032C73C4-3A6F-4502-872B-4F5B41161C6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBlacklistCallTaggingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBlacklistCallTaggingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBlacklistCallTaggingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBlacklistCallTaggingsResponseBody) GetData() []*ListBlacklistCallTaggingsResponseBodyData {
	return s.Data
}

func (s *ListBlacklistCallTaggingsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBlacklistCallTaggingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBlacklistCallTaggingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBlacklistCallTaggingsResponseBody) SetCode(v string) *ListBlacklistCallTaggingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListBlacklistCallTaggingsResponseBody) SetData(v []*ListBlacklistCallTaggingsResponseBodyData) *ListBlacklistCallTaggingsResponseBody {
	s.Data = v
	return s
}

func (s *ListBlacklistCallTaggingsResponseBody) SetHttpStatusCode(v int32) *ListBlacklistCallTaggingsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBlacklistCallTaggingsResponseBody) SetMessage(v string) *ListBlacklistCallTaggingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListBlacklistCallTaggingsResponseBody) SetRequestId(v string) *ListBlacklistCallTaggingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBlacklistCallTaggingsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBlacklistCallTaggingsResponseBodyData struct {
	// example:
	//
	// true
	Blacklisted *bool `json:"Blacklisted,omitempty" xml:"Blacklisted,omitempty"`
	// example:
	//
	// job-481841171213393920
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1521083xxxx
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s ListBlacklistCallTaggingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBlacklistCallTaggingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBlacklistCallTaggingsResponseBodyData) GetBlacklisted() *bool {
	return s.Blacklisted
}

func (s *ListBlacklistCallTaggingsResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *ListBlacklistCallTaggingsResponseBodyData) GetNumber() *string {
	return s.Number
}

func (s *ListBlacklistCallTaggingsResponseBodyData) SetBlacklisted(v bool) *ListBlacklistCallTaggingsResponseBodyData {
	s.Blacklisted = &v
	return s
}

func (s *ListBlacklistCallTaggingsResponseBodyData) SetJobId(v string) *ListBlacklistCallTaggingsResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ListBlacklistCallTaggingsResponseBodyData) SetNumber(v string) *ListBlacklistCallTaggingsResponseBodyData {
	s.Number = &v
	return s
}

func (s *ListBlacklistCallTaggingsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
