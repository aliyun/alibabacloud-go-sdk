// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllMediaBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaBucketList(v *ListAllMediaBucketResponseBodyMediaBucketList) *ListAllMediaBucketResponseBody
	GetMediaBucketList() *ListAllMediaBucketResponseBodyMediaBucketList
	SetNextPageToken(v string) *ListAllMediaBucketResponseBody
	GetNextPageToken() *string
	SetRequestId(v string) *ListAllMediaBucketResponseBody
	GetRequestId() *string
}

type ListAllMediaBucketResponseBody struct {
	// The media buckets returned.
	MediaBucketList *ListAllMediaBucketResponseBodyMediaBucketList `json:"MediaBucketList,omitempty" xml:"MediaBucketList,omitempty" type:"Struct"`
	// The returned value of NextPageToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// P2Zqo1PLGhZdygo-ajSsjUX5zrBHCgXy6j4hEvv****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 79760D91-D3CF-4165-****-B7E2836EF62A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAllMediaBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllMediaBucketResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllMediaBucketResponseBody) GetMediaBucketList() *ListAllMediaBucketResponseBodyMediaBucketList {
	return s.MediaBucketList
}

func (s *ListAllMediaBucketResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListAllMediaBucketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllMediaBucketResponseBody) SetMediaBucketList(v *ListAllMediaBucketResponseBodyMediaBucketList) *ListAllMediaBucketResponseBody {
	s.MediaBucketList = v
	return s
}

func (s *ListAllMediaBucketResponseBody) SetNextPageToken(v string) *ListAllMediaBucketResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListAllMediaBucketResponseBody) SetRequestId(v string) *ListAllMediaBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllMediaBucketResponseBody) Validate() error {
	if s.MediaBucketList != nil {
		if err := s.MediaBucketList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAllMediaBucketResponseBodyMediaBucketList struct {
	MediaBucket []*ListAllMediaBucketResponseBodyMediaBucketListMediaBucket `json:"MediaBucket,omitempty" xml:"MediaBucket,omitempty" type:"Repeated"`
}

func (s ListAllMediaBucketResponseBodyMediaBucketList) String() string {
	return dara.Prettify(s)
}

func (s ListAllMediaBucketResponseBodyMediaBucketList) GoString() string {
	return s.String()
}

func (s *ListAllMediaBucketResponseBodyMediaBucketList) GetMediaBucket() []*ListAllMediaBucketResponseBodyMediaBucketListMediaBucket {
	return s.MediaBucket
}

func (s *ListAllMediaBucketResponseBodyMediaBucketList) SetMediaBucket(v []*ListAllMediaBucketResponseBodyMediaBucketListMediaBucket) *ListAllMediaBucketResponseBodyMediaBucketList {
	s.MediaBucket = v
	return s
}

func (s *ListAllMediaBucketResponseBodyMediaBucketList) Validate() error {
	if s.MediaBucket != nil {
		for _, item := range s.MediaBucket {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllMediaBucketResponseBodyMediaBucketListMediaBucket struct {
	// The name of the media bucket.
	//
	// example:
	//
	// example-bucket-****
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The settings of Object Storage Service (OSS) hotlink protection. For more information, see [Hotlink protection](https://help.aliyun.com/document_detail/31869.html).
	//
	// example:
	//
	// http://www.example.com
	Referer *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
	// The type of the media bucket. Valid values:
	//
	// 	- Input: input media bucket
	//
	// 	- Output: output media bucket
	//
	// example:
	//
	// Input
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAllMediaBucketResponseBodyMediaBucketListMediaBucket) String() string {
	return dara.Prettify(s)
}

func (s ListAllMediaBucketResponseBodyMediaBucketListMediaBucket) GoString() string {
	return s.String()
}

func (s *ListAllMediaBucketResponseBodyMediaBucketListMediaBucket) GetBucket() *string {
	return s.Bucket
}

func (s *ListAllMediaBucketResponseBodyMediaBucketListMediaBucket) GetReferer() *string {
	return s.Referer
}

func (s *ListAllMediaBucketResponseBodyMediaBucketListMediaBucket) GetType() *string {
	return s.Type
}

func (s *ListAllMediaBucketResponseBodyMediaBucketListMediaBucket) SetBucket(v string) *ListAllMediaBucketResponseBodyMediaBucketListMediaBucket {
	s.Bucket = &v
	return s
}

func (s *ListAllMediaBucketResponseBodyMediaBucketListMediaBucket) SetReferer(v string) *ListAllMediaBucketResponseBodyMediaBucketListMediaBucket {
	s.Referer = &v
	return s
}

func (s *ListAllMediaBucketResponseBodyMediaBucketListMediaBucket) SetType(v string) *ListAllMediaBucketResponseBodyMediaBucketListMediaBucket {
	s.Type = &v
	return s
}

func (s *ListAllMediaBucketResponseBodyMediaBucketListMediaBucket) Validate() error {
	return dara.Validate(s)
}
