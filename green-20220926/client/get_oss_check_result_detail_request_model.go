// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckResultDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *GetOssCheckResultDetailRequest
	GetBucket() *string
	SetMediaType(v int32) *GetOssCheckResultDetailRequest
	GetMediaType() *int32
	SetObject(v string) *GetOssCheckResultDetailRequest
	GetObject() *string
	SetParentTaskId(v string) *GetOssCheckResultDetailRequest
	GetParentTaskId() *string
	SetQueryRequestId(v string) *GetOssCheckResultDetailRequest
	GetQueryRequestId() *string
	SetRegionId(v string) *GetOssCheckResultDetailRequest
	GetRegionId() *string
	SetServiceCode(v string) *GetOssCheckResultDetailRequest
	GetServiceCode() *string
}

type GetOssCheckResultDetailRequest struct {
	// Bucket name.
	//
	// example:
	//
	// oss-tmp
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Media type.
	//
	// example:
	//
	// 1
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Object name.
	//
	// example:
	//
	// 1748396909030.jpg
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// Parent task ID.
	//
	// example:
	//
	// P_RZQ66T
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	// Query request ID.
	//
	// example:
	//
	// 62E97001-1255-50A9-8E1E-4FD05473D952
	QueryRequestId *string `json:"QueryRequestId,omitempty" xml:"QueryRequestId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service code.
	//
	// example:
	//
	// audio_media_detection_01
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetOssCheckResultDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckResultDetailRequest) GoString() string {
	return s.String()
}

func (s *GetOssCheckResultDetailRequest) GetBucket() *string {
	return s.Bucket
}

func (s *GetOssCheckResultDetailRequest) GetMediaType() *int32 {
	return s.MediaType
}

func (s *GetOssCheckResultDetailRequest) GetObject() *string {
	return s.Object
}

func (s *GetOssCheckResultDetailRequest) GetParentTaskId() *string {
	return s.ParentTaskId
}

func (s *GetOssCheckResultDetailRequest) GetQueryRequestId() *string {
	return s.QueryRequestId
}

func (s *GetOssCheckResultDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOssCheckResultDetailRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetOssCheckResultDetailRequest) SetBucket(v string) *GetOssCheckResultDetailRequest {
	s.Bucket = &v
	return s
}

func (s *GetOssCheckResultDetailRequest) SetMediaType(v int32) *GetOssCheckResultDetailRequest {
	s.MediaType = &v
	return s
}

func (s *GetOssCheckResultDetailRequest) SetObject(v string) *GetOssCheckResultDetailRequest {
	s.Object = &v
	return s
}

func (s *GetOssCheckResultDetailRequest) SetParentTaskId(v string) *GetOssCheckResultDetailRequest {
	s.ParentTaskId = &v
	return s
}

func (s *GetOssCheckResultDetailRequest) SetQueryRequestId(v string) *GetOssCheckResultDetailRequest {
	s.QueryRequestId = &v
	return s
}

func (s *GetOssCheckResultDetailRequest) SetRegionId(v string) *GetOssCheckResultDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetOssCheckResultDetailRequest) SetServiceCode(v string) *GetOssCheckResultDetailRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetOssCheckResultDetailRequest) Validate() error {
	return dara.Validate(s)
}
