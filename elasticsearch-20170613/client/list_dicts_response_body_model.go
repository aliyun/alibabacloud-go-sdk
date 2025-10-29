// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDictsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListDictsResponseBodyHeaders) *ListDictsResponseBody
	GetHeaders() *ListDictsResponseBodyHeaders
	SetRequestId(v string) *ListDictsResponseBody
	GetRequestId() *string
	SetResult(v []*ListDictsResponseBodyResult) *ListDictsResponseBody
	GetResult() []*ListDictsResponseBodyResult
}

type ListDictsResponseBody struct {
	// The header of the response.
	Headers *ListDictsResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2937F832-F39E-41EF-89BA-B528342A2A3A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*ListDictsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDictsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDictsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDictsResponseBody) GetHeaders() *ListDictsResponseBodyHeaders {
	return s.Headers
}

func (s *ListDictsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDictsResponseBody) GetResult() []*ListDictsResponseBodyResult {
	return s.Result
}

func (s *ListDictsResponseBody) SetHeaders(v *ListDictsResponseBodyHeaders) *ListDictsResponseBody {
	s.Headers = v
	return s
}

func (s *ListDictsResponseBody) SetRequestId(v string) *ListDictsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDictsResponseBody) SetResult(v []*ListDictsResponseBodyResult) *ListDictsResponseBody {
	s.Result = v
	return s
}

func (s *ListDictsResponseBody) Validate() error {
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDictsResponseBodyHeaders struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListDictsResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDictsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListDictsResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListDictsResponseBodyHeaders) SetXTotalCount(v int32) *ListDictsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListDictsResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListDictsResponseBodyResult struct {
	// The link that is used to download the dictionary over the Internet. The link is valid for 90s.
	//
	// example:
	//
	// http://test_bucket.oss-cn-hangzhou.aliyuncs.com/AliyunEs/test.dic?Expires=162573****&OSSAccessKeyId=LTAI*****V9&Signature=PNPO********BBGsJDO4V3VfU4sE%3D
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// The size of the dictionary file. Unit: byte.
	//
	// example:
	//
	// 2782602
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The name of the dictionary file.
	//
	// example:
	//
	// SYSTEM_MAIN.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source type.
	//
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The type of the IK dictionary. Valid values:
	//
	// 	- MAIN: main dictionary
	//
	// 	- STOP: stopword list
	//
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDictsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDictsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDictsResponseBodyResult) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListDictsResponseBodyResult) GetFileSize() *int64 {
	return s.FileSize
}

func (s *ListDictsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListDictsResponseBodyResult) GetSourceType() *string {
	return s.SourceType
}

func (s *ListDictsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListDictsResponseBodyResult) SetDownloadUrl(v string) *ListDictsResponseBodyResult {
	s.DownloadUrl = &v
	return s
}

func (s *ListDictsResponseBodyResult) SetFileSize(v int64) *ListDictsResponseBodyResult {
	s.FileSize = &v
	return s
}

func (s *ListDictsResponseBodyResult) SetName(v string) *ListDictsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDictsResponseBodyResult) SetSourceType(v string) *ListDictsResponseBodyResult {
	s.SourceType = &v
	return s
}

func (s *ListDictsResponseBodyResult) SetType(v string) *ListDictsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListDictsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
