// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncRecordingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SyncRecordingRulesResponseBody
	GetCode() *int32
	SetData(v string) *SyncRecordingRulesResponseBody
	GetData() *string
	SetMessage(v string) *SyncRecordingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncRecordingRulesResponseBody
	GetRequestId() *string
}

type SyncRecordingRulesResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response struct.
	//
	// example:
	//
	// {     "data":[         "c06ca68cd16f14f52bb07772eda***",         "c33dd70a0ac184c1b879d807ab2***",         "c384cf7e4dcb543e6ac8c7d4dd3***",         "ce30f833bc4a04a56a06b070319***"     ],     "message":"设置RecordingRule失败的clusterIds",     "success":true }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned with the result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncRecordingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncRecordingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *SyncRecordingRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SyncRecordingRulesResponseBody) GetData() *string {
	return s.Data
}

func (s *SyncRecordingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncRecordingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncRecordingRulesResponseBody) SetCode(v int32) *SyncRecordingRulesResponseBody {
	s.Code = &v
	return s
}

func (s *SyncRecordingRulesResponseBody) SetData(v string) *SyncRecordingRulesResponseBody {
	s.Data = &v
	return s
}

func (s *SyncRecordingRulesResponseBody) SetMessage(v string) *SyncRecordingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *SyncRecordingRulesResponseBody) SetRequestId(v string) *SyncRecordingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncRecordingRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
