// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleForProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductName(v string) *CheckServiceLinkedRoleForProductRequest
	GetProductName() *string
}

type CheckServiceLinkedRoleForProductRequest struct {
	// example:
	//
	// AliyunServiceRoleForEventBridgeConnectVPC
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s CheckServiceLinkedRoleForProductRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleForProductRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForProductRequest) GetProductName() *string {
	return s.ProductName
}

func (s *CheckServiceLinkedRoleForProductRequest) SetProductName(v string) *CheckServiceLinkedRoleForProductRequest {
	s.ProductName = &v
	return s
}

func (s *CheckServiceLinkedRoleForProductRequest) Validate() error {
	return dara.Validate(s)
}
