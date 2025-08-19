// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDictsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDictsResponseBody
	GetRequestId() *string
	SetResult(v []*ListDictsResponseBodyResult) *ListDictsResponseBody
	GetResult() []*ListDictsResponseBodyResult
	SetTotalCount(v int32) *ListDictsResponseBody
	GetTotalCount() *int32
}

type ListDictsResponseBody struct {
	// example:
	//
	// E92BCBB9-3CFE-58DD-8D8C-56DF46AB3BF3
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListDictsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDictsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDictsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDictsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDictsResponseBody) GetResult() []*ListDictsResponseBodyResult {
	return s.Result
}

func (s *ListDictsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDictsResponseBody) SetRequestId(v string) *ListDictsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDictsResponseBody) SetResult(v []*ListDictsResponseBodyResult) *ListDictsResponseBody {
	s.Result = v
	return s
}

func (s *ListDictsResponseBody) SetTotalCount(v int32) *ListDictsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDictsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDictsResponseBodyResult struct {
	// example:
	//
	// http://es-serverless-****.oss-cn-hangzhou.aliyuncs.com/app/es7**190/0/config/analysis-ik/stopword.dic?Expires=1705923089&OSSAccessKeyId=STS.NV18q****UkVp6LNj&Signat
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// example:
	//
	// a.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OSS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
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
