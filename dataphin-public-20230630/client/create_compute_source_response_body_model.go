// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateComputeSourceResponseBody
	GetCode() *string
	SetCreateResult(v *CreateComputeSourceResponseBodyCreateResult) *CreateComputeSourceResponseBody
	GetCreateResult() *CreateComputeSourceResponseBodyCreateResult
	SetHttpStatusCode(v int32) *CreateComputeSourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateComputeSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateComputeSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateComputeSourceResponseBody
	GetSuccess() *bool
}

type CreateComputeSourceResponseBody struct {
	// example:
	//
	// OK
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateResult *CreateComputeSourceResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateComputeSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateComputeSourceResponseBody) GetCreateResult() *CreateComputeSourceResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreateComputeSourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateComputeSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateComputeSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComputeSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateComputeSourceResponseBody) SetCode(v string) *CreateComputeSourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComputeSourceResponseBody) SetCreateResult(v *CreateComputeSourceResponseBodyCreateResult) *CreateComputeSourceResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateComputeSourceResponseBody) SetHttpStatusCode(v int32) *CreateComputeSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateComputeSourceResponseBody) SetMessage(v string) *CreateComputeSourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateComputeSourceResponseBody) SetRequestId(v string) *CreateComputeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComputeSourceResponseBody) SetSuccess(v bool) *CreateComputeSourceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateComputeSourceResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateComputeSourceResponseBodyCreateResult struct {
	// example:
	//
	// 12356
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateComputeSourceResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeSourceResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateComputeSourceResponseBodyCreateResult) GetId() *int64 {
	return s.Id
}

func (s *CreateComputeSourceResponseBodyCreateResult) SetId(v int64) *CreateComputeSourceResponseBodyCreateResult {
	s.Id = &v
	return s
}

func (s *CreateComputeSourceResponseBodyCreateResult) Validate() error {
	return dara.Validate(s)
}
