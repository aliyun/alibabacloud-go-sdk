// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDictInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDictInformationResponseBody
	GetRequestId() *string
	SetResult(v *ListDictInformationResponseBodyResult) *ListDictInformationResponseBody
	GetResult() *ListDictInformationResponseBodyResult
}

type ListDictInformationResponseBody struct {
	// example:
	//
	// 7C4334EA-D22B-48BD-AE28-08EE68******
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListDictInformationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListDictInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDictInformationResponseBody) GoString() string {
	return s.String()
}

func (s *ListDictInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDictInformationResponseBody) GetResult() *ListDictInformationResponseBodyResult {
	return s.Result
}

func (s *ListDictInformationResponseBody) SetRequestId(v string) *ListDictInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDictInformationResponseBody) SetResult(v *ListDictInformationResponseBodyResult) *ListDictInformationResponseBody {
	s.Result = v
	return s
}

func (s *ListDictInformationResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDictInformationResponseBodyResult struct {
	// example:
	//
	// 2202301
	FileSize  *int64                                          `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	OssObject *ListDictInformationResponseBodyResultOssObject `json:"ossObject,omitempty" xml:"ossObject,omitempty" type:"Struct"`
	// example:
	//
	// STOP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDictInformationResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDictInformationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDictInformationResponseBodyResult) GetFileSize() *int64 {
	return s.FileSize
}

func (s *ListDictInformationResponseBodyResult) GetOssObject() *ListDictInformationResponseBodyResultOssObject {
	return s.OssObject
}

func (s *ListDictInformationResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListDictInformationResponseBodyResult) SetFileSize(v int64) *ListDictInformationResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *ListDictInformationResponseBodyResult) SetOssObject(v *ListDictInformationResponseBodyResultOssObject) *ListDictInformationResponseBodyResult {
	s.OssObject = v
	return s
}

func (s *ListDictInformationResponseBodyResult) SetType(v string) *ListDictInformationResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListDictInformationResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListDictInformationResponseBodyResultOssObject struct {
	// example:
	//
	// es-osstest*
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// example:
	//
	// 2ABAB5E70BBF631145647F6BE533****
	Etag *string `json:"etag,omitempty" xml:"etag,omitempty"`
	// example:
	//
	// oss/dict_0*.dic
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ListDictInformationResponseBodyResultOssObject) String() string {
	return dara.Prettify(s)
}

func (s ListDictInformationResponseBodyResultOssObject) GoString() string {
	return s.String()
}

func (s *ListDictInformationResponseBodyResultOssObject) GetBucketName() *string {
	return s.BucketName
}

func (s *ListDictInformationResponseBodyResultOssObject) GetEtag() *string {
	return s.Etag
}

func (s *ListDictInformationResponseBodyResultOssObject) GetKey() *string {
	return s.Key
}

func (s *ListDictInformationResponseBodyResultOssObject) SetBucketName(v string) *ListDictInformationResponseBodyResultOssObject {
	s.BucketName = &v
	return s
}

func (s *ListDictInformationResponseBodyResultOssObject) SetEtag(v string) *ListDictInformationResponseBodyResultOssObject {
	s.Etag = &v
	return s
}

func (s *ListDictInformationResponseBodyResultOssObject) SetKey(v string) *ListDictInformationResponseBodyResultOssObject {
	s.Key = &v
	return s
}

func (s *ListDictInformationResponseBodyResultOssObject) Validate() error {
	return dara.Validate(s)
}
