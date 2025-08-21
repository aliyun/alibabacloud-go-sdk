// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExtendfilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListExtendfilesResponseBody
	GetRequestId() *string
	SetResult(v []*ListExtendfilesResponseBodyResult) *ListExtendfilesResponseBody
	GetResult() []*ListExtendfilesResponseBodyResult
}

type ListExtendfilesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*ListExtendfilesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListExtendfilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExtendfilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExtendfilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExtendfilesResponseBody) GetResult() []*ListExtendfilesResponseBodyResult {
	return s.Result
}

func (s *ListExtendfilesResponseBody) SetRequestId(v string) *ListExtendfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExtendfilesResponseBody) SetResult(v []*ListExtendfilesResponseBodyResult) *ListExtendfilesResponseBody {
	s.Result = v
	return s
}

func (s *ListExtendfilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExtendfilesResponseBodyResult struct {
	// The path of the driver file.
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// The size of the driver file.
	//
	// example:
	//
	// 968668
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The name of the driver file.
	//
	// example:
	//
	// mysql-connector-java-5.1.35.jar
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source type.
	//
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListExtendfilesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListExtendfilesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListExtendfilesResponseBodyResult) GetFilePath() *string {
	return s.FilePath
}

func (s *ListExtendfilesResponseBodyResult) GetFileSize() *int64 {
	return s.FileSize
}

func (s *ListExtendfilesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListExtendfilesResponseBodyResult) GetSourceType() *string {
	return s.SourceType
}

func (s *ListExtendfilesResponseBodyResult) SetFilePath(v string) *ListExtendfilesResponseBodyResult {
	s.FilePath = &v
	return s
}

func (s *ListExtendfilesResponseBodyResult) SetFileSize(v int64) *ListExtendfilesResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *ListExtendfilesResponseBodyResult) SetName(v string) *ListExtendfilesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListExtendfilesResponseBodyResult) SetSourceType(v string) *ListExtendfilesResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *ListExtendfilesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
