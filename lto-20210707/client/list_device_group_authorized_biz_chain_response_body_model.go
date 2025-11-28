// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceGroupAuthorizedBizChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDeviceGroupAuthorizedBizChainResponseBody
	GetCode() *string
	SetData(v []*ListDeviceGroupAuthorizedBizChainResponseBodyData) *ListDeviceGroupAuthorizedBizChainResponseBody
	GetData() []*ListDeviceGroupAuthorizedBizChainResponseBodyData
	SetHttpStatusCode(v int32) *ListDeviceGroupAuthorizedBizChainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDeviceGroupAuthorizedBizChainResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDeviceGroupAuthorizedBizChainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDeviceGroupAuthorizedBizChainResponseBody
	GetSuccess() *bool
}

type ListDeviceGroupAuthorizedBizChainResponseBody struct {
	Code           *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListDeviceGroupAuthorizedBizChainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDeviceGroupAuthorizedBizChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupAuthorizedBizChainResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) GetData() []*ListDeviceGroupAuthorizedBizChainResponseBodyData {
	return s.Data
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) SetCode(v string) *ListDeviceGroupAuthorizedBizChainResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) SetData(v []*ListDeviceGroupAuthorizedBizChainResponseBodyData) *ListDeviceGroupAuthorizedBizChainResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) SetHttpStatusCode(v int32) *ListDeviceGroupAuthorizedBizChainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) SetMessage(v string) *ListDeviceGroupAuthorizedBizChainResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) SetRequestId(v string) *ListDeviceGroupAuthorizedBizChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) SetSuccess(v bool) *ListDeviceGroupAuthorizedBizChainResponseBody {
	s.Success = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBody) Validate() error {
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

type ListDeviceGroupAuthorizedBizChainResponseBodyData struct {
	Authorized     *bool   `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	BizChainId     *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	BizChainName   *string `json:"BizChainName,omitempty" xml:"BizChainName,omitempty"`
	BlockChainType *string `json:"BlockChainType,omitempty" xml:"BlockChainType,omitempty"`
}

func (s ListDeviceGroupAuthorizedBizChainResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupAuthorizedBizChainResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBodyData) GetAuthorized() *bool {
	return s.Authorized
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBodyData) GetBizChainId() *string {
	return s.BizChainId
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBodyData) GetBizChainName() *string {
	return s.BizChainName
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBodyData) GetBlockChainType() *string {
	return s.BlockChainType
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBodyData) SetAuthorized(v bool) *ListDeviceGroupAuthorizedBizChainResponseBodyData {
	s.Authorized = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBodyData) SetBizChainId(v string) *ListDeviceGroupAuthorizedBizChainResponseBodyData {
	s.BizChainId = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBodyData) SetBizChainName(v string) *ListDeviceGroupAuthorizedBizChainResponseBodyData {
	s.BizChainName = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBodyData) SetBlockChainType(v string) *ListDeviceGroupAuthorizedBizChainResponseBodyData {
	s.BlockChainType = &v
	return s
}

func (s *ListDeviceGroupAuthorizedBizChainResponseBodyData) Validate() error {
	return dara.Validate(s)
}
