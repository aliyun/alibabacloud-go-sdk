// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRecallManagementTableRequest
	GetInstanceId() *string
}

type GetRecallManagementTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRecallManagementTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementTableRequest) GoString() string {
	return s.String()
}

func (s *GetRecallManagementTableRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRecallManagementTableRequest) SetInstanceId(v string) *GetRecallManagementTableRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRecallManagementTableRequest) Validate() error {
	return dara.Validate(s)
}
