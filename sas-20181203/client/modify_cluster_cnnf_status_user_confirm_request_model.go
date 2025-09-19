// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterCnnfStatusUserConfirmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v []*string) *ModifyClusterCnnfStatusUserConfirmRequest
	GetClusterIds() []*string
	SetUserConfirm(v bool) *ModifyClusterCnnfStatusUserConfirmRequest
	GetUserConfirm() *bool
}

type ModifyClusterCnnfStatusUserConfirmRequest struct {
	// The cluster IDs.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
	// Specifies whether to fix the blocking status of the cluster. Valid values:
	//
	// 	- true: yes
	//
	// 	- false: no
	//
	// example:
	//
	// true
	UserConfirm *bool `json:"UserConfirm,omitempty" xml:"UserConfirm,omitempty"`
}

func (s ModifyClusterCnnfStatusUserConfirmRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterCnnfStatusUserConfirmRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterCnnfStatusUserConfirmRequest) GetClusterIds() []*string {
	return s.ClusterIds
}

func (s *ModifyClusterCnnfStatusUserConfirmRequest) GetUserConfirm() *bool {
	return s.UserConfirm
}

func (s *ModifyClusterCnnfStatusUserConfirmRequest) SetClusterIds(v []*string) *ModifyClusterCnnfStatusUserConfirmRequest {
	s.ClusterIds = v
	return s
}

func (s *ModifyClusterCnnfStatusUserConfirmRequest) SetUserConfirm(v bool) *ModifyClusterCnnfStatusUserConfirmRequest {
	s.UserConfirm = &v
	return s
}

func (s *ModifyClusterCnnfStatusUserConfirmRequest) Validate() error {
	return dara.Validate(s)
}
