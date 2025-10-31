// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v *GetImageResponseBodyImage) *GetImageResponseBody
	GetImage() *GetImageResponseBodyImage
	SetRequestId(v string) *GetImageResponseBody
	GetRequestId() *string
}

type GetImageResponseBody struct {
	// The image information.
	Image *GetImageResponseBodyImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) GetImage() *GetImageResponseBodyImage {
	return s.Image
}

func (s *GetImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageResponseBody) SetImage(v *GetImageResponseBodyImage) *GetImageResponseBody {
	s.Image = v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseBody) Validate() error {
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageResponseBodyImage struct {
	// The ID of the backup.
	//
	// example:
	//
	// backup-1618017313
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The time when the image was copied. The value is accurate to the millisecond. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1641275680000
	CopyTime *int64 `json:"CopyTime,omitempty" xml:"CopyTime,omitempty"`
	// The time when the image was generated. The value is accurate to the millisecond. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1786776567788
	ExportTime *int64 `json:"ExportTime,omitempty" xml:"ExportTime,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// image-wz9c5ths5dfuwx47****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the hardware security module (HSM).
	//
	// example:
	//
	// hsm-cn-9lb32vll****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The image generation mode. Valid values:
	//
	// 	- PERIODIC
	//
	// 	- MANUAL
	//
	// example:
	//
	// MANUAL
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the backup.
	//
	// example:
	//
	// hsm-test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the source backup.
	//
	// example:
	//
	// backup-gfuiasdfa****
	SourceBackupUid *string `json:"SourceBackupUid,omitempty" xml:"SourceBackupUid,omitempty"`
	// The ID of the source image.
	//
	// example:
	//
	// image-kklhhhh****
	SourceImageUid *string `json:"SourceImageUid,omitempty" xml:"SourceImageUid,omitempty"`
	// The ID of the source HSM.
	//
	// example:
	//
	// hsm-wz9fnmvx190shfbk****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The region ID of the source image.
	//
	// example:
	//
	// cn-beijing
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The status of the image. Valid values:
	//
	// 	- NEW
	//
	// 	- DELETED
	//
	// 	- CREATING
	//
	// 	- NORMAL
	//
	// example:
	//
	// NEW
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The digest of the HSM.
	//
	// example:
	//
	// 3kGeHnmQzXwSsfF0Jk9eJYhe2gP6An0/HlYIiZh1****
	VsmDigest *string `json:"VsmDigest,omitempty" xml:"VsmDigest,omitempty"`
}

func (s GetImageResponseBodyImage) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyImage) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImage) GetBackupId() *string {
	return s.BackupId
}

func (s *GetImageResponseBodyImage) GetCopyTime() *int64 {
	return s.CopyTime
}

func (s *GetImageResponseBodyImage) GetExportTime() *int64 {
	return s.ExportTime
}

func (s *GetImageResponseBodyImage) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageResponseBodyImage) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetImageResponseBodyImage) GetMode() *string {
	return s.Mode
}

func (s *GetImageResponseBodyImage) GetRegionId() *string {
	return s.RegionId
}

func (s *GetImageResponseBodyImage) GetRemark() *string {
	return s.Remark
}

func (s *GetImageResponseBodyImage) GetSourceBackupUid() *string {
	return s.SourceBackupUid
}

func (s *GetImageResponseBodyImage) GetSourceImageUid() *string {
	return s.SourceImageUid
}

func (s *GetImageResponseBodyImage) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *GetImageResponseBodyImage) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *GetImageResponseBodyImage) GetStatus() *string {
	return s.Status
}

func (s *GetImageResponseBodyImage) GetVsmDigest() *string {
	return s.VsmDigest
}

func (s *GetImageResponseBodyImage) SetBackupId(v string) *GetImageResponseBodyImage {
	s.BackupId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetCopyTime(v int64) *GetImageResponseBodyImage {
	s.CopyTime = &v
	return s
}

func (s *GetImageResponseBodyImage) SetExportTime(v int64) *GetImageResponseBodyImage {
	s.ExportTime = &v
	return s
}

func (s *GetImageResponseBodyImage) SetImageId(v string) *GetImageResponseBodyImage {
	s.ImageId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetInstanceId(v string) *GetImageResponseBodyImage {
	s.InstanceId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetMode(v string) *GetImageResponseBodyImage {
	s.Mode = &v
	return s
}

func (s *GetImageResponseBodyImage) SetRegionId(v string) *GetImageResponseBodyImage {
	s.RegionId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetRemark(v string) *GetImageResponseBodyImage {
	s.Remark = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSourceBackupUid(v string) *GetImageResponseBodyImage {
	s.SourceBackupUid = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSourceImageUid(v string) *GetImageResponseBodyImage {
	s.SourceImageUid = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSourceInstanceId(v string) *GetImageResponseBodyImage {
	s.SourceInstanceId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSourceRegionId(v string) *GetImageResponseBodyImage {
	s.SourceRegionId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetStatus(v string) *GetImageResponseBodyImage {
	s.Status = &v
	return s
}

func (s *GetImageResponseBodyImage) SetVsmDigest(v string) *GetImageResponseBodyImage {
	s.VsmDigest = &v
	return s
}

func (s *GetImageResponseBodyImage) Validate() error {
	return dara.Validate(s)
}
