// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DetachKeyPairRequest
	GetInstanceIds() []*string
	SetKeyPairId(v string) *DetachKeyPairRequest
	GetKeyPairId() *string
}

type DetachKeyPairRequest struct {
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

func (s DetachKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairRequest) GoString() string {
	return s.String()
}

func (s *DetachKeyPairRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DetachKeyPairRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *DetachKeyPairRequest) SetInstanceIds(v []*string) *DetachKeyPairRequest {
	s.InstanceIds = v
	return s
}

func (s *DetachKeyPairRequest) SetKeyPairId(v string) *DetachKeyPairRequest {
	s.KeyPairId = &v
	return s
}

func (s *DetachKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
