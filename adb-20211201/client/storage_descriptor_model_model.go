// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStorageDescriptorModel interface {
	dara.Model
	String() string
	GoString() string
	SetCompressed(v bool) *StorageDescriptorModel
	GetCompressed() *bool
	SetInputFormat(v string) *StorageDescriptorModel
	GetInputFormat() *string
	SetLocation(v string) *StorageDescriptorModel
	GetLocation() *string
	SetNumBuckets(v int64) *StorageDescriptorModel
	GetNumBuckets() *int64
	SetOutputFormat(v string) *StorageDescriptorModel
	GetOutputFormat() *string
	SetParameters(v map[string]*string) *StorageDescriptorModel
	GetParameters() map[string]*string
	SetSdId(v int64) *StorageDescriptorModel
	GetSdId() *int64
	SetSerDeInfo(v *SerDeInfoModel) *StorageDescriptorModel
	GetSerDeInfo() *SerDeInfoModel
	SetStoredAsSubDirectories(v bool) *StorageDescriptorModel
	GetStoredAsSubDirectories() *bool
}

type StorageDescriptorModel struct {
	Compressed             *bool              `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	InputFormat            *string            `json:"InputFormat,omitempty" xml:"InputFormat,omitempty"`
	Location               *string            `json:"Location,omitempty" xml:"Location,omitempty"`
	NumBuckets             *int64             `json:"NumBuckets,omitempty" xml:"NumBuckets,omitempty"`
	OutputFormat           *string            `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	Parameters             map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SdId                   *int64             `json:"SdId,omitempty" xml:"SdId,omitempty"`
	SerDeInfo              *SerDeInfoModel    `json:"SerDeInfo,omitempty" xml:"SerDeInfo,omitempty"`
	StoredAsSubDirectories *bool              `json:"StoredAsSubDirectories,omitempty" xml:"StoredAsSubDirectories,omitempty"`
}

func (s StorageDescriptorModel) String() string {
	return dara.Prettify(s)
}

func (s StorageDescriptorModel) GoString() string {
	return s.String()
}

func (s *StorageDescriptorModel) GetCompressed() *bool {
	return s.Compressed
}

func (s *StorageDescriptorModel) GetInputFormat() *string {
	return s.InputFormat
}

func (s *StorageDescriptorModel) GetLocation() *string {
	return s.Location
}

func (s *StorageDescriptorModel) GetNumBuckets() *int64 {
	return s.NumBuckets
}

func (s *StorageDescriptorModel) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *StorageDescriptorModel) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *StorageDescriptorModel) GetSdId() *int64 {
	return s.SdId
}

func (s *StorageDescriptorModel) GetSerDeInfo() *SerDeInfoModel {
	return s.SerDeInfo
}

func (s *StorageDescriptorModel) GetStoredAsSubDirectories() *bool {
	return s.StoredAsSubDirectories
}

func (s *StorageDescriptorModel) SetCompressed(v bool) *StorageDescriptorModel {
	s.Compressed = &v
	return s
}

func (s *StorageDescriptorModel) SetInputFormat(v string) *StorageDescriptorModel {
	s.InputFormat = &v
	return s
}

func (s *StorageDescriptorModel) SetLocation(v string) *StorageDescriptorModel {
	s.Location = &v
	return s
}

func (s *StorageDescriptorModel) SetNumBuckets(v int64) *StorageDescriptorModel {
	s.NumBuckets = &v
	return s
}

func (s *StorageDescriptorModel) SetOutputFormat(v string) *StorageDescriptorModel {
	s.OutputFormat = &v
	return s
}

func (s *StorageDescriptorModel) SetParameters(v map[string]*string) *StorageDescriptorModel {
	s.Parameters = v
	return s
}

func (s *StorageDescriptorModel) SetSdId(v int64) *StorageDescriptorModel {
	s.SdId = &v
	return s
}

func (s *StorageDescriptorModel) SetSerDeInfo(v *SerDeInfoModel) *StorageDescriptorModel {
	s.SerDeInfo = v
	return s
}

func (s *StorageDescriptorModel) SetStoredAsSubDirectories(v bool) *StorageDescriptorModel {
	s.StoredAsSubDirectories = &v
	return s
}

func (s *StorageDescriptorModel) Validate() error {
	return dara.Validate(s)
}
