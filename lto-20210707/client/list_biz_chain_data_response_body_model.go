// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizChainDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBizChainDataResponseBody
	GetCode() *string
	SetData(v *ListBizChainDataResponseBodyData) *ListBizChainDataResponseBody
	GetData() *ListBizChainDataResponseBodyData
	SetHttpStatusCode(v int32) *ListBizChainDataResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBizChainDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBizChainDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBizChainDataResponseBody
	GetSuccess() *bool
}

type ListBizChainDataResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListBizChainDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBizChainDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBizChainDataResponseBody) GoString() string {
	return s.String()
}

func (s *ListBizChainDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBizChainDataResponseBody) GetData() *ListBizChainDataResponseBodyData {
	return s.Data
}

func (s *ListBizChainDataResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBizChainDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBizChainDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBizChainDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBizChainDataResponseBody) SetCode(v string) *ListBizChainDataResponseBody {
	s.Code = &v
	return s
}

func (s *ListBizChainDataResponseBody) SetData(v *ListBizChainDataResponseBodyData) *ListBizChainDataResponseBody {
	s.Data = v
	return s
}

func (s *ListBizChainDataResponseBody) SetHttpStatusCode(v int32) *ListBizChainDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBizChainDataResponseBody) SetMessage(v string) *ListBizChainDataResponseBody {
	s.Message = &v
	return s
}

func (s *ListBizChainDataResponseBody) SetRequestId(v string) *ListBizChainDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBizChainDataResponseBody) SetSuccess(v bool) *ListBizChainDataResponseBody {
	s.Success = &v
	return s
}

func (s *ListBizChainDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBizChainDataResponseBodyData struct {
	Num      *int32                                      `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListBizChainDataResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                      `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListBizChainDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBizChainDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBizChainDataResponseBodyData) GetNum() *int32 {
	return s.Num
}

func (s *ListBizChainDataResponseBodyData) GetPageData() []*ListBizChainDataResponseBodyDataPageData {
	return s.PageData
}

func (s *ListBizChainDataResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListBizChainDataResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListBizChainDataResponseBodyData) SetNum(v int32) *ListBizChainDataResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListBizChainDataResponseBodyData) SetPageData(v []*ListBizChainDataResponseBodyDataPageData) *ListBizChainDataResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListBizChainDataResponseBodyData) SetSize(v int32) *ListBizChainDataResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListBizChainDataResponseBodyData) SetTotal(v int32) *ListBizChainDataResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListBizChainDataResponseBodyData) Validate() error {
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBizChainDataResponseBodyDataPageData struct {
	BlockHash  *string `json:"BlockHash,omitempty" xml:"BlockHash,omitempty"`
	BlockNum   *string `json:"BlockNum,omitempty" xml:"BlockNum,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotDataDID *string `json:"IotDataDID,omitempty" xml:"IotDataDID,omitempty"`
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Timestamp  *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TxHash     *string `json:"TxHash,omitempty" xml:"TxHash,omitempty"`
}

func (s ListBizChainDataResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListBizChainDataResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListBizChainDataResponseBodyDataPageData) GetBlockHash() *string {
	return s.BlockHash
}

func (s *ListBizChainDataResponseBodyDataPageData) GetBlockNum() *string {
	return s.BlockNum
}

func (s *ListBizChainDataResponseBodyDataPageData) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ListBizChainDataResponseBodyDataPageData) GetIotDataDID() *string {
	return s.IotDataDID
}

func (s *ListBizChainDataResponseBodyDataPageData) GetMemberName() *string {
	return s.MemberName
}

func (s *ListBizChainDataResponseBodyDataPageData) GetProductKey() *string {
	return s.ProductKey
}

func (s *ListBizChainDataResponseBodyDataPageData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListBizChainDataResponseBodyDataPageData) GetTxHash() *string {
	return s.TxHash
}

func (s *ListBizChainDataResponseBodyDataPageData) SetBlockHash(v string) *ListBizChainDataResponseBodyDataPageData {
	s.BlockHash = &v
	return s
}

func (s *ListBizChainDataResponseBodyDataPageData) SetBlockNum(v string) *ListBizChainDataResponseBodyDataPageData {
	s.BlockNum = &v
	return s
}

func (s *ListBizChainDataResponseBodyDataPageData) SetDeviceName(v string) *ListBizChainDataResponseBodyDataPageData {
	s.DeviceName = &v
	return s
}

func (s *ListBizChainDataResponseBodyDataPageData) SetIotDataDID(v string) *ListBizChainDataResponseBodyDataPageData {
	s.IotDataDID = &v
	return s
}

func (s *ListBizChainDataResponseBodyDataPageData) SetMemberName(v string) *ListBizChainDataResponseBodyDataPageData {
	s.MemberName = &v
	return s
}

func (s *ListBizChainDataResponseBodyDataPageData) SetProductKey(v string) *ListBizChainDataResponseBodyDataPageData {
	s.ProductKey = &v
	return s
}

func (s *ListBizChainDataResponseBodyDataPageData) SetTimestamp(v int64) *ListBizChainDataResponseBodyDataPageData {
	s.Timestamp = &v
	return s
}

func (s *ListBizChainDataResponseBodyDataPageData) SetTxHash(v string) *ListBizChainDataResponseBodyDataPageData {
	s.TxHash = &v
	return s
}

func (s *ListBizChainDataResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
