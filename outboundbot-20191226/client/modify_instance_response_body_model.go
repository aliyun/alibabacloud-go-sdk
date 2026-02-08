// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyInstanceResponseBody
	GetHttpStatusCode() *int32
	SetInstance(v *ModifyInstanceResponseBodyInstance) *ModifyInstanceResponseBody
	GetInstance() *ModifyInstanceResponseBodyInstance
	SetMessage(v string) *ModifyInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyInstanceResponseBody
	GetSuccess() *bool
}

type ModifyInstanceResponseBody struct {
	// Status code.
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
	//
	// example:
	//
	// {\\"InstanceId\\": \\"ob369xifpi2074\\", \\"AutoUpgradeObVersion\\": False}
	Instance *ModifyInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
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

func (s ModifyInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyInstanceResponseBody) GetInstance() *ModifyInstanceResponseBodyInstance {
	return s.Instance
}

func (s *ModifyInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyInstanceResponseBody) SetCode(v string) *ModifyInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetHttpStatusCode(v int32) *ModifyInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetInstance(v *ModifyInstanceResponseBodyInstance) *ModifyInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *ModifyInstanceResponseBody) SetMessage(v string) *ModifyInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetRequestId(v string) *ModifyInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetSuccess(v bool) *ModifyInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyInstanceResponseBodyInstance struct {
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
	// 123123
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
	// 这是第一个实例
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
	// Maximum concurrent conversations.
	//
	// example:
	//
	// 5
	MaxConcurrentConversation *int32 `json:"MaxConcurrentConversation,omitempty" xml:"MaxConcurrentConversation,omitempty"`
	// Owner name.
	//
	// example:
	//
	// xxx
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
}

func (s ModifyInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBodyInstance) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ModifyInstanceResponseBodyInstance) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *ModifyInstanceResponseBodyInstance) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ModifyInstanceResponseBodyInstance) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *ModifyInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceResponseBodyInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceResponseBodyInstance) GetMaxConcurrentConversation() *int32 {
	return s.MaxConcurrentConversation
}

func (s *ModifyInstanceResponseBodyInstance) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ModifyInstanceResponseBodyInstance) SetCreationTime(v int64) *ModifyInstanceResponseBodyInstance {
	s.CreationTime = &v
	return s
}

func (s *ModifyInstanceResponseBodyInstance) SetCreatorId(v int64) *ModifyInstanceResponseBodyInstance {
	s.CreatorId = &v
	return s
}

func (s *ModifyInstanceResponseBodyInstance) SetCreatorName(v string) *ModifyInstanceResponseBodyInstance {
	s.CreatorName = &v
	return s
}

func (s *ModifyInstanceResponseBodyInstance) SetInstanceDescription(v string) *ModifyInstanceResponseBodyInstance {
	s.InstanceDescription = &v
	return s
}

func (s *ModifyInstanceResponseBodyInstance) SetInstanceId(v string) *ModifyInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceResponseBodyInstance) SetInstanceName(v string) *ModifyInstanceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceResponseBodyInstance) SetMaxConcurrentConversation(v int32) *ModifyInstanceResponseBodyInstance {
	s.MaxConcurrentConversation = &v
	return s
}

func (s *ModifyInstanceResponseBodyInstance) SetOwnerName(v string) *ModifyInstanceResponseBodyInstance {
	s.OwnerName = &v
	return s
}

func (s *ModifyInstanceResponseBodyInstance) Validate() error {
	return dara.Validate(s)
}
