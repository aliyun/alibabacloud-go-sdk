// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadIoTDataToBlockchainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIotAuthType(v string) *UploadIoTDataToBlockchainRequest
	GetIotAuthType() *string
	SetIotDataDID(v string) *UploadIoTDataToBlockchainRequest
	GetIotDataDID() *string
	SetIotDataDigest(v string) *UploadIoTDataToBlockchainRequest
	GetIotDataDigest() *string
	SetIotDataToken(v string) *UploadIoTDataToBlockchainRequest
	GetIotDataToken() *string
	SetIotId(v string) *UploadIoTDataToBlockchainRequest
	GetIotId() *string
	SetIotIdServiceProvider(v string) *UploadIoTDataToBlockchainRequest
	GetIotIdServiceProvider() *string
	SetIotIdSource(v string) *UploadIoTDataToBlockchainRequest
	GetIotIdSource() *string
	SetPlainData(v string) *UploadIoTDataToBlockchainRequest
	GetPlainData() *string
	SetPrivacyData(v string) *UploadIoTDataToBlockchainRequest
	GetPrivacyData() *string
	SetRegionId(v string) *UploadIoTDataToBlockchainRequest
	GetRegionId() *string
}

type UploadIoTDataToBlockchainRequest struct {
	// This parameter is required.
	IotAuthType *string `json:"IotAuthType,omitempty" xml:"IotAuthType,omitempty"`
	// This parameter is required.
	IotDataDID *string `json:"IotDataDID,omitempty" xml:"IotDataDID,omitempty"`
	// This parameter is required.
	IotDataDigest *string `json:"IotDataDigest,omitempty" xml:"IotDataDigest,omitempty"`
	// This parameter is required.
	IotDataToken *string `json:"IotDataToken,omitempty" xml:"IotDataToken,omitempty"`
	// This parameter is required.
	IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	// This parameter is required.
	IotIdServiceProvider *string `json:"IotIdServiceProvider,omitempty" xml:"IotIdServiceProvider,omitempty"`
	// This parameter is required.
	IotIdSource *string `json:"IotIdSource,omitempty" xml:"IotIdSource,omitempty"`
	// This parameter is required.
	PlainData *string `json:"PlainData,omitempty" xml:"PlainData,omitempty"`
	// This parameter is required.
	PrivacyData *string `json:"PrivacyData,omitempty" xml:"PrivacyData,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UploadIoTDataToBlockchainRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadIoTDataToBlockchainRequest) GoString() string {
	return s.String()
}

func (s *UploadIoTDataToBlockchainRequest) GetIotAuthType() *string {
	return s.IotAuthType
}

func (s *UploadIoTDataToBlockchainRequest) GetIotDataDID() *string {
	return s.IotDataDID
}

func (s *UploadIoTDataToBlockchainRequest) GetIotDataDigest() *string {
	return s.IotDataDigest
}

func (s *UploadIoTDataToBlockchainRequest) GetIotDataToken() *string {
	return s.IotDataToken
}

func (s *UploadIoTDataToBlockchainRequest) GetIotId() *string {
	return s.IotId
}

func (s *UploadIoTDataToBlockchainRequest) GetIotIdServiceProvider() *string {
	return s.IotIdServiceProvider
}

func (s *UploadIoTDataToBlockchainRequest) GetIotIdSource() *string {
	return s.IotIdSource
}

func (s *UploadIoTDataToBlockchainRequest) GetPlainData() *string {
	return s.PlainData
}

func (s *UploadIoTDataToBlockchainRequest) GetPrivacyData() *string {
	return s.PrivacyData
}

func (s *UploadIoTDataToBlockchainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UploadIoTDataToBlockchainRequest) SetIotAuthType(v string) *UploadIoTDataToBlockchainRequest {
	s.IotAuthType = &v
	return s
}

func (s *UploadIoTDataToBlockchainRequest) SetIotDataDID(v string) *UploadIoTDataToBlockchainRequest {
	s.IotDataDID = &v
	return s
}

func (s *UploadIoTDataToBlockchainRequest) SetIotDataDigest(v string) *UploadIoTDataToBlockchainRequest {
	s.IotDataDigest = &v
	return s
}

func (s *UploadIoTDataToBlockchainRequest) SetIotDataToken(v string) *UploadIoTDataToBlockchainRequest {
	s.IotDataToken = &v
	return s
}

func (s *UploadIoTDataToBlockchainRequest) SetIotId(v string) *UploadIoTDataToBlockchainRequest {
	s.IotId = &v
	return s
}

func (s *UploadIoTDataToBlockchainRequest) SetIotIdServiceProvider(v string) *UploadIoTDataToBlockchainRequest {
	s.IotIdServiceProvider = &v
	return s
}

func (s *UploadIoTDataToBlockchainRequest) SetIotIdSource(v string) *UploadIoTDataToBlockchainRequest {
	s.IotIdSource = &v
	return s
}

func (s *UploadIoTDataToBlockchainRequest) SetPlainData(v string) *UploadIoTDataToBlockchainRequest {
	s.PlainData = &v
	return s
}

func (s *UploadIoTDataToBlockchainRequest) SetPrivacyData(v string) *UploadIoTDataToBlockchainRequest {
	s.PrivacyData = &v
	return s
}

func (s *UploadIoTDataToBlockchainRequest) SetRegionId(v string) *UploadIoTDataToBlockchainRequest {
	s.RegionId = &v
	return s
}

func (s *UploadIoTDataToBlockchainRequest) Validate() error {
	return dara.Validate(s)
}
