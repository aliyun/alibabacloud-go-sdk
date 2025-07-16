// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiDimTableFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateMultiDimTableFieldResponseBody
	GetId() *string
	SetName(v string) *CreateMultiDimTableFieldResponseBody
	GetName() *string
	SetProperty(v map[string]interface{}) *CreateMultiDimTableFieldResponseBody
	GetProperty() map[string]interface{}
	SetType(v string) *CreateMultiDimTableFieldResponseBody
	GetType() *string
	SetRequestId(v string) *CreateMultiDimTableFieldResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *CreateMultiDimTableFieldResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CreateMultiDimTableFieldResponseBody
	GetVendorType() *string
}

type CreateMultiDimTableFieldResponseBody struct {
	// example:
	//
	// UhfysgH
	Id       *string                `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Property map[string]interface{} `json:"Property,omitempty" xml:"Property,omitempty"`
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s CreateMultiDimTableFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiDimTableFieldResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMultiDimTableFieldResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateMultiDimTableFieldResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateMultiDimTableFieldResponseBody) GetProperty() map[string]interface{} {
	return s.Property
}

func (s *CreateMultiDimTableFieldResponseBody) GetType() *string {
	return s.Type
}

func (s *CreateMultiDimTableFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMultiDimTableFieldResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CreateMultiDimTableFieldResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CreateMultiDimTableFieldResponseBody) SetId(v string) *CreateMultiDimTableFieldResponseBody {
	s.Id = &v
	return s
}

func (s *CreateMultiDimTableFieldResponseBody) SetName(v string) *CreateMultiDimTableFieldResponseBody {
	s.Name = &v
	return s
}

func (s *CreateMultiDimTableFieldResponseBody) SetProperty(v map[string]interface{}) *CreateMultiDimTableFieldResponseBody {
	s.Property = v
	return s
}

func (s *CreateMultiDimTableFieldResponseBody) SetType(v string) *CreateMultiDimTableFieldResponseBody {
	s.Type = &v
	return s
}

func (s *CreateMultiDimTableFieldResponseBody) SetRequestId(v string) *CreateMultiDimTableFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMultiDimTableFieldResponseBody) SetVendorRequestId(v string) *CreateMultiDimTableFieldResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CreateMultiDimTableFieldResponseBody) SetVendorType(v string) *CreateMultiDimTableFieldResponseBody {
	s.VendorType = &v
	return s
}

func (s *CreateMultiDimTableFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
