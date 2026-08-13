// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeDetails(v bool) *GetSourceRequest
	GetIncludeDetails() *bool
	SetSourceId(v string) *GetSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *GetSourceRequest
	GetTenantId() *string
}

type GetSourceRequest struct {
	// 是否返回大体积明细字段（settings / notes / structuredTables / unstructuredDocs）。默认 False，仅返回元信息。
	//
	// example:
	//
	// false
	IncludeDetails *bool `json:"includeDetails,omitempty" xml:"includeDetails,omitempty"`
	// 数据源 ID（租户内唯一）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSourceRequest) GoString() string {
	return s.String()
}

func (s *GetSourceRequest) GetIncludeDetails() *bool {
	return s.IncludeDetails
}

func (s *GetSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *GetSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSourceRequest) SetIncludeDetails(v bool) *GetSourceRequest {
	s.IncludeDetails = &v
	return s
}

func (s *GetSourceRequest) SetSourceId(v string) *GetSourceRequest {
	s.SourceId = &v
	return s
}

func (s *GetSourceRequest) SetTenantId(v string) *GetSourceRequest {
	s.TenantId = &v
	return s
}

func (s *GetSourceRequest) Validate() error {
	return dara.Validate(s)
}
