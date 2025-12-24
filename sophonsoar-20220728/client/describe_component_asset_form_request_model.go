// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentAssetFormRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentName(v string) *DescribeComponentAssetFormRequest
	GetComponentName() *string
	SetLang(v string) *DescribeComponentAssetFormRequest
	GetLang() *string
}

type DescribeComponentAssetFormRequest struct {
	// The component name.
	//
	// This parameter is required.
	//
	// example:
	//
	// python3
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeComponentAssetFormRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentAssetFormRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetFormRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *DescribeComponentAssetFormRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeComponentAssetFormRequest) SetComponentName(v string) *DescribeComponentAssetFormRequest {
	s.ComponentName = &v
	return s
}

func (s *DescribeComponentAssetFormRequest) SetLang(v string) *DescribeComponentAssetFormRequest {
	s.Lang = &v
	return s
}

func (s *DescribeComponentAssetFormRequest) Validate() error {
	return dara.Validate(s)
}
