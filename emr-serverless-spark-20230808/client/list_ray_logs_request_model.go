// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *ListRayLogsRequest
	GetBucketName() *string
	SetDelimiter(v string) *ListRayLogsRequest
	GetDelimiter() *string
	SetMarker(v string) *ListRayLogsRequest
	GetMarker() *string
	SetMaxKeys(v int64) *ListRayLogsRequest
	GetMaxKeys() *int64
	SetPrefix(v string) *ListRayLogsRequest
	GetPrefix() *string
}

type ListRayLogsRequest struct {
	// The bucket name.
	//
	// example:
	//
	// mybucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The character used to group object names. All objects whose names contain the specified prefix and between which the delimiter character appears for the first time are grouped as a set of elements (CommonPrefixes).
	//
	// example:
	//
	// /
	Delimiter *string `json:"delimiter,omitempty" xml:"delimiter,omitempty"`
	// The marker after which the returned objects are listed in alphabetical order.
	//
	// example:
	//
	// test1.txt
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of objects to return. If the listing cannot be completed in a single request due to the max-keys setting, a NextMarker element is included in the response as the marker for the next listing request.
	//
	// Valid values: greater than 0 and less than 1000.
	//
	// Default value: 100.
	//
	// example:
	//
	// 50
	MaxKeys *int64 `json:"maxKeys,omitempty" xml:"maxKeys,omitempty"`
	// The prefix that the keys of the returned files must start with.
	//
	// example:
	//
	// /w-xxxxxxx/ray/logs/rj-xxxxxxxxxx_default/
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListRayLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRayLogsRequest) GoString() string {
	return s.String()
}

func (s *ListRayLogsRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *ListRayLogsRequest) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListRayLogsRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListRayLogsRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListRayLogsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListRayLogsRequest) SetBucketName(v string) *ListRayLogsRequest {
	s.BucketName = &v
	return s
}

func (s *ListRayLogsRequest) SetDelimiter(v string) *ListRayLogsRequest {
	s.Delimiter = &v
	return s
}

func (s *ListRayLogsRequest) SetMarker(v string) *ListRayLogsRequest {
	s.Marker = &v
	return s
}

func (s *ListRayLogsRequest) SetMaxKeys(v int64) *ListRayLogsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListRayLogsRequest) SetPrefix(v string) *ListRayLogsRequest {
	s.Prefix = &v
	return s
}

func (s *ListRayLogsRequest) Validate() error {
	return dara.Validate(s)
}
