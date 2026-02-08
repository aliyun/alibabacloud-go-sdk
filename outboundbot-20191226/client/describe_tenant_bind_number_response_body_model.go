// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTenantBindNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTenantBindNumberResponseBody
	GetCode() *string
	SetData(v *DescribeTenantBindNumberResponseBodyData) *DescribeTenantBindNumberResponseBody
	GetData() *DescribeTenantBindNumberResponseBodyData
	SetHttpStatusCode(v int32) *DescribeTenantBindNumberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeTenantBindNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTenantBindNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTenantBindNumberResponseBody
	GetSuccess() *bool
}

type DescribeTenantBindNumberResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *DescribeTenantBindNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 1364f208-982d-4d0c-89aa-d56e22b47589
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTenantBindNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTenantBindNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantBindNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTenantBindNumberResponseBody) GetData() *DescribeTenantBindNumberResponseBodyData {
	return s.Data
}

func (s *DescribeTenantBindNumberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeTenantBindNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTenantBindNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTenantBindNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTenantBindNumberResponseBody) SetCode(v string) *DescribeTenantBindNumberResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTenantBindNumberResponseBody) SetData(v *DescribeTenantBindNumberResponseBodyData) *DescribeTenantBindNumberResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTenantBindNumberResponseBody) SetHttpStatusCode(v int32) *DescribeTenantBindNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTenantBindNumberResponseBody) SetMessage(v string) *DescribeTenantBindNumberResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTenantBindNumberResponseBody) SetRequestId(v string) *DescribeTenantBindNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantBindNumberResponseBody) SetSuccess(v bool) *DescribeTenantBindNumberResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTenantBindNumberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTenantBindNumberResponseBodyData struct {
	// Job group description
	List []*DescribeTenantBindNumberResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeTenantBindNumberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTenantBindNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTenantBindNumberResponseBodyData) GetList() []*DescribeTenantBindNumberResponseBodyDataList {
	return s.List
}

func (s *DescribeTenantBindNumberResponseBodyData) SetList(v []*DescribeTenantBindNumberResponseBodyDataList) *DescribeTenantBindNumberResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeTenantBindNumberResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTenantBindNumberResponseBodyDataList struct {
	// instance ID
	//
	// example:
	//
	// e2d7a184-7d6c-45d4-ac24-34ab48f54669
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance name
	//
	// example:
	//
	// xxxx
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether the instance is in the attached state.
	//
	// example:
	//
	// true
	IsBinding *bool `json:"IsBinding,omitempty" xml:"IsBinding,omitempty"`
}

func (s DescribeTenantBindNumberResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTenantBindNumberResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeTenantBindNumberResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTenantBindNumberResponseBodyDataList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeTenantBindNumberResponseBodyDataList) GetIsBinding() *bool {
	return s.IsBinding
}

func (s *DescribeTenantBindNumberResponseBodyDataList) SetInstanceId(v string) *DescribeTenantBindNumberResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *DescribeTenantBindNumberResponseBodyDataList) SetInstanceName(v string) *DescribeTenantBindNumberResponseBodyDataList {
	s.InstanceName = &v
	return s
}

func (s *DescribeTenantBindNumberResponseBodyDataList) SetIsBinding(v bool) *DescribeTenantBindNumberResponseBodyDataList {
	s.IsBinding = &v
	return s
}

func (s *DescribeTenantBindNumberResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
