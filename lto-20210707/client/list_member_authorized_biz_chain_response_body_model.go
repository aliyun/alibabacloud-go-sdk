// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemberAuthorizedBizChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemberAuthorizedBizChainResponseBody
	GetCode() *string
	SetData(v []*ListMemberAuthorizedBizChainResponseBodyData) *ListMemberAuthorizedBizChainResponseBody
	GetData() []*ListMemberAuthorizedBizChainResponseBodyData
	SetHttpStatusCode(v int32) *ListMemberAuthorizedBizChainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListMemberAuthorizedBizChainResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMemberAuthorizedBizChainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMemberAuthorizedBizChainResponseBody
	GetSuccess() *bool
}

type ListMemberAuthorizedBizChainResponseBody struct {
	Code           *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListMemberAuthorizedBizChainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMemberAuthorizedBizChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemberAuthorizedBizChainResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemberAuthorizedBizChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemberAuthorizedBizChainResponseBody) GetData() []*ListMemberAuthorizedBizChainResponseBodyData {
	return s.Data
}

func (s *ListMemberAuthorizedBizChainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMemberAuthorizedBizChainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMemberAuthorizedBizChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemberAuthorizedBizChainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMemberAuthorizedBizChainResponseBody) SetCode(v string) *ListMemberAuthorizedBizChainResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBody) SetData(v []*ListMemberAuthorizedBizChainResponseBodyData) *ListMemberAuthorizedBizChainResponseBody {
	s.Data = v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBody) SetHttpStatusCode(v int32) *ListMemberAuthorizedBizChainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBody) SetMessage(v string) *ListMemberAuthorizedBizChainResponseBody {
	s.Message = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBody) SetRequestId(v string) *ListMemberAuthorizedBizChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBody) SetSuccess(v bool) *ListMemberAuthorizedBizChainResponseBody {
	s.Success = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBody) Validate() error {
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

type ListMemberAuthorizedBizChainResponseBodyData struct {
	Authorized   *bool                                                   `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	BizChainId   *string                                                 `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	BizChainName *string                                                 `json:"BizChainName,omitempty" xml:"BizChainName,omitempty"`
	BizChainType *string                                                 `json:"BizChainType,omitempty" xml:"BizChainType,omitempty"`
	PeerList     []*ListMemberAuthorizedBizChainResponseBodyDataPeerList `json:"PeerList,omitempty" xml:"PeerList,omitempty" type:"Repeated"`
}

func (s ListMemberAuthorizedBizChainResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMemberAuthorizedBizChainResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMemberAuthorizedBizChainResponseBodyData) GetAuthorized() *bool {
	return s.Authorized
}

func (s *ListMemberAuthorizedBizChainResponseBodyData) GetBizChainId() *string {
	return s.BizChainId
}

func (s *ListMemberAuthorizedBizChainResponseBodyData) GetBizChainName() *string {
	return s.BizChainName
}

func (s *ListMemberAuthorizedBizChainResponseBodyData) GetBizChainType() *string {
	return s.BizChainType
}

func (s *ListMemberAuthorizedBizChainResponseBodyData) GetPeerList() []*ListMemberAuthorizedBizChainResponseBodyDataPeerList {
	return s.PeerList
}

func (s *ListMemberAuthorizedBizChainResponseBodyData) SetAuthorized(v bool) *ListMemberAuthorizedBizChainResponseBodyData {
	s.Authorized = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBodyData) SetBizChainId(v string) *ListMemberAuthorizedBizChainResponseBodyData {
	s.BizChainId = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBodyData) SetBizChainName(v string) *ListMemberAuthorizedBizChainResponseBodyData {
	s.BizChainName = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBodyData) SetBizChainType(v string) *ListMemberAuthorizedBizChainResponseBodyData {
	s.BizChainType = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBodyData) SetPeerList(v []*ListMemberAuthorizedBizChainResponseBodyDataPeerList) *ListMemberAuthorizedBizChainResponseBodyData {
	s.PeerList = v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBodyData) Validate() error {
	if s.PeerList != nil {
		for _, item := range s.PeerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMemberAuthorizedBizChainResponseBodyDataPeerList struct {
	Authorized *bool   `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	PeerName   *string `json:"PeerName,omitempty" xml:"PeerName,omitempty"`
}

func (s ListMemberAuthorizedBizChainResponseBodyDataPeerList) String() string {
	return dara.Prettify(s)
}

func (s ListMemberAuthorizedBizChainResponseBodyDataPeerList) GoString() string {
	return s.String()
}

func (s *ListMemberAuthorizedBizChainResponseBodyDataPeerList) GetAuthorized() *bool {
	return s.Authorized
}

func (s *ListMemberAuthorizedBizChainResponseBodyDataPeerList) GetPeerName() *string {
	return s.PeerName
}

func (s *ListMemberAuthorizedBizChainResponseBodyDataPeerList) SetAuthorized(v bool) *ListMemberAuthorizedBizChainResponseBodyDataPeerList {
	s.Authorized = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBodyDataPeerList) SetPeerName(v string) *ListMemberAuthorizedBizChainResponseBodyDataPeerList {
	s.PeerName = &v
	return s
}

func (s *ListMemberAuthorizedBizChainResponseBodyDataPeerList) Validate() error {
	return dara.Validate(s)
}
