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
	// Specifies whether the vulnerability has been handled. Valid values:
	//
	// - **y**: handled
	//
	// - **n**: not handled.
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The container search field name.
	//
	// example:
	//
	// namespace
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The container search field value.
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
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The vulnerability level. Separate multiple levels with commas (,). Valid values:
	//
	// - **high**: high
	//
	// - **medium**: medium
	//
	// - **low**: low.
	//
	// example:
	//
	// high,low
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The priority level of vulnerability fixing. Separate multiple levels with commas (,). Valid values:
	//
	// - **asap**: high
	//
	// - **later**: medium
	//
	// - **nntf**: low.
	//
	// example:
	//
	// asap,later,nntf
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// The asset information for the vulnerability query. You can set this parameter to the asset name, public IP address, or private IP address. Fuzzy match is supported.
	//
	// example:
	//
	// 10.7.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The tag for querying vulnerabilities.
	//
	// example:
	//
	// oval
	SearchTags *string `json:"SearchTags,omitempty" xml:"SearchTags,omitempty"`
	// The fix status of the vulnerability. Separate multiple statuses with commas (,). Valid values:
	//
	// - **1**: unfixed
	//
	// - **2**: fix failed.
	//
	// example:
	//
	// 1,2
	StatusList *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	// The vulnerability tag.
	//
	// example:
	//
	// oval
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The query type. Valid values:
	//
	// - **containerId**: container ID
	//
	// - **uuid**: asset ID.
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of vulnerability to query. Valid values:
	//
	// - **cve**: Linux software vulnerability
	//
	// - **sys**: Windows system vulnerability.
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The instance IDs of VPC-connected instances to query for vulnerabilities. Separate multiple instance IDs with commas (,).
	//
	// example:
	//
	// vpc-uf6ssrvbrwe37ekw****,vpc-bp1aevy8sofi8mh1q****
	VpcInstanceIds *string `json:"VpcInstanceIds,omitempty" xml:"VpcInstanceIds,omitempty"`
	// The collection of vulnerability names.
	//
	// > You can call the [DescribeGroupedVul](~~DescribeGroupedVul~~) operation to obtain this parameter.
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
