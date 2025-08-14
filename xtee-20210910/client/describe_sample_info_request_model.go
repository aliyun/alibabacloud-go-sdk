// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSampleInfoRequest
	GetLang() *string
	SetId(v int64) *DescribeSampleInfoRequest
	GetId() *int64
	SetRegId(v string) *DescribeSampleInfoRequest
	GetRegId() *string
	SetVersions(v int32) *DescribeSampleInfoRequest
	GetVersions() *int32
}

type DescribeSampleInfoRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Version number.
	//
	// example:
	//
	// 1
	Versions *int32 `json:"versions,omitempty" xml:"versions,omitempty"`
}

func (s DescribeSampleInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSampleInfoRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeSampleInfoRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSampleInfoRequest) GetVersions() *int32 {
	return s.Versions
}

func (s *DescribeSampleInfoRequest) SetLang(v string) *DescribeSampleInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSampleInfoRequest) SetId(v int64) *DescribeSampleInfoRequest {
	s.Id = &v
	return s
}

func (s *DescribeSampleInfoRequest) SetRegId(v string) *DescribeSampleInfoRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSampleInfoRequest) SetVersions(v int32) *DescribeSampleInfoRequest {
	s.Versions = &v
	return s
}

func (s *DescribeSampleInfoRequest) Validate() error {
	return dara.Validate(s)
}
