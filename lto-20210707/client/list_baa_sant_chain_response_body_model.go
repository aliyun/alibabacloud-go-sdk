// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSAntChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBaaSAntChainResponseBody
	GetCode() *string
	SetData(v []*ListBaaSAntChainResponseBodyData) *ListBaaSAntChainResponseBody
	GetData() []*ListBaaSAntChainResponseBodyData
	SetHttpStatusCode(v int32) *ListBaaSAntChainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBaaSAntChainResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBaaSAntChainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBaaSAntChainResponseBody
	GetSuccess() *bool
}

type ListBaaSAntChainResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListBaaSAntChainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBaaSAntChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainResponseBody) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBaaSAntChainResponseBody) GetData() []*ListBaaSAntChainResponseBodyData {
	return s.Data
}

func (s *ListBaaSAntChainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBaaSAntChainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBaaSAntChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBaaSAntChainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBaaSAntChainResponseBody) SetCode(v string) *ListBaaSAntChainResponseBody {
	s.Code = &v
	return s
}

func (s *ListBaaSAntChainResponseBody) SetData(v []*ListBaaSAntChainResponseBodyData) *ListBaaSAntChainResponseBody {
	s.Data = v
	return s
}

func (s *ListBaaSAntChainResponseBody) SetHttpStatusCode(v int32) *ListBaaSAntChainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBaaSAntChainResponseBody) SetMessage(v string) *ListBaaSAntChainResponseBody {
	s.Message = &v
	return s
}

func (s *ListBaaSAntChainResponseBody) SetRequestId(v string) *ListBaaSAntChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBaaSAntChainResponseBody) SetSuccess(v bool) *ListBaaSAntChainResponseBody {
	s.Success = &v
	return s
}

func (s *ListBaaSAntChainResponseBody) Validate() error {
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

type ListBaaSAntChainResponseBodyData struct {
	BaaSAntChainChainId   *string `json:"BaaSAntChainChainId,omitempty" xml:"BaaSAntChainChainId,omitempty"`
	BaaSAntChainChainName *string `json:"BaaSAntChainChainName,omitempty" xml:"BaaSAntChainChainName,omitempty"`
}

func (s ListBaaSAntChainResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainResponseBodyData) GetBaaSAntChainChainId() *string {
	return s.BaaSAntChainChainId
}

func (s *ListBaaSAntChainResponseBodyData) GetBaaSAntChainChainName() *string {
	return s.BaaSAntChainChainName
}

func (s *ListBaaSAntChainResponseBodyData) SetBaaSAntChainChainId(v string) *ListBaaSAntChainResponseBodyData {
	s.BaaSAntChainChainId = &v
	return s
}

func (s *ListBaaSAntChainResponseBodyData) SetBaaSAntChainChainName(v string) *ListBaaSAntChainResponseBodyData {
	s.BaaSAntChainChainName = &v
	return s
}

func (s *ListBaaSAntChainResponseBodyData) Validate() error {
	return dara.Validate(s)
}
