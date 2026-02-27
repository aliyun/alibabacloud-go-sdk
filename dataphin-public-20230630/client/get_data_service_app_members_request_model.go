// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int32) *GetDataServiceAppMembersRequest
	GetAppId() *int32
	SetOpTenantId(v int64) *GetDataServiceAppMembersRequest
	GetOpTenantId() *int64
}

type GetDataServiceAppMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetDataServiceAppMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppMembersRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppMembersRequest) GetAppId() *int32 {
	return s.AppId
}

func (s *GetDataServiceAppMembersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceAppMembersRequest) SetAppId(v int32) *GetDataServiceAppMembersRequest {
	s.AppId = &v
	return s
}

func (s *GetDataServiceAppMembersRequest) SetOpTenantId(v int64) *GetDataServiceAppMembersRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceAppMembersRequest) Validate() error {
	return dara.Validate(s)
}
