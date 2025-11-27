// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPoolName(v string) *GetPoolRequest
	GetPoolName() *string
}

type GetPoolRequest struct {
	// The name of the resource pool.
	//
	// 	- The value can be up to 15 characters in length.
	//
	// 	- It can contain digits, uppercase letters, lowercase letters, underscores (_), and dots (.).
	//
	// This parameter is required.
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
}

func (s GetPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPoolRequest) GoString() string {
	return s.String()
}

func (s *GetPoolRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *GetPoolRequest) SetPoolName(v string) *GetPoolRequest {
	s.PoolName = &v
	return s
}

func (s *GetPoolRequest) Validate() error {
	return dara.Validate(s)
}
