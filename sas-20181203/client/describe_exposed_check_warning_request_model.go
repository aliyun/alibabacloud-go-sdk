// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedCheckWarningRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeExposedCheckWarningRequest
	GetLang() *string
	SetTypeName(v string) *DescribeExposedCheckWarningRequest
	GetTypeName() *string
	SetUuids(v string) *DescribeExposedCheckWarningRequest
	GetUuids() *string
}

type DescribeExposedCheckWarningRequest struct {
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
	// The type of the baseline.
	//
	// >  You can call the [DescribeRiskType](~~DescribeRiskType~~) operation to obtain the types of baselines from the response parameter **TypeName**.
	//
	// example:
	//
	// weak_password
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// The UUID of the server. Separate multiple UUIDs with commas (,).
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// 6541631a-7d47-41fd-9fef-9518113f****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeExposedCheckWarningRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedCheckWarningRequest) GoString() string {
	return s.String()
}

func (s *DescribeExposedCheckWarningRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExposedCheckWarningRequest) GetTypeName() *string {
	return s.TypeName
}

func (s *DescribeExposedCheckWarningRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeExposedCheckWarningRequest) SetLang(v string) *DescribeExposedCheckWarningRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExposedCheckWarningRequest) SetTypeName(v string) *DescribeExposedCheckWarningRequest {
	s.TypeName = &v
	return s
}

func (s *DescribeExposedCheckWarningRequest) SetUuids(v string) *DescribeExposedCheckWarningRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeExposedCheckWarningRequest) Validate() error {
	return dara.Validate(s)
}
