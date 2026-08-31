// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateKgEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchCreateKgEntityResponseBody
	GetCode() *string
	SetCreateResult(v *BatchCreateKgEntityResponseBodyCreateResult) *BatchCreateKgEntityResponseBody
	GetCreateResult() *BatchCreateKgEntityResponseBodyCreateResult
	SetHttpStatusCode(v int32) *BatchCreateKgEntityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchCreateKgEntityResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchCreateKgEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchCreateKgEntityResponseBody
	GetSuccess() *bool
}

type BatchCreateKgEntityResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of creating entity records in batches.
	CreateResult *BatchCreateKgEntityResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
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

func (s BatchCreateKgEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgEntityResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateKgEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchCreateKgEntityResponseBody) GetCreateResult() *BatchCreateKgEntityResponseBodyCreateResult {
	return s.CreateResult
}

func (s *BatchCreateKgEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchCreateKgEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchCreateKgEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateKgEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchCreateKgEntityResponseBody) SetCode(v string) *BatchCreateKgEntityResponseBody {
	s.Code = &v
	return s
}

func (s *BatchCreateKgEntityResponseBody) SetCreateResult(v *BatchCreateKgEntityResponseBodyCreateResult) *BatchCreateKgEntityResponseBody {
	s.CreateResult = v
	return s
}

func (s *BatchCreateKgEntityResponseBody) SetHttpStatusCode(v int32) *BatchCreateKgEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchCreateKgEntityResponseBody) SetMessage(v string) *BatchCreateKgEntityResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCreateKgEntityResponseBody) SetRequestId(v string) *BatchCreateKgEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateKgEntityResponseBody) SetSuccess(v bool) *BatchCreateKgEntityResponseBody {
	s.Success = &v
	return s
}

func (s *BatchCreateKgEntityResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCreateKgEntityResponseBodyCreateResult struct {
	// The number of entity records that failed to be created.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The number of successfully created entity records.
	//
	// example:
	//
	// 10
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The list of IDs of successfully created entity records.
	SuccessEntityList []*BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList `json:"SuccessEntityList,omitempty" xml:"SuccessEntityList,omitempty" type:"Repeated"`
}

func (s BatchCreateKgEntityResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgEntityResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *BatchCreateKgEntityResponseBodyCreateResult) GetFailCount() *int32 {
	return s.FailCount
}

func (s *BatchCreateKgEntityResponseBodyCreateResult) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *BatchCreateKgEntityResponseBodyCreateResult) GetSuccessEntityList() []*BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList {
	return s.SuccessEntityList
}

func (s *BatchCreateKgEntityResponseBodyCreateResult) SetFailCount(v int32) *BatchCreateKgEntityResponseBodyCreateResult {
	s.FailCount = &v
	return s
}

func (s *BatchCreateKgEntityResponseBodyCreateResult) SetSuccessCount(v int32) *BatchCreateKgEntityResponseBodyCreateResult {
	s.SuccessCount = &v
	return s
}

func (s *BatchCreateKgEntityResponseBodyCreateResult) SetSuccessEntityList(v []*BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList) *BatchCreateKgEntityResponseBodyCreateResult {
	s.SuccessEntityList = v
	return s
}

func (s *BatchCreateKgEntityResponseBodyCreateResult) Validate() error {
	if s.SuccessEntityList != nil {
		for _, item := range s.SuccessEntityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList struct {
	// The entity record ID.
	//
	// example:
	//
	// abc-xxx
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity type code.
	//
	// example:
	//
	// Company
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList) GoString() string {
	return s.String()
}

func (s *BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList) GetEntityId() *string {
	return s.EntityId
}

func (s *BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList) GetEntityType() *string {
	return s.EntityType
}

func (s *BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList) SetEntityId(v string) *BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList {
	s.EntityId = &v
	return s
}

func (s *BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList) SetEntityType(v string) *BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList {
	s.EntityType = &v
	return s
}

func (s *BatchCreateKgEntityResponseBodyCreateResultSuccessEntityList) Validate() error {
	return dara.Validate(s)
}
