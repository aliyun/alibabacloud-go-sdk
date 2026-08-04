// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBindsByPkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *QueryBindsByPkRequest
	GetAppName() *string
	SetPk(v string) *QueryBindsByPkRequest
	GetPk() *string
	SetTenantIds(v map[string]interface{}) *QueryBindsByPkRequest
	GetTenantIds() map[string]interface{}
}

type QueryBindsByPkRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	TenantIds map[string]interface{} `json:"TenantIds,omitempty" xml:"TenantIds,omitempty"`
}

func (s QueryBindsByPkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBindsByPkRequest) GoString() string {
	return s.String()
}

func (s *QueryBindsByPkRequest) GetAppName() *string {
	return s.AppName
}

func (s *QueryBindsByPkRequest) GetPk() *string {
	return s.Pk
}

func (s *QueryBindsByPkRequest) GetTenantIds() map[string]interface{} {
	return s.TenantIds
}

func (s *QueryBindsByPkRequest) SetAppName(v string) *QueryBindsByPkRequest {
	s.AppName = &v
	return s
}

func (s *QueryBindsByPkRequest) SetPk(v string) *QueryBindsByPkRequest {
	s.Pk = &v
	return s
}

func (s *QueryBindsByPkRequest) SetTenantIds(v map[string]interface{}) *QueryBindsByPkRequest {
	s.TenantIds = v
	return s
}

func (s *QueryBindsByPkRequest) Validate() error {
	return dara.Validate(s)
}
