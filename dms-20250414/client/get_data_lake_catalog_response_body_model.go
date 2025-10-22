// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v *DLCatalog) *GetDataLakeCatalogResponseBody
	GetCatalog() *DLCatalog
	SetErrorCode(v string) *GetDataLakeCatalogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataLakeCatalogResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataLakeCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataLakeCatalogResponseBody
	GetSuccess() *bool
}

type GetDataLakeCatalogResponseBody struct {
	Catalog *DLCatalog `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// code: 404, can not find catalog, name : hiv request id: 6090E571-E5B1-1E6D-BF44-F9E10E8B****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// FE8EE2F1-4880-46BC-A704-5CF63EAF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataLakeCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataLakeCatalogResponseBody) GetCatalog() *DLCatalog {
	return s.Catalog
}

func (s *GetDataLakeCatalogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataLakeCatalogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataLakeCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataLakeCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataLakeCatalogResponseBody) SetCatalog(v *DLCatalog) *GetDataLakeCatalogResponseBody {
	s.Catalog = v
	return s
}

func (s *GetDataLakeCatalogResponseBody) SetErrorCode(v string) *GetDataLakeCatalogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataLakeCatalogResponseBody) SetErrorMessage(v string) *GetDataLakeCatalogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataLakeCatalogResponseBody) SetRequestId(v string) *GetDataLakeCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataLakeCatalogResponseBody) SetSuccess(v bool) *GetDataLakeCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataLakeCatalogResponseBody) Validate() error {
	if s.Catalog != nil {
		if err := s.Catalog.Validate(); err != nil {
			return err
		}
	}
	return nil
}
