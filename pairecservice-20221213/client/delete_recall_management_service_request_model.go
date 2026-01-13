// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteRecallManagementServiceRequest
	GetInstanceId() *string
}

type DeleteRecallManagementServiceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteRecallManagementServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteRecallManagementServiceRequest) SetInstanceId(v string) *DeleteRecallManagementServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRecallManagementServiceRequest) Validate() error {
	return dara.Validate(s)
}
