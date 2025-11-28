// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllBizChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllBizChainResponseBody
	GetCode() *string
	SetData(v []*ListAllBizChainResponseBodyData) *ListAllBizChainResponseBody
	GetData() []*ListAllBizChainResponseBodyData
	SetHttpStatusCode(v int32) *ListAllBizChainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAllBizChainResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllBizChainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllBizChainResponseBody
	GetSuccess() *bool
}

type ListAllBizChainResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListAllBizChainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllBizChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllBizChainResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllBizChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllBizChainResponseBody) GetData() []*ListAllBizChainResponseBodyData {
	return s.Data
}

func (s *ListAllBizChainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAllBizChainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllBizChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllBizChainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllBizChainResponseBody) SetCode(v string) *ListAllBizChainResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllBizChainResponseBody) SetData(v []*ListAllBizChainResponseBodyData) *ListAllBizChainResponseBody {
	s.Data = v
	return s
}

func (s *ListAllBizChainResponseBody) SetHttpStatusCode(v int32) *ListAllBizChainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAllBizChainResponseBody) SetMessage(v string) *ListAllBizChainResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllBizChainResponseBody) SetRequestId(v string) *ListAllBizChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllBizChainResponseBody) SetSuccess(v bool) *ListAllBizChainResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllBizChainResponseBody) Validate() error {
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

type ListAllBizChainResponseBodyData struct {
	BizChainId       *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UsedOnchainCount *int64  `json:"UsedOnchainCount,omitempty" xml:"UsedOnchainCount,omitempty"`
}

func (s ListAllBizChainResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllBizChainResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllBizChainResponseBodyData) GetBizChainId() *string {
	return s.BizChainId
}

func (s *ListAllBizChainResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAllBizChainResponseBodyData) GetUsedOnchainCount() *int64 {
	return s.UsedOnchainCount
}

func (s *ListAllBizChainResponseBodyData) SetBizChainId(v string) *ListAllBizChainResponseBodyData {
	s.BizChainId = &v
	return s
}

func (s *ListAllBizChainResponseBodyData) SetName(v string) *ListAllBizChainResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAllBizChainResponseBodyData) SetUsedOnchainCount(v int64) *ListAllBizChainResponseBodyData {
	s.UsedOnchainCount = &v
	return s
}

func (s *ListAllBizChainResponseBodyData) Validate() error {
	return dara.Validate(s)
}
