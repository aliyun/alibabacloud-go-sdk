// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int32) *DeleteDataServiceAppRequest
	GetAppId() *int32
	SetOpTenantId(v int64) *DeleteDataServiceAppRequest
	GetOpTenantId() *int64
}

type DeleteDataServiceAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// -535093682933
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteDataServiceAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceAppRequest) GetAppId() *int32 {
	return s.AppId
}

func (s *DeleteDataServiceAppRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteDataServiceAppRequest) SetAppId(v int32) *DeleteDataServiceAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteDataServiceAppRequest) SetOpTenantId(v int64) *DeleteDataServiceAppRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteDataServiceAppRequest) Validate() error {
	return dara.Validate(s)
}
