// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceTDEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceTDEResponseBodyData) *DescribeDBInstanceTDEResponseBody
	GetData() *DescribeDBInstanceTDEResponseBodyData
	SetRequestId(v string) *DescribeDBInstanceTDEResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceTDEResponseBody struct {
	Data *DescribeDBInstanceTDEResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceTDEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTDEResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEResponseBody) GetData() *DescribeDBInstanceTDEResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceTDEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceTDEResponseBody) SetData(v *DescribeDBInstanceTDEResponseBodyData) *DescribeDBInstanceTDEResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceTDEResponseBody) SetRequestId(v string) *DescribeDBInstanceTDEResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceTDEResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceTDEResponseBodyData struct {
	// example:
	//
	// 0
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s DescribeDBInstanceTDEResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTDEResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEResponseBodyData) GetTDEStatus() *string {
	return s.TDEStatus
}

func (s *DescribeDBInstanceTDEResponseBodyData) SetTDEStatus(v string) *DescribeDBInstanceTDEResponseBodyData {
	s.TDEStatus = &v
	return s
}

func (s *DescribeDBInstanceTDEResponseBodyData) Validate() error {
	return dara.Validate(s)
}
