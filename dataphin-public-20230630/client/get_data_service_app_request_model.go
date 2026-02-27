// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int32) *GetDataServiceAppRequest
	GetAppId() *int32
	SetOpTenantId(v int64) *GetDataServiceAppRequest
	GetOpTenantId() *int64
}

type GetDataServiceAppRequest struct {
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

func (s GetDataServiceAppRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppRequest) GetAppId() *int32 {
	return s.AppId
}

func (s *GetDataServiceAppRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceAppRequest) SetAppId(v int32) *GetDataServiceAppRequest {
	s.AppId = &v
	return s
}

func (s *GetDataServiceAppRequest) SetOpTenantId(v int64) *GetDataServiceAppRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceAppRequest) Validate() error {
	return dara.Validate(s)
}
