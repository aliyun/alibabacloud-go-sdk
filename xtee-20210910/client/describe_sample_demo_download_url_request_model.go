// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDemoDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSampleDemoDownloadUrlRequest
	GetLang() *string
	SetRegId(v string) *DescribeSampleDemoDownloadUrlRequest
	GetRegId() *string
	SetScene(v string) *DescribeSampleDemoDownloadUrlRequest
	GetScene() *string
}

type DescribeSampleDemoDownloadUrlRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Scene code
	//
	// This parameter is required.
	//
	// example:
	//
	// coupon_abuse_detection
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
}

func (s DescribeSampleDemoDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDemoDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleDemoDownloadUrlRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSampleDemoDownloadUrlRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSampleDemoDownloadUrlRequest) GetScene() *string {
	return s.Scene
}

func (s *DescribeSampleDemoDownloadUrlRequest) SetLang(v string) *DescribeSampleDemoDownloadUrlRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSampleDemoDownloadUrlRequest) SetRegId(v string) *DescribeSampleDemoDownloadUrlRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSampleDemoDownloadUrlRequest) SetScene(v string) *DescribeSampleDemoDownloadUrlRequest {
	s.Scene = &v
	return s
}

func (s *DescribeSampleDemoDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
