// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeInstanceResponseBody
	GetHttpStatusCode() *int32
	SetInstance(v *DescribeInstanceResponseBodyInstance) *DescribeInstanceResponseBody
	GetInstance() *DescribeInstanceResponseBodyInstance
	SetMessage(v string) *DescribeInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceResponseBody
	GetSuccess() *bool
}

type DescribeInstanceResponseBody struct {
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
	Instance *DescribeInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// API message.
	//
	// example:
	//
	// Success
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

func (s DescribeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeInstanceResponseBody) GetInstance() *DescribeInstanceResponseBodyInstance {
	return s.Instance
}

func (s *DescribeInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceResponseBody) SetCode(v string) *DescribeInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetHttpStatusCode(v int32) *DescribeInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstance(v *DescribeInstanceResponseBodyInstance) *DescribeInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *DescribeInstanceResponseBody) SetMessage(v string) *DescribeInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetSuccess(v bool) *DescribeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceResponseBodyInstance struct {
	// Creation Time.
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
	// Business instance owner name.
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

func (s DescribeInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstance) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribeInstanceResponseBodyInstance) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *DescribeInstanceResponseBodyInstance) GetCreatorName() *string {
	return s.CreatorName
}

func (s *DescribeInstanceResponseBodyInstance) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *DescribeInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceResponseBodyInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceResponseBodyInstance) GetMaxConcurrentConversation() *int32 {
	return s.MaxConcurrentConversation
}

func (s *DescribeInstanceResponseBodyInstance) GetOwnerName() *string {
	return s.OwnerName
}

func (s *DescribeInstanceResponseBodyInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceResponseBodyInstance) SetCreationTime(v int64) *DescribeInstanceResponseBodyInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetCreatorId(v int64) *DescribeInstanceResponseBodyInstance {
	s.CreatorId = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetCreatorName(v string) *DescribeInstanceResponseBodyInstance {
	s.CreatorName = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetInstanceDescription(v string) *DescribeInstanceResponseBodyInstance {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetInstanceId(v string) *DescribeInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetInstanceName(v string) *DescribeInstanceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetMaxConcurrentConversation(v int32) *DescribeInstanceResponseBodyInstance {
	s.MaxConcurrentConversation = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetOwnerName(v string) *DescribeInstanceResponseBodyInstance {
	s.OwnerName = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetResourceGroupId(v string) *DescribeInstanceResponseBodyInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) Validate() error {
	return dara.Validate(s)
}
