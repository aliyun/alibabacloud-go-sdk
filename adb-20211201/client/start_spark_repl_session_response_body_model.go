// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSparkReplSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StartSparkReplSessionResponseBodyData) *StartSparkReplSessionResponseBody
	GetData() *StartSparkReplSessionResponseBodyData
	SetRequestId(v string) *StartSparkReplSessionResponseBody
	GetRequestId() *string
}

type StartSparkReplSessionResponseBody struct {
	// The returned data.
	Data *StartSparkReplSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C1797FEA-B7D6-5ED6-A24B-2C8C5F4D7361
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartSparkReplSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSparkReplSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StartSparkReplSessionResponseBody) GetData() *StartSparkReplSessionResponseBodyData {
	return s.Data
}

func (s *StartSparkReplSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSparkReplSessionResponseBody) SetData(v *StartSparkReplSessionResponseBodyData) *StartSparkReplSessionResponseBody {
	s.Data = v
	return s
}

func (s *StartSparkReplSessionResponseBody) SetRequestId(v string) *StartSparkReplSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSparkReplSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartSparkReplSessionResponseBodyData struct {
	// The ID of the Alibaba Cloud account that owns the cluster.
	//
	// example:
	//
	// 178157466******
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The attempt ID of the Spark application.
	//
	// example:
	//
	// s202301061000h****
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	// The error message.
	//
	// example:
	//
	// session time out
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
	// IDLE
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The URL of the web UI for the Spark application.
	//
	// example:
	//
	// https://adbsparkui-cn-hangzhou.aliyuncs.com/?token=****
	WebUiAddress *string `json:"WebUiAddress,omitempty" xml:"WebUiAddress,omitempty"`
}

func (s StartSparkReplSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartSparkReplSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartSparkReplSessionResponseBodyData) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *StartSparkReplSessionResponseBodyData) GetAttemptId() *string {
	return s.AttemptId
}

func (s *StartSparkReplSessionResponseBodyData) GetError() *string {
	return s.Error
}

func (s *StartSparkReplSessionResponseBodyData) GetSessionId() *int64 {
	return s.SessionId
}

func (s *StartSparkReplSessionResponseBodyData) GetState() *string {
	return s.State
}

func (s *StartSparkReplSessionResponseBodyData) GetWebUiAddress() *string {
	return s.WebUiAddress
}

func (s *StartSparkReplSessionResponseBodyData) SetAliyunUid(v string) *StartSparkReplSessionResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *StartSparkReplSessionResponseBodyData) SetAttemptId(v string) *StartSparkReplSessionResponseBodyData {
	s.AttemptId = &v
	return s
}

func (s *StartSparkReplSessionResponseBodyData) SetError(v string) *StartSparkReplSessionResponseBodyData {
	s.Error = &v
	return s
}

func (s *StartSparkReplSessionResponseBodyData) SetSessionId(v int64) *StartSparkReplSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *StartSparkReplSessionResponseBodyData) SetState(v string) *StartSparkReplSessionResponseBodyData {
	s.State = &v
	return s
}

func (s *StartSparkReplSessionResponseBodyData) SetWebUiAddress(v string) *StartSparkReplSessionResponseBodyData {
	s.WebUiAddress = &v
	return s
}

func (s *StartSparkReplSessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
