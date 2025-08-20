// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkReplSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSparkReplSessionResponseBodyData) *GetSparkReplSessionResponseBody
	GetData() *GetSparkReplSessionResponseBodyData
	SetRequestId(v string) *GetSparkReplSessionResponseBody
	GetRequestId() *string
}

type GetSparkReplSessionResponseBody struct {
	// The returned data.
	Data *GetSparkReplSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4CE6DF97-AEA4-484F-906F-C407EE3770EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkReplSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkReplSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkReplSessionResponseBody) GetData() *GetSparkReplSessionResponseBodyData {
	return s.Data
}

func (s *GetSparkReplSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkReplSessionResponseBody) SetData(v *GetSparkReplSessionResponseBodyData) *GetSparkReplSessionResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkReplSessionResponseBody) SetRequestId(v string) *GetSparkReplSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkReplSessionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSparkReplSessionResponseBodyData struct {
	// Indicates whether the session is active. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Active *string `json:"Active,omitempty" xml:"Active,omitempty"`
	// The ID of the Alibaba Cloud account that owns the cluster.
	//
	// example:
	//
	// 178157466101****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The attempt ID of the Spark application.
	//
	// example:
	//
	// s202301061000hz57d797b00002****
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	// The error message.
	//
	// example:
	//
	// Session timed out
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The ID of the session that executes the code.
	//
	// example:
	//
	// 1
	SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The status of the session. Valid values:
	//
	// 	- IDLE
	//
	// 	- BUSY
	//
	// 	- DEAD
	//
	// example:
	//
	// IDEL
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The URL of the web UI for the Spark application.
	//
	// example:
	//
	// https://adbsparkui-cn-hangzhou.aliyuncs.com/?token=****
	WebUiAddress *string `json:"WebUiAddress,omitempty" xml:"WebUiAddress,omitempty"`
}

func (s GetSparkReplSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSparkReplSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkReplSessionResponseBodyData) GetActive() *string {
	return s.Active
}

func (s *GetSparkReplSessionResponseBodyData) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *GetSparkReplSessionResponseBodyData) GetAttemptId() *string {
	return s.AttemptId
}

func (s *GetSparkReplSessionResponseBodyData) GetError() *string {
	return s.Error
}

func (s *GetSparkReplSessionResponseBodyData) GetSessionId() *int64 {
	return s.SessionId
}

func (s *GetSparkReplSessionResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetSparkReplSessionResponseBodyData) GetWebUiAddress() *string {
	return s.WebUiAddress
}

func (s *GetSparkReplSessionResponseBodyData) SetActive(v string) *GetSparkReplSessionResponseBodyData {
	s.Active = &v
	return s
}

func (s *GetSparkReplSessionResponseBodyData) SetAliyunUid(v string) *GetSparkReplSessionResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *GetSparkReplSessionResponseBodyData) SetAttemptId(v string) *GetSparkReplSessionResponseBodyData {
	s.AttemptId = &v
	return s
}

func (s *GetSparkReplSessionResponseBodyData) SetError(v string) *GetSparkReplSessionResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetSparkReplSessionResponseBodyData) SetSessionId(v int64) *GetSparkReplSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetSparkReplSessionResponseBodyData) SetState(v string) *GetSparkReplSessionResponseBodyData {
	s.State = &v
	return s
}

func (s *GetSparkReplSessionResponseBodyData) SetWebUiAddress(v string) *GetSparkReplSessionResponseBodyData {
	s.WebUiAddress = &v
	return s
}

func (s *GetSparkReplSessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
