// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSFabricChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBaaSFabricChannelResponseBody
	GetCode() *string
	SetData(v []*ListBaaSFabricChannelResponseBodyData) *ListBaaSFabricChannelResponseBody
	GetData() []*ListBaaSFabricChannelResponseBodyData
	SetHttpStatusCode(v int32) *ListBaaSFabricChannelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBaaSFabricChannelResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBaaSFabricChannelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBaaSFabricChannelResponseBody
	GetSuccess() *bool
}

type ListBaaSFabricChannelResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListBaaSFabricChannelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBaaSFabricChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricChannelResponseBody) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricChannelResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBaaSFabricChannelResponseBody) GetData() []*ListBaaSFabricChannelResponseBodyData {
	return s.Data
}

func (s *ListBaaSFabricChannelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBaaSFabricChannelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBaaSFabricChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBaaSFabricChannelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBaaSFabricChannelResponseBody) SetCode(v string) *ListBaaSFabricChannelResponseBody {
	s.Code = &v
	return s
}

func (s *ListBaaSFabricChannelResponseBody) SetData(v []*ListBaaSFabricChannelResponseBodyData) *ListBaaSFabricChannelResponseBody {
	s.Data = v
	return s
}

func (s *ListBaaSFabricChannelResponseBody) SetHttpStatusCode(v int32) *ListBaaSFabricChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBaaSFabricChannelResponseBody) SetMessage(v string) *ListBaaSFabricChannelResponseBody {
	s.Message = &v
	return s
}

func (s *ListBaaSFabricChannelResponseBody) SetRequestId(v string) *ListBaaSFabricChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBaaSFabricChannelResponseBody) SetSuccess(v bool) *ListBaaSFabricChannelResponseBody {
	s.Success = &v
	return s
}

func (s *ListBaaSFabricChannelResponseBody) Validate() error {
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

type ListBaaSFabricChannelResponseBodyData struct {
	BaaSFabricChannelId   *string `json:"BaaSFabricChannelId,omitempty" xml:"BaaSFabricChannelId,omitempty"`
	BaaSFabricChannelName *string `json:"BaaSFabricChannelName,omitempty" xml:"BaaSFabricChannelName,omitempty"`
}

func (s ListBaaSFabricChannelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricChannelResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricChannelResponseBodyData) GetBaaSFabricChannelId() *string {
	return s.BaaSFabricChannelId
}

func (s *ListBaaSFabricChannelResponseBodyData) GetBaaSFabricChannelName() *string {
	return s.BaaSFabricChannelName
}

func (s *ListBaaSFabricChannelResponseBodyData) SetBaaSFabricChannelId(v string) *ListBaaSFabricChannelResponseBodyData {
	s.BaaSFabricChannelId = &v
	return s
}

func (s *ListBaaSFabricChannelResponseBodyData) SetBaaSFabricChannelName(v string) *ListBaaSFabricChannelResponseBodyData {
	s.BaaSFabricChannelName = &v
	return s
}

func (s *ListBaaSFabricChannelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
