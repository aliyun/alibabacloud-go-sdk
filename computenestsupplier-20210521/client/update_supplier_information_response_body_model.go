// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSupplierInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSupplierInformationResponseBody
	GetRequestId() *string
	SetSupplierDesc(v string) *UpdateSupplierInformationResponseBody
	GetSupplierDesc() *string
	SetSupplierName(v string) *UpdateSupplierInformationResponseBody
	GetSupplierName() *string
	SetSupplierUrl(v string) *UpdateSupplierInformationResponseBody
	GetSupplierUrl() *string
}

type UpdateSupplierInformationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The description of service provider.
	//
	// example:
	//
	// Test supplier
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://www.xxx.xxx.cn
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
}

func (s UpdateSupplierInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSupplierInformationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSupplierInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSupplierInformationResponseBody) GetSupplierDesc() *string {
	return s.SupplierDesc
}

func (s *UpdateSupplierInformationResponseBody) GetSupplierName() *string {
	return s.SupplierName
}

func (s *UpdateSupplierInformationResponseBody) GetSupplierUrl() *string {
	return s.SupplierUrl
}

func (s *UpdateSupplierInformationResponseBody) SetRequestId(v string) *UpdateSupplierInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSupplierInformationResponseBody) SetSupplierDesc(v string) *UpdateSupplierInformationResponseBody {
	s.SupplierDesc = &v
	return s
}

func (s *UpdateSupplierInformationResponseBody) SetSupplierName(v string) *UpdateSupplierInformationResponseBody {
	s.SupplierName = &v
	return s
}

func (s *UpdateSupplierInformationResponseBody) SetSupplierUrl(v string) *UpdateSupplierInformationResponseBody {
	s.SupplierUrl = &v
	return s
}

func (s *UpdateSupplierInformationResponseBody) Validate() error {
	return dara.Validate(s)
}
