// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDataServiceApiDocumentRequest
	GetId() *int64
	SetOpTenantId(v int64) *GetDataServiceApiDocumentRequest
	GetOpTenantId() *int64
	SetVersionId(v string) *GetDataServiceApiDocumentRequest
	GetVersionId() *string
}

type GetDataServiceApiDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 102101
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetDataServiceApiDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiDocumentRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiDocumentRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDataServiceApiDocumentRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataServiceApiDocumentRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetDataServiceApiDocumentRequest) SetId(v int64) *GetDataServiceApiDocumentRequest {
	s.Id = &v
	return s
}

func (s *GetDataServiceApiDocumentRequest) SetOpTenantId(v int64) *GetDataServiceApiDocumentRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataServiceApiDocumentRequest) SetVersionId(v string) *GetDataServiceApiDocumentRequest {
	s.VersionId = &v
	return s
}

func (s *GetDataServiceApiDocumentRequest) Validate() error {
	return dara.Validate(s)
}
