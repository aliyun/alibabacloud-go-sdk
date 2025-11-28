// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRouteRuleResponseBody
	GetCode() *string
	SetData(v *ListRouteRuleResponseBodyData) *ListRouteRuleResponseBody
	GetData() *ListRouteRuleResponseBodyData
	SetHttpStatusCode(v int32) *ListRouteRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRouteRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRouteRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRouteRuleResponseBody
	GetSuccess() *bool
}

type ListRouteRuleResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListRouteRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRouteRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRouteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListRouteRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRouteRuleResponseBody) GetData() *ListRouteRuleResponseBodyData {
	return s.Data
}

func (s *ListRouteRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRouteRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRouteRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRouteRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRouteRuleResponseBody) SetCode(v string) *ListRouteRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListRouteRuleResponseBody) SetData(v *ListRouteRuleResponseBodyData) *ListRouteRuleResponseBody {
	s.Data = v
	return s
}

func (s *ListRouteRuleResponseBody) SetHttpStatusCode(v int32) *ListRouteRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRouteRuleResponseBody) SetMessage(v string) *ListRouteRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ListRouteRuleResponseBody) SetRequestId(v string) *ListRouteRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRouteRuleResponseBody) SetSuccess(v bool) *ListRouteRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ListRouteRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRouteRuleResponseBodyData struct {
	Num      *int32                                   `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListRouteRuleResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                   `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListRouteRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRouteRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRouteRuleResponseBodyData) GetNum() *int32 {
	return s.Num
}

func (s *ListRouteRuleResponseBodyData) GetPageData() []*ListRouteRuleResponseBodyDataPageData {
	return s.PageData
}

func (s *ListRouteRuleResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListRouteRuleResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListRouteRuleResponseBodyData) SetNum(v int32) *ListRouteRuleResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListRouteRuleResponseBodyData) SetPageData(v []*ListRouteRuleResponseBodyDataPageData) *ListRouteRuleResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListRouteRuleResponseBodyData) SetSize(v int32) *ListRouteRuleResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListRouteRuleResponseBodyData) SetTotal(v int32) *ListRouteRuleResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListRouteRuleResponseBodyData) Validate() error {
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

type ListRouteRuleResponseBodyDataPageData struct {
	BizChainId         *int64  `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	BizChainName       *string `json:"BizChainName,omitempty" xml:"BizChainName,omitempty"`
	BlockChainType     *string `json:"BlockChainType,omitempty" xml:"BlockChainType,omitempty"`
	ChainUpMode        *string `json:"ChainUpMode,omitempty" xml:"ChainUpMode,omitempty"`
	ContractName       *string `json:"ContractName,omitempty" xml:"ContractName,omitempty"`
	ContractTemplateId *string `json:"ContractTemplateId,omitempty" xml:"ContractTemplateId,omitempty"`
	DeviceGroupId      *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceGroupName    *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	InvokeType         *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	PrivacyRuleId      *string `json:"PrivacyRuleId,omitempty" xml:"PrivacyRuleId,omitempty"`
	PrivacyRuleName    *string `json:"PrivacyRuleName,omitempty" xml:"PrivacyRuleName,omitempty"`
	Remark             *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RouteRuleId        *string `json:"RouteRuleId,omitempty" xml:"RouteRuleId,omitempty"`
}

func (s ListRouteRuleResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListRouteRuleResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListRouteRuleResponseBodyDataPageData) GetBizChainId() *int64 {
	return s.BizChainId
}

func (s *ListRouteRuleResponseBodyDataPageData) GetBizChainName() *string {
	return s.BizChainName
}

func (s *ListRouteRuleResponseBodyDataPageData) GetBlockChainType() *string {
	return s.BlockChainType
}

func (s *ListRouteRuleResponseBodyDataPageData) GetChainUpMode() *string {
	return s.ChainUpMode
}

func (s *ListRouteRuleResponseBodyDataPageData) GetContractName() *string {
	return s.ContractName
}

func (s *ListRouteRuleResponseBodyDataPageData) GetContractTemplateId() *string {
	return s.ContractTemplateId
}

func (s *ListRouteRuleResponseBodyDataPageData) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *ListRouteRuleResponseBodyDataPageData) GetDeviceGroupName() *string {
	return s.DeviceGroupName
}

func (s *ListRouteRuleResponseBodyDataPageData) GetInvokeType() *string {
	return s.InvokeType
}

func (s *ListRouteRuleResponseBodyDataPageData) GetPrivacyRuleId() *string {
	return s.PrivacyRuleId
}

func (s *ListRouteRuleResponseBodyDataPageData) GetPrivacyRuleName() *string {
	return s.PrivacyRuleName
}

func (s *ListRouteRuleResponseBodyDataPageData) GetRemark() *string {
	return s.Remark
}

func (s *ListRouteRuleResponseBodyDataPageData) GetRouteRuleId() *string {
	return s.RouteRuleId
}

func (s *ListRouteRuleResponseBodyDataPageData) SetBizChainId(v int64) *ListRouteRuleResponseBodyDataPageData {
	s.BizChainId = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetBizChainName(v string) *ListRouteRuleResponseBodyDataPageData {
	s.BizChainName = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetBlockChainType(v string) *ListRouteRuleResponseBodyDataPageData {
	s.BlockChainType = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetChainUpMode(v string) *ListRouteRuleResponseBodyDataPageData {
	s.ChainUpMode = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetContractName(v string) *ListRouteRuleResponseBodyDataPageData {
	s.ContractName = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetContractTemplateId(v string) *ListRouteRuleResponseBodyDataPageData {
	s.ContractTemplateId = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetDeviceGroupId(v string) *ListRouteRuleResponseBodyDataPageData {
	s.DeviceGroupId = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetDeviceGroupName(v string) *ListRouteRuleResponseBodyDataPageData {
	s.DeviceGroupName = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetInvokeType(v string) *ListRouteRuleResponseBodyDataPageData {
	s.InvokeType = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetPrivacyRuleId(v string) *ListRouteRuleResponseBodyDataPageData {
	s.PrivacyRuleId = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetPrivacyRuleName(v string) *ListRouteRuleResponseBodyDataPageData {
	s.PrivacyRuleName = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetRemark(v string) *ListRouteRuleResponseBodyDataPageData {
	s.Remark = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) SetRouteRuleId(v string) *ListRouteRuleResponseBodyDataPageData {
	s.RouteRuleId = &v
	return s
}

func (s *ListRouteRuleResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
