// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSslCertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSslCertsResponseBody
	GetCode() *string
	SetData(v *ListSslCertsResponseBodyData) *ListSslCertsResponseBody
	GetData() *ListSslCertsResponseBodyData
	SetMessage(v string) *ListSslCertsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSslCertsResponseBody
	GetRequestId() *string
}

type ListSslCertsResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned data
	Data *ListSslCertsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AADF7197-3384-52AF-A2DE-A66696734129
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSslCertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSslCertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSslCertsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSslCertsResponseBody) GetData() *ListSslCertsResponseBodyData {
	return s.Data
}

func (s *ListSslCertsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSslCertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSslCertsResponseBody) SetCode(v string) *ListSslCertsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSslCertsResponseBody) SetData(v *ListSslCertsResponseBodyData) *ListSslCertsResponseBody {
	s.Data = v
	return s
}

func (s *ListSslCertsResponseBody) SetMessage(v string) *ListSslCertsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSslCertsResponseBody) SetRequestId(v string) *ListSslCertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSslCertsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSslCertsResponseBodyData struct {
	// List of certificate information.
	Items []*SslCertMetaInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total count.
	//
	// example:
	//
	// 2
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListSslCertsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSslCertsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSslCertsResponseBodyData) GetItems() []*SslCertMetaInfo {
	return s.Items
}

func (s *ListSslCertsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSslCertsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSslCertsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListSslCertsResponseBodyData) SetItems(v []*SslCertMetaInfo) *ListSslCertsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListSslCertsResponseBodyData) SetPageNumber(v int32) *ListSslCertsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSslCertsResponseBodyData) SetPageSize(v int32) *ListSslCertsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSslCertsResponseBodyData) SetTotalSize(v int32) *ListSslCertsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListSslCertsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
