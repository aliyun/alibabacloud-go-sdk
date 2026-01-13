// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementServiceVersionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteRecallManagementServiceVersionConfigRequest
	GetInstanceId() *string
}

type DeleteRecallManagementServiceVersionConfigRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteRecallManagementServiceVersionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementServiceVersionConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementServiceVersionConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteRecallManagementServiceVersionConfigRequest) SetInstanceId(v string) *DeleteRecallManagementServiceVersionConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRecallManagementServiceVersionConfigRequest) Validate() error {
	return dara.Validate(s)
}
