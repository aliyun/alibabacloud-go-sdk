// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitMultipartFileUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InitMultipartFileUploadResponseBody
	GetRequestId() *string
	SetStorageDriver(v string) *InitMultipartFileUploadResponseBody
	GetStorageDriver() *string
	SetUploadKey(v string) *InitMultipartFileUploadResponseBody
	GetUploadKey() *string
	SetVendorRequestId(v string) *InitMultipartFileUploadResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *InitMultipartFileUploadResponseBody
	GetVendorType() *string
}

type InitMultipartFileUploadResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"storageDriver,omitempty" xml:"storageDriver,omitempty"`
	// example:
	//
	// xhy89xxxxx
	UploadKey *string `json:"uploadKey,omitempty" xml:"uploadKey,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s InitMultipartFileUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitMultipartFileUploadResponseBody) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitMultipartFileUploadResponseBody) GetStorageDriver() *string {
	return s.StorageDriver
}

func (s *InitMultipartFileUploadResponseBody) GetUploadKey() *string {
	return s.UploadKey
}

func (s *InitMultipartFileUploadResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *InitMultipartFileUploadResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *InitMultipartFileUploadResponseBody) SetRequestId(v string) *InitMultipartFileUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitMultipartFileUploadResponseBody) SetStorageDriver(v string) *InitMultipartFileUploadResponseBody {
	s.StorageDriver = &v
	return s
}

func (s *InitMultipartFileUploadResponseBody) SetUploadKey(v string) *InitMultipartFileUploadResponseBody {
	s.UploadKey = &v
	return s
}

func (s *InitMultipartFileUploadResponseBody) SetVendorRequestId(v string) *InitMultipartFileUploadResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *InitMultipartFileUploadResponseBody) SetVendorType(v string) *InitMultipartFileUploadResponseBody {
	s.VendorType = &v
	return s
}

func (s *InitMultipartFileUploadResponseBody) Validate() error {
	return dara.Validate(s)
}
