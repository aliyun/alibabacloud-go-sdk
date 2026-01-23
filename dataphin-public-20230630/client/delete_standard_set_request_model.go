// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteStandardSetRequest
	GetId() *int64
	SetOpTenantId(v int64) *DeleteStandardSetRequest
	GetOpTenantId() *int64
}

type DeleteStandardSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteStandardSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardSetRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteStandardSetRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteStandardSetRequest) SetId(v int64) *DeleteStandardSetRequest {
	s.Id = &v
	return s
}

func (s *DeleteStandardSetRequest) SetOpTenantId(v int64) *DeleteStandardSetRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteStandardSetRequest) Validate() error {
	return dara.Validate(s)
}
