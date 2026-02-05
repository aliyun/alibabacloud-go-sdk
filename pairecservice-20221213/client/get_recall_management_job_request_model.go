// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRecallManagementJobRequest
	GetInstanceId() *string
}

type GetRecallManagementJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-1324***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRecallManagementJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementJobRequest) GoString() string {
	return s.String()
}

func (s *GetRecallManagementJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRecallManagementJobRequest) SetInstanceId(v string) *GetRecallManagementJobRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRecallManagementJobRequest) Validate() error {
	return dara.Validate(s)
}
