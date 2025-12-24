// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponents(v string) *DescribeComponentListResponseBody
	GetComponents() *string
	SetRequestId(v string) *DescribeComponentListResponseBody
	GetRequestId() *string
}

type DescribeComponentListResponseBody struct {
	// The information about the components. The value is a JSON array.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "actions": [
	//
	//             {
	//
	//                 "description": "mysql component",
	//
	//                 "name": "storeIdb",
	//
	//                 "parameters": [
	//
	//                     {
	//
	//                         "description": "update the mysql db",
	//
	//                         "name": "updateSql",
	//
	//                         "required": false
	//
	//                     }
	//
	//                 ]
	//
	//             }
	//
	//         ],
	//
	//         "basic": {
	//
	//             "description": "mysq sql component for 5.6",
	//
	//             "logo": "https://img.alicdn.com/tfs/TB1H89IpH3nBKNjSZFMXXaUSFXa-200-200.svg",
	//
	//             "name": "Mysql"
	//
	//         }
	//
	//     }
	//
	// ]
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B0A255B3-495C-56FB-8B6B-DB073F80388A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentListResponseBody) GetComponents() *string {
	return s.Components
}

func (s *DescribeComponentListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComponentListResponseBody) SetComponents(v string) *DescribeComponentListResponseBody {
	s.Components = &v
	return s
}

func (s *DescribeComponentListResponseBody) SetRequestId(v string) *DescribeComponentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentListResponseBody) Validate() error {
	return dara.Validate(s)
}
