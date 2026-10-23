// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBTenantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *CreateAgenticDBTenantResponseBody
	GetCreateTime() *string
	SetDBClusterId(v string) *CreateAgenticDBTenantResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *CreateAgenticDBTenantResponseBody
	GetRequestId() *string
	SetTenantId(v string) *CreateAgenticDBTenantResponseBody
	GetTenantId() *string
	SetTenantName(v string) *CreateAgenticDBTenantResponseBody
	GetTenantName() *string
}

type CreateAgenticDBTenantResponseBody struct {
	// The creation time.
	//
	// example:
	//
	// 2026-10-15T08:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The AgenticDB cluster ID.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01A13E8C-3773-1A76-BA95-96EA81D260DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The tenant name.
	//
	// example:
	//
	// my-tenant
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s CreateAgenticDBTenantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBTenantResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBTenantResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAgenticDBTenantResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAgenticDBTenantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgenticDBTenantResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateAgenticDBTenantResponseBody) GetTenantName() *string {
	return s.TenantName
}

func (s *CreateAgenticDBTenantResponseBody) SetCreateTime(v string) *CreateAgenticDBTenantResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateAgenticDBTenantResponseBody) SetDBClusterId(v string) *CreateAgenticDBTenantResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateAgenticDBTenantResponseBody) SetRequestId(v string) *CreateAgenticDBTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgenticDBTenantResponseBody) SetTenantId(v string) *CreateAgenticDBTenantResponseBody {
	s.TenantId = &v
	return s
}

func (s *CreateAgenticDBTenantResponseBody) SetTenantName(v string) *CreateAgenticDBTenantResponseBody {
	s.TenantName = &v
	return s
}

func (s *CreateAgenticDBTenantResponseBody) Validate() error {
	return dara.Validate(s)
}
