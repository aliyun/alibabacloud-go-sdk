// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRepoTagResponseBody
	GetCode() *string
	SetDigest(v string) *GetRepoTagResponseBody
	GetDigest() *string
	SetImageCreate(v int64) *GetRepoTagResponseBody
	GetImageCreate() *int64
	SetImageId(v string) *GetRepoTagResponseBody
	GetImageId() *string
	SetImageSize(v int64) *GetRepoTagResponseBody
	GetImageSize() *int64
	SetImageUpdate(v int64) *GetRepoTagResponseBody
	GetImageUpdate() *int64
	SetIsSuccess(v bool) *GetRepoTagResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetRepoTagResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetRepoTagResponseBody
	GetStatus() *string
	SetTag(v string) *GetRepoTagResponseBody
	GetTag() *string
}

type GetRepoTagResponseBody struct {
	// The ID of the image.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The size of the image. Unit: Bytes.
	//
	// example:
	//
	// 67bfbcc12b67936ec7f867927817cbb071832b873dbcaed312a1930ba5f1****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// crr-tquyps22md8p****
	//
	// example:
	//
	// 1572839125000
	ImageCreate *int64 `json:"ImageCreate,omitempty" xml:"ImageCreate,omitempty"`
	// example:
	//
	// 45023655bf39c382e26a8607d057c27871dee163c1ecf48cc1ebf2a1****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The number of milliseconds that have elapsed since the image was last updated.
	//
	// example:
	//
	// 27107966
	ImageSize *int64 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1572875608000
	ImageUpdate *int64 `json:"ImageUpdate,omitempty" xml:"ImageUpdate,omitempty"`
	// The status of the image. Valid values:
	//
	// 	- `NORMAL`: The image is normal.
	//
	// 	- `DELETING`: The image is being deleted.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// 1.0
	//
	// example:
	//
	// 031572FA-7D8F-4C05-B790-1071E0E05DE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version of the repository.
	//
	// example:
	//
	// 1.0
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetRepoTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRepoTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRepoTagResponseBody) GetDigest() *string {
	return s.Digest
}

func (s *GetRepoTagResponseBody) GetImageCreate() *int64 {
	return s.ImageCreate
}

func (s *GetRepoTagResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *GetRepoTagResponseBody) GetImageSize() *int64 {
	return s.ImageSize
}

func (s *GetRepoTagResponseBody) GetImageUpdate() *int64 {
	return s.ImageUpdate
}

func (s *GetRepoTagResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetRepoTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRepoTagResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRepoTagResponseBody) GetTag() *string {
	return s.Tag
}

func (s *GetRepoTagResponseBody) SetCode(v string) *GetRepoTagResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoTagResponseBody) SetDigest(v string) *GetRepoTagResponseBody {
	s.Digest = &v
	return s
}

func (s *GetRepoTagResponseBody) SetImageCreate(v int64) *GetRepoTagResponseBody {
	s.ImageCreate = &v
	return s
}

func (s *GetRepoTagResponseBody) SetImageId(v string) *GetRepoTagResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetRepoTagResponseBody) SetImageSize(v int64) *GetRepoTagResponseBody {
	s.ImageSize = &v
	return s
}

func (s *GetRepoTagResponseBody) SetImageUpdate(v int64) *GetRepoTagResponseBody {
	s.ImageUpdate = &v
	return s
}

func (s *GetRepoTagResponseBody) SetIsSuccess(v bool) *GetRepoTagResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoTagResponseBody) SetRequestId(v string) *GetRepoTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoTagResponseBody) SetStatus(v string) *GetRepoTagResponseBody {
	s.Status = &v
	return s
}

func (s *GetRepoTagResponseBody) SetTag(v string) *GetRepoTagResponseBody {
	s.Tag = &v
	return s
}

func (s *GetRepoTagResponseBody) Validate() error {
	return dara.Validate(s)
}
