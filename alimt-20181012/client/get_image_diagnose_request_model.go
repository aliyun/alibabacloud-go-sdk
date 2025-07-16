// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDiagnoseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtra(v string) *GetImageDiagnoseRequest
	GetExtra() *string
	SetUrl(v string) *GetImageDiagnoseRequest
	GetUrl() *string
}

type GetImageDiagnoseRequest struct {
	// example:
	//
	// {   "product_id": "1",   "platform": "ae" }
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxxxx.oss-cn-shenzhen.aliyuncs.com/jd/41209/xxxxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetImageDiagnoseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *GetImageDiagnoseRequest) GetExtra() *string {
	return s.Extra
}

func (s *GetImageDiagnoseRequest) GetUrl() *string {
	return s.Url
}

func (s *GetImageDiagnoseRequest) SetExtra(v string) *GetImageDiagnoseRequest {
	s.Extra = &v
	return s
}

func (s *GetImageDiagnoseRequest) SetUrl(v string) *GetImageDiagnoseRequest {
	s.Url = &v
	return s
}

func (s *GetImageDiagnoseRequest) Validate() error {
	return dara.Validate(s)
}
