// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectWhiteListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetProjectWhiteListsRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetProjectWhiteListsRequest
	GetOpTenantId() *int64
}

type GetProjectWhiteListsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1030111021
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetProjectWhiteListsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectWhiteListsRequest) GoString() string {
	return s.String()
}

func (s *GetProjectWhiteListsRequest) GetId() *int64 {
	return s.Id
}

func (s *GetProjectWhiteListsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetProjectWhiteListsRequest) SetId(v int64) *GetProjectWhiteListsRequest {
	s.Id = &v
	return s
}

func (s *GetProjectWhiteListsRequest) SetOpTenantId(v int64) *GetProjectWhiteListsRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetProjectWhiteListsRequest) Validate() error {
	return dara.Validate(s)
}
