// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateKgRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchCreateKgRelationResponseBody
	GetCode() *string
	SetCreateResult(v *BatchCreateKgRelationResponseBodyCreateResult) *BatchCreateKgRelationResponseBody
	GetCreateResult() *BatchCreateKgRelationResponseBodyCreateResult
	SetHttpStatusCode(v int32) *BatchCreateKgRelationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchCreateKgRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchCreateKgRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchCreateKgRelationResponseBody
	GetSuccess() *bool
}

type BatchCreateKgRelationResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of batch relationship record creation.
	CreateResult *BatchCreateKgRelationResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchCreateKgRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgRelationResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateKgRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchCreateKgRelationResponseBody) GetCreateResult() *BatchCreateKgRelationResponseBodyCreateResult {
	return s.CreateResult
}

func (s *BatchCreateKgRelationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchCreateKgRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchCreateKgRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateKgRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchCreateKgRelationResponseBody) SetCode(v string) *BatchCreateKgRelationResponseBody {
	s.Code = &v
	return s
}

func (s *BatchCreateKgRelationResponseBody) SetCreateResult(v *BatchCreateKgRelationResponseBodyCreateResult) *BatchCreateKgRelationResponseBody {
	s.CreateResult = v
	return s
}

func (s *BatchCreateKgRelationResponseBody) SetHttpStatusCode(v int32) *BatchCreateKgRelationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchCreateKgRelationResponseBody) SetMessage(v string) *BatchCreateKgRelationResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCreateKgRelationResponseBody) SetRequestId(v string) *BatchCreateKgRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateKgRelationResponseBody) SetSuccess(v bool) *BatchCreateKgRelationResponseBody {
	s.Success = &v
	return s
}

func (s *BatchCreateKgRelationResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCreateKgRelationResponseBodyCreateResult struct {
	// The number of failed records.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The number of successfully created records.
	//
	// example:
	//
	// 10
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The list of successfully created entity records.
	SuccessRelationList []*BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList `json:"SuccessRelationList,omitempty" xml:"SuccessRelationList,omitempty" type:"Repeated"`
}

func (s BatchCreateKgRelationResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgRelationResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *BatchCreateKgRelationResponseBodyCreateResult) GetFailCount() *int32 {
	return s.FailCount
}

func (s *BatchCreateKgRelationResponseBodyCreateResult) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *BatchCreateKgRelationResponseBodyCreateResult) GetSuccessRelationList() []*BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList {
	return s.SuccessRelationList
}

func (s *BatchCreateKgRelationResponseBodyCreateResult) SetFailCount(v int32) *BatchCreateKgRelationResponseBodyCreateResult {
	s.FailCount = &v
	return s
}

func (s *BatchCreateKgRelationResponseBodyCreateResult) SetSuccessCount(v int32) *BatchCreateKgRelationResponseBodyCreateResult {
	s.SuccessCount = &v
	return s
}

func (s *BatchCreateKgRelationResponseBodyCreateResult) SetSuccessRelationList(v []*BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList) *BatchCreateKgRelationResponseBodyCreateResult {
	s.SuccessRelationList = v
	return s
}

func (s *BatchCreateKgRelationResponseBodyCreateResult) Validate() error {
	if s.SuccessRelationList != nil {
		for _, item := range s.SuccessRelationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList struct {
	// The relationship record ID.
	//
	// example:
	//
	// abc-xxx
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The relationship type code.
	//
	// example:
	//
	// BELONG_TO
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
}

func (s BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList) GoString() string {
	return s.String()
}

func (s *BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList) GetRelationId() *string {
	return s.RelationId
}

func (s *BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList) GetRelationType() *string {
	return s.RelationType
}

func (s *BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList) SetRelationId(v string) *BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList {
	s.RelationId = &v
	return s
}

func (s *BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList) SetRelationType(v string) *BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList {
	s.RelationType = &v
	return s
}

func (s *BatchCreateKgRelationResponseBodyCreateResultSuccessRelationList) Validate() error {
	return dara.Validate(s)
}
