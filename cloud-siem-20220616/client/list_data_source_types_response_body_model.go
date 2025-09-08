// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataSourceTypesResponseBodyData) *ListDataSourceTypesResponseBody
	GetData() []*ListDataSourceTypesResponseBodyData
	SetRequestId(v string) *ListDataSourceTypesResponseBody
	GetRequestId() *string
}

type ListDataSourceTypesResponseBody struct {
	// The data returned.
	Data []*ListDataSourceTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTypesResponseBody) GetData() []*ListDataSourceTypesResponseBodyData {
	return s.Data
}

func (s *ListDataSourceTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourceTypesResponseBody) SetData(v []*ListDataSourceTypesResponseBodyData) *ListDataSourceTypesResponseBody {
	s.Data = v
	return s
}

func (s *ListDataSourceTypesResponseBody) SetRequestId(v string) *ListDataSourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataSourceTypesResponseBodyData struct {
	// The code of the third-party cloud service.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- obs: Huawei Cloud Object Storage Service (OBS)
	//
	// 	- wafApi: download API of Tencent Cloud Web Application Firewall (WAF)
	//
	// 	- ckafka: Tencent Cloud Kafka (CKafka)
	//
	// example:
	//
	// obs
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
}

func (s ListDataSourceTypesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataSourceTypesResponseBodyData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListDataSourceTypesResponseBodyData) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListDataSourceTypesResponseBodyData) SetCloudCode(v string) *ListDataSourceTypesResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *ListDataSourceTypesResponseBodyData) SetDataSourceType(v string) *ListDataSourceTypesResponseBodyData {
	s.DataSourceType = &v
	return s
}

func (s *ListDataSourceTypesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
