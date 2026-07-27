// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBasicProjectResponseBody
	GetCode() *string
	SetCreateResult(v *CreateBasicProjectResponseBodyCreateResult) *CreateBasicProjectResponseBody
	GetCreateResult() *CreateBasicProjectResponseBodyCreateResult
	SetHttpStatusCode(v int32) *CreateBasicProjectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBasicProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBasicProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBasicProjectResponseBody
	GetSuccess() *bool
}

type CreateBasicProjectResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The creation result.
	CreateResult *CreateBasicProjectResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
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

func (s CreateBasicProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBasicProjectResponseBody) GetCreateResult() *CreateBasicProjectResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreateBasicProjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBasicProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBasicProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBasicProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBasicProjectResponseBody) SetCode(v string) *CreateBasicProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBasicProjectResponseBody) SetCreateResult(v *CreateBasicProjectResponseBodyCreateResult) *CreateBasicProjectResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateBasicProjectResponseBody) SetHttpStatusCode(v int32) *CreateBasicProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBasicProjectResponseBody) SetMessage(v string) *CreateBasicProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBasicProjectResponseBody) SetRequestId(v string) *CreateBasicProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBasicProjectResponseBody) SetSuccess(v bool) *CreateBasicProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBasicProjectResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBasicProjectResponseBodyCreateResult struct {
	// The project ID.
	//
	// example:
	//
	// 102111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateBasicProjectResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicProjectResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateBasicProjectResponseBodyCreateResult) GetId() *int64 {
	return s.Id
}

func (s *CreateBasicProjectResponseBodyCreateResult) SetId(v int64) *CreateBasicProjectResponseBodyCreateResult {
	s.Id = &v
	return s
}

func (s *CreateBasicProjectResponseBodyCreateResult) Validate() error {
	return dara.Validate(s)
}
