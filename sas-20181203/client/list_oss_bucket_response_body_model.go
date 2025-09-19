// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListOssBucketResponseBodyData) *ListOssBucketResponseBody
	GetData() []*ListOssBucketResponseBodyData
	SetRequestId(v string) *ListOssBucketResponseBody
	GetRequestId() *string
}

type ListOssBucketResponseBody struct {
	// The data returned.
	Data []*ListOssBucketResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3956048F-9D73-5EDB-834B-4827BB48****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOssBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOssBucketResponseBody) GoString() string {
	return s.String()
}

func (s *ListOssBucketResponseBody) GetData() []*ListOssBucketResponseBodyData {
	return s.Data
}

func (s *ListOssBucketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOssBucketResponseBody) SetData(v []*ListOssBucketResponseBodyData) *ListOssBucketResponseBody {
	s.Data = v
	return s
}

func (s *ListOssBucketResponseBody) SetRequestId(v string) *ListOssBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOssBucketResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOssBucketResponseBodyData struct {
	// The name of the bucket.
	//
	// example:
	//
	// iboxpublic****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The reason why the bucket cannot be checked.
	//
	// example:
	//
	// Unsupported Region.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The storage class of the bucket. Valid values:
	//
	// 	- Standard (default)
	//
	// 	- IA
	//
	// 	- Archive
	//
	// 	- ColdArchive
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// Indicates whether the bucket can be checked. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Support *bool `json:"Support,omitempty" xml:"Support,omitempty"`
	// Whether to support config. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SupportConfig *string `json:"SupportConfig,omitempty" xml:"SupportConfig,omitempty"`
}

func (s ListOssBucketResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOssBucketResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOssBucketResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *ListOssBucketResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListOssBucketResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOssBucketResponseBodyData) GetStorageClass() *string {
	return s.StorageClass
}

func (s *ListOssBucketResponseBodyData) GetSupport() *bool {
	return s.Support
}

func (s *ListOssBucketResponseBodyData) GetSupportConfig() *string {
	return s.SupportConfig
}

func (s *ListOssBucketResponseBodyData) SetBucketName(v string) *ListOssBucketResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *ListOssBucketResponseBodyData) SetMessage(v string) *ListOssBucketResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListOssBucketResponseBodyData) SetRegionId(v string) *ListOssBucketResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListOssBucketResponseBodyData) SetStorageClass(v string) *ListOssBucketResponseBodyData {
	s.StorageClass = &v
	return s
}

func (s *ListOssBucketResponseBodyData) SetSupport(v bool) *ListOssBucketResponseBodyData {
	s.Support = &v
	return s
}

func (s *ListOssBucketResponseBodyData) SetSupportConfig(v string) *ListOssBucketResponseBodyData {
	s.SupportConfig = &v
	return s
}

func (s *ListOssBucketResponseBodyData) Validate() error {
	return dara.Validate(s)
}
