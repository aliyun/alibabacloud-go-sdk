// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationRequest
	GetAppId() *string
	SetAppVersions(v string) *DescribeApplicationRequest
	GetAppVersions() *string
	SetLevel(v string) *DescribeApplicationRequest
	GetLevel() *string
	SetOutDetailStatParams(v string) *DescribeApplicationRequest
	GetOutDetailStatParams() *string
	SetResourceSelector(v string) *DescribeApplicationRequest
	GetResourceSelector() *string
}

type DescribeApplicationRequest struct {
	// The ID of the application. You can obtain the ID by calling the ListApplications operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2bac6f4-75dc-455e-8389-2dc8e47526d3
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The business service version numbers of the application. Multiple values are supported and separated by commas (,). If you enter All, data of all versions is returned. By default, only data associated with the stable version number is returned.
	//
	// example:
	//
	// v1,v2
	AppVersions *string `json:"AppVersions,omitempty" xml:"AppVersions,omitempty"`
	// The regional statistical level of the application edge resources. Valid values:
	//
	// - National: nationwide
	//
	// - Big: region
	//
	// - Middle: province
	//
	// - Small: city
	//
	// - RegionId: edge node
	//
	// Default value: National.
	//
	// example:
	//
	// National
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Specifies whether to return information about statistical resource instances and container pods. The value is a JSON string. By default, all information is returned.
	//
	// example:
	//
	// {\\"appInfo\\":true,\\"detailStat\\": true, \\"appVersionStat\\": true, \\"districtStat\\":true ,\\"instanceStat\\": true, \\"podCountStat\\": true}
	OutDetailStatParams *string `json:"OutDetailStatParams,omitempty" xml:"OutDetailStatParams,omitempty"`
	// The resource filter conditions.
	//
	// example:
	//
	// [{\\"regionCode\\": \\"cn-wuxi-telecom_unicom_cmcc-3\\",    \\"ispCode\\": \\"telecom\\",    \\"count\\": 2	},{    \\"regionCode\\": \\"cn-shanghai-cmcc\\",    \\"count\\": 4	}]
	ResourceSelector *string `json:"ResourceSelector,omitempty" xml:"ResourceSelector,omitempty"`
}

func (s DescribeApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationRequest) GetAppVersions() *string {
	return s.AppVersions
}

func (s *DescribeApplicationRequest) GetLevel() *string {
	return s.Level
}

func (s *DescribeApplicationRequest) GetOutDetailStatParams() *string {
	return s.OutDetailStatParams
}

func (s *DescribeApplicationRequest) GetResourceSelector() *string {
	return s.ResourceSelector
}

func (s *DescribeApplicationRequest) SetAppId(v string) *DescribeApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationRequest) SetAppVersions(v string) *DescribeApplicationRequest {
	s.AppVersions = &v
	return s
}

func (s *DescribeApplicationRequest) SetLevel(v string) *DescribeApplicationRequest {
	s.Level = &v
	return s
}

func (s *DescribeApplicationRequest) SetOutDetailStatParams(v string) *DescribeApplicationRequest {
	s.OutDetailStatParams = &v
	return s
}

func (s *DescribeApplicationRequest) SetResourceSelector(v string) *DescribeApplicationRequest {
	s.ResourceSelector = &v
	return s
}

func (s *DescribeApplicationRequest) Validate() error {
	return dara.Validate(s)
}
