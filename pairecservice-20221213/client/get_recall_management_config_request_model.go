// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRecallManagementConfigRequest
	GetInstanceId() *string
}

type GetRecallManagementConfigRequest struct {
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRecallManagementConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementConfigRequest) GoString() string {
	return s.String()
}

func (s *GetRecallManagementConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRecallManagementConfigRequest) SetInstanceId(v string) *GetRecallManagementConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRecallManagementConfigRequest) Validate() error {
	return dara.Validate(s)
}
