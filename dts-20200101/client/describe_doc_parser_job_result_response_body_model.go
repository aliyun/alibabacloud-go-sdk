// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocParserJobResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContentList(v []*DescribeDocParserJobResultResponseBodyContentList) *DescribeDocParserJobResultResponseBody
	GetContentList() []*DescribeDocParserJobResultResponseBodyContentList
	SetDynamicCode(v string) *DescribeDocParserJobResultResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeDocParserJobResultResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeDocParserJobResultResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDocParserJobResultResponseBody
	GetErrMessage() *string
	SetFileUrl(v string) *DescribeDocParserJobResultResponseBody
	GetFileUrl() *string
	SetHttpStatusCode(v int32) *DescribeDocParserJobResultResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeDocParserJobResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDocParserJobResultResponseBody
	GetSuccess() *bool
}

type DescribeDocParserJobResultResponseBody struct {
	ContentList []*DescribeDocParserJobResultResponseBodyContentList `json:"ContentList,omitempty" xml:"ContentList,omitempty" type:"Repeated"`
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// https://oss-cn-hangzhou.aliyuncs.com/806a_209584525031252870_078f1180f27b4c069c0f271758aa****
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// C166D79D-436B-45F0-B5A5-25E1959F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDocParserJobResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobResultResponseBody) GetContentList() []*DescribeDocParserJobResultResponseBodyContentList {
	return s.ContentList
}

func (s *DescribeDocParserJobResultResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeDocParserJobResultResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDocParserJobResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDocParserJobResultResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDocParserJobResultResponseBody) GetFileUrl() *string {
	return s.FileUrl
}

func (s *DescribeDocParserJobResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDocParserJobResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDocParserJobResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDocParserJobResultResponseBody) SetContentList(v []*DescribeDocParserJobResultResponseBodyContentList) *DescribeDocParserJobResultResponseBody {
	s.ContentList = v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetDynamicCode(v string) *DescribeDocParserJobResultResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetDynamicMessage(v string) *DescribeDocParserJobResultResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetErrCode(v string) *DescribeDocParserJobResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetErrMessage(v string) *DescribeDocParserJobResultResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetFileUrl(v string) *DescribeDocParserJobResultResponseBody {
	s.FileUrl = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetHttpStatusCode(v int32) *DescribeDocParserJobResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetRequestId(v string) *DescribeDocParserJobResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) SetSuccess(v bool) *DescribeDocParserJobResultResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBody) Validate() error {
	if s.ContentList != nil {
		for _, item := range s.ContentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDocParserJobResultResponseBodyContentList struct {
	// example:
	//
	// 	- Demo 	- *	- Demo title **
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDocParserJobResultResponseBodyContentList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobResultResponseBodyContentList) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobResultResponseBodyContentList) GetContent() *string {
	return s.Content
}

func (s *DescribeDocParserJobResultResponseBodyContentList) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDocParserJobResultResponseBodyContentList) SetContent(v string) *DescribeDocParserJobResultResponseBodyContentList {
	s.Content = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBodyContentList) SetPageNumber(v int32) *DescribeDocParserJobResultResponseBodyContentList {
	s.PageNumber = &v
	return s
}

func (s *DescribeDocParserJobResultResponseBodyContentList) Validate() error {
	return dara.Validate(s)
}
