// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPoolName(v string) *DeletePoolRequest
	GetPoolName() *string
}

type DeletePoolRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
}

func (s DeletePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePoolRequest) GoString() string {
	return s.String()
}

func (s *DeletePoolRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *DeletePoolRequest) SetPoolName(v string) *DeletePoolRequest {
	s.PoolName = &v
	return s
}

func (s *DeletePoolRequest) Validate() error {
	return dara.Validate(s)
}
