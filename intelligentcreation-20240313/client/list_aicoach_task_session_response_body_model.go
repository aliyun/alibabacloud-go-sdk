// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICoachTaskSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListAICoachTaskSessionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAICoachTaskSessionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListAICoachTaskSessionResponseBody
	GetRequestId() *string
	SetSessionList(v []*ListAICoachTaskSessionResponseBodySessionList) *ListAICoachTaskSessionResponseBody
	GetSessionList() []*ListAICoachTaskSessionResponseBodySessionList
	SetSuccess(v bool) *ListAICoachTaskSessionResponseBody
	GetSuccess() *bool
}

type ListAICoachTaskSessionResponseBody struct {
	ErrorCode    *string                                          `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                          `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SessionList  []*ListAICoachTaskSessionResponseBodySessionList `json:"sessionList,omitempty" xml:"sessionList,omitempty" type:"Repeated"`
	Success      *bool                                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListAICoachTaskSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachTaskSessionResponseBody) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskSessionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAICoachTaskSessionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAICoachTaskSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAICoachTaskSessionResponseBody) GetSessionList() []*ListAICoachTaskSessionResponseBodySessionList {
	return s.SessionList
}

func (s *ListAICoachTaskSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAICoachTaskSessionResponseBody) SetErrorCode(v string) *ListAICoachTaskSessionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAICoachTaskSessionResponseBody) SetErrorMessage(v string) *ListAICoachTaskSessionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAICoachTaskSessionResponseBody) SetRequestId(v string) *ListAICoachTaskSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAICoachTaskSessionResponseBody) SetSessionList(v []*ListAICoachTaskSessionResponseBodySessionList) *ListAICoachTaskSessionResponseBody {
	s.SessionList = v
	return s
}

func (s *ListAICoachTaskSessionResponseBody) SetSuccess(v bool) *ListAICoachTaskSessionResponseBody {
	s.Success = &v
	return s
}

func (s *ListAICoachTaskSessionResponseBody) Validate() error {
	if s.SessionList != nil {
		for _, item := range s.SessionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAICoachTaskSessionResponseBodySessionList struct {
	SessionCreateTime *string `json:"sessionCreateTime,omitempty" xml:"sessionCreateTime,omitempty"`
	SessionDuration   *int64  `json:"sessionDuration,omitempty" xml:"sessionDuration,omitempty"`
	SessionId         *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	SessionStatus     *int32  `json:"sessionStatus,omitempty" xml:"sessionStatus,omitempty"`
}

func (s ListAICoachTaskSessionResponseBodySessionList) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachTaskSessionResponseBodySessionList) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskSessionResponseBodySessionList) GetSessionCreateTime() *string {
	return s.SessionCreateTime
}

func (s *ListAICoachTaskSessionResponseBodySessionList) GetSessionDuration() *int64 {
	return s.SessionDuration
}

func (s *ListAICoachTaskSessionResponseBodySessionList) GetSessionId() *string {
	return s.SessionId
}

func (s *ListAICoachTaskSessionResponseBodySessionList) GetSessionStatus() *int32 {
	return s.SessionStatus
}

func (s *ListAICoachTaskSessionResponseBodySessionList) SetSessionCreateTime(v string) *ListAICoachTaskSessionResponseBodySessionList {
	s.SessionCreateTime = &v
	return s
}

func (s *ListAICoachTaskSessionResponseBodySessionList) SetSessionDuration(v int64) *ListAICoachTaskSessionResponseBodySessionList {
	s.SessionDuration = &v
	return s
}

func (s *ListAICoachTaskSessionResponseBodySessionList) SetSessionId(v string) *ListAICoachTaskSessionResponseBodySessionList {
	s.SessionId = &v
	return s
}

func (s *ListAICoachTaskSessionResponseBodySessionList) SetSessionStatus(v int32) *ListAICoachTaskSessionResponseBodySessionList {
	s.SessionStatus = &v
	return s
}

func (s *ListAICoachTaskSessionResponseBodySessionList) Validate() error {
	return dara.Validate(s)
}
