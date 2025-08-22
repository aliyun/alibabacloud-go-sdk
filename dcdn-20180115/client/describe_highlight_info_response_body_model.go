// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighlightInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataModule(v []*DescribeHighlightInfoResponseBodyDataModule) *DescribeHighlightInfoResponseBody
	GetDataModule() []*DescribeHighlightInfoResponseBodyDataModule
	SetRequestId(v string) *DescribeHighlightInfoResponseBody
	GetRequestId() *string
}

type DescribeHighlightInfoResponseBody struct {
	// The data model of the highlighted data.
	DataModule []*DescribeHighlightInfoResponseBodyDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHighlightInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighlightInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHighlightInfoResponseBody) GetDataModule() []*DescribeHighlightInfoResponseBodyDataModule {
	return s.DataModule
}

func (s *DescribeHighlightInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHighlightInfoResponseBody) SetDataModule(v []*DescribeHighlightInfoResponseBodyDataModule) *DescribeHighlightInfoResponseBody {
	s.DataModule = v
	return s
}

func (s *DescribeHighlightInfoResponseBody) SetRequestId(v string) *DescribeHighlightInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHighlightInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHighlightInfoResponseBodyDataModule struct {
	// The highlighted data.
	//
	// example:
	//
	// [\\"data:image/php;base64\\"]
	Hit *string `json:"Hit,omitempty" xml:"Hit,omitempty"`
	// The type of the highlighted data.
	//
	// example:
	//
	// URL
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The complete data.
	//
	// example:
	//
	// data:image/php;base64,PD9waHAXXXXXXanVzdHR0dHXXXXXB0ZXN0Ijs/Pg==
	Raw *string `json:"Raw,omitempty" xml:"Raw,omitempty"`
}

func (s DescribeHighlightInfoResponseBodyDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighlightInfoResponseBodyDataModule) GoString() string {
	return s.String()
}

func (s *DescribeHighlightInfoResponseBodyDataModule) GetHit() *string {
	return s.Hit
}

func (s *DescribeHighlightInfoResponseBodyDataModule) GetKey() *string {
	return s.Key
}

func (s *DescribeHighlightInfoResponseBodyDataModule) GetRaw() *string {
	return s.Raw
}

func (s *DescribeHighlightInfoResponseBodyDataModule) SetHit(v string) *DescribeHighlightInfoResponseBodyDataModule {
	s.Hit = &v
	return s
}

func (s *DescribeHighlightInfoResponseBodyDataModule) SetKey(v string) *DescribeHighlightInfoResponseBodyDataModule {
	s.Key = &v
	return s
}

func (s *DescribeHighlightInfoResponseBodyDataModule) SetRaw(v string) *DescribeHighlightInfoResponseBodyDataModule {
	s.Raw = &v
	return s
}

func (s *DescribeHighlightInfoResponseBodyDataModule) Validate() error {
	return dara.Validate(s)
}
