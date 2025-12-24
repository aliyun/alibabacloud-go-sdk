// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefineMaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RefineMaskResponseBodyData) *RefineMaskResponseBody
	GetData() *RefineMaskResponseBodyData
	SetRequestId(v string) *RefineMaskResponseBody
	GetRequestId() *string
}

type RefineMaskResponseBody struct {
	Data *RefineMaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// E948F80B-86D9-54E0-9FF9-ACF3B1DA83C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefineMaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefineMaskResponseBody) GoString() string {
	return s.String()
}

func (s *RefineMaskResponseBody) GetData() *RefineMaskResponseBodyData {
	return s.Data
}

func (s *RefineMaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefineMaskResponseBody) SetData(v *RefineMaskResponseBodyData) *RefineMaskResponseBody {
	s.Data = v
	return s
}

func (s *RefineMaskResponseBody) SetRequestId(v string) *RefineMaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefineMaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefineMaskResponseBodyData struct {
	Elements []*RefineMaskResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RefineMaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RefineMaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefineMaskResponseBodyData) GetElements() []*RefineMaskResponseBodyDataElements {
	return s.Elements
}

func (s *RefineMaskResponseBodyData) SetElements(v []*RefineMaskResponseBodyDataElements) *RefineMaskResponseBodyData {
	s.Elements = v
	return s
}

func (s *RefineMaskResponseBodyData) Validate() error {
	if s.Elements != nil {
		for _, item := range s.Elements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RefineMaskResponseBodyDataElements struct {
	// example:
	//
	// http://algo-app-taobao-mm-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/pixelai-portrait-beauty%2F2020_03_04%2F61f544a1a5004c88a2bf29452db494e9.jpeg?OSSAccessKeyId=LTAI4Fmdm1gQonFLrghJ****&Expires=1583406122&Signature=Heet1ivG0xFP3YlO6usvd0pmrH****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RefineMaskResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s RefineMaskResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RefineMaskResponseBodyDataElements) GetImageURL() *string {
	return s.ImageURL
}

func (s *RefineMaskResponseBodyDataElements) SetImageURL(v string) *RefineMaskResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *RefineMaskResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}
