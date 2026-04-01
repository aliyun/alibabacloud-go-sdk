// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachKeyPairShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *AttachKeyPairShrinkRequest
	GetInstanceIdsShrink() *string
	SetKeyPairId(v string) *AttachKeyPairShrinkRequest
	GetKeyPairId() *string
	SetKeyPairName(v string) *AttachKeyPairShrinkRequest
	GetKeyPairName() *string
}

type AttachKeyPairShrinkRequest struct {
	// The instance IDs.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
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

func (s AttachKeyPairShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachKeyPairShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *AttachKeyPairShrinkRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *AttachKeyPairShrinkRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *AttachKeyPairShrinkRequest) SetInstanceIdsShrink(v string) *AttachKeyPairShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *AttachKeyPairShrinkRequest) SetKeyPairId(v string) *AttachKeyPairShrinkRequest {
	s.KeyPairId = &v
	return s
}

func (s *AttachKeyPairShrinkRequest) SetKeyPairName(v string) *AttachKeyPairShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *AttachKeyPairShrinkRequest) Validate() error {
	return dara.Validate(s)
}
