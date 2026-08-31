// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityArchiveTableProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetQualityArchiveTableProgressRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetQualityArchiveTableProgressRequest
	GetOpUserId() *string
	SetProgressId(v string) *GetQualityArchiveTableProgressRequest
	GetProgressId() *string
}

type GetQualityArchiveTableProgressRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The asynchronous task progress ID returned by the UpsertQualityArchiveTable operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// d78f0b5c9a1e4f2ab3c6d5e4f7a8b9c0
	ProgressId *string `json:"ProgressId,omitempty" xml:"ProgressId,omitempty"`
}

func (s GetQualityArchiveTableProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityArchiveTableProgressRequest) GoString() string {
	return s.String()
}

func (s *GetQualityArchiveTableProgressRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualityArchiveTableProgressRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetQualityArchiveTableProgressRequest) GetProgressId() *string {
	return s.ProgressId
}

func (s *GetQualityArchiveTableProgressRequest) SetOpTenantId(v int64) *GetQualityArchiveTableProgressRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualityArchiveTableProgressRequest) SetOpUserId(v string) *GetQualityArchiveTableProgressRequest {
	s.OpUserId = &v
	return s
}

func (s *GetQualityArchiveTableProgressRequest) SetProgressId(v string) *GetQualityArchiveTableProgressRequest {
	s.ProgressId = &v
	return s
}

func (s *GetQualityArchiveTableProgressRequest) Validate() error {
	return dara.Validate(s)
}
