// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleSceneListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSampleSceneListRequest
	GetLang() *string
	SetRegId(v string) *DescribeSampleSceneListRequest
	GetRegId() *string
}

type DescribeSampleSceneListRequest struct {
	// Set the language type for request and response messages, default value is **zh**. Values:
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
}

func (s DescribeSampleSceneListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleSceneListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleSceneListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSampleSceneListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSampleSceneListRequest) SetLang(v string) *DescribeSampleSceneListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSampleSceneListRequest) SetRegId(v string) *DescribeSampleSceneListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSampleSceneListRequest) Validate() error {
	return dara.Validate(s)
}
