// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSliderCaptchaPagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPages(v []*ListSliderCaptchaPagesResponseBodyPages) *ListSliderCaptchaPagesResponseBody
	GetPages() []*ListSliderCaptchaPagesResponseBodyPages
	SetRequestId(v string) *ListSliderCaptchaPagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListSliderCaptchaPagesResponseBody
	GetTotalCount() *int64
}

type ListSliderCaptchaPagesResponseBody struct {
	// The list of custom slider CAPTCHA pages.
	Pages []*ListSliderCaptchaPagesResponseBodyPages `json:"Pages,omitempty" xml:"Pages,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSliderCaptchaPagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSliderCaptchaPagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSliderCaptchaPagesResponseBody) GetPages() []*ListSliderCaptchaPagesResponseBodyPages {
	return s.Pages
}

func (s *ListSliderCaptchaPagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSliderCaptchaPagesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSliderCaptchaPagesResponseBody) SetPages(v []*ListSliderCaptchaPagesResponseBodyPages) *ListSliderCaptchaPagesResponseBody {
	s.Pages = v
	return s
}

func (s *ListSliderCaptchaPagesResponseBody) SetRequestId(v string) *ListSliderCaptchaPagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSliderCaptchaPagesResponseBody) SetTotalCount(v int64) *ListSliderCaptchaPagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSliderCaptchaPagesResponseBody) Validate() error {
	if s.Pages != nil {
		for _, item := range s.Pages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSliderCaptchaPagesResponseBodyPages struct {
	// The ID of the custom slider CAPTCHA page.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the custom slider CAPTCHA page.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of associated site IDs.
	SiteIds []*int64 `json:"SiteIds,omitempty" xml:"SiteIds,omitempty" type:"Repeated"`
}

func (s ListSliderCaptchaPagesResponseBodyPages) String() string {
	return dara.Prettify(s)
}

func (s ListSliderCaptchaPagesResponseBodyPages) GoString() string {
	return s.String()
}

func (s *ListSliderCaptchaPagesResponseBodyPages) GetId() *int64 {
	return s.Id
}

func (s *ListSliderCaptchaPagesResponseBodyPages) GetName() *string {
	return s.Name
}

func (s *ListSliderCaptchaPagesResponseBodyPages) GetSiteIds() []*int64 {
	return s.SiteIds
}

func (s *ListSliderCaptchaPagesResponseBodyPages) SetId(v int64) *ListSliderCaptchaPagesResponseBodyPages {
	s.Id = &v
	return s
}

func (s *ListSliderCaptchaPagesResponseBodyPages) SetName(v string) *ListSliderCaptchaPagesResponseBodyPages {
	s.Name = &v
	return s
}

func (s *ListSliderCaptchaPagesResponseBodyPages) SetSiteIds(v []*int64) *ListSliderCaptchaPagesResponseBodyPages {
	s.SiteIds = v
	return s
}

func (s *ListSliderCaptchaPagesResponseBodyPages) Validate() error {
	return dara.Validate(s)
}
