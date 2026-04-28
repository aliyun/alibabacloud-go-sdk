// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPartInfo interface {
	dara.Model
	String() string
	GoString() string
	SetEtag(v string) *UploadPartInfo
	GetEtag() *string
	SetInternalUploadUrl(v string) *UploadPartInfo
	GetInternalUploadUrl() *string
	SetParallelSha1Ctx(v *UploadPartInfoParallelSha1Ctx) *UploadPartInfo
	GetParallelSha1Ctx() *UploadPartInfoParallelSha1Ctx
	SetParallelSha256Ctx(v *UploadPartInfoParallelSha256Ctx) *UploadPartInfo
	GetParallelSha256Ctx() *UploadPartInfoParallelSha256Ctx
	SetPartNumber(v int32) *UploadPartInfo
	GetPartNumber() *int32
	SetPartSize(v int64) *UploadPartInfo
	GetPartSize() *int64
	SetUploadUrl(v string) *UploadPartInfo
	GetUploadUrl() *string
}

type UploadPartInfo struct {
	// This parameter is discontinued.
	//
	// example:
	//
	// "0CC175B9C0F1B6A831C399E269772661"
	Etag *string `json:"etag,omitempty" xml:"etag,omitempty"`
	// The internal upload URL that is used for internal access over a virtual private cloud (VPC). If the intelligent domain name feature is enabled, this parameter is not required. This parameter is returned in the upload_url parameter based on the request. If you want to use this parameter, contact Photo and Drive Service (PDS) technical support.
	//
	// example:
	//
	// https://data-vpc.aliyunpds.com/xxx/xxx?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx&partNumber=1&uploadId=0CC175B9C0F1B6A831C399E269772661
	InternalUploadUrl *string `json:"internal_upload_url,omitempty" xml:"internal_upload_url,omitempty"`
	// The Secure Hash Algorithm 1 (SHA-1) context of the previous part. This parameter takes effect only if the parallel upload feature is enabled.
	ParallelSha1Ctx *UploadPartInfoParallelSha1Ctx `json:"parallel_sha1_ctx,omitempty" xml:"parallel_sha1_ctx,omitempty" type:"Struct"`
	// The SHA-256 context of the previous part.
	ParallelSha256Ctx *UploadPartInfoParallelSha256Ctx `json:"parallel_sha256_ctx,omitempty" xml:"parallel_sha256_ctx,omitempty" type:"Struct"`
	// The serial number of the file part.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PartNumber *int32 `json:"part_number,omitempty" xml:"part_number,omitempty"`
	// This parameter is discontinued.
	//
	// example:
	//
	// 1024
	PartSize *int64 `json:"part_size,omitempty" xml:"part_size,omitempty"`
	// The upload URL. By default, the validity period of the URL is 15 minutes. If the URL expires, you must call the GetUploadUrl operation to obtain another URL. If the intelligent domain name feature is enabled, the internal_upload_url value is returned within the parameter based on the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://data.aliyunpds.com/xxx/xxx?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx&partNumber=1&uploadId=0CC175B9C0F1B6A831C399E269772661
	UploadUrl *string `json:"upload_url,omitempty" xml:"upload_url,omitempty"`
}

func (s UploadPartInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadPartInfo) GoString() string {
	return s.String()
}

func (s *UploadPartInfo) GetEtag() *string {
	return s.Etag
}

func (s *UploadPartInfo) GetInternalUploadUrl() *string {
	return s.InternalUploadUrl
}

func (s *UploadPartInfo) GetParallelSha1Ctx() *UploadPartInfoParallelSha1Ctx {
	return s.ParallelSha1Ctx
}

func (s *UploadPartInfo) GetParallelSha256Ctx() *UploadPartInfoParallelSha256Ctx {
	return s.ParallelSha256Ctx
}

func (s *UploadPartInfo) GetPartNumber() *int32 {
	return s.PartNumber
}

func (s *UploadPartInfo) GetPartSize() *int64 {
	return s.PartSize
}

func (s *UploadPartInfo) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *UploadPartInfo) SetEtag(v string) *UploadPartInfo {
	s.Etag = &v
	return s
}

func (s *UploadPartInfo) SetInternalUploadUrl(v string) *UploadPartInfo {
	s.InternalUploadUrl = &v
	return s
}

func (s *UploadPartInfo) SetParallelSha1Ctx(v *UploadPartInfoParallelSha1Ctx) *UploadPartInfo {
	s.ParallelSha1Ctx = v
	return s
}

func (s *UploadPartInfo) SetParallelSha256Ctx(v *UploadPartInfoParallelSha256Ctx) *UploadPartInfo {
	s.ParallelSha256Ctx = v
	return s
}

func (s *UploadPartInfo) SetPartNumber(v int32) *UploadPartInfo {
	s.PartNumber = &v
	return s
}

func (s *UploadPartInfo) SetPartSize(v int64) *UploadPartInfo {
	s.PartSize = &v
	return s
}

func (s *UploadPartInfo) SetUploadUrl(v string) *UploadPartInfo {
	s.UploadUrl = &v
	return s
}

func (s *UploadPartInfo) Validate() error {
	if s.ParallelSha1Ctx != nil {
		if err := s.ParallelSha1Ctx.Validate(); err != nil {
			return err
		}
	}
	if s.ParallelSha256Ctx != nil {
		if err := s.ParallelSha256Ctx.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadPartInfoParallelSha1Ctx struct {
	// The first to fifth 32-bit variables in the SHA-1 context of the previous part. This parameter takes effect only if the parallel upload feature is enabled.
	H []*int64 `json:"h,omitempty" xml:"h,omitempty" type:"Repeated"`
	// The total size of all the previous parts. Unit: bytes. The value must be a multiple of 64. This parameter takes effect only if the parallel upload feature is enabled.
	//
	// example:
	//
	// 10240
	PartOffset *int64 `json:"part_offset,omitempty" xml:"part_offset,omitempty"`
}

func (s UploadPartInfoParallelSha1Ctx) String() string {
	return dara.Prettify(s)
}

func (s UploadPartInfoParallelSha1Ctx) GoString() string {
	return s.String()
}

func (s *UploadPartInfoParallelSha1Ctx) GetH() []*int64 {
	return s.H
}

func (s *UploadPartInfoParallelSha1Ctx) GetPartOffset() *int64 {
	return s.PartOffset
}

func (s *UploadPartInfoParallelSha1Ctx) SetH(v []*int64) *UploadPartInfoParallelSha1Ctx {
	s.H = v
	return s
}

func (s *UploadPartInfoParallelSha1Ctx) SetPartOffset(v int64) *UploadPartInfoParallelSha1Ctx {
	s.PartOffset = &v
	return s
}

func (s *UploadPartInfoParallelSha1Ctx) Validate() error {
	return dara.Validate(s)
}

type UploadPartInfoParallelSha256Ctx struct {
	// The first to eighth 32-bit variables in the SHA-256 context of the previous part.
	H []*int64 `json:"h,omitempty" xml:"h,omitempty" type:"Repeated"`
	// The total size of all the previous parts. Unit: bytes. The value must be a multiple of 64.
	PartOffset *int64 `json:"part_offset,omitempty" xml:"part_offset,omitempty"`
}

func (s UploadPartInfoParallelSha256Ctx) String() string {
	return dara.Prettify(s)
}

func (s UploadPartInfoParallelSha256Ctx) GoString() string {
	return s.String()
}

func (s *UploadPartInfoParallelSha256Ctx) GetH() []*int64 {
	return s.H
}

func (s *UploadPartInfoParallelSha256Ctx) GetPartOffset() *int64 {
	return s.PartOffset
}

func (s *UploadPartInfoParallelSha256Ctx) SetH(v []*int64) *UploadPartInfoParallelSha256Ctx {
	s.H = v
	return s
}

func (s *UploadPartInfoParallelSha256Ctx) SetPartOffset(v int64) *UploadPartInfoParallelSha256Ctx {
	s.PartOffset = &v
	return s
}

func (s *UploadPartInfoParallelSha256Ctx) Validate() error {
	return dara.Validate(s)
}
