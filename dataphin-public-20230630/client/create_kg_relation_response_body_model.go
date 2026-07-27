// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKgRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKgRelationResponseBody
	GetCode() *string
	SetCreateResult(v *CreateKgRelationResponseBodyCreateResult) *CreateKgRelationResponseBody
	GetCreateResult() *CreateKgRelationResponseBodyCreateResult
	SetHttpStatusCode(v int32) *CreateKgRelationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateKgRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateKgRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateKgRelationResponseBody
	GetSuccess() *bool
}

type CreateKgRelationResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of creating the relationship record.
	CreateResult *CreateKgRelationResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
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

func (s CreateKgRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKgRelationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKgRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKgRelationResponseBody) GetCreateResult() *CreateKgRelationResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreateKgRelationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateKgRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKgRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKgRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateKgRelationResponseBody) SetCode(v string) *CreateKgRelationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKgRelationResponseBody) SetCreateResult(v *CreateKgRelationResponseBodyCreateResult) *CreateKgRelationResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateKgRelationResponseBody) SetHttpStatusCode(v int32) *CreateKgRelationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateKgRelationResponseBody) SetMessage(v string) *CreateKgRelationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKgRelationResponseBody) SetRequestId(v string) *CreateKgRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKgRelationResponseBody) SetSuccess(v bool) *CreateKgRelationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateKgRelationResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKgRelationResponseBodyCreateResult struct {
	// The relationship record ID.
	//
	// example:
	//
	// abc-xxx
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
}

func (s CreateKgRelationResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreateKgRelationResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateKgRelationResponseBodyCreateResult) GetRelationId() *string {
	return s.RelationId
}

func (s *CreateKgRelationResponseBodyCreateResult) SetRelationId(v string) *CreateKgRelationResponseBodyCreateResult {
	s.RelationId = &v
	return s
}

func (s *CreateKgRelationResponseBodyCreateResult) Validate() error {
	return dara.Validate(s)
}
