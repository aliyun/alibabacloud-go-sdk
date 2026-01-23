// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardWordRootRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteStandardWordRootRequest
	GetName() *string
	SetOpTenantId(v int64) *DeleteStandardWordRootRequest
	GetOpTenantId() *int64
}

type DeleteStandardWordRootRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 平均值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteStandardWordRootRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardWordRootRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardWordRootRequest) GetName() *string {
	return s.Name
}

func (s *DeleteStandardWordRootRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteStandardWordRootRequest) SetName(v string) *DeleteStandardWordRootRequest {
	s.Name = &v
	return s
}

func (s *DeleteStandardWordRootRequest) SetOpTenantId(v int64) *DeleteStandardWordRootRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteStandardWordRootRequest) Validate() error {
	return dara.Validate(s)
}
