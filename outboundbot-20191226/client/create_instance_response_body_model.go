// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateInstanceResponseBody
	GetHttpStatusCode() *int32
	SetInstance(v *CreateInstanceResponseBodyInstance) *CreateInstanceResponseBody
	GetInstance() *CreateInstanceResponseBodyInstance
	SetMessage(v string) *CreateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceResponseBody
	GetSuccess() *bool
}

type CreateInstanceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Business instance information.
	Instance *CreateInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// API message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateInstanceResponseBody) GetInstance() *CreateInstanceResponseBodyInstance {
	return s.Instance
}

func (s *CreateInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceResponseBody) SetCode(v string) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetHttpStatusCode(v int32) *CreateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstance(v *CreateInstanceResponseBodyInstance) *CreateInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *CreateInstanceResponseBody) SetMessage(v string) *CreateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceResponseBodyInstance struct {
	// Creation Time, in milliseconds.
	//
	// example:
	//
	// 1578469042851
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Creator ID.
	//
	// example:
	//
	// 435986
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// Creator name.
	//
	// example:
	//
	// xxx
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// Business instance description.
	//
	// example:
	//
	// 这个是第一个实例
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// Business instance ID.
	//
	// example:
	//
	// 90515b5-6115-4ccf-83e2-52d5bfaf2ddf
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Business instance name.
	//
	// example:
	//
	// 第一个实例
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Maximum concurrent conversations for the instance.
	//
	// example:
	//
	// 4
	MaxConcurrentConversation *int32 `json:"MaxConcurrentConversation,omitempty" xml:"MaxConcurrentConversation,omitempty"`
	// Owner name.
	//
	// example:
	//
	// xxxx
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// 90515b5-6115-4ccf-83e2-52d5bfaf2ddf
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyInstance) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *CreateInstanceResponseBodyInstance) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *CreateInstanceResponseBodyInstance) GetCreatorName() *string {
	return s.CreatorName
}

func (s *CreateInstanceResponseBodyInstance) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *CreateInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBodyInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceResponseBodyInstance) GetMaxConcurrentConversation() *int32 {
	return s.MaxConcurrentConversation
}

func (s *CreateInstanceResponseBodyInstance) GetOwnerName() *string {
	return s.OwnerName
}

func (s *CreateInstanceResponseBodyInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceResponseBodyInstance) SetCreationTime(v int64) *CreateInstanceResponseBodyInstance {
	s.CreationTime = &v
	return s
}

func (s *CreateInstanceResponseBodyInstance) SetCreatorId(v int64) *CreateInstanceResponseBodyInstance {
	s.CreatorId = &v
	return s
}

func (s *CreateInstanceResponseBodyInstance) SetCreatorName(v string) *CreateInstanceResponseBodyInstance {
	s.CreatorName = &v
	return s
}

func (s *CreateInstanceResponseBodyInstance) SetInstanceDescription(v string) *CreateInstanceResponseBodyInstance {
	s.InstanceDescription = &v
	return s
}

func (s *CreateInstanceResponseBodyInstance) SetInstanceId(v string) *CreateInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyInstance) SetInstanceName(v string) *CreateInstanceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceResponseBodyInstance) SetMaxConcurrentConversation(v int32) *CreateInstanceResponseBodyInstance {
	s.MaxConcurrentConversation = &v
	return s
}

func (s *CreateInstanceResponseBodyInstance) SetOwnerName(v string) *CreateInstanceResponseBodyInstance {
	s.OwnerName = &v
	return s
}

func (s *CreateInstanceResponseBodyInstance) SetResourceGroupId(v string) *CreateInstanceResponseBodyInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceResponseBodyInstance) Validate() error {
	return dara.Validate(s)
}
