// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllBizChainContractResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllBizChainContractResponseBody
	GetCode() *string
	SetData(v []*ListAllBizChainContractResponseBodyData) *ListAllBizChainContractResponseBody
	GetData() []*ListAllBizChainContractResponseBodyData
	SetHttpStatusCode(v int32) *ListAllBizChainContractResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAllBizChainContractResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllBizChainContractResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllBizChainContractResponseBody
	GetSuccess() *bool
}

type ListAllBizChainContractResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListAllBizChainContractResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllBizChainContractResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllBizChainContractResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllBizChainContractResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllBizChainContractResponseBody) GetData() []*ListAllBizChainContractResponseBodyData {
	return s.Data
}

func (s *ListAllBizChainContractResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAllBizChainContractResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllBizChainContractResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllBizChainContractResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllBizChainContractResponseBody) SetCode(v string) *ListAllBizChainContractResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllBizChainContractResponseBody) SetData(v []*ListAllBizChainContractResponseBodyData) *ListAllBizChainContractResponseBody {
	s.Data = v
	return s
}

func (s *ListAllBizChainContractResponseBody) SetHttpStatusCode(v int32) *ListAllBizChainContractResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAllBizChainContractResponseBody) SetMessage(v string) *ListAllBizChainContractResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllBizChainContractResponseBody) SetRequestId(v string) *ListAllBizChainContractResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllBizChainContractResponseBody) SetSuccess(v bool) *ListAllBizChainContractResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllBizChainContractResponseBody) Validate() error {
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

type ListAllBizChainContractResponseBodyData struct {
	ContractName       *string `json:"ContractName,omitempty" xml:"ContractName,omitempty"`
	ContractTemplateId *string `json:"ContractTemplateId,omitempty" xml:"ContractTemplateId,omitempty"`
	InvokeType         *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
}

func (s ListAllBizChainContractResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllBizChainContractResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllBizChainContractResponseBodyData) GetContractName() *string {
	return s.ContractName
}

func (s *ListAllBizChainContractResponseBodyData) GetContractTemplateId() *string {
	return s.ContractTemplateId
}

func (s *ListAllBizChainContractResponseBodyData) GetInvokeType() *string {
	return s.InvokeType
}

func (s *ListAllBizChainContractResponseBodyData) SetContractName(v string) *ListAllBizChainContractResponseBodyData {
	s.ContractName = &v
	return s
}

func (s *ListAllBizChainContractResponseBodyData) SetContractTemplateId(v string) *ListAllBizChainContractResponseBodyData {
	s.ContractTemplateId = &v
	return s
}

func (s *ListAllBizChainContractResponseBodyData) SetInvokeType(v string) *ListAllBizChainContractResponseBodyData {
	s.InvokeType = &v
	return s
}

func (s *ListAllBizChainContractResponseBodyData) Validate() error {
	return dara.Validate(s)
}
