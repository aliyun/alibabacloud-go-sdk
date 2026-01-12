// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClothTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListClothTypesResponseBody
	GetCode() *string
	SetData(v []*ListClothTypesResponseBodyData) *ListClothTypesResponseBody
	GetData() []*ListClothTypesResponseBodyData
	SetErrorName(v string) *ListClothTypesResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *ListClothTypesResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *ListClothTypesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListClothTypesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListClothTypesResponseBody
	GetSuccess() *bool
}

type ListClothTypesResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListClothTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorName *string                           `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                            `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListClothTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClothTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClothTypesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListClothTypesResponseBody) GetData() []*ListClothTypesResponseBodyData {
	return s.Data
}

func (s *ListClothTypesResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *ListClothTypesResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListClothTypesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClothTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClothTypesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListClothTypesResponseBody) SetCode(v string) *ListClothTypesResponseBody {
	s.Code = &v
	return s
}

func (s *ListClothTypesResponseBody) SetData(v []*ListClothTypesResponseBodyData) *ListClothTypesResponseBody {
	s.Data = v
	return s
}

func (s *ListClothTypesResponseBody) SetErrorName(v string) *ListClothTypesResponseBody {
	s.ErrorName = &v
	return s
}

func (s *ListClothTypesResponseBody) SetHttpCode(v int32) *ListClothTypesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListClothTypesResponseBody) SetMessage(v string) *ListClothTypesResponseBody {
	s.Message = &v
	return s
}

func (s *ListClothTypesResponseBody) SetRequestId(v string) *ListClothTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClothTypesResponseBody) SetSuccess(v bool) *ListClothTypesResponseBody {
	s.Success = &v
	return s
}

func (s *ListClothTypesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClothTypesResponseBodyData struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListClothTypesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClothTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClothTypesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListClothTypesResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListClothTypesResponseBodyData) SetName(v string) *ListClothTypesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListClothTypesResponseBodyData) SetType(v string) *ListClothTypesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListClothTypesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
