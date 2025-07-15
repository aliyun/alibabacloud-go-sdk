// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeyPairsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairIds(v []*string) *DeleteKeyPairsRequest
	GetKeyPairIds() []*string
}

type DeleteKeyPairsRequest struct {
	// The IDs of the ADB key pairs.
	KeyPairIds []*string `json:"KeyPairIds,omitempty" xml:"KeyPairIds,omitempty" type:"Repeated"`
}

func (s DeleteKeyPairsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsRequest) GetKeyPairIds() []*string {
	return s.KeyPairIds
}

func (s *DeleteKeyPairsRequest) SetKeyPairIds(v []*string) *DeleteKeyPairsRequest {
	s.KeyPairIds = v
	return s
}

func (s *DeleteKeyPairsRequest) Validate() error {
	return dara.Validate(s)
}
