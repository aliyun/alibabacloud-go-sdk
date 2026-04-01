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
	SetKeyPairName(v string) *AttachKeyPairRequest
	GetKeyPairName() *string
}

type AttachKeyPairRequest struct {
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
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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

func (s *AttachKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *AttachKeyPairRequest) SetInstanceIds(v []*string) *AttachKeyPairRequest {
	s.InstanceIds = v
	return s
}

func (s *AttachKeyPairRequest) SetKeyPairId(v string) *AttachKeyPairRequest {
	s.KeyPairId = &v
	return s
}

func (s *AttachKeyPairRequest) SetKeyPairName(v string) *AttachKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *AttachKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
