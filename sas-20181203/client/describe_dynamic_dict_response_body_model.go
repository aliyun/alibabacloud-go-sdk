// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDynamicDictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeDynamicDictResponseBody
	GetCount() *int32
	SetFileName(v string) *DescribeDynamicDictResponseBody
	GetFileName() *string
	SetFileSize(v int64) *DescribeDynamicDictResponseBody
	GetFileSize() *int64
	SetLimit(v int32) *DescribeDynamicDictResponseBody
	GetLimit() *int32
	SetOssUrl(v string) *DescribeDynamicDictResponseBody
	GetOssUrl() *string
	SetRequestId(v string) *DescribeDynamicDictResponseBody
	GetRequestId() *string
	SetUploadTime(v int64) *DescribeDynamicDictResponseBody
	GetUploadTime() *int64
}

type DescribeDynamicDictResponseBody struct {
	// The number of weak password rules that are added.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the OSS object that contains custom weak passwords.
	//
	// example:
	//
	// test_dict.plain
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The size of the OSS object. Unit: bytes.
	//
	// example:
	//
	// 40
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The maximum number of weak password rules that can be added.
	//
	// example:
	//
	// 9
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The IP address of the Object Storage Service (OSS) object.
	//
	// example:
	//
	// cloudtest01/661767e1-5ae3-4ec5-865f-03039436893a/sacc2*****
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3D7C47D-3F11-57BB-90E8-E5C20C61****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timestamp when the OSS object was uploaded. Unit: milliseconds.
	//
	// example:
	//
	// 1719919421
	UploadTime *int64 `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s DescribeDynamicDictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicDictResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDynamicDictResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeDynamicDictResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *DescribeDynamicDictResponseBody) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeDynamicDictResponseBody) GetLimit() *int32 {
	return s.Limit
}

func (s *DescribeDynamicDictResponseBody) GetOssUrl() *string {
	return s.OssUrl
}

func (s *DescribeDynamicDictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDynamicDictResponseBody) GetUploadTime() *int64 {
	return s.UploadTime
}

func (s *DescribeDynamicDictResponseBody) SetCount(v int32) *DescribeDynamicDictResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeDynamicDictResponseBody) SetFileName(v string) *DescribeDynamicDictResponseBody {
	s.FileName = &v
	return s
}

func (s *DescribeDynamicDictResponseBody) SetFileSize(v int64) *DescribeDynamicDictResponseBody {
	s.FileSize = &v
	return s
}

func (s *DescribeDynamicDictResponseBody) SetLimit(v int32) *DescribeDynamicDictResponseBody {
	s.Limit = &v
	return s
}

func (s *DescribeDynamicDictResponseBody) SetOssUrl(v string) *DescribeDynamicDictResponseBody {
	s.OssUrl = &v
	return s
}

func (s *DescribeDynamicDictResponseBody) SetRequestId(v string) *DescribeDynamicDictResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDynamicDictResponseBody) SetUploadTime(v int64) *DescribeDynamicDictResponseBody {
	s.UploadTime = &v
	return s
}

func (s *DescribeDynamicDictResponseBody) Validate() error {
	return dara.Validate(s)
}
