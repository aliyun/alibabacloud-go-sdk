// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachKeyPairShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *DetachKeyPairShrinkRequest
	GetInstanceIdsShrink() *string
	SetKeyPairId(v string) *DetachKeyPairShrinkRequest
	GetKeyPairId() *string
	SetKeyPairName(v string) *DetachKeyPairShrinkRequest
	GetKeyPairName() *string
}

type DetachKeyPairShrinkRequest struct {
	// The instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["i-50f8q9mbfjzlkuk9znjglnne3"]
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The key pair ID.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// ssh-xxx
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the SSH key pair.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// test-kp
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s DetachKeyPairShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachKeyPairShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DetachKeyPairShrinkRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *DetachKeyPairShrinkRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DetachKeyPairShrinkRequest) SetInstanceIdsShrink(v string) *DetachKeyPairShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DetachKeyPairShrinkRequest) SetKeyPairId(v string) *DetachKeyPairShrinkRequest {
	s.KeyPairId = &v
	return s
}

func (s *DetachKeyPairShrinkRequest) SetKeyPairName(v string) *DetachKeyPairShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *DetachKeyPairShrinkRequest) Validate() error {
	return dara.Validate(s)
}
