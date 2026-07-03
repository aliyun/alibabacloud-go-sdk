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
	// The ID of the request.
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
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourceTypesResponseBodyData struct {
	// The code of the multicloud service.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The type of the data source. Valid values:
	//
	// - obs: Huawei Cloud OBS.
	//
	// - wafApi: Tencent Cloud WAF download API.
	//
	// - ckafka: Tencent Cloud CKafka.
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
