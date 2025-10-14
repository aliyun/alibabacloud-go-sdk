// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeParameterGroupsResponseBodyData) *DescribeParameterGroupsResponseBody
	GetData() *DescribeParameterGroupsResponseBodyData
	SetRequestId(v string) *DescribeParameterGroupsResponseBody
	GetRequestId() *string
}

type DescribeParameterGroupsResponseBody struct {
	Data *DescribeParameterGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponseBody) GetData() *DescribeParameterGroupsResponseBodyData {
	return s.Data
}

func (s *DescribeParameterGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterGroupsResponseBody) SetData(v *DescribeParameterGroupsResponseBodyData) *DescribeParameterGroupsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeParameterGroupsResponseBody) SetRequestId(v string) *DescribeParameterGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterGroupsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeParameterGroupsResponseBodyData struct {
	ParameterGroups []*DescribeParameterGroupsResponseBodyDataParameterGroups `json:"ParameterGroups,omitempty" xml:"ParameterGroups,omitempty" type:"Repeated"`
	// example:
	//
	// 1E5DCFFC-A00D-****-836E-73318F8CA577
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParameterGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponseBodyData) GetParameterGroups() []*DescribeParameterGroupsResponseBodyDataParameterGroups {
	return s.ParameterGroups
}

func (s *DescribeParameterGroupsResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterGroupsResponseBodyData) SetParameterGroups(v []*DescribeParameterGroupsResponseBodyDataParameterGroups) *DescribeParameterGroupsResponseBodyData {
	s.ParameterGroups = v
	return s
}

func (s *DescribeParameterGroupsResponseBodyData) SetRequestId(v string) *DescribeParameterGroupsResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyData) Validate() error {
	if s.ParameterGroups != nil {
		for _, item := range s.ParameterGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeParameterGroupsResponseBodyDataParameterGroups struct {
	CnForceRestart *bool `json:"CnForceRestart,omitempty" xml:"CnForceRestart,omitempty"`
	// example:
	//
	// 10
	CnParamCount *int32 `json:"CnParamCount,omitempty" xml:"CnParamCount,omitempty"`
	// example:
	//
	// polarx
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// 5.7
	DbVersion      *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	DnForceRestart *bool   `json:"DnForceRestart,omitempty" xml:"DnForceRestart,omitempty"`
	// example:
	//
	// 10
	DnParamCount *int32 `json:"DnParamCount,omitempty" xml:"DnParamCount,omitempty"`
	// example:
	//
	// 2024-12-19T16:41:31+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 1605079985000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// rpg-s1y1xy06****fqs7y
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// example:
	//
	// dstest_api_new
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// example:
	//
	// 0
	ParameterGroupType *string `json:"ParameterGroupType,omitempty" xml:"ParameterGroupType,omitempty"`
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
}

func (s DescribeParameterGroupsResponseBodyDataParameterGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsResponseBodyDataParameterGroups) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetCnForceRestart() *bool {
	return s.CnForceRestart
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetCnParamCount() *int32 {
	return s.CnParamCount
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetDbType() *string {
	return s.DbType
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetDbVersion() *string {
	return s.DbVersion
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetDnForceRestart() *bool {
	return s.DnForceRestart
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetDnParamCount() *int32 {
	return s.DnParamCount
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetParameterGroupType() *string {
	return s.ParameterGroupType
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) GetSeries() *string {
	return s.Series
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetCnForceRestart(v bool) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.CnForceRestart = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetCnParamCount(v int32) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.CnParamCount = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetDbType(v string) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.DbType = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetDbVersion(v string) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.DbVersion = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetDnForceRestart(v bool) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.DnForceRestart = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetDnParamCount(v int32) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.DnParamCount = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetGmtCreated(v string) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.GmtCreated = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetGmtModified(v string) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.GmtModified = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetParameterGroupId(v string) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetParameterGroupName(v string) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetParameterGroupType(v string) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.ParameterGroupType = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) SetSeries(v string) *DescribeParameterGroupsResponseBodyDataParameterGroups {
	s.Series = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyDataParameterGroups) Validate() error {
	return dara.Validate(s)
}
