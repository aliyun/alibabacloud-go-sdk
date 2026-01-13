// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecallManagementTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteRecallManagementTableRequest
	GetInstanceId() *string
}

type DeleteRecallManagementTableRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteRecallManagementTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecallManagementTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecallManagementTableRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteRecallManagementTableRequest) SetInstanceId(v string) *DeleteRecallManagementTableRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRecallManagementTableRequest) Validate() error {
	return dara.Validate(s)
}
