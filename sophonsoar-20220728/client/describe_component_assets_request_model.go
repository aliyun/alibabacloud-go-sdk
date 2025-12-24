// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentName(v string) *DescribeComponentAssetsRequest
	GetComponentName() *string
	SetLang(v string) *DescribeComponentAssetsRequest
	GetLang() *string
}

type DescribeComponentAssetsRequest struct {
	// The name of the component.
	//
	// This parameter is required.
	//
	// example:
	//
	// python3
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The language of the content within the request and response. Valid values:
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

func (s DescribeComponentAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *DescribeComponentAssetsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeComponentAssetsRequest) SetComponentName(v string) *DescribeComponentAssetsRequest {
	s.ComponentName = &v
	return s
}

func (s *DescribeComponentAssetsRequest) SetLang(v string) *DescribeComponentAssetsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeComponentAssetsRequest) Validate() error {
	return dara.Validate(s)
}
