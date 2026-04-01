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
	SetKeyPairName(v string) *DetachKeyPairRequest
	GetKeyPairName() *string
}

type DetachKeyPairRequest struct {
	// The instance IDs.
	//
	// This parameter is required.
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

func (s *DetachKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DetachKeyPairRequest) SetInstanceIds(v []*string) *DetachKeyPairRequest {
	s.InstanceIds = v
	return s
}

func (s *DetachKeyPairRequest) SetKeyPairId(v string) *DetachKeyPairRequest {
	s.KeyPairId = &v
	return s
}

func (s *DetachKeyPairRequest) SetKeyPairName(v string) *DetachKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *DetachKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
