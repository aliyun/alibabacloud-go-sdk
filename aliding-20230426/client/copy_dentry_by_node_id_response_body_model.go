// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryByNodeIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *CopyDentryByNodeIdResponseBody
	GetContentType() *string
	SetCreatedTime(v int64) *CopyDentryByNodeIdResponseBody
	GetCreatedTime() *int64
	SetDentryUuid(v string) *CopyDentryByNodeIdResponseBody
	GetDentryUuid() *string
	SetExtension(v string) *CopyDentryByNodeIdResponseBody
	GetExtension() *string
	SetRequestId(v string) *CopyDentryByNodeIdResponseBody
	GetRequestId() *string
	SetSpaceId(v string) *CopyDentryByNodeIdResponseBody
	GetSpaceId() *string
	SetVendorRequestId(v string) *CopyDentryByNodeIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CopyDentryByNodeIdResponseBody
	GetVendorType() *string
}

type CopyDentryByNodeIdResponseBody struct {
	// example:
	//
	// doc
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// example:
	//
	// 12345678
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// cdefg
	DentryUuid *string `json:"dentryUuid,omitempty" xml:"dentryUuid,omitempty"`
	// example:
	//
	// alidoc
	Extension *string `json:"extension,omitempty" xml:"extension,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// bcd
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s CopyDentryByNodeIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryByNodeIdResponseBody) GoString() string {
	return s.String()
}

func (s *CopyDentryByNodeIdResponseBody) GetContentType() *string {
	return s.ContentType
}

func (s *CopyDentryByNodeIdResponseBody) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *CopyDentryByNodeIdResponseBody) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *CopyDentryByNodeIdResponseBody) GetExtension() *string {
	return s.Extension
}

func (s *CopyDentryByNodeIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyDentryByNodeIdResponseBody) GetSpaceId() *string {
	return s.SpaceId
}

func (s *CopyDentryByNodeIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CopyDentryByNodeIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CopyDentryByNodeIdResponseBody) SetContentType(v string) *CopyDentryByNodeIdResponseBody {
	s.ContentType = &v
	return s
}

func (s *CopyDentryByNodeIdResponseBody) SetCreatedTime(v int64) *CopyDentryByNodeIdResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CopyDentryByNodeIdResponseBody) SetDentryUuid(v string) *CopyDentryByNodeIdResponseBody {
	s.DentryUuid = &v
	return s
}

func (s *CopyDentryByNodeIdResponseBody) SetExtension(v string) *CopyDentryByNodeIdResponseBody {
	s.Extension = &v
	return s
}

func (s *CopyDentryByNodeIdResponseBody) SetRequestId(v string) *CopyDentryByNodeIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyDentryByNodeIdResponseBody) SetSpaceId(v string) *CopyDentryByNodeIdResponseBody {
	s.SpaceId = &v
	return s
}

func (s *CopyDentryByNodeIdResponseBody) SetVendorRequestId(v string) *CopyDentryByNodeIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CopyDentryByNodeIdResponseBody) SetVendorType(v string) *CopyDentryByNodeIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *CopyDentryByNodeIdResponseBody) Validate() error {
	return dara.Validate(s)
}
