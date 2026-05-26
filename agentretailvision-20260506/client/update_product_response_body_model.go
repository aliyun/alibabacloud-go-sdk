// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateProductResponseBody
	GetCode() *string
	SetData(v *UpdateProductResponseBodyData) *UpdateProductResponseBody
	GetData() *UpdateProductResponseBodyData
	SetMessage(v string) *UpdateProductResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateProductResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateProductResponseBody
	GetSuccess() *bool
}

type UpdateProductResponseBody struct {
	// example:
	//
	// 200
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateProductResponseBody) GetData() *UpdateProductResponseBodyData {
	return s.Data
}

func (s *UpdateProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProductResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateProductResponseBody) SetCode(v string) *UpdateProductResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateProductResponseBody) SetData(v *UpdateProductResponseBodyData) *UpdateProductResponseBody {
	s.Data = v
	return s
}

func (s *UpdateProductResponseBody) SetMessage(v string) *UpdateProductResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateProductResponseBody) SetRequestId(v string) *UpdateProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProductResponseBody) SetSuccess(v bool) *UpdateProductResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProductResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProductResponseBodyData struct {
	// example:
	//
	// ITEM_001
	ItemUniqueId *string `json:"ItemUniqueId,omitempty" xml:"ItemUniqueId,omitempty"`
	// example:
	//
	// PLAT_001
	PlatformItemId *string `json:"PlatformItemId,omitempty" xml:"PlatformItemId,omitempty"`
}

func (s UpdateProductResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateProductResponseBodyData) GetItemUniqueId() *string {
	return s.ItemUniqueId
}

func (s *UpdateProductResponseBodyData) GetPlatformItemId() *string {
	return s.PlatformItemId
}

func (s *UpdateProductResponseBodyData) SetItemUniqueId(v string) *UpdateProductResponseBodyData {
	s.ItemUniqueId = &v
	return s
}

func (s *UpdateProductResponseBodyData) SetPlatformItemId(v string) *UpdateProductResponseBodyData {
	s.PlatformItemId = &v
	return s
}

func (s *UpdateProductResponseBodyData) Validate() error {
	return dara.Validate(s)
}
