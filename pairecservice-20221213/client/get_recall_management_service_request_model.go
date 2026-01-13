// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRecallManagementServiceRequest
	GetInstanceId() *string
}

type GetRecallManagementServiceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRecallManagementServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceRequest) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRecallManagementServiceRequest) SetInstanceId(v string) *GetRecallManagementServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRecallManagementServiceRequest) Validate() error {
	return dara.Validate(s)
}
