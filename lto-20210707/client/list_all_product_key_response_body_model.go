// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllProductKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllProductKeyResponseBody
	GetCode() *string
	SetData(v []*ListAllProductKeyResponseBodyData) *ListAllProductKeyResponseBody
	GetData() []*ListAllProductKeyResponseBodyData
	SetHttpStatusCode(v int32) *ListAllProductKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAllProductKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllProductKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllProductKeyResponseBody
	GetSuccess() *bool
}

type ListAllProductKeyResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListAllProductKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllProductKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllProductKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllProductKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllProductKeyResponseBody) GetData() []*ListAllProductKeyResponseBodyData {
	return s.Data
}

func (s *ListAllProductKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAllProductKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllProductKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllProductKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllProductKeyResponseBody) SetCode(v string) *ListAllProductKeyResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllProductKeyResponseBody) SetData(v []*ListAllProductKeyResponseBodyData) *ListAllProductKeyResponseBody {
	s.Data = v
	return s
}

func (s *ListAllProductKeyResponseBody) SetHttpStatusCode(v int32) *ListAllProductKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAllProductKeyResponseBody) SetMessage(v string) *ListAllProductKeyResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllProductKeyResponseBody) SetRequestId(v string) *ListAllProductKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllProductKeyResponseBody) SetSuccess(v bool) *ListAllProductKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllProductKeyResponseBody) Validate() error {
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

type ListAllProductKeyResponseBodyData struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s ListAllProductKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllProductKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllProductKeyResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAllProductKeyResponseBodyData) GetProductKey() *string {
	return s.ProductKey
}

func (s *ListAllProductKeyResponseBodyData) SetName(v string) *ListAllProductKeyResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAllProductKeyResponseBodyData) SetProductKey(v string) *ListAllProductKeyResponseBodyData {
	s.ProductKey = &v
	return s
}

func (s *ListAllProductKeyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
