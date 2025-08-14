// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchMediaByFaceResponseBody
	GetCode() *string
	SetMediaInfoList(v []*SearchMediaByFaceResponseBodyMediaInfoList) *SearchMediaByFaceResponseBody
	GetMediaInfoList() []*SearchMediaByFaceResponseBodyMediaInfoList
	SetRequestId(v string) *SearchMediaByFaceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SearchMediaByFaceResponseBody
	GetSuccess() *string
	SetTotal(v int64) *SearchMediaByFaceResponseBody
	GetTotal() *int64
}

type SearchMediaByFaceResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The media assets that meet the conditions.
	MediaInfoList []*SearchMediaByFaceResponseBodyMediaInfoList `json:"MediaInfoList,omitempty" xml:"MediaInfoList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7CA7D615-CFB1-5437-9A12-2D185C3EE6CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of data records that meet the specified filter condition.
	//
	// example:
	//
	// 163
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchMediaByFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByFaceResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMediaByFaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchMediaByFaceResponseBody) GetMediaInfoList() []*SearchMediaByFaceResponseBodyMediaInfoList {
	return s.MediaInfoList
}

func (s *SearchMediaByFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMediaByFaceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SearchMediaByFaceResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *SearchMediaByFaceResponseBody) SetCode(v string) *SearchMediaByFaceResponseBody {
	s.Code = &v
	return s
}

func (s *SearchMediaByFaceResponseBody) SetMediaInfoList(v []*SearchMediaByFaceResponseBodyMediaInfoList) *SearchMediaByFaceResponseBody {
	s.MediaInfoList = v
	return s
}

func (s *SearchMediaByFaceResponseBody) SetRequestId(v string) *SearchMediaByFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMediaByFaceResponseBody) SetSuccess(v string) *SearchMediaByFaceResponseBody {
	s.Success = &v
	return s
}

func (s *SearchMediaByFaceResponseBody) SetTotal(v int64) *SearchMediaByFaceResponseBody {
	s.Total = &v
	return s
}

func (s *SearchMediaByFaceResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchMediaByFaceResponseBodyMediaInfoList struct {
	// The ID of the media asset.
	//
	// example:
	//
	// 3b187b3620c8490886cfc2a9578c****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s SearchMediaByFaceResponseBodyMediaInfoList) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByFaceResponseBodyMediaInfoList) GoString() string {
	return s.String()
}

func (s *SearchMediaByFaceResponseBodyMediaInfoList) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaByFaceResponseBodyMediaInfoList) SetMediaId(v string) *SearchMediaByFaceResponseBodyMediaInfoList {
	s.MediaId = &v
	return s
}

func (s *SearchMediaByFaceResponseBodyMediaInfoList) Validate() error {
	return dara.Validate(s)
}
