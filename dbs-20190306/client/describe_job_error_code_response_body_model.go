// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobErrorCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeJobErrorCodeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeJobErrorCodeResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeJobErrorCodeResponseBody
	GetHttpStatusCode() *int32
	SetItem(v *DescribeJobErrorCodeResponseBodyItem) *DescribeJobErrorCodeResponseBody
	GetItem() *DescribeJobErrorCodeResponseBodyItem
	SetRequestId(v string) *DescribeJobErrorCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeJobErrorCodeResponseBody
	GetSuccess() *bool
}

type DescribeJobErrorCodeResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error information.
	Item *DescribeJobErrorCodeResponseBodyItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1FC2F86D-AFF4-4ED9-BB25-ADDE196CB2B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeJobErrorCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobErrorCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobErrorCodeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeJobErrorCodeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeJobErrorCodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeJobErrorCodeResponseBody) GetItem() *DescribeJobErrorCodeResponseBodyItem {
	return s.Item
}

func (s *DescribeJobErrorCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobErrorCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeJobErrorCodeResponseBody) SetErrCode(v string) *DescribeJobErrorCodeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBody) SetErrMessage(v string) *DescribeJobErrorCodeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBody) SetHttpStatusCode(v int32) *DescribeJobErrorCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBody) SetItem(v *DescribeJobErrorCodeResponseBodyItem) *DescribeJobErrorCodeResponseBody {
	s.Item = v
	return s
}

func (s *DescribeJobErrorCodeResponseBody) SetRequestId(v string) *DescribeJobErrorCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBody) SetSuccess(v bool) *DescribeJobErrorCodeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBody) Validate() error {
	if s.Item != nil {
		if err := s.Item.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeJobErrorCodeResponseBodyItem struct {
	// The error code.
	//
	// example:
	//
	// failed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The standard error message.
	//
	// example:
	//
	// Describe preCheck progress failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the full backup or restore task.
	//
	// example:
	//
	// tooi0****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// init
	JobState *string `json:"JobState,omitempty" xml:"JobState,omitempty"`
	// The internal ID of the DBS task type.
	//
	// example:
	//
	// testId
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The language of the error message.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s DescribeJobErrorCodeResponseBodyItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobErrorCodeResponseBodyItem) GoString() string {
	return s.String()
}

func (s *DescribeJobErrorCodeResponseBodyItem) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeJobErrorCodeResponseBodyItem) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeJobErrorCodeResponseBodyItem) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobErrorCodeResponseBodyItem) GetJobState() *string {
	return s.JobState
}

func (s *DescribeJobErrorCodeResponseBodyItem) GetJobType() *string {
	return s.JobType
}

func (s *DescribeJobErrorCodeResponseBodyItem) GetLanguage() *string {
	return s.Language
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetErrorCode(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.ErrorCode = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetErrorMessage(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetJobId(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.JobId = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetJobState(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.JobState = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetJobType(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.JobType = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBodyItem) SetLanguage(v string) *DescribeJobErrorCodeResponseBodyItem {
	s.Language = &v
	return s
}

func (s *DescribeJobErrorCodeResponseBodyItem) Validate() error {
	return dara.Validate(s)
}
