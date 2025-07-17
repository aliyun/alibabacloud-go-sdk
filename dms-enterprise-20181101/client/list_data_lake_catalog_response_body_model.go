// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCataLogList(v []*DLCatalog) *ListDataLakeCatalogResponseBody
	GetCataLogList() []*DLCatalog
	SetErrorCode(v string) *ListDataLakeCatalogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataLakeCatalogResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDataLakeCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataLakeCatalogResponseBody
	GetSuccess() *bool
}

type ListDataLakeCatalogResponseBody struct {
	CataLogList []*DLCatalog `json:"CataLogList,omitempty" xml:"CataLogList,omitempty" type:"Repeated"`
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// E76DD2E7-EBAC-5724-B163-19AAC233F8F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataLakeCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakeCatalogResponseBody) GetCataLogList() []*DLCatalog {
	return s.CataLogList
}

func (s *ListDataLakeCatalogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataLakeCatalogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataLakeCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataLakeCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataLakeCatalogResponseBody) SetCataLogList(v []*DLCatalog) *ListDataLakeCatalogResponseBody {
	s.CataLogList = v
	return s
}

func (s *ListDataLakeCatalogResponseBody) SetErrorCode(v string) *ListDataLakeCatalogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataLakeCatalogResponseBody) SetErrorMessage(v string) *ListDataLakeCatalogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataLakeCatalogResponseBody) SetRequestId(v string) *ListDataLakeCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLakeCatalogResponseBody) SetSuccess(v bool) *ListDataLakeCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataLakeCatalogResponseBody) Validate() error {
	return dara.Validate(s)
}
