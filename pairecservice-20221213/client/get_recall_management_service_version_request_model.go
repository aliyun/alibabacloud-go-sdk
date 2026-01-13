// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRecallManagementServiceVersionRequest
	GetInstanceId() *string
}

type GetRecallManagementServiceVersionRequest struct {
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRecallManagementServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRecallManagementServiceVersionRequest) SetInstanceId(v string) *GetRecallManagementServiceVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRecallManagementServiceVersionRequest) Validate() error {
	return dara.Validate(s)
}
