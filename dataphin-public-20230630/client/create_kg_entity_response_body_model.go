// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKgEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKgEntityResponseBody
	GetCode() *string
	SetCreateResult(v *CreateKgEntityResponseBodyCreateResult) *CreateKgEntityResponseBody
	GetCreateResult() *CreateKgEntityResponseBodyCreateResult
	SetHttpStatusCode(v int32) *CreateKgEntityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateKgEntityResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateKgEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateKgEntityResponseBody
	GetSuccess() *bool
}

type CreateKgEntityResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The entity record creation result.
	CreateResult *CreateKgEntityResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
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

func (s CreateKgEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKgEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKgEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKgEntityResponseBody) GetCreateResult() *CreateKgEntityResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreateKgEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateKgEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKgEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKgEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateKgEntityResponseBody) SetCode(v string) *CreateKgEntityResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKgEntityResponseBody) SetCreateResult(v *CreateKgEntityResponseBodyCreateResult) *CreateKgEntityResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateKgEntityResponseBody) SetHttpStatusCode(v int32) *CreateKgEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateKgEntityResponseBody) SetMessage(v string) *CreateKgEntityResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKgEntityResponseBody) SetRequestId(v string) *CreateKgEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKgEntityResponseBody) SetSuccess(v bool) *CreateKgEntityResponseBody {
	s.Success = &v
	return s
}

func (s *CreateKgEntityResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKgEntityResponseBodyCreateResult struct {
	// The entity record ID.
	//
	// example:
	//
	// abc-xxx
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s CreateKgEntityResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreateKgEntityResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateKgEntityResponseBodyCreateResult) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateKgEntityResponseBodyCreateResult) SetEntityId(v string) *CreateKgEntityResponseBodyCreateResult {
	s.EntityId = &v
	return s
}

func (s *CreateKgEntityResponseBodyCreateResult) Validate() error {
	return dara.Validate(s)
}
