// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePageDocumentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDocs(v []*DescribePageDocumentsResponseBodyDocs) *DescribePageDocumentsResponseBody
	GetDocs() []*DescribePageDocumentsResponseBodyDocs
	SetImageUrl(v string) *DescribePageDocumentsResponseBody
	GetImageUrl() *string
	SetModule(v string) *DescribePageDocumentsResponseBody
	GetModule() *string
	SetMore(v *DescribePageDocumentsResponseBodyMore) *DescribePageDocumentsResponseBody
	GetMore() *DescribePageDocumentsResponseBodyMore
	SetRequestId(v string) *DescribePageDocumentsResponseBody
	GetRequestId() *string
}

type DescribePageDocumentsResponseBody struct {
	Docs []*DescribePageDocumentsResponseBodyDocs `json:"Docs,omitempty" xml:"Docs,omitempty" type:"Repeated"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1E4FCAuT2gK0jSZFvXXXnFXXa-3399-662.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// sg_server
	Module *string                                `json:"Module,omitempty" xml:"Module,omitempty"`
	More   *DescribePageDocumentsResponseBodyMore `json:"More,omitempty" xml:"More,omitempty" type:"Struct"`
	// example:
	//
	// 7C81E1AD-08C0-5E09-853B-FDC77B90****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePageDocumentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePageDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePageDocumentsResponseBody) GetDocs() []*DescribePageDocumentsResponseBodyDocs {
	return s.Docs
}

func (s *DescribePageDocumentsResponseBody) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribePageDocumentsResponseBody) GetModule() *string {
	return s.Module
}

func (s *DescribePageDocumentsResponseBody) GetMore() *DescribePageDocumentsResponseBodyMore {
	return s.More
}

func (s *DescribePageDocumentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePageDocumentsResponseBody) SetDocs(v []*DescribePageDocumentsResponseBodyDocs) *DescribePageDocumentsResponseBody {
	s.Docs = v
	return s
}

func (s *DescribePageDocumentsResponseBody) SetImageUrl(v string) *DescribePageDocumentsResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *DescribePageDocumentsResponseBody) SetModule(v string) *DescribePageDocumentsResponseBody {
	s.Module = &v
	return s
}

func (s *DescribePageDocumentsResponseBody) SetMore(v *DescribePageDocumentsResponseBodyMore) *DescribePageDocumentsResponseBody {
	s.More = v
	return s
}

func (s *DescribePageDocumentsResponseBody) SetRequestId(v string) *DescribePageDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePageDocumentsResponseBody) Validate() error {
	if s.Docs != nil {
		for _, item := range s.Docs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.More != nil {
		if err := s.More.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePageDocumentsResponseBodyDocs struct {
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://notify-center-test.ybaobx.com/webhook/arms
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribePageDocumentsResponseBodyDocs) String() string {
	return dara.Prettify(s)
}

func (s DescribePageDocumentsResponseBodyDocs) GoString() string {
	return s.String()
}

func (s *DescribePageDocumentsResponseBodyDocs) GetTitle() *string {
	return s.Title
}

func (s *DescribePageDocumentsResponseBodyDocs) GetUrl() *string {
	return s.Url
}

func (s *DescribePageDocumentsResponseBodyDocs) SetTitle(v string) *DescribePageDocumentsResponseBodyDocs {
	s.Title = &v
	return s
}

func (s *DescribePageDocumentsResponseBodyDocs) SetUrl(v string) *DescribePageDocumentsResponseBodyDocs {
	s.Url = &v
	return s
}

func (s *DescribePageDocumentsResponseBodyDocs) Validate() error {
	return dara.Validate(s)
}

type DescribePageDocumentsResponseBodyMore struct {
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://zjysfy.womanhospital.cn/pub/hos/0/noneBackGround.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribePageDocumentsResponseBodyMore) String() string {
	return dara.Prettify(s)
}

func (s DescribePageDocumentsResponseBodyMore) GoString() string {
	return s.String()
}

func (s *DescribePageDocumentsResponseBodyMore) GetTitle() *string {
	return s.Title
}

func (s *DescribePageDocumentsResponseBodyMore) GetUrl() *string {
	return s.Url
}

func (s *DescribePageDocumentsResponseBodyMore) SetTitle(v string) *DescribePageDocumentsResponseBodyMore {
	s.Title = &v
	return s
}

func (s *DescribePageDocumentsResponseBodyMore) SetUrl(v string) *DescribePageDocumentsResponseBodyMore {
	s.Url = &v
	return s
}

func (s *DescribePageDocumentsResponseBodyMore) Validate() error {
	return dara.Validate(s)
}
