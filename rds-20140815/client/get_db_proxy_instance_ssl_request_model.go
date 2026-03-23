// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDbProxyInstanceSslRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBProxyEngineType(v string) *GetDbProxyInstanceSslRequest
	GetDBProxyEngineType() *string
	SetDbInstanceId(v string) *GetDbProxyInstanceSslRequest
	GetDbInstanceId() *string
	SetRegionId(v string) *GetDbProxyInstanceSslRequest
	GetRegionId() *string
}

type GetDbProxyInstanceSslRequest struct {
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// This parameter is required.
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDbProxyInstanceSslRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDbProxyInstanceSslRequest) GoString() string {
	return s.String()
}

func (s *GetDbProxyInstanceSslRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *GetDbProxyInstanceSslRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *GetDbProxyInstanceSslRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDbProxyInstanceSslRequest) SetDBProxyEngineType(v string) *GetDbProxyInstanceSslRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *GetDbProxyInstanceSslRequest) SetDbInstanceId(v string) *GetDbProxyInstanceSslRequest {
	s.DbInstanceId = &v
	return s
}

func (s *GetDbProxyInstanceSslRequest) SetRegionId(v string) *GetDbProxyInstanceSslRequest {
	s.RegionId = &v
	return s
}

func (s *GetDbProxyInstanceSslRequest) Validate() error {
	return dara.Validate(s)
}
