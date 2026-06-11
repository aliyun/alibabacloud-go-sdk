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
	// A list of data catalogs.
	CataLogList []*DLCatalog `json:"CataLogList,omitempty" xml:"CataLogList,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, used for log tracing and troubleshooting.
	//
	// example:
	//
	// E76DD2E7-EBAC-5724-B163-19AAC233****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. Valid values:
	//
	// - **true**: The request succeeded.
	//
	// - ******false**: The request failed.
	//
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
	if s.CataLogList != nil {
		for _, item := range s.CataLogList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
