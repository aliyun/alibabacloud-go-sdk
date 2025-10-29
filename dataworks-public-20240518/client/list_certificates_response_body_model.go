// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListCertificatesResponseBodyPagingInfo) *ListCertificatesResponseBody
	GetPagingInfo() *ListCertificatesResponseBodyPagingInfo
	SetRequestId(v string) *ListCertificatesResponseBody
	GetRequestId() *string
}

type ListCertificatesResponseBody struct {
	// The pagination information.
	PagingInfo *ListCertificatesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ecb967ec-c137-48****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBody) GetPagingInfo() *ListCertificatesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCertificatesResponseBody) SetPagingInfo(v *ListCertificatesResponseBodyPagingInfo) *ListCertificatesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListCertificatesResponseBody) SetRequestId(v string) *ListCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertificatesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCertificatesResponseBodyPagingInfo struct {
	// The certificate files.
	Certificates []*ListCertificatesResponseBodyPagingInfoCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertificatesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBodyPagingInfo) GetCertificates() []*ListCertificatesResponseBodyPagingInfoCertificates {
	return s.Certificates
}

func (s *ListCertificatesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCertificatesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCertificatesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCertificatesResponseBodyPagingInfo) SetCertificates(v []*ListCertificatesResponseBodyPagingInfoCertificates) *ListCertificatesResponseBodyPagingInfo {
	s.Certificates = v
	return s
}

func (s *ListCertificatesResponseBodyPagingInfo) SetPageNumber(v int32) *ListCertificatesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListCertificatesResponseBodyPagingInfo) SetPageSize(v int32) *ListCertificatesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListCertificatesResponseBodyPagingInfo) SetTotalCount(v int32) *ListCertificatesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCertificatesResponseBodyPagingInfo) Validate() error {
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCertificatesResponseBodyPagingInfoCertificates struct {
	// The time when the certificate file was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1730217600000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the certificate file.
	//
	// example:
	//
	// 1107550004253538
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a file
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The size of the certificate file, in bytes.
	//
	// example:
	//
	// 1024
	FileSizeInBytes *int64 `json:"FileSizeInBytes,omitempty" xml:"FileSizeInBytes,omitempty"`
	// The ID of the certificate file.
	//
	// example:
	//
	// 676303114031776
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the certificate file.
	//
	// example:
	//
	// ca1.crt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListCertificatesResponseBodyPagingInfoCertificates) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesResponseBodyPagingInfoCertificates) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) GetDescription() *string {
	return s.Description
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) GetFileSizeInBytes() *int64 {
	return s.FileSizeInBytes
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) GetId() *int64 {
	return s.Id
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) GetName() *string {
	return s.Name
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) SetCreateTime(v int64) *ListCertificatesResponseBodyPagingInfoCertificates {
	s.CreateTime = &v
	return s
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) SetCreateUser(v string) *ListCertificatesResponseBodyPagingInfoCertificates {
	s.CreateUser = &v
	return s
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) SetDescription(v string) *ListCertificatesResponseBodyPagingInfoCertificates {
	s.Description = &v
	return s
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) SetFileSizeInBytes(v int64) *ListCertificatesResponseBodyPagingInfoCertificates {
	s.FileSizeInBytes = &v
	return s
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) SetId(v int64) *ListCertificatesResponseBodyPagingInfoCertificates {
	s.Id = &v
	return s
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) SetName(v string) *ListCertificatesResponseBodyPagingInfoCertificates {
	s.Name = &v
	return s
}

func (s *ListCertificatesResponseBodyPagingInfoCertificates) Validate() error {
	return dara.Validate(s)
}
