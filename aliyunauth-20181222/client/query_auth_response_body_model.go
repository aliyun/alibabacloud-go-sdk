// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryAuthResponseBody
	GetCode() *int32
	SetData(v *QueryAuthResponseBodyData) *QueryAuthResponseBody
	GetData() *QueryAuthResponseBodyData
	SetMessage(v string) *QueryAuthResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAuthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAuthResponseBody
	GetSuccess() *bool
}

type QueryAuthResponseBody struct {
	Code      *int32                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryAuthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAuthResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryAuthResponseBody) GetData() *QueryAuthResponseBodyData {
	return s.Data
}

func (s *QueryAuthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAuthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAuthResponseBody) SetCode(v int32) *QueryAuthResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAuthResponseBody) SetData(v *QueryAuthResponseBodyData) *QueryAuthResponseBody {
	s.Data = v
	return s
}

func (s *QueryAuthResponseBody) SetMessage(v string) *QueryAuthResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAuthResponseBody) SetRequestId(v string) *QueryAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAuthResponseBody) SetSuccess(v bool) *QueryAuthResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAuthResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAuthResponseBodyData struct {
	InfoDTOList *QueryAuthResponseBodyDataInfoDTOList `json:"InfoDTOList,omitempty" xml:"InfoDTOList,omitempty" type:"Struct"`
	InstanceId  *string                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProductCode *string                               `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s QueryAuthResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAuthResponseBodyData) GetInfoDTOList() *QueryAuthResponseBodyDataInfoDTOList {
	return s.InfoDTOList
}

func (s *QueryAuthResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryAuthResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryAuthResponseBodyData) SetInfoDTOList(v *QueryAuthResponseBodyDataInfoDTOList) *QueryAuthResponseBodyData {
	s.InfoDTOList = v
	return s
}

func (s *QueryAuthResponseBodyData) SetInstanceId(v string) *QueryAuthResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryAuthResponseBodyData) SetProductCode(v string) *QueryAuthResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *QueryAuthResponseBodyData) Validate() error {
	if s.InfoDTOList != nil {
		if err := s.InfoDTOList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAuthResponseBodyDataInfoDTOList struct {
	InfoDTOList []*QueryAuthResponseBodyDataInfoDTOListInfoDTOList `json:"InfoDTOList,omitempty" xml:"InfoDTOList,omitempty" type:"Repeated"`
}

func (s QueryAuthResponseBodyDataInfoDTOList) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthResponseBodyDataInfoDTOList) GoString() string {
	return s.String()
}

func (s *QueryAuthResponseBodyDataInfoDTOList) GetInfoDTOList() []*QueryAuthResponseBodyDataInfoDTOListInfoDTOList {
	return s.InfoDTOList
}

func (s *QueryAuthResponseBodyDataInfoDTOList) SetInfoDTOList(v []*QueryAuthResponseBodyDataInfoDTOListInfoDTOList) *QueryAuthResponseBodyDataInfoDTOList {
	s.InfoDTOList = v
	return s
}

func (s *QueryAuthResponseBodyDataInfoDTOList) Validate() error {
	if s.InfoDTOList != nil {
		for _, item := range s.InfoDTOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAuthResponseBodyDataInfoDTOListInfoDTOList struct {
	AuthOrderVid  *string `json:"AuthOrderVid,omitempty" xml:"AuthOrderVid,omitempty"`
	ItemName      *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	ItemRecordVid *string `json:"ItemRecordVid,omitempty" xml:"ItemRecordVid,omitempty"`
	OperateCode   *string `json:"OperateCode,omitempty" xml:"OperateCode,omitempty"`
}

func (s QueryAuthResponseBodyDataInfoDTOListInfoDTOList) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthResponseBodyDataInfoDTOListInfoDTOList) GoString() string {
	return s.String()
}

func (s *QueryAuthResponseBodyDataInfoDTOListInfoDTOList) GetAuthOrderVid() *string {
	return s.AuthOrderVid
}

func (s *QueryAuthResponseBodyDataInfoDTOListInfoDTOList) GetItemName() *string {
	return s.ItemName
}

func (s *QueryAuthResponseBodyDataInfoDTOListInfoDTOList) GetItemRecordVid() *string {
	return s.ItemRecordVid
}

func (s *QueryAuthResponseBodyDataInfoDTOListInfoDTOList) GetOperateCode() *string {
	return s.OperateCode
}

func (s *QueryAuthResponseBodyDataInfoDTOListInfoDTOList) SetAuthOrderVid(v string) *QueryAuthResponseBodyDataInfoDTOListInfoDTOList {
	s.AuthOrderVid = &v
	return s
}

func (s *QueryAuthResponseBodyDataInfoDTOListInfoDTOList) SetItemName(v string) *QueryAuthResponseBodyDataInfoDTOListInfoDTOList {
	s.ItemName = &v
	return s
}

func (s *QueryAuthResponseBodyDataInfoDTOListInfoDTOList) SetItemRecordVid(v string) *QueryAuthResponseBodyDataInfoDTOListInfoDTOList {
	s.ItemRecordVid = &v
	return s
}

func (s *QueryAuthResponseBodyDataInfoDTOListInfoDTOList) SetOperateCode(v string) *QueryAuthResponseBodyDataInfoDTOListInfoDTOList {
	s.OperateCode = &v
	return s
}

func (s *QueryAuthResponseBodyDataInfoDTOListInfoDTOList) Validate() error {
	return dara.Validate(s)
}
