// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteRecallManagementServiceVersionRequest
	GetInstanceId() *string
}

type DeleteRecallManagementServiceVersionRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteRecallManagementServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementServiceVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteRecallManagementServiceVersionRequest) SetInstanceId(v string) *DeleteRecallManagementServiceVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRecallManagementServiceVersionRequest) Validate() error {
	return dara.Validate(s)
}
