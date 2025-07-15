// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *AttachKeyPairRequest
	GetInstanceIds() []*string
	SetKeyPairId(v string) *AttachKeyPairRequest
	GetKeyPairId() *string
}

type AttachKeyPairRequest struct {
	// The IDs of the cloud phone instances. You can specify a maximum of 50 cloud phone instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the ADB key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
}

func (s AttachKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairRequest) GoString() string {
	return s.String()
}

func (s *AttachKeyPairRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *AttachKeyPairRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *AttachKeyPairRequest) SetInstanceIds(v []*string) *AttachKeyPairRequest {
	s.InstanceIds = v
	return s
}

func (s *AttachKeyPairRequest) SetKeyPairId(v string) *AttachKeyPairRequest {
	s.KeyPairId = &v
	return s
}

func (s *AttachKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
