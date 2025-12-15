// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupSupportParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeParameterGroupSupportParamResponseBody
	GetRequestId() *string
	SetResourceList(v []*DescribeParameterGroupSupportParamResponseBodyResourceList) *DescribeParameterGroupSupportParamResponseBody
	GetResourceList() []*DescribeParameterGroupSupportParamResponseBodyResourceList
}

type DescribeParameterGroupSupportParamResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BB73740C-23E2-4392-9DA4-2660C74C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The parameters.
	ResourceList []*DescribeParameterGroupSupportParamResponseBodyResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
}

func (s DescribeParameterGroupSupportParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupSupportParamResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupSupportParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterGroupSupportParamResponseBody) GetResourceList() []*DescribeParameterGroupSupportParamResponseBodyResourceList {
	return s.ResourceList
}

func (s *DescribeParameterGroupSupportParamResponseBody) SetRequestId(v string) *DescribeParameterGroupSupportParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterGroupSupportParamResponseBody) SetResourceList(v []*DescribeParameterGroupSupportParamResponseBodyResourceList) *DescribeParameterGroupSupportParamResponseBody {
	s.ResourceList = v
	return s
}

func (s *DescribeParameterGroupSupportParamResponseBody) Validate() error {
	if s.ResourceList != nil {
		for _, item := range s.ResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeParameterGroupSupportParamResponseBodyResourceList struct {
	// The service category.
	//
	// example:
	//
	// standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The engine type.
	//
	// example:
	//
	// redis
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 5
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// rt_threshold_ms
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
}

func (s DescribeParameterGroupSupportParamResponseBodyResourceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupSupportParamResponseBodyResourceList) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupSupportParamResponseBodyResourceList) GetCategory() *string {
	return s.Category
}

func (s *DescribeParameterGroupSupportParamResponseBodyResourceList) GetDbType() *string {
	return s.DbType
}

func (s *DescribeParameterGroupSupportParamResponseBodyResourceList) GetDbVersion() *string {
	return s.DbVersion
}

func (s *DescribeParameterGroupSupportParamResponseBodyResourceList) GetParamName() *string {
	return s.ParamName
}

func (s *DescribeParameterGroupSupportParamResponseBodyResourceList) SetCategory(v string) *DescribeParameterGroupSupportParamResponseBodyResourceList {
	s.Category = &v
	return s
}

func (s *DescribeParameterGroupSupportParamResponseBodyResourceList) SetDbType(v string) *DescribeParameterGroupSupportParamResponseBodyResourceList {
	s.DbType = &v
	return s
}

func (s *DescribeParameterGroupSupportParamResponseBodyResourceList) SetDbVersion(v string) *DescribeParameterGroupSupportParamResponseBodyResourceList {
	s.DbVersion = &v
	return s
}

func (s *DescribeParameterGroupSupportParamResponseBodyResourceList) SetParamName(v string) *DescribeParameterGroupSupportParamResponseBodyResourceList {
	s.ParamName = &v
	return s
}

func (s *DescribeParameterGroupSupportParamResponseBodyResourceList) Validate() error {
	return dara.Validate(s)
}
