// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEnsInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AttachEnsInstancesRequest
	GetInstanceId() *string
	SetScripts(v string) *AttachEnsInstancesRequest
	GetScripts() *string
}

type AttachEnsInstancesRequest struct {
	// The instance ID. Only one ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// testInstacneId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The command to run on the instance, encoded in Base64 using UTF-8.
	//
	// This parameter is required.
	//
	// example:
	//
	// wget d2dldCBodHRwOi8vYWxpYWNzLWs4cy1jbxxxx
	Scripts *string `json:"Scripts,omitempty" xml:"Scripts,omitempty"`
}

func (s AttachEnsInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachEnsInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachEnsInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachEnsInstancesRequest) GetScripts() *string {
	return s.Scripts
}

func (s *AttachEnsInstancesRequest) SetInstanceId(v string) *AttachEnsInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachEnsInstancesRequest) SetScripts(v string) *AttachEnsInstancesRequest {
	s.Scripts = &v
	return s
}

func (s *AttachEnsInstancesRequest) Validate() error {
	return dara.Validate(s)
}
