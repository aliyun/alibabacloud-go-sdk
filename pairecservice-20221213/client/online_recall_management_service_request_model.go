// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineRecallManagementServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnlineRecallManagementServiceRequest
	GetInstanceId() *string
}

type OnlineRecallManagementServiceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnlineRecallManagementServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineRecallManagementServiceRequest) GoString() string {
	return s.String()
}

func (s *OnlineRecallManagementServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnlineRecallManagementServiceRequest) SetInstanceId(v string) *OnlineRecallManagementServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *OnlineRecallManagementServiceRequest) Validate() error {
	return dara.Validate(s)
}
