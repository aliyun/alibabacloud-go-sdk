// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageTextsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOCRContents(v []*OCRContents) *DetectImageTextsResponseBody
	GetOCRContents() []*OCRContents
	SetOCRTexts(v string) *DetectImageTextsResponseBody
	GetOCRTexts() *string
	SetRequestId(v string) *DetectImageTextsResponseBody
	GetRequestId() *string
}

type DetectImageTextsResponseBody struct {
	// OCR text blocks.
	OCRContents []*OCRContents `json:"OCRContents,omitempty" xml:"OCRContents,omitempty" type:"Repeated"`
	// The full Optical Character Recognition (OCR) text, which is spliced by using the content of OCRContents.
	OCRTexts *string `json:"OCRTexts,omitempty" xml:"OCRTexts,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1B3D5E0A-D8B8-4DA0-8127-ED32C851****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageTextsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectImageTextsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageTextsResponseBody) GetOCRContents() []*OCRContents {
	return s.OCRContents
}

func (s *DetectImageTextsResponseBody) GetOCRTexts() *string {
	return s.OCRTexts
}

func (s *DetectImageTextsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectImageTextsResponseBody) SetOCRContents(v []*OCRContents) *DetectImageTextsResponseBody {
	s.OCRContents = v
	return s
}

func (s *DetectImageTextsResponseBody) SetOCRTexts(v string) *DetectImageTextsResponseBody {
	s.OCRTexts = &v
	return s
}

func (s *DetectImageTextsResponseBody) SetRequestId(v string) *DetectImageTextsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageTextsResponseBody) Validate() error {
	if s.OCRContents != nil {
		for _, item := range s.OCRContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
