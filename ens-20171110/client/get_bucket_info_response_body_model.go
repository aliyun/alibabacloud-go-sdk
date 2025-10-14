// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketInfo(v *GetBucketInfoResponseBodyBucketInfo) *GetBucketInfoResponseBody
	GetBucketInfo() *GetBucketInfoResponseBodyBucketInfo
	SetRequestId(v string) *GetBucketInfoResponseBody
	GetRequestId() *string
}

type GetBucketInfoResponseBody struct {
	// The list of bucket information.
	BucketInfo *GetBucketInfoResponseBodyBucketInfo `json:"BucketInfo,omitempty" xml:"BucketInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C5831388-2D4B-46F4-A96B-D4E6BD06E7521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBucketInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketInfoResponseBody) GetBucketInfo() *GetBucketInfoResponseBodyBucketInfo {
	return s.BucketInfo
}

func (s *GetBucketInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBucketInfoResponseBody) SetBucketInfo(v *GetBucketInfoResponseBodyBucketInfo) *GetBucketInfoResponseBody {
	s.BucketInfo = v
	return s
}

func (s *GetBucketInfoResponseBody) SetRequestId(v string) *GetBucketInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBucketInfoResponseBody) Validate() error {
	if s.BucketInfo != nil {
		if err := s.BucketInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketInfoResponseBodyBucketInfo struct {
	// The ACL of the bucket.
	//
	// 	- **public-read-write**
	//
	// 	- **public-read**
	//
	// 	- **private*	- (default)
	//
	// example:
	//
	// private
	BucketAcl *string `json:"BucketAcl,omitempty" xml:"BucketAcl,omitempty"`
	// The name of the bucket.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The remarks.
	//
	// example:
	//
	// das
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the bucket was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-10-12T05:45:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Single-node storage. Set the value to sink.
	//
	// example:
	//
	// sink
	LogicalBucketType *string `json:"LogicalBucketType,omitempty" xml:"LogicalBucketType,omitempty"`
	// The time when the bucket was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-10-12T06:45:00Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s GetBucketInfoResponseBodyBucketInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBucketInfoResponseBodyBucketInfo) GoString() string {
	return s.String()
}

func (s *GetBucketInfoResponseBodyBucketInfo) GetBucketAcl() *string {
	return s.BucketAcl
}

func (s *GetBucketInfoResponseBodyBucketInfo) GetBucketName() *string {
	return s.BucketName
}

func (s *GetBucketInfoResponseBodyBucketInfo) GetComment() *string {
	return s.Comment
}

func (s *GetBucketInfoResponseBodyBucketInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetBucketInfoResponseBodyBucketInfo) GetLogicalBucketType() *string {
	return s.LogicalBucketType
}

func (s *GetBucketInfoResponseBodyBucketInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetBucketAcl(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.BucketAcl = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetBucketName(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.BucketName = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetComment(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.Comment = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetCreateTime(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.CreateTime = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetLogicalBucketType(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.LogicalBucketType = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetModifyTime(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) Validate() error {
	return dara.Validate(s)
}
