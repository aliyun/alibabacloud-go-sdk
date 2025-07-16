// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceConfigResponseBodyData) *DescribeDBInstanceConfigResponseBody
	GetData() *DescribeDBInstanceConfigResponseBodyData
	SetRequestId(v string) *DescribeDBInstanceConfigResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceConfigResponseBody struct {
	Data *DescribeDBInstanceConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigResponseBody) GetData() *DescribeDBInstanceConfigResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceConfigResponseBody) SetData(v *DescribeDBInstanceConfigResponseBodyData) *DescribeDBInstanceConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceConfigResponseBody) SetRequestId(v string) *DescribeDBInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceConfigResponseBodyData struct {
	// example:
	//
	// htap
	ConfigName     *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	ConfigValue    *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
}

func (s DescribeDBInstanceConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigResponseBodyData) GetConfigName() *string {
	return s.ConfigName
}

func (s *DescribeDBInstanceConfigResponseBodyData) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeDBInstanceConfigResponseBodyData) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeDBInstanceConfigResponseBodyData) SetConfigName(v string) *DescribeDBInstanceConfigResponseBodyData {
	s.ConfigName = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyData) SetConfigValue(v string) *DescribeDBInstanceConfigResponseBodyData {
	s.ConfigValue = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyData) SetDbInstanceName(v string) *DescribeDBInstanceConfigResponseBodyData {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
