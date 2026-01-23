// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardLookupTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteStandardLookupTableRequest
	GetId() *int64
	SetOpTenantId(v int64) *DeleteStandardLookupTableRequest
	GetOpTenantId() *int64
}

type DeleteStandardLookupTableRequest struct {
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

func (s DeleteStandardLookupTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardLookupTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardLookupTableRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteStandardLookupTableRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteStandardLookupTableRequest) SetId(v int64) *DeleteStandardLookupTableRequest {
	s.Id = &v
	return s
}

func (s *DeleteStandardLookupTableRequest) SetOpTenantId(v int64) *DeleteStandardLookupTableRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteStandardLookupTableRequest) Validate() error {
	return dara.Validate(s)
}
