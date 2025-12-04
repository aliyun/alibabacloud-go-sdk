// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeVccRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *InitializeVccRequest
	GetResourceGroupId() *string
}

type InitializeVccRequest struct {
	// The resource group ID.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-acfmxhucx5ewuwy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s InitializeVccRequest) String() string {
	return dara.Prettify(s)
}

func (s InitializeVccRequest) GoString() string {
	return s.String()
}

func (s *InitializeVccRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *InitializeVccRequest) SetResourceGroupId(v string) *InitializeVccRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *InitializeVccRequest) Validate() error {
	return dara.Validate(s)
}
