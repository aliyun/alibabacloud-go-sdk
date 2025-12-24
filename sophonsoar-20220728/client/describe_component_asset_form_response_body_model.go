// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentAssetFormResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponentAssetForm(v string) *DescribeComponentAssetFormResponseBody
	GetComponentAssetForm() *string
	SetRequestId(v string) *DescribeComponentAssetFormResponseBody
	GetRequestId() *string
}

type DescribeComponentAssetFormResponseBody struct {
	// The metadata of the asset in the component. The value is a JSON array and contains the following fields:
	//
	// 	- **name**: the parameter name.
	//
	// 	- **defaultValue**: the default parameter value.
	//
	// 	- **description**: the parameter description.
	//
	// 	- **required**: indicates whether the parameter is required. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "defaultValue": "",
	//
	//         "description": "assetname",
	//
	//         "name": "assetname",
	//
	//         "required": true
	//
	//     }
	//
	// ]
	ComponentAssetForm *string `json:"ComponentAssetForm,omitempty" xml:"ComponentAssetForm,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9D1651AC-31CC-5CC4-A14E-626B3FCC1022
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentAssetFormResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentAssetFormResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetFormResponseBody) GetComponentAssetForm() *string {
	return s.ComponentAssetForm
}

func (s *DescribeComponentAssetFormResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComponentAssetFormResponseBody) SetComponentAssetForm(v string) *DescribeComponentAssetFormResponseBody {
	s.ComponentAssetForm = &v
	return s
}

func (s *DescribeComponentAssetFormResponseBody) SetRequestId(v string) *DescribeComponentAssetFormResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentAssetFormResponseBody) Validate() error {
	return dara.Validate(s)
}
