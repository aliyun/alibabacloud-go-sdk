// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineRecallManagementServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OfflineRecallManagementServiceRequest
	GetInstanceId() *string
}

type OfflineRecallManagementServiceRequest struct {
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OfflineRecallManagementServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineRecallManagementServiceRequest) GoString() string {
	return s.String()
}

func (s *OfflineRecallManagementServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OfflineRecallManagementServiceRequest) SetInstanceId(v string) *OfflineRecallManagementServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *OfflineRecallManagementServiceRequest) Validate() error {
	return dara.Validate(s)
}
