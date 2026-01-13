// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateRecallManagementServiceRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateRecallManagementServiceRequest
	GetInstanceId() *string
}

type UpdateRecallManagementServiceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateRecallManagementServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRecallManagementServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRecallManagementServiceRequest) SetDescription(v string) *UpdateRecallManagementServiceRequest {
	s.Description = &v
	return s
}

func (s *UpdateRecallManagementServiceRequest) SetInstanceId(v string) *UpdateRecallManagementServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRecallManagementServiceRequest) Validate() error {
	return dara.Validate(s)
}
