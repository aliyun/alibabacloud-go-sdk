// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRayLogsResponseBody
	GetCode() *string
	SetData(v *ListRayLogsResponseBodyData) *ListRayLogsResponseBody
	GetData() *ListRayLogsResponseBodyData
	SetMessage(v string) *ListRayLogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRayLogsResponseBody
	GetRequestId() *string
}

type ListRayLogsResponseBody struct {
	// The response status code. The value 1000000 indicates success.
	//
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListRayLogsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRayLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRayLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRayLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRayLogsResponseBody) GetData() *ListRayLogsResponseBodyData {
	return s.Data
}

func (s *ListRayLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRayLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRayLogsResponseBody) SetCode(v string) *ListRayLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRayLogsResponseBody) SetData(v *ListRayLogsResponseBodyData) *ListRayLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListRayLogsResponseBody) SetMessage(v string) *ListRayLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRayLogsResponseBody) SetRequestId(v string) *ListRayLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRayLogsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRayLogsResponseBodyData struct {
	// The OSS bucket name.
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
	// Indicates whether the results returned in the request are truncated.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"isTruncated,omitempty" xml:"isTruncated,omitempty"`
	// The marker after which the returned objects are listed in alphabetical order.
	//
	// example:
	//
	// test1.txt
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of objects to return.
	//
	// example:
	//
	// 50
	MaxKeys *int64 `json:"maxKeys,omitempty" xml:"maxKeys,omitempty"`
	// The marker from which the next listing of files starts.
	//
	// example:
	//
	// test2.txt
	NextMarker *string `json:"nextMarker,omitempty" xml:"nextMarker,omitempty"`
	// The list of object metadata.
	ObjectList []*ListRayLogsResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// The prefix that the keys of the returned files must start with.
	//
	// example:
	//
	// /w-xxxxxxx/ray/logs/rj-xxxxxxxxxx_default/
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListRayLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRayLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRayLogsResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *ListRayLogsResponseBodyData) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListRayLogsResponseBodyData) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListRayLogsResponseBodyData) GetMarker() *string {
	return s.Marker
}

func (s *ListRayLogsResponseBodyData) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListRayLogsResponseBodyData) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListRayLogsResponseBodyData) GetObjectList() []*ListRayLogsResponseBodyDataObjectList {
	return s.ObjectList
}

func (s *ListRayLogsResponseBodyData) GetPrefix() *string {
	return s.Prefix
}

func (s *ListRayLogsResponseBodyData) SetBucketName(v string) *ListRayLogsResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *ListRayLogsResponseBodyData) SetDelimiter(v string) *ListRayLogsResponseBodyData {
	s.Delimiter = &v
	return s
}

func (s *ListRayLogsResponseBodyData) SetIsTruncated(v bool) *ListRayLogsResponseBodyData {
	s.IsTruncated = &v
	return s
}

func (s *ListRayLogsResponseBodyData) SetMarker(v string) *ListRayLogsResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListRayLogsResponseBodyData) SetMaxKeys(v int64) *ListRayLogsResponseBodyData {
	s.MaxKeys = &v
	return s
}

func (s *ListRayLogsResponseBodyData) SetNextMarker(v string) *ListRayLogsResponseBodyData {
	s.NextMarker = &v
	return s
}

func (s *ListRayLogsResponseBodyData) SetObjectList(v []*ListRayLogsResponseBodyDataObjectList) *ListRayLogsResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListRayLogsResponseBodyData) SetPrefix(v string) *ListRayLogsResponseBodyData {
	s.Prefix = &v
	return s
}

func (s *ListRayLogsResponseBodyData) Validate() error {
	if s.ObjectList != nil {
		for _, item := range s.ObjectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRayLogsResponseBodyDataObjectList struct {
	// Indicates whether the entry is a directory.
	//
	// example:
	//
	// false
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// The name.
	//
	// example:
	//
	// test1.txt
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The file path.
	//
	// example:
	//
	// /w-xxxxxxx/ray/logs/rj-xxxxxxxxxx_default/test1.txt
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The file size, in bytes.
	//
	// example:
	//
	// 535345
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The modification time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 17344656363434
	TimeModified *int64 `json:"timeModified,omitempty" xml:"timeModified,omitempty"`
}

func (s ListRayLogsResponseBodyDataObjectList) String() string {
	return dara.Prettify(s)
}

func (s ListRayLogsResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListRayLogsResponseBodyDataObjectList) GetIsDir() *bool {
	return s.IsDir
}

func (s *ListRayLogsResponseBodyDataObjectList) GetName() *string {
	return s.Name
}

func (s *ListRayLogsResponseBodyDataObjectList) GetPath() *string {
	return s.Path
}

func (s *ListRayLogsResponseBodyDataObjectList) GetSize() *int64 {
	return s.Size
}

func (s *ListRayLogsResponseBodyDataObjectList) GetTimeModified() *int64 {
	return s.TimeModified
}

func (s *ListRayLogsResponseBodyDataObjectList) SetIsDir(v bool) *ListRayLogsResponseBodyDataObjectList {
	s.IsDir = &v
	return s
}

func (s *ListRayLogsResponseBodyDataObjectList) SetName(v string) *ListRayLogsResponseBodyDataObjectList {
	s.Name = &v
	return s
}

func (s *ListRayLogsResponseBodyDataObjectList) SetPath(v string) *ListRayLogsResponseBodyDataObjectList {
	s.Path = &v
	return s
}

func (s *ListRayLogsResponseBodyDataObjectList) SetSize(v int64) *ListRayLogsResponseBodyDataObjectList {
	s.Size = &v
	return s
}

func (s *ListRayLogsResponseBodyDataObjectList) SetTimeModified(v int64) *ListRayLogsResponseBodyDataObjectList {
	s.TimeModified = &v
	return s
}

func (s *ListRayLogsResponseBodyDataObjectList) Validate() error {
	return dara.Validate(s)
}
