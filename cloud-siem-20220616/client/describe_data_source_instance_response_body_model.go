// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDataSourceInstanceResponseBodyData) *DescribeDataSourceInstanceResponseBody
	GetData() *DescribeDataSourceInstanceResponseBodyData
	SetRequestId(v string) *DescribeDataSourceInstanceResponseBody
	GetRequestId() *string
}

type DescribeDataSourceInstanceResponseBody struct {
	// The data returned.
	Data *DescribeDataSourceInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataSourceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceInstanceResponseBody) GetData() *DescribeDataSourceInstanceResponseBodyData {
	return s.Data
}

func (s *DescribeDataSourceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataSourceInstanceResponseBody) SetData(v *DescribeDataSourceInstanceResponseBodyData) *DescribeDataSourceInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDataSourceInstanceResponseBody) SetRequestId(v string) *DescribeDataSourceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataSourceInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataSourceInstanceResponseBodyData struct {
	// The ID of the cloud account.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The code of the cloud service provider. Valid values:
	//
	// 	- qcloud: Tencent Cloud
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The ID of the data source. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
	// The parameters of the data source.
	DataSourceInstanceParams []*DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams `json:"DataSourceInstanceParams,omitempty" xml:"DataSourceInstanceParams,omitempty" type:"Repeated"`
}

func (s DescribeDataSourceInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceInstanceResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeDataSourceInstanceResponseBodyData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DescribeDataSourceInstanceResponseBodyData) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *DescribeDataSourceInstanceResponseBodyData) GetDataSourceInstanceParams() []*DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams {
	return s.DataSourceInstanceParams
}

func (s *DescribeDataSourceInstanceResponseBodyData) SetAccountId(v string) *DescribeDataSourceInstanceResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *DescribeDataSourceInstanceResponseBodyData) SetCloudCode(v string) *DescribeDataSourceInstanceResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *DescribeDataSourceInstanceResponseBodyData) SetDataSourceInstanceId(v string) *DescribeDataSourceInstanceResponseBodyData {
	s.DataSourceInstanceId = &v
	return s
}

func (s *DescribeDataSourceInstanceResponseBodyData) SetDataSourceInstanceParams(v []*DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) *DescribeDataSourceInstanceResponseBodyData {
	s.DataSourceInstanceParams = v
	return s
}

func (s *DescribeDataSourceInstanceResponseBodyData) Validate() error {
	if s.DataSourceInstanceParams != nil {
		for _, item := range s.DataSourceInstanceParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams struct {
	// The code of the parameter.
	//
	// example:
	//
	// region_code
	ParaCode *string `json:"ParaCode,omitempty" xml:"ParaCode,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// ap-guangzhou
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
}

func (s DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) GetParaCode() *string {
	return s.ParaCode
}

func (s *DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) GetParaValue() *string {
	return s.ParaValue
}

func (s *DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) SetParaCode(v string) *DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams {
	s.ParaCode = &v
	return s
}

func (s *DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) SetParaValue(v string) *DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams {
	s.ParaValue = &v
	return s
}

func (s *DescribeDataSourceInstanceResponseBodyDataDataSourceInstanceParams) Validate() error {
	return dara.Validate(s)
}
