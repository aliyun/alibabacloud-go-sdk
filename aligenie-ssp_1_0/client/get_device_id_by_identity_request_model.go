// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceIdByIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncodeKey(v string) *GetDeviceIdByIdentityRequest
	GetEncodeKey() *string
	SetEncodeType(v string) *GetDeviceIdByIdentityRequest
	GetEncodeType() *string
	SetIdentityId(v string) *GetDeviceIdByIdentityRequest
	GetIdentityId() *string
	SetIdentityType(v string) *GetDeviceIdByIdentityRequest
	GetIdentityType() *string
	SetProductKey(v string) *GetDeviceIdByIdentityRequest
	GetProductKey() *string
}

type GetDeviceIdByIdentityRequest struct {
	// The value corresponding to the encoding type. Enter the Project ID of the project to which this product belongs. You can view it in the Tmall Genie AI platform console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 129****0946
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. Enter **PROJECT_ID*	- here.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Authentication identifier. Enter the MAC address or the SN value.
	//
	// This parameter is required.
	//
	// example:
	//
	// b4:xx:xx:xx:65:2b
	IdentityId *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	// Device authentication type. Enter **MAC**, **SN**, or **CTEI**.
	//
	// This parameter is required.
	//
	// example:
	//
	// MAC
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The unique product identifier ProductKey, which is a globally unique identity issued by the platform when the product is created in the Tmall Genie AI platform. This parameter is optional when IdentityType is **CTEI**.
	//
	// example:
	//
	// Mm*****XnZ8
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s GetDeviceIdByIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceIdByIdentityRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityRequest) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetDeviceIdByIdentityRequest) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetDeviceIdByIdentityRequest) GetIdentityId() *string {
	return s.IdentityId
}

func (s *GetDeviceIdByIdentityRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *GetDeviceIdByIdentityRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *GetDeviceIdByIdentityRequest) SetEncodeKey(v string) *GetDeviceIdByIdentityRequest {
	s.EncodeKey = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetEncodeType(v string) *GetDeviceIdByIdentityRequest {
	s.EncodeType = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetIdentityId(v string) *GetDeviceIdByIdentityRequest {
	s.IdentityId = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetIdentityType(v string) *GetDeviceIdByIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) SetProductKey(v string) *GetDeviceIdByIdentityRequest {
	s.ProductKey = &v
	return s
}

func (s *GetDeviceIdByIdentityRequest) Validate() error {
	return dara.Validate(s)
}
