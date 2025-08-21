// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *ListObjectsRequest
	GetBucketName() *string
	SetContinuationToken(v string) *ListObjectsRequest
	GetContinuationToken() *string
	SetEncodingType(v string) *ListObjectsRequest
	GetEncodingType() *string
	SetMarker(v string) *ListObjectsRequest
	GetMarker() *string
	SetMaxKeys(v int64) *ListObjectsRequest
	GetMaxKeys() *int64
	SetPrefix(v string) *ListObjectsRequest
	GetPrefix() *string
	SetStartAfter(v string) *ListObjectsRequest
	GetStartAfter() *string
}

type ListObjectsRequest struct {
	// The name of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The token used in this list operation. If the number of objects exceeds the value of MaxKeys, the NextContinuationToken is included in the response as the token for the next list operation.
	//
	// example:
	//
	// test1.txt
	ContinuationToken *string `json:"ContinuationToken,omitempty" xml:"ContinuationToken,omitempty"`
	// The encoding type of the object names in the response. Only URL encoding is supported.
	//
	// example:
	//
	// url
	EncodingType *string `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	// The position from which the list operation starts. If this parameter is specified, objects whose names are alphabetically greater than value of Marker are returned. The Marker parameter is used to list the returned objects by page, and its value must be smaller than 1,024 bytes in length.
	//
	// Even if the value specified for Marker does not exist in the list during a conditional query, the list starts from the object whose name is alphabetically greater than the value of Marker.
	//
	// example:
	//
	// a
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The maximum number of objects to return. Valid values: 0 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxKeys *int64 `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	// The prefix that must be included in the names of objects you want to list. If you specify a prefix to query objects, the returned object names contain the prefix.
	//
	// The value of the parameter must be less than 1,000 bytes in length.
	//
	// example:
	//
	// b
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The position from which the list operation starts. If this parameter is specified, objects whose names are alphabetically greater than the value of StartAfter are returned. The StartAfter parameter is used to list the returned objects by page, and its value must be less than 1,000 bytes in length. Even if the value specified for StartAfter does not exist in the list during a conditional query, the list starts from the object whose name is alphabetically greater than the value of StartAfter.
	//
	// example:
	//
	// b
	StartAfter *string `json:"StartAfter,omitempty" xml:"StartAfter,omitempty"`
}

func (s ListObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListObjectsRequest) GoString() string {
	return s.String()
}

func (s *ListObjectsRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *ListObjectsRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListObjectsRequest) GetEncodingType() *string {
	return s.EncodingType
}

func (s *ListObjectsRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListObjectsRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListObjectsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListObjectsRequest) GetStartAfter() *string {
	return s.StartAfter
}

func (s *ListObjectsRequest) SetBucketName(v string) *ListObjectsRequest {
	s.BucketName = &v
	return s
}

func (s *ListObjectsRequest) SetContinuationToken(v string) *ListObjectsRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListObjectsRequest) SetEncodingType(v string) *ListObjectsRequest {
	s.EncodingType = &v
	return s
}

func (s *ListObjectsRequest) SetMarker(v string) *ListObjectsRequest {
	s.Marker = &v
	return s
}

func (s *ListObjectsRequest) SetMaxKeys(v int64) *ListObjectsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectsRequest) SetPrefix(v string) *ListObjectsRequest {
	s.Prefix = &v
	return s
}

func (s *ListObjectsRequest) SetStartAfter(v string) *ListObjectsRequest {
	s.StartAfter = &v
	return s
}

func (s *ListObjectsRequest) Validate() error {
	return dara.Validate(s)
}
