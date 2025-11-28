// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSFabricConsortiumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBaaSFabricConsortiumResponseBody
	GetCode() *string
	SetData(v []*ListBaaSFabricConsortiumResponseBodyData) *ListBaaSFabricConsortiumResponseBody
	GetData() []*ListBaaSFabricConsortiumResponseBodyData
	SetHttpStatusCode(v int32) *ListBaaSFabricConsortiumResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBaaSFabricConsortiumResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBaaSFabricConsortiumResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBaaSFabricConsortiumResponseBody
	GetSuccess() *bool
}

type ListBaaSFabricConsortiumResponseBody struct {
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListBaaSFabricConsortiumResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBaaSFabricConsortiumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricConsortiumResponseBody) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricConsortiumResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBaaSFabricConsortiumResponseBody) GetData() []*ListBaaSFabricConsortiumResponseBodyData {
	return s.Data
}

func (s *ListBaaSFabricConsortiumResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBaaSFabricConsortiumResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBaaSFabricConsortiumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBaaSFabricConsortiumResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBaaSFabricConsortiumResponseBody) SetCode(v string) *ListBaaSFabricConsortiumResponseBody {
	s.Code = &v
	return s
}

func (s *ListBaaSFabricConsortiumResponseBody) SetData(v []*ListBaaSFabricConsortiumResponseBodyData) *ListBaaSFabricConsortiumResponseBody {
	s.Data = v
	return s
}

func (s *ListBaaSFabricConsortiumResponseBody) SetHttpStatusCode(v int32) *ListBaaSFabricConsortiumResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBaaSFabricConsortiumResponseBody) SetMessage(v string) *ListBaaSFabricConsortiumResponseBody {
	s.Message = &v
	return s
}

func (s *ListBaaSFabricConsortiumResponseBody) SetRequestId(v string) *ListBaaSFabricConsortiumResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBaaSFabricConsortiumResponseBody) SetSuccess(v bool) *ListBaaSFabricConsortiumResponseBody {
	s.Success = &v
	return s
}

func (s *ListBaaSFabricConsortiumResponseBody) Validate() error {
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

type ListBaaSFabricConsortiumResponseBodyData struct {
	BaaSFabricConsortiumId   *string `json:"BaaSFabricConsortiumId,omitempty" xml:"BaaSFabricConsortiumId,omitempty"`
	BaaSFabricConsortiumName *string `json:"BaaSFabricConsortiumName,omitempty" xml:"BaaSFabricConsortiumName,omitempty"`
}

func (s ListBaaSFabricConsortiumResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricConsortiumResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricConsortiumResponseBodyData) GetBaaSFabricConsortiumId() *string {
	return s.BaaSFabricConsortiumId
}

func (s *ListBaaSFabricConsortiumResponseBodyData) GetBaaSFabricConsortiumName() *string {
	return s.BaaSFabricConsortiumName
}

func (s *ListBaaSFabricConsortiumResponseBodyData) SetBaaSFabricConsortiumId(v string) *ListBaaSFabricConsortiumResponseBodyData {
	s.BaaSFabricConsortiumId = &v
	return s
}

func (s *ListBaaSFabricConsortiumResponseBodyData) SetBaaSFabricConsortiumName(v string) *ListBaaSFabricConsortiumResponseBodyData {
	s.BaaSFabricConsortiumName = &v
	return s
}

func (s *ListBaaSFabricConsortiumResponseBodyData) Validate() error {
	return dara.Validate(s)
}
