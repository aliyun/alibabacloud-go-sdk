// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOuterAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOuterAccountId(v string) *DeleteOuterAccountRequest
	GetOuterAccountId() *string
	SetOuterAccountType(v string) *DeleteOuterAccountRequest
	GetOuterAccountType() *string
}

type DeleteOuterAccountRequest struct {
	// External Account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OuterAccountId *string `json:"OuterAccountId,omitempty" xml:"OuterAccountId,omitempty"`
	// Type of the external account.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun
	OuterAccountType *string `json:"OuterAccountType,omitempty" xml:"OuterAccountType,omitempty"`
}

func (s DeleteOuterAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOuterAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteOuterAccountRequest) GetOuterAccountId() *string {
	return s.OuterAccountId
}

func (s *DeleteOuterAccountRequest) GetOuterAccountType() *string {
	return s.OuterAccountType
}

func (s *DeleteOuterAccountRequest) SetOuterAccountId(v string) *DeleteOuterAccountRequest {
	s.OuterAccountId = &v
	return s
}

func (s *DeleteOuterAccountRequest) SetOuterAccountType(v string) *DeleteOuterAccountRequest {
	s.OuterAccountType = &v
	return s
}

func (s *DeleteOuterAccountRequest) Validate() error {
	return dara.Validate(s)
}
