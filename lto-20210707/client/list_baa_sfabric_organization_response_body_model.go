// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSFabricOrganizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBaaSFabricOrganizationResponseBody
	GetCode() *string
	SetData(v []*ListBaaSFabricOrganizationResponseBodyData) *ListBaaSFabricOrganizationResponseBody
	GetData() []*ListBaaSFabricOrganizationResponseBodyData
	SetHttpStatusCode(v int32) *ListBaaSFabricOrganizationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBaaSFabricOrganizationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBaaSFabricOrganizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBaaSFabricOrganizationResponseBody
	GetSuccess() *bool
}

type ListBaaSFabricOrganizationResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListBaaSFabricOrganizationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBaaSFabricOrganizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricOrganizationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBaaSFabricOrganizationResponseBody) GetData() []*ListBaaSFabricOrganizationResponseBodyData {
	return s.Data
}

func (s *ListBaaSFabricOrganizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBaaSFabricOrganizationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBaaSFabricOrganizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBaaSFabricOrganizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBaaSFabricOrganizationResponseBody) SetCode(v string) *ListBaaSFabricOrganizationResponseBody {
	s.Code = &v
	return s
}

func (s *ListBaaSFabricOrganizationResponseBody) SetData(v []*ListBaaSFabricOrganizationResponseBodyData) *ListBaaSFabricOrganizationResponseBody {
	s.Data = v
	return s
}

func (s *ListBaaSFabricOrganizationResponseBody) SetHttpStatusCode(v int32) *ListBaaSFabricOrganizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBaaSFabricOrganizationResponseBody) SetMessage(v string) *ListBaaSFabricOrganizationResponseBody {
	s.Message = &v
	return s
}

func (s *ListBaaSFabricOrganizationResponseBody) SetRequestId(v string) *ListBaaSFabricOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBaaSFabricOrganizationResponseBody) SetSuccess(v bool) *ListBaaSFabricOrganizationResponseBody {
	s.Success = &v
	return s
}

func (s *ListBaaSFabricOrganizationResponseBody) Validate() error {
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

type ListBaaSFabricOrganizationResponseBodyData struct {
	BaaSFabricOrganizationId   *string `json:"BaaSFabricOrganizationId,omitempty" xml:"BaaSFabricOrganizationId,omitempty"`
	BaaSFabricOrganizationName *string `json:"BaaSFabricOrganizationName,omitempty" xml:"BaaSFabricOrganizationName,omitempty"`
}

func (s ListBaaSFabricOrganizationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricOrganizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricOrganizationResponseBodyData) GetBaaSFabricOrganizationId() *string {
	return s.BaaSFabricOrganizationId
}

func (s *ListBaaSFabricOrganizationResponseBodyData) GetBaaSFabricOrganizationName() *string {
	return s.BaaSFabricOrganizationName
}

func (s *ListBaaSFabricOrganizationResponseBodyData) SetBaaSFabricOrganizationId(v string) *ListBaaSFabricOrganizationResponseBodyData {
	s.BaaSFabricOrganizationId = &v
	return s
}

func (s *ListBaaSFabricOrganizationResponseBodyData) SetBaaSFabricOrganizationName(v string) *ListBaaSFabricOrganizationResponseBodyData {
	s.BaaSFabricOrganizationName = &v
	return s
}

func (s *ListBaaSFabricOrganizationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
