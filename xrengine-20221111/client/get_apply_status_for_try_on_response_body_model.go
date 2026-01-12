// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplyStatusForTryOnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetApplyStatusForTryOnResponseBody
	GetCode() *string
	SetData(v *GetApplyStatusForTryOnResponseBodyData) *GetApplyStatusForTryOnResponseBody
	GetData() *GetApplyStatusForTryOnResponseBodyData
	SetErrorName(v string) *GetApplyStatusForTryOnResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *GetApplyStatusForTryOnResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *GetApplyStatusForTryOnResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApplyStatusForTryOnResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetApplyStatusForTryOnResponseBody
	GetSuccess() *bool
}

type GetApplyStatusForTryOnResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetApplyStatusForTryOnResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                                 `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                                  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetApplyStatusForTryOnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplyStatusForTryOnResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplyStatusForTryOnResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApplyStatusForTryOnResponseBody) GetData() *GetApplyStatusForTryOnResponseBodyData {
	return s.Data
}

func (s *GetApplyStatusForTryOnResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *GetApplyStatusForTryOnResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetApplyStatusForTryOnResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApplyStatusForTryOnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplyStatusForTryOnResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetApplyStatusForTryOnResponseBody) SetCode(v string) *GetApplyStatusForTryOnResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplyStatusForTryOnResponseBody) SetData(v *GetApplyStatusForTryOnResponseBodyData) *GetApplyStatusForTryOnResponseBody {
	s.Data = v
	return s
}

func (s *GetApplyStatusForTryOnResponseBody) SetErrorName(v string) *GetApplyStatusForTryOnResponseBody {
	s.ErrorName = &v
	return s
}

func (s *GetApplyStatusForTryOnResponseBody) SetHttpCode(v int32) *GetApplyStatusForTryOnResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetApplyStatusForTryOnResponseBody) SetMessage(v string) *GetApplyStatusForTryOnResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplyStatusForTryOnResponseBody) SetRequestId(v string) *GetApplyStatusForTryOnResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplyStatusForTryOnResponseBody) SetSuccess(v bool) *GetApplyStatusForTryOnResponseBody {
	s.Success = &v
	return s
}

func (s *GetApplyStatusForTryOnResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplyStatusForTryOnResponseBodyData struct {
	Agreement *bool   `json:"Agreement,omitempty" xml:"Agreement,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetApplyStatusForTryOnResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApplyStatusForTryOnResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplyStatusForTryOnResponseBodyData) GetAgreement() *bool {
	return s.Agreement
}

func (s *GetApplyStatusForTryOnResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetApplyStatusForTryOnResponseBodyData) SetAgreement(v bool) *GetApplyStatusForTryOnResponseBodyData {
	s.Agreement = &v
	return s
}

func (s *GetApplyStatusForTryOnResponseBodyData) SetStatus(v string) *GetApplyStatusForTryOnResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetApplyStatusForTryOnResponseBodyData) Validate() error {
	return dara.Validate(s)
}
