// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileRecommendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MobileRecommendResponseBody
	GetCode() *string
	SetMessage(v string) *MobileRecommendResponseBody
	GetMessage() *string
	SetRequestId(v string) *MobileRecommendResponseBody
	GetRequestId() *string
	SetResult(v []*MobileRecommendResponseBodyResult) *MobileRecommendResponseBody
	GetResult() []*MobileRecommendResponseBodyResult
}

type MobileRecommendResponseBody struct {
	// Return code of the invocation
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// If an error occurs, the error message will be output.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5C5809B4-F465-52E0-9A8B-61396F9E593B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result.
	Result []*MobileRecommendResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s MobileRecommendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MobileRecommendResponseBody) GoString() string {
	return s.String()
}

func (s *MobileRecommendResponseBody) GetCode() *string {
	return s.Code
}

func (s *MobileRecommendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MobileRecommendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MobileRecommendResponseBody) GetResult() []*MobileRecommendResponseBodyResult {
	return s.Result
}

func (s *MobileRecommendResponseBody) SetCode(v string) *MobileRecommendResponseBody {
	s.Code = &v
	return s
}

func (s *MobileRecommendResponseBody) SetMessage(v string) *MobileRecommendResponseBody {
	s.Message = &v
	return s
}

func (s *MobileRecommendResponseBody) SetRequestId(v string) *MobileRecommendResponseBody {
	s.RequestId = &v
	return s
}

func (s *MobileRecommendResponseBody) SetResult(v []*MobileRecommendResponseBodyResult) *MobileRecommendResponseBody {
	s.Result = v
	return s
}

func (s *MobileRecommendResponseBody) Validate() error {
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

type MobileRecommendResponseBodyResult struct {
	// Author information
	Authors []*string `json:"Authors,omitempty" xml:"Authors,omitempty" type:"Repeated"`
	// Album thumbnail image
	//
	// example:
	//
	// http://img4.kuwo.cn/star/albumcover/120/78/77/1688821132.jpg
	Cover *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// Third-party ID of the content
	//
	// example:
	//
	// 550144364
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// Source of the content
	//
	// example:
	//
	// KG
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Content title
	//
	// example:
	//
	// 稻香
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s MobileRecommendResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s MobileRecommendResponseBodyResult) GoString() string {
	return s.String()
}

func (s *MobileRecommendResponseBodyResult) GetAuthors() []*string {
	return s.Authors
}

func (s *MobileRecommendResponseBodyResult) GetCover() *string {
	return s.Cover
}

func (s *MobileRecommendResponseBodyResult) GetRawId() *string {
	return s.RawId
}

func (s *MobileRecommendResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *MobileRecommendResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *MobileRecommendResponseBodyResult) SetAuthors(v []*string) *MobileRecommendResponseBodyResult {
	s.Authors = v
	return s
}

func (s *MobileRecommendResponseBodyResult) SetCover(v string) *MobileRecommendResponseBodyResult {
	s.Cover = &v
	return s
}

func (s *MobileRecommendResponseBodyResult) SetRawId(v string) *MobileRecommendResponseBodyResult {
	s.RawId = &v
	return s
}

func (s *MobileRecommendResponseBodyResult) SetSource(v string) *MobileRecommendResponseBodyResult {
	s.Source = &v
	return s
}

func (s *MobileRecommendResponseBodyResult) SetTitle(v string) *MobileRecommendResponseBodyResult {
	s.Title = &v
	return s
}

func (s *MobileRecommendResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
