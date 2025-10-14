// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdDataBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeColdDataBasicInfoResponseBodyData) *DescribeColdDataBasicInfoResponseBody
	GetData() *DescribeColdDataBasicInfoResponseBodyData
	SetRequestId(v string) *DescribeColdDataBasicInfoResponseBody
	GetRequestId() *string
}

type DescribeColdDataBasicInfoResponseBody struct {
	Data      *DescribeColdDataBasicInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColdDataBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdDataBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColdDataBasicInfoResponseBody) GetData() *DescribeColdDataBasicInfoResponseBodyData {
	return s.Data
}

func (s *DescribeColdDataBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeColdDataBasicInfoResponseBody) SetData(v *DescribeColdDataBasicInfoResponseBodyData) *DescribeColdDataBasicInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBody) SetRequestId(v string) *DescribeColdDataBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeColdDataBasicInfoResponseBodyData struct {
	BackupSetCount     *int32   `json:"BackupSetCount,omitempty" xml:"BackupSetCount,omitempty"`
	BackupSetSpaceSize *float64 `json:"BackupSetSpaceSize,omitempty" xml:"BackupSetSpaceSize,omitempty"`
	CloudProduct       *string  `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	CurrentSpaceSize   *float64 `json:"CurrentSpaceSize,omitempty" xml:"CurrentSpaceSize,omitempty"`
	DataRedundancyType *string  `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	EnableStatus       *bool    `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	ReadAccessNum      *int64   `json:"ReadAccessNum,omitempty" xml:"ReadAccessNum,omitempty"`
	RegionId           *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VolumeName         *string  `json:"VolumeName,omitempty" xml:"VolumeName,omitempty"`
	WriteAccessNum     *float64 `json:"WriteAccessNum,omitempty" xml:"WriteAccessNum,omitempty"`
}

func (s DescribeColdDataBasicInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdDataBasicInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeColdDataBasicInfoResponseBodyData) GetBackupSetCount() *int32 {
	return s.BackupSetCount
}

func (s *DescribeColdDataBasicInfoResponseBodyData) GetBackupSetSpaceSize() *float64 {
	return s.BackupSetSpaceSize
}

func (s *DescribeColdDataBasicInfoResponseBodyData) GetCloudProduct() *string {
	return s.CloudProduct
}

func (s *DescribeColdDataBasicInfoResponseBodyData) GetCurrentSpaceSize() *float64 {
	return s.CurrentSpaceSize
}

func (s *DescribeColdDataBasicInfoResponseBodyData) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *DescribeColdDataBasicInfoResponseBodyData) GetEnableStatus() *bool {
	return s.EnableStatus
}

func (s *DescribeColdDataBasicInfoResponseBodyData) GetReadAccessNum() *int64 {
	return s.ReadAccessNum
}

func (s *DescribeColdDataBasicInfoResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeColdDataBasicInfoResponseBodyData) GetVolumeName() *string {
	return s.VolumeName
}

func (s *DescribeColdDataBasicInfoResponseBodyData) GetWriteAccessNum() *float64 {
	return s.WriteAccessNum
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetBackupSetCount(v int32) *DescribeColdDataBasicInfoResponseBodyData {
	s.BackupSetCount = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetBackupSetSpaceSize(v float64) *DescribeColdDataBasicInfoResponseBodyData {
	s.BackupSetSpaceSize = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetCloudProduct(v string) *DescribeColdDataBasicInfoResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetCurrentSpaceSize(v float64) *DescribeColdDataBasicInfoResponseBodyData {
	s.CurrentSpaceSize = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetDataRedundancyType(v string) *DescribeColdDataBasicInfoResponseBodyData {
	s.DataRedundancyType = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetEnableStatus(v bool) *DescribeColdDataBasicInfoResponseBodyData {
	s.EnableStatus = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetReadAccessNum(v int64) *DescribeColdDataBasicInfoResponseBodyData {
	s.ReadAccessNum = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetRegionId(v string) *DescribeColdDataBasicInfoResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetVolumeName(v string) *DescribeColdDataBasicInfoResponseBodyData {
	s.VolumeName = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) SetWriteAccessNum(v float64) *DescribeColdDataBasicInfoResponseBodyData {
	s.WriteAccessNum = &v
	return s
}

func (s *DescribeColdDataBasicInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
