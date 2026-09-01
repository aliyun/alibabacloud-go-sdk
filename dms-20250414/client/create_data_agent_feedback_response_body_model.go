// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentFeedbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDataAgentFeedbackResponseBody
	GetCode() *string
	SetData(v *CreateDataAgentFeedbackResponseBodyData) *CreateDataAgentFeedbackResponseBody
	GetData() *CreateDataAgentFeedbackResponseBodyData
	SetErrorCode(v string) *CreateDataAgentFeedbackResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *CreateDataAgentFeedbackResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDataAgentFeedbackResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDataAgentFeedbackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataAgentFeedbackResponseBody
	GetSuccess() *bool
}

type CreateDataAgentFeedbackResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateDataAgentFeedbackResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Specified parameter Tid is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 89a07eac-96ff-48be-983b-f22c55*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataAgentFeedbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataAgentFeedbackResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDataAgentFeedbackResponseBody) GetData() *CreateDataAgentFeedbackResponseBodyData {
	return s.Data
}

func (s *CreateDataAgentFeedbackResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataAgentFeedbackResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDataAgentFeedbackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDataAgentFeedbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataAgentFeedbackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataAgentFeedbackResponseBody) SetCode(v string) *CreateDataAgentFeedbackResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBody) SetData(v *CreateDataAgentFeedbackResponseBodyData) *CreateDataAgentFeedbackResponseBody {
	s.Data = v
	return s
}

func (s *CreateDataAgentFeedbackResponseBody) SetErrorCode(v string) *CreateDataAgentFeedbackResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBody) SetHttpStatusCode(v int32) *CreateDataAgentFeedbackResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBody) SetMessage(v string) *CreateDataAgentFeedbackResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBody) SetRequestId(v string) *CreateDataAgentFeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBody) SetSuccess(v bool) *CreateDataAgentFeedbackResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataAgentFeedbackResponseBodyData struct {
	// example:
	//
	// {"feedback_type":"PRODUCT_SUGGESTION","user_feedback": "test","email":"yourname@example.com","is_authorized":"Y"}
	FeedbackContent *string `json:"FeedbackContent,omitempty" xml:"FeedbackContent,omitempty"`
	// example:
	//
	// ISSUE_REPORT
	FeedbackType *string `json:"FeedbackType,omitempty" xml:"FeedbackType,omitempty"`
	// example:
	//
	// 1
	LikeValue *int32 `json:"LikeValue,omitempty" xml:"LikeValue,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// h8r********4fch
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// h8r********4fch_sdesfews
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// example:
	//
	// SESSION
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateDataAgentFeedbackResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentFeedbackResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDataAgentFeedbackResponseBodyData) GetFeedbackContent() *string {
	return s.FeedbackContent
}

func (s *CreateDataAgentFeedbackResponseBodyData) GetFeedbackType() *string {
	return s.FeedbackType
}

func (s *CreateDataAgentFeedbackResponseBodyData) GetLikeValue() *int32 {
	return s.LikeValue
}

func (s *CreateDataAgentFeedbackResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDataAgentFeedbackResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateDataAgentFeedbackResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateDataAgentFeedbackResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateDataAgentFeedbackResponseBodyData) SetFeedbackContent(v string) *CreateDataAgentFeedbackResponseBodyData {
	s.FeedbackContent = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBodyData) SetFeedbackType(v string) *CreateDataAgentFeedbackResponseBodyData {
	s.FeedbackType = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBodyData) SetLikeValue(v int32) *CreateDataAgentFeedbackResponseBodyData {
	s.LikeValue = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBodyData) SetRegionId(v string) *CreateDataAgentFeedbackResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBodyData) SetSessionId(v string) *CreateDataAgentFeedbackResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBodyData) SetTargetId(v string) *CreateDataAgentFeedbackResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBodyData) SetTargetType(v string) *CreateDataAgentFeedbackResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *CreateDataAgentFeedbackResponseBodyData) Validate() error {
	return dara.Validate(s)
}
