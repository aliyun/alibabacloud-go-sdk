// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRecallManagementServiceRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateRecallManagementServiceRequest
	GetInstanceId() *string
	SetName(v string) *CreateRecallManagementServiceRequest
	GetName() *string
}

type CreateRecallManagementServiceRequest struct {
	// example:
	//
	// this is a test recall
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// hot_global_recall
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateRecallManagementServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRecallManagementServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRecallManagementServiceRequest) GetName() *string {
	return s.Name
}

func (s *CreateRecallManagementServiceRequest) SetDescription(v string) *CreateRecallManagementServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateRecallManagementServiceRequest) SetInstanceId(v string) *CreateRecallManagementServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRecallManagementServiceRequest) SetName(v string) *CreateRecallManagementServiceRequest {
	s.Name = &v
	return s
}

func (s *CreateRecallManagementServiceRequest) Validate() error {
	return dara.Validate(s)
}
