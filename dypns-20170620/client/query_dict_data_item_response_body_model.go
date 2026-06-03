// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDictDataItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryDictDataItemResponseBody
	GetCode() *string
	SetData(v []*QueryDictDataItemResponseBodyData) *QueryDictDataItemResponseBody
	GetData() []*QueryDictDataItemResponseBodyData
	SetMessage(v string) *QueryDictDataItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryDictDataItemResponseBody
	GetRequestId() *string
}

type QueryDictDataItemResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryDictDataItemResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDictDataItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDictDataItemResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDictDataItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryDictDataItemResponseBody) GetData() []*QueryDictDataItemResponseBodyData {
	return s.Data
}

func (s *QueryDictDataItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryDictDataItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDictDataItemResponseBody) SetCode(v string) *QueryDictDataItemResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDictDataItemResponseBody) SetData(v []*QueryDictDataItemResponseBodyData) *QueryDictDataItemResponseBody {
	s.Data = v
	return s
}

func (s *QueryDictDataItemResponseBody) SetMessage(v string) *QueryDictDataItemResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDictDataItemResponseBody) SetRequestId(v string) *QueryDictDataItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDictDataItemResponseBody) Validate() error {
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

type QueryDictDataItemResponseBodyData struct {
	ClassifyItemCode  *string                                         `json:"ClassifyItemCode,omitempty" xml:"ClassifyItemCode,omitempty"`
	ClassifyItemValue *string                                         `json:"ClassifyItemValue,omitempty" xml:"ClassifyItemValue,omitempty"`
	DataItemCode      *string                                         `json:"DataItemCode,omitempty" xml:"DataItemCode,omitempty"`
	DataItemOrder     *int32                                          `json:"DataItemOrder,omitempty" xml:"DataItemOrder,omitempty"`
	DataItemParent    *string                                         `json:"DataItemParent,omitempty" xml:"DataItemParent,omitempty"`
	DataItemValue     *string                                         `json:"DataItemValue,omitempty" xml:"DataItemValue,omitempty"`
	ItemType          *int32                                          `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	SubListData       []*QueryDictDataItemResponseBodyDataSubListData `json:"SubListData,omitempty" xml:"SubListData,omitempty" type:"Repeated"`
}

func (s QueryDictDataItemResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDictDataItemResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDictDataItemResponseBodyData) GetClassifyItemCode() *string {
	return s.ClassifyItemCode
}

func (s *QueryDictDataItemResponseBodyData) GetClassifyItemValue() *string {
	return s.ClassifyItemValue
}

func (s *QueryDictDataItemResponseBodyData) GetDataItemCode() *string {
	return s.DataItemCode
}

func (s *QueryDictDataItemResponseBodyData) GetDataItemOrder() *int32 {
	return s.DataItemOrder
}

func (s *QueryDictDataItemResponseBodyData) GetDataItemParent() *string {
	return s.DataItemParent
}

func (s *QueryDictDataItemResponseBodyData) GetDataItemValue() *string {
	return s.DataItemValue
}

func (s *QueryDictDataItemResponseBodyData) GetItemType() *int32 {
	return s.ItemType
}

func (s *QueryDictDataItemResponseBodyData) GetSubListData() []*QueryDictDataItemResponseBodyDataSubListData {
	return s.SubListData
}

func (s *QueryDictDataItemResponseBodyData) SetClassifyItemCode(v string) *QueryDictDataItemResponseBodyData {
	s.ClassifyItemCode = &v
	return s
}

func (s *QueryDictDataItemResponseBodyData) SetClassifyItemValue(v string) *QueryDictDataItemResponseBodyData {
	s.ClassifyItemValue = &v
	return s
}

func (s *QueryDictDataItemResponseBodyData) SetDataItemCode(v string) *QueryDictDataItemResponseBodyData {
	s.DataItemCode = &v
	return s
}

func (s *QueryDictDataItemResponseBodyData) SetDataItemOrder(v int32) *QueryDictDataItemResponseBodyData {
	s.DataItemOrder = &v
	return s
}

func (s *QueryDictDataItemResponseBodyData) SetDataItemParent(v string) *QueryDictDataItemResponseBodyData {
	s.DataItemParent = &v
	return s
}

func (s *QueryDictDataItemResponseBodyData) SetDataItemValue(v string) *QueryDictDataItemResponseBodyData {
	s.DataItemValue = &v
	return s
}

func (s *QueryDictDataItemResponseBodyData) SetItemType(v int32) *QueryDictDataItemResponseBodyData {
	s.ItemType = &v
	return s
}

func (s *QueryDictDataItemResponseBodyData) SetSubListData(v []*QueryDictDataItemResponseBodyDataSubListData) *QueryDictDataItemResponseBodyData {
	s.SubListData = v
	return s
}

func (s *QueryDictDataItemResponseBodyData) Validate() error {
	if s.SubListData != nil {
		for _, item := range s.SubListData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDictDataItemResponseBodyDataSubListData struct {
	ClassifyItemCode  *string `json:"ClassifyItemCode,omitempty" xml:"ClassifyItemCode,omitempty"`
	ClassifyItemValue *string `json:"ClassifyItemValue,omitempty" xml:"ClassifyItemValue,omitempty"`
	DataItemCode      *string `json:"DataItemCode,omitempty" xml:"DataItemCode,omitempty"`
	DataItemOrder     *int32  `json:"DataItemOrder,omitempty" xml:"DataItemOrder,omitempty"`
	DataItemParent    *string `json:"DataItemParent,omitempty" xml:"DataItemParent,omitempty"`
	DataItemValue     *string `json:"DataItemValue,omitempty" xml:"DataItemValue,omitempty"`
	ItemType          *int32  `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
}

func (s QueryDictDataItemResponseBodyDataSubListData) String() string {
	return dara.Prettify(s)
}

func (s QueryDictDataItemResponseBodyDataSubListData) GoString() string {
	return s.String()
}

func (s *QueryDictDataItemResponseBodyDataSubListData) GetClassifyItemCode() *string {
	return s.ClassifyItemCode
}

func (s *QueryDictDataItemResponseBodyDataSubListData) GetClassifyItemValue() *string {
	return s.ClassifyItemValue
}

func (s *QueryDictDataItemResponseBodyDataSubListData) GetDataItemCode() *string {
	return s.DataItemCode
}

func (s *QueryDictDataItemResponseBodyDataSubListData) GetDataItemOrder() *int32 {
	return s.DataItemOrder
}

func (s *QueryDictDataItemResponseBodyDataSubListData) GetDataItemParent() *string {
	return s.DataItemParent
}

func (s *QueryDictDataItemResponseBodyDataSubListData) GetDataItemValue() *string {
	return s.DataItemValue
}

func (s *QueryDictDataItemResponseBodyDataSubListData) GetItemType() *int32 {
	return s.ItemType
}

func (s *QueryDictDataItemResponseBodyDataSubListData) SetClassifyItemCode(v string) *QueryDictDataItemResponseBodyDataSubListData {
	s.ClassifyItemCode = &v
	return s
}

func (s *QueryDictDataItemResponseBodyDataSubListData) SetClassifyItemValue(v string) *QueryDictDataItemResponseBodyDataSubListData {
	s.ClassifyItemValue = &v
	return s
}

func (s *QueryDictDataItemResponseBodyDataSubListData) SetDataItemCode(v string) *QueryDictDataItemResponseBodyDataSubListData {
	s.DataItemCode = &v
	return s
}

func (s *QueryDictDataItemResponseBodyDataSubListData) SetDataItemOrder(v int32) *QueryDictDataItemResponseBodyDataSubListData {
	s.DataItemOrder = &v
	return s
}

func (s *QueryDictDataItemResponseBodyDataSubListData) SetDataItemParent(v string) *QueryDictDataItemResponseBodyDataSubListData {
	s.DataItemParent = &v
	return s
}

func (s *QueryDictDataItemResponseBodyDataSubListData) SetDataItemValue(v string) *QueryDictDataItemResponseBodyDataSubListData {
	s.DataItemValue = &v
	return s
}

func (s *QueryDictDataItemResponseBodyDataSubListData) SetItemType(v int32) *QueryDictDataItemResponseBodyDataSubListData {
	s.ItemType = &v
	return s
}

func (s *QueryDictDataItemResponseBodyDataSubListData) Validate() error {
	return dara.Validate(s)
}
