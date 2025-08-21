// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *ListObjectsResponseBody
	GetBucketName() *string
	SetCommonPrefixes(v []*string) *ListObjectsResponseBody
	GetCommonPrefixes() []*string
	SetContents(v []*ListObjectsResponseBodyContents) *ListObjectsResponseBody
	GetContents() []*ListObjectsResponseBodyContents
	SetContinuationToken(v string) *ListObjectsResponseBody
	GetContinuationToken() *string
	SetDelimiter(v string) *ListObjectsResponseBody
	GetDelimiter() *string
	SetEncodingType(v string) *ListObjectsResponseBody
	GetEncodingType() *string
	SetIsTruncated(v bool) *ListObjectsResponseBody
	GetIsTruncated() *bool
	SetKeyCount(v int64) *ListObjectsResponseBody
	GetKeyCount() *int64
	SetMarker(v string) *ListObjectsResponseBody
	GetMarker() *string
	SetMaxKeys(v int64) *ListObjectsResponseBody
	GetMaxKeys() *int64
	SetNextContinuationToken(v string) *ListObjectsResponseBody
	GetNextContinuationToken() *string
	SetNextMarker(v string) *ListObjectsResponseBody
	GetNextMarker() *string
	SetPrefix(v string) *ListObjectsResponseBody
	GetPrefix() *string
	SetRequestId(v string) *ListObjectsResponseBody
	GetRequestId() *string
}

type ListObjectsResponseBody struct {
	// The name of the bucket.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// If the delimiter parameter is specified in the request, the response contains CommonPrefixes. Objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
	CommonPrefixes []*string `json:"CommonPrefixes,omitempty" xml:"CommonPrefixes,omitempty" type:"Repeated"`
	// The list of object metadata.
	Contents []*ListObjectsResponseBodyContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// The token used in this list operation.
	//
	// example:
	//
	// test1.txt
	ContinuationToken *string `json:"ContinuationToken,omitempty" xml:"ContinuationToken,omitempty"`
	// The character used to group objects by name.
	//
	// example:
	//
	// /
	Delimiter *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	// The encoding type of the object names in the response.
	//
	// example:
	//
	// N/A
	EncodingType *string `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	// Indicates whether the listed objects are truncated. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The number of keys returned for this request.
	//
	// example:
	//
	// 10
	KeyCount *int64 `json:"KeyCount,omitempty" xml:"KeyCount,omitempty"`
	// The position from which the list operation starts.
	//
	// example:
	//
	// ceshi.txt1617853707991
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The maximum number of objects returned.
	//
	// example:
	//
	// 10
	MaxKeys *int64 `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	// The token used in the next list operation.
	//
	// example:
	//
	// CgJiYw--
	NextContinuationToken *string `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
	// The position from which the next list operation starts.
	//
	// example:
	//
	// ceshi.txt1617853707991
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// The prefix contained in the names of returned objects.
	//
	// example:
	//
	// b
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4833C4AC-9396-458C-8F25-1D701334E560
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListObjectsResponseBody) GetBucketName() *string {
	return s.BucketName
}

func (s *ListObjectsResponseBody) GetCommonPrefixes() []*string {
	return s.CommonPrefixes
}

func (s *ListObjectsResponseBody) GetContents() []*ListObjectsResponseBodyContents {
	return s.Contents
}

func (s *ListObjectsResponseBody) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListObjectsResponseBody) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListObjectsResponseBody) GetEncodingType() *string {
	return s.EncodingType
}

func (s *ListObjectsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListObjectsResponseBody) GetKeyCount() *int64 {
	return s.KeyCount
}

func (s *ListObjectsResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListObjectsResponseBody) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListObjectsResponseBody) GetNextContinuationToken() *string {
	return s.NextContinuationToken
}

func (s *ListObjectsResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListObjectsResponseBody) GetPrefix() *string {
	return s.Prefix
}

func (s *ListObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListObjectsResponseBody) SetBucketName(v string) *ListObjectsResponseBody {
	s.BucketName = &v
	return s
}

func (s *ListObjectsResponseBody) SetCommonPrefixes(v []*string) *ListObjectsResponseBody {
	s.CommonPrefixes = v
	return s
}

func (s *ListObjectsResponseBody) SetContents(v []*ListObjectsResponseBodyContents) *ListObjectsResponseBody {
	s.Contents = v
	return s
}

func (s *ListObjectsResponseBody) SetContinuationToken(v string) *ListObjectsResponseBody {
	s.ContinuationToken = &v
	return s
}

func (s *ListObjectsResponseBody) SetDelimiter(v string) *ListObjectsResponseBody {
	s.Delimiter = &v
	return s
}

func (s *ListObjectsResponseBody) SetEncodingType(v string) *ListObjectsResponseBody {
	s.EncodingType = &v
	return s
}

func (s *ListObjectsResponseBody) SetIsTruncated(v bool) *ListObjectsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListObjectsResponseBody) SetKeyCount(v int64) *ListObjectsResponseBody {
	s.KeyCount = &v
	return s
}

func (s *ListObjectsResponseBody) SetMarker(v string) *ListObjectsResponseBody {
	s.Marker = &v
	return s
}

func (s *ListObjectsResponseBody) SetMaxKeys(v int64) *ListObjectsResponseBody {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectsResponseBody) SetNextContinuationToken(v string) *ListObjectsResponseBody {
	s.NextContinuationToken = &v
	return s
}

func (s *ListObjectsResponseBody) SetNextMarker(v string) *ListObjectsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListObjectsResponseBody) SetPrefix(v string) *ListObjectsResponseBody {
	s.Prefix = &v
	return s
}

func (s *ListObjectsResponseBody) SetRequestId(v string) *ListObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListObjectsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListObjectsResponseBodyContents struct {
	// The entity tag (ETag). When an object is created, an ETag is created to identify the content of the object.
	//
	// 	- For an object that is created by calling the PutObject operation, the ETag value of the object is the MD5 hash of the object content.
	//
	// 	- For an object that is not created by calling the PutObject operation, the ETag value of the object is the UUID of the object content.
	//
	// 	- The ETag of an object can be used to check whether the object content is modified. However, we recommend that you use the MD5 hash of an object rather than the ETag value of the object to verify data integrity.
	//
	// example:
	//
	// 5B3C1A2E053D763E1B002CC607C5****
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// The name of the object.
	//
	// example:
	//
	// ceshi.txt1617853706546
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The time when the object was last modified.
	//
	// example:
	//
	// 2021-04-08T03:48:47.488Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The size of the returned object. Unit: bytes.
	//
	// example:
	//
	// 15
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListObjectsResponseBodyContents) String() string {
	return dara.Prettify(s)
}

func (s ListObjectsResponseBodyContents) GoString() string {
	return s.String()
}

func (s *ListObjectsResponseBodyContents) GetETag() *string {
	return s.ETag
}

func (s *ListObjectsResponseBodyContents) GetKey() *string {
	return s.Key
}

func (s *ListObjectsResponseBodyContents) GetLastModified() *string {
	return s.LastModified
}

func (s *ListObjectsResponseBodyContents) GetSize() *int64 {
	return s.Size
}

func (s *ListObjectsResponseBodyContents) SetETag(v string) *ListObjectsResponseBodyContents {
	s.ETag = &v
	return s
}

func (s *ListObjectsResponseBodyContents) SetKey(v string) *ListObjectsResponseBodyContents {
	s.Key = &v
	return s
}

func (s *ListObjectsResponseBodyContents) SetLastModified(v string) *ListObjectsResponseBodyContents {
	s.LastModified = &v
	return s
}

func (s *ListObjectsResponseBodyContents) SetSize(v int64) *ListObjectsResponseBodyContents {
	s.Size = &v
	return s
}

func (s *ListObjectsResponseBodyContents) Validate() error {
	return dara.Validate(s)
}
