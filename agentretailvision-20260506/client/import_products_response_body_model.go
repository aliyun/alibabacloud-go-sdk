// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportProductsResponseBody
	GetCode() *string
	SetData(v *ImportProductsResponseBodyData) *ImportProductsResponseBody
	GetData() *ImportProductsResponseBodyData
	SetMessage(v string) *ImportProductsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportProductsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportProductsResponseBody
	GetSuccess() *bool
}

type ImportProductsResponseBody struct {
	// example:
	//
	// 200
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ImportProductsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ImportProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ImportProductsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportProductsResponseBody) GetData() *ImportProductsResponseBodyData {
	return s.Data
}

func (s *ImportProductsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportProductsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportProductsResponseBody) SetCode(v string) *ImportProductsResponseBody {
	s.Code = &v
	return s
}

func (s *ImportProductsResponseBody) SetData(v *ImportProductsResponseBodyData) *ImportProductsResponseBody {
	s.Data = v
	return s
}

func (s *ImportProductsResponseBody) SetMessage(v string) *ImportProductsResponseBody {
	s.Message = &v
	return s
}

func (s *ImportProductsResponseBody) SetRequestId(v string) *ImportProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportProductsResponseBody) SetSuccess(v bool) *ImportProductsResponseBody {
	s.Success = &v
	return s
}

func (s *ImportProductsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportProductsResponseBodyData struct {
	// example:
	//
	// ITEM_001
	ItemUniqueId *string `json:"ItemUniqueId,omitempty" xml:"ItemUniqueId,omitempty"`
	// example:
	//
	// PLAT_001
	PlatformItemId *string `json:"PlatformItemId,omitempty" xml:"PlatformItemId,omitempty"`
}

func (s ImportProductsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportProductsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportProductsResponseBodyData) GetItemUniqueId() *string {
	return s.ItemUniqueId
}

func (s *ImportProductsResponseBodyData) GetPlatformItemId() *string {
	return s.PlatformItemId
}

func (s *ImportProductsResponseBodyData) SetItemUniqueId(v string) *ImportProductsResponseBodyData {
	s.ItemUniqueId = &v
	return s
}

func (s *ImportProductsResponseBodyData) SetPlatformItemId(v string) *ImportProductsResponseBodyData {
	s.PlatformItemId = &v
	return s
}

func (s *ImportProductsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
