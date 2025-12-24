// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListK8sResourceResponseBody
	GetCode() *int32
	SetData(v []*ListK8sResourceResponseBodyData) *ListK8sResourceResponseBody
	GetData() []*ListK8sResourceResponseBodyData
	SetMessage(v string) *ListK8sResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListK8sResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListK8sResourceResponseBody
	GetSuccess() *bool
}

type ListK8sResourceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*ListK8sResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2ECA6FC9-7557-5576-AF5F-FC3E7BCC9C21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListK8sResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListK8sResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListK8sResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListK8sResourceResponseBody) GetData() []*ListK8sResourceResponseBodyData {
	return s.Data
}

func (s *ListK8sResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListK8sResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListK8sResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListK8sResourceResponseBody) SetCode(v int32) *ListK8sResourceResponseBody {
	s.Code = &v
	return s
}

func (s *ListK8sResourceResponseBody) SetData(v []*ListK8sResourceResponseBodyData) *ListK8sResourceResponseBody {
	s.Data = v
	return s
}

func (s *ListK8sResourceResponseBody) SetMessage(v string) *ListK8sResourceResponseBody {
	s.Message = &v
	return s
}

func (s *ListK8sResourceResponseBody) SetRequestId(v string) *ListK8sResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListK8sResourceResponseBody) SetSuccess(v bool) *ListK8sResourceResponseBody {
	s.Success = &v
	return s
}

func (s *ListK8sResourceResponseBody) Validate() error {
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

type ListK8sResourceResponseBodyData struct {
	// example:
	//
	// xxljob-01632622cda2f
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// q_ecs_enterprise_spot_c
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s ListK8sResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListK8sResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListK8sResourceResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListK8sResourceResponseBodyData) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListK8sResourceResponseBodyData) SetResourceId(v string) *ListK8sResourceResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ListK8sResourceResponseBodyData) SetResourceName(v string) *ListK8sResourceResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *ListK8sResourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
