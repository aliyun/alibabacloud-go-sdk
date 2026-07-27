// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDevProdProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDevProdProjectResponseBody
	GetCode() *string
	SetCreateResult(v *CreateDevProdProjectResponseBodyCreateResult) *CreateDevProdProjectResponseBody
	GetCreateResult() *CreateDevProdProjectResponseBodyCreateResult
	SetHttpStatusCode(v int32) *CreateDevProdProjectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDevProdProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDevProdProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDevProdProjectResponseBody
	GetSuccess() *bool
}

type CreateDevProdProjectResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The creation result.
	CreateResult *CreateDevProdProjectResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
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

func (s CreateDevProdProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDevProdProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevProdProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDevProdProjectResponseBody) GetCreateResult() *CreateDevProdProjectResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreateDevProdProjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDevProdProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDevProdProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDevProdProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDevProdProjectResponseBody) SetCode(v string) *CreateDevProdProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDevProdProjectResponseBody) SetCreateResult(v *CreateDevProdProjectResponseBodyCreateResult) *CreateDevProdProjectResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateDevProdProjectResponseBody) SetHttpStatusCode(v int32) *CreateDevProdProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDevProdProjectResponseBody) SetMessage(v string) *CreateDevProdProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDevProdProjectResponseBody) SetRequestId(v string) *CreateDevProdProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDevProdProjectResponseBody) SetSuccess(v bool) *CreateDevProdProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDevProdProjectResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDevProdProjectResponseBodyCreateResult struct {
	// The project ID.
	//
	// example:
	//
	// 102111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateDevProdProjectResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreateDevProdProjectResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateDevProdProjectResponseBodyCreateResult) GetId() *int64 {
	return s.Id
}

func (s *CreateDevProdProjectResponseBodyCreateResult) SetId(v int64) *CreateDevProdProjectResponseBodyCreateResult {
	s.Id = &v
	return s
}

func (s *CreateDevProdProjectResponseBodyCreateResult) Validate() error {
	return dara.Validate(s)
}
