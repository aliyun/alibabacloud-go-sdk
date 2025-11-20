// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableAllFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMultiDimTableAllFieldsResponseBody
	GetRequestId() *string
	SetValue(v []*GetMultiDimTableAllFieldsResponseBodyValue) *GetMultiDimTableAllFieldsResponseBody
	GetValue() []*GetMultiDimTableAllFieldsResponseBodyValue
	SetVendorRequestId(v string) *GetMultiDimTableAllFieldsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetMultiDimTableAllFieldsResponseBody
	GetVendorType() *string
}

type GetMultiDimTableAllFieldsResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// []
	Value []*GetMultiDimTableAllFieldsResponseBodyValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetMultiDimTableAllFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiDimTableAllFieldsResponseBody) GetValue() []*GetMultiDimTableAllFieldsResponseBodyValue {
	return s.Value
}

func (s *GetMultiDimTableAllFieldsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetMultiDimTableAllFieldsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetMultiDimTableAllFieldsResponseBody) SetRequestId(v string) *GetMultiDimTableAllFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiDimTableAllFieldsResponseBody) SetValue(v []*GetMultiDimTableAllFieldsResponseBodyValue) *GetMultiDimTableAllFieldsResponseBody {
	s.Value = v
	return s
}

func (s *GetMultiDimTableAllFieldsResponseBody) SetVendorRequestId(v string) *GetMultiDimTableAllFieldsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetMultiDimTableAllFieldsResponseBody) SetVendorType(v string) *GetMultiDimTableAllFieldsResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetMultiDimTableAllFieldsResponseBody) Validate() error {
	if s.Value != nil {
		for _, item := range s.Value {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMultiDimTableAllFieldsResponseBodyValue struct {
	// example:
	//
	// stxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Sheet1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// xxx
	Property map[string]interface{} `json:"Property,omitempty" xml:"Property,omitempty"`
	// example:
	//
	// xxx
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMultiDimTableAllFieldsResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllFieldsResponseBodyValue) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllFieldsResponseBodyValue) GetId() *string {
	return s.Id
}

func (s *GetMultiDimTableAllFieldsResponseBodyValue) GetName() *string {
	return s.Name
}

func (s *GetMultiDimTableAllFieldsResponseBodyValue) GetProperty() map[string]interface{} {
	return s.Property
}

func (s *GetMultiDimTableAllFieldsResponseBodyValue) GetType() *string {
	return s.Type
}

func (s *GetMultiDimTableAllFieldsResponseBodyValue) SetId(v string) *GetMultiDimTableAllFieldsResponseBodyValue {
	s.Id = &v
	return s
}

func (s *GetMultiDimTableAllFieldsResponseBodyValue) SetName(v string) *GetMultiDimTableAllFieldsResponseBodyValue {
	s.Name = &v
	return s
}

func (s *GetMultiDimTableAllFieldsResponseBodyValue) SetProperty(v map[string]interface{}) *GetMultiDimTableAllFieldsResponseBodyValue {
	s.Property = v
	return s
}

func (s *GetMultiDimTableAllFieldsResponseBodyValue) SetType(v string) *GetMultiDimTableAllFieldsResponseBodyValue {
	s.Type = &v
	return s
}

func (s *GetMultiDimTableAllFieldsResponseBodyValue) Validate() error {
	return dara.Validate(s)
}
