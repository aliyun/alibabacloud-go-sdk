// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSAntChainPeerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBaaSAntChainPeerResponseBody
	GetCode() *string
	SetData(v []*ListBaaSAntChainPeerResponseBodyData) *ListBaaSAntChainPeerResponseBody
	GetData() []*ListBaaSAntChainPeerResponseBodyData
	SetHttpStatusCode(v int32) *ListBaaSAntChainPeerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBaaSAntChainPeerResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBaaSAntChainPeerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBaaSAntChainPeerResponseBody
	GetSuccess() *bool
}

type ListBaaSAntChainPeerResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListBaaSAntChainPeerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBaaSAntChainPeerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainPeerResponseBody) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainPeerResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBaaSAntChainPeerResponseBody) GetData() []*ListBaaSAntChainPeerResponseBodyData {
	return s.Data
}

func (s *ListBaaSAntChainPeerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBaaSAntChainPeerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBaaSAntChainPeerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBaaSAntChainPeerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBaaSAntChainPeerResponseBody) SetCode(v string) *ListBaaSAntChainPeerResponseBody {
	s.Code = &v
	return s
}

func (s *ListBaaSAntChainPeerResponseBody) SetData(v []*ListBaaSAntChainPeerResponseBodyData) *ListBaaSAntChainPeerResponseBody {
	s.Data = v
	return s
}

func (s *ListBaaSAntChainPeerResponseBody) SetHttpStatusCode(v int32) *ListBaaSAntChainPeerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBaaSAntChainPeerResponseBody) SetMessage(v string) *ListBaaSAntChainPeerResponseBody {
	s.Message = &v
	return s
}

func (s *ListBaaSAntChainPeerResponseBody) SetRequestId(v string) *ListBaaSAntChainPeerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBaaSAntChainPeerResponseBody) SetSuccess(v bool) *ListBaaSAntChainPeerResponseBody {
	s.Success = &v
	return s
}

func (s *ListBaaSAntChainPeerResponseBody) Validate() error {
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

type ListBaaSAntChainPeerResponseBodyData struct {
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s ListBaaSAntChainPeerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainPeerResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainPeerResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *ListBaaSAntChainPeerResponseBodyData) SetNodeName(v string) *ListBaaSAntChainPeerResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *ListBaaSAntChainPeerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
