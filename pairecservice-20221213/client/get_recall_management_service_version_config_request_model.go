// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementServiceVersionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRecallManagementServiceVersionConfigRequest
	GetInstanceId() *string
}

type GetRecallManagementServiceVersionConfigRequest struct {
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRecallManagementServiceVersionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionConfigRequest) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRecallManagementServiceVersionConfigRequest) SetInstanceId(v string) *GetRecallManagementServiceVersionConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigRequest) Validate() error {
	return dara.Validate(s)
}
