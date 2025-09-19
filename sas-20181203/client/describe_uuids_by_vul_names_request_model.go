// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUuidsByVulNamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDealed(v string) *DescribeUuidsByVulNamesRequest
	GetDealed() *string
	SetFieldName(v string) *DescribeUuidsByVulNamesRequest
	GetFieldName() *string
	SetFieldValue(v string) *DescribeUuidsByVulNamesRequest
	GetFieldValue() *string
	SetGroupId(v int64) *DescribeUuidsByVulNamesRequest
	GetGroupId() *int64
	SetLang(v string) *DescribeUuidsByVulNamesRequest
	GetLang() *string
	SetLevel(v string) *DescribeUuidsByVulNamesRequest
	GetLevel() *string
	SetNecessity(v string) *DescribeUuidsByVulNamesRequest
	GetNecessity() *string
	SetRemark(v string) *DescribeUuidsByVulNamesRequest
	GetRemark() *string
	SetSearchTags(v string) *DescribeUuidsByVulNamesRequest
	GetSearchTags() *string
	SetStatusList(v string) *DescribeUuidsByVulNamesRequest
	GetStatusList() *string
	SetTag(v string) *DescribeUuidsByVulNamesRequest
	GetTag() *string
	SetTargetType(v string) *DescribeUuidsByVulNamesRequest
	GetTargetType() *string
	SetType(v string) *DescribeUuidsByVulNamesRequest
	GetType() *string
	SetVpcInstanceIds(v string) *DescribeUuidsByVulNamesRequest
	GetVpcInstanceIds() *string
	SetVulNames(v []*string) *DescribeUuidsByVulNamesRequest
	GetVulNames() []*string
}

type DescribeUuidsByVulNamesRequest struct {
	// Specifies whether the vulnerability is fixed. Valid values:
	//
	// 	- **y**: the vulnerability is fixed.
	//
	// 	- **n**: the vulnerability is not fixed.
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The name of the search field that is used to query containers.
	//
	// example:
	//
	// namespace
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The value of the search field that is used to query containers.
	//
	// example:
	//
	// cas-adad-qeqwe
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The ID of the asset group.
	//
	// example:
	//
	// 11286014
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The severity of the vulnerability. Separate multiple severities with commas (,). Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// high,low
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The priority based on which the vulnerability is fixed. Separate multiple priorities with commas (,). Valid values:
	//
	// 	- **asap**: high
	//
	// 	- **later**: medium
	//
	// 	- **nntf**: low
	//
	// example:
	//
	// asap,later,nntf
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// The remarks for the asset affected by the vulnerability. The value can be the private IP address, public IP address, or name of the asset. Fuzzy match is supported.
	//
	// example:
	//
	// 10.7.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The tags that are used to search for the vulnerability.
	//
	// example:
	//
	// oval
	SearchTags *string `json:"SearchTags,omitempty" xml:"SearchTags,omitempty"`
	// The status of the vulnerability. Separate multiple states with commas (,). Valid values:
	//
	// 	- **1**: unfixed
	//
	// 	- **2**: fix failed
	//
	// example:
	//
	// 1,4
	StatusList *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	// The tag of the vulnerability.
	//
	// example:
	//
	// oval
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The type of the query condition. Valid values:
	//
	// 	- **containerId**: the ID of the container
	//
	// 	- **uuid**: the ID of the asset
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the vulnerability is detected. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// vpc-uf6ssrvbrwe37ekw****,vpc-bp1aevy8sofi8mh1q****
	VpcInstanceIds *string `json:"VpcInstanceIds,omitempty" xml:"VpcInstanceIds,omitempty"`
	// An array that consists of the names of vulnerabilities.
	//
	// >  You can call the [DescribeGroupedVul](~~DescribeGroupedVul~~) operation to obtain the names of vulnerabilities.
	//
	// This parameter is required.
	VulNames []*string `json:"VulNames,omitempty" xml:"VulNames,omitempty" type:"Repeated"`
}

func (s DescribeUuidsByVulNamesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUuidsByVulNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUuidsByVulNamesRequest) GetDealed() *string {
	return s.Dealed
}

func (s *DescribeUuidsByVulNamesRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeUuidsByVulNamesRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeUuidsByVulNamesRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeUuidsByVulNamesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeUuidsByVulNamesRequest) GetLevel() *string {
	return s.Level
}

func (s *DescribeUuidsByVulNamesRequest) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeUuidsByVulNamesRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeUuidsByVulNamesRequest) GetSearchTags() *string {
	return s.SearchTags
}

func (s *DescribeUuidsByVulNamesRequest) GetStatusList() *string {
	return s.StatusList
}

func (s *DescribeUuidsByVulNamesRequest) GetTag() *string {
	return s.Tag
}

func (s *DescribeUuidsByVulNamesRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeUuidsByVulNamesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeUuidsByVulNamesRequest) GetVpcInstanceIds() *string {
	return s.VpcInstanceIds
}

func (s *DescribeUuidsByVulNamesRequest) GetVulNames() []*string {
	return s.VulNames
}

func (s *DescribeUuidsByVulNamesRequest) SetDealed(v string) *DescribeUuidsByVulNamesRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetFieldName(v string) *DescribeUuidsByVulNamesRequest {
	s.FieldName = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetFieldValue(v string) *DescribeUuidsByVulNamesRequest {
	s.FieldValue = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetGroupId(v int64) *DescribeUuidsByVulNamesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetLang(v string) *DescribeUuidsByVulNamesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetLevel(v string) *DescribeUuidsByVulNamesRequest {
	s.Level = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetNecessity(v string) *DescribeUuidsByVulNamesRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetRemark(v string) *DescribeUuidsByVulNamesRequest {
	s.Remark = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetSearchTags(v string) *DescribeUuidsByVulNamesRequest {
	s.SearchTags = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetStatusList(v string) *DescribeUuidsByVulNamesRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetTag(v string) *DescribeUuidsByVulNamesRequest {
	s.Tag = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetTargetType(v string) *DescribeUuidsByVulNamesRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetType(v string) *DescribeUuidsByVulNamesRequest {
	s.Type = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetVpcInstanceIds(v string) *DescribeUuidsByVulNamesRequest {
	s.VpcInstanceIds = &v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) SetVulNames(v []*string) *DescribeUuidsByVulNamesRequest {
	s.VulNames = v
	return s
}

func (s *DescribeUuidsByVulNamesRequest) Validate() error {
	return dara.Validate(s)
}
