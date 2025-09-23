// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSampleApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *UploadSampleApiRequest
	GetDataType() *string
	SetDataValue(v string) *UploadSampleApiRequest
	GetDataValue() *string
	SetSampleType(v string) *UploadSampleApiRequest
	GetSampleType() *string
	SetService(v string) *UploadSampleApiRequest
	GetService() *string
}

type UploadSampleApiRequest struct {
	// The data type of the sample
	//
	// - Phone number: mobile
	//
	// - MD5 of phone number: mobileMd5
	//
	// - IP: ip
	//
	// - Unique device ID: umid
	//
	// - Account ID: accountId
	//
	// - IMEI: imei
	//
	// - MD5 of IMEI: imeiMd5
	//
	// - OAID: oaid
	//
	// - MD5 of OAID: oaidMd5
	//
	// - Android ID: androidId
	//
	// - MD5 of Android ID: androidIdMd5
	//
	// This parameter is required.
	//
	// example:
	//
	// ip
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Specific value of the sample, to be passed in JSON format. Do not exceed 1000 entries at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["123.124.125.126","123.124.125.127"]
	DataValue *string `json:"DataValue,omitempty" xml:"DataValue,omitempty"`
	// The type of the sample
	//
	// - Blacklist: block
	//
	// - Whitelist: pass
	//
	// This parameter is required.
	//
	// example:
	//
	// block
	SampleType *string `json:"SampleType,omitempty" xml:"SampleType,omitempty"`
	// List of effective services, separate multiple services with commas
	//
	// - Basic/Enhanced Registration Risk Identification: account_abuse
	//
	// - Basic/Enhanced Marketing Risk Identification: coupon_abuse
	//
	// - Basic/Enhanced Login Risk Identification: account_takeover
	//
	// This parameter is required.
	//
	// example:
	//
	// account_abuse,coupon_abuse,account_takeover
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s UploadSampleApiRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadSampleApiRequest) GoString() string {
	return s.String()
}

func (s *UploadSampleApiRequest) GetDataType() *string {
	return s.DataType
}

func (s *UploadSampleApiRequest) GetDataValue() *string {
	return s.DataValue
}

func (s *UploadSampleApiRequest) GetSampleType() *string {
	return s.SampleType
}

func (s *UploadSampleApiRequest) GetService() *string {
	return s.Service
}

func (s *UploadSampleApiRequest) SetDataType(v string) *UploadSampleApiRequest {
	s.DataType = &v
	return s
}

func (s *UploadSampleApiRequest) SetDataValue(v string) *UploadSampleApiRequest {
	s.DataValue = &v
	return s
}

func (s *UploadSampleApiRequest) SetSampleType(v string) *UploadSampleApiRequest {
	s.SampleType = &v
	return s
}

func (s *UploadSampleApiRequest) SetService(v string) *UploadSampleApiRequest {
	s.Service = &v
	return s
}

func (s *UploadSampleApiRequest) Validate() error {
	return dara.Validate(s)
}
