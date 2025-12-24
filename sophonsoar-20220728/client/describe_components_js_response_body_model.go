// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentsJsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponentsJs(v string) *DescribeComponentsJsResponseBody
	GetComponentsJs() *string
	SetRequestId(v string) *DescribeComponentsJsResponseBody
	GetRequestId() *string
}

type DescribeComponentsJsResponseBody struct {
	// The configuration of the JavaScript file for the component.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "js": "https://xxxxx.oss-cn-zhangjiakou.aliyuncs.com/componentUpload/xxxxx",
	//
	//         "name": "python3",
	//
	//         "ownType": "sys"
	//
	//     }
	//
	// ]
	ComponentsJs *string `json:"ComponentsJs,omitempty" xml:"ComponentsJs,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 58A518BC-E4A8-5BD7-AFEA-366046ED9073
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentsJsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentsJsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentsJsResponseBody) GetComponentsJs() *string {
	return s.ComponentsJs
}

func (s *DescribeComponentsJsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComponentsJsResponseBody) SetComponentsJs(v string) *DescribeComponentsJsResponseBody {
	s.ComponentsJs = &v
	return s
}

func (s *DescribeComponentsJsResponseBody) SetRequestId(v string) *DescribeComponentsJsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentsJsResponseBody) Validate() error {
	return dara.Validate(s)
}
