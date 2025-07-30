// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountJobByConditionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *CountJobByConditionResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CountJobByConditionResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *CountJobByConditionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CountJobByConditionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CountJobByConditionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CountJobByConditionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CountJobByConditionResponseBody
	GetSuccess() *bool
	SetTotalRecordCount(v int64) *CountJobByConditionResponseBody
	GetTotalRecordCount() *int64
}

type CountJobByConditionResponseBody struct {
	// The internal error code. This parameter will be removed soon.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the  **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// > If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// present environment is not support,so skip.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FC3BAAF2-74E3-4471-8EB5-96202D6A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of data synchronization instances that meet the requirements and belong to your Alibaba Cloud account.
	//
	// example:
	//
	// 100
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s CountJobByConditionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CountJobByConditionResponseBody) GoString() string {
	return s.String()
}

func (s *CountJobByConditionResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CountJobByConditionResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CountJobByConditionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CountJobByConditionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CountJobByConditionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CountJobByConditionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CountJobByConditionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CountJobByConditionResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *CountJobByConditionResponseBody) SetDynamicCode(v string) *CountJobByConditionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CountJobByConditionResponseBody) SetDynamicMessage(v string) *CountJobByConditionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CountJobByConditionResponseBody) SetErrCode(v string) *CountJobByConditionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CountJobByConditionResponseBody) SetErrMessage(v string) *CountJobByConditionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CountJobByConditionResponseBody) SetHttpStatusCode(v int32) *CountJobByConditionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CountJobByConditionResponseBody) SetRequestId(v string) *CountJobByConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountJobByConditionResponseBody) SetSuccess(v bool) *CountJobByConditionResponseBody {
	s.Success = &v
	return s
}

func (s *CountJobByConditionResponseBody) SetTotalRecordCount(v int64) *CountJobByConditionResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *CountJobByConditionResponseBody) Validate() error {
	return dara.Validate(s)
}
