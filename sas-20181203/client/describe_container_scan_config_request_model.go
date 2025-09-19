// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerScanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeContainerScanConfigRequest
	GetLang() *string
}

type DescribeContainerScanConfigRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeContainerScanConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerScanConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerScanConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeContainerScanConfigRequest) SetLang(v string) *DescribeContainerScanConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeContainerScanConfigRequest) Validate() error {
	return dara.Validate(s)
}
