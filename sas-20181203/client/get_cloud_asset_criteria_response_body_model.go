// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAssetCriteriaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaList(v []*GetCloudAssetCriteriaResponseBodyCriteriaList) *GetCloudAssetCriteriaResponseBody
	GetCriteriaList() []*GetCloudAssetCriteriaResponseBodyCriteriaList
	SetRequestId(v string) *GetCloudAssetCriteriaResponseBody
	GetRequestId() *string
}

type GetCloudAssetCriteriaResponseBody struct {
	// An array consisting of the information about the filter conditions that are used to search for cloud assets.
	CriteriaList []*GetCloudAssetCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCloudAssetCriteriaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudAssetCriteriaResponseBody) GetCriteriaList() []*GetCloudAssetCriteriaResponseBodyCriteriaList {
	return s.CriteriaList
}

func (s *GetCloudAssetCriteriaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudAssetCriteriaResponseBody) SetCriteriaList(v []*GetCloudAssetCriteriaResponseBodyCriteriaList) *GetCloudAssetCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *GetCloudAssetCriteriaResponseBody) SetRequestId(v string) *GetCloudAssetCriteriaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudAssetCriteriaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCloudAssetCriteriaResponseBodyCriteriaList struct {
	// The structured attribute values of the assets that match the keyword. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **vendor**: providers
	//
	// 	- **regionIds**: IDs of supported regions
	//
	// example:
	//
	// [{"vendor":0,"regionIds":{"default":["ap-southeast-1","ap-northeast-2","ap-southeast-3","ap-southeast-5","ap-southeast-7","me-central-1"]}},{"vendor":1,"regionIds":{"default":["outside-of-aliyun"]}}]
	MultiValues *string `json:"MultiValues,omitempty" xml:"MultiValues,omitempty"`
	// The name of the filter condition. Valid values:
	//
	// 	- **instanceId**: the ID of the instance
	//
	// 	- **instanceName**: the name of an instance
	//
	// 	- **internetIp**: the public IP address
	//
	// 	- **riskStatus**: the risk status
	//
	// 	- **vendorRegionId**: the region ID by service provider
	//
	// example:
	//
	// instanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the filter condition. Valid values:
	//
	// 	- **input**: The filter condition needs to be specified.
	//
	// 	- **select**: The filter condition is an option that can be selected from the drop-down list.
	//
	// example:
	//
	// select
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The values of the search condition. This parameter is returned only if the value of **Type*	- is **select**.
	//
	// >  If the value of **Type*	- is **input**, the value of this parameter is an empty string.
	//
	// example:
	//
	// fvt*
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s GetCloudAssetCriteriaResponseBodyCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *GetCloudAssetCriteriaResponseBodyCriteriaList) GetMultiValues() *string {
	return s.MultiValues
}

func (s *GetCloudAssetCriteriaResponseBodyCriteriaList) GetName() *string {
	return s.Name
}

func (s *GetCloudAssetCriteriaResponseBodyCriteriaList) GetType() *string {
	return s.Type
}

func (s *GetCloudAssetCriteriaResponseBodyCriteriaList) GetValues() *string {
	return s.Values
}

func (s *GetCloudAssetCriteriaResponseBodyCriteriaList) SetMultiValues(v string) *GetCloudAssetCriteriaResponseBodyCriteriaList {
	s.MultiValues = &v
	return s
}

func (s *GetCloudAssetCriteriaResponseBodyCriteriaList) SetName(v string) *GetCloudAssetCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

func (s *GetCloudAssetCriteriaResponseBodyCriteriaList) SetType(v string) *GetCloudAssetCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *GetCloudAssetCriteriaResponseBodyCriteriaList) SetValues(v string) *GetCloudAssetCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *GetCloudAssetCriteriaResponseBodyCriteriaList) Validate() error {
	return dara.Validate(s)
}
