// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsSqlAuditStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDrdsSqlAuditStatusResponseBodyData) *DescribeDrdsSqlAuditStatusResponseBody
	GetData() *DescribeDrdsSqlAuditStatusResponseBodyData
	SetRequestId(v string) *DescribeDrdsSqlAuditStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsSqlAuditStatusResponseBody
	GetSuccess() *bool
}

type DescribeDrdsSqlAuditStatusResponseBody struct {
	// The returned data set.
	Data *DescribeDrdsSqlAuditStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// DC3ABA3E-0F8A-4596-9104-F5155C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) GetData() *DescribeDrdsSqlAuditStatusResponseBodyData {
	return s.Data
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetData(v *DescribeDrdsSqlAuditStatusResponseBodyData) *DescribeDrdsSqlAuditStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetRequestId(v string) *DescribeDrdsSqlAuditStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) SetSuccess(v bool) *DescribeDrdsSqlAuditStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsSqlAuditStatusResponseBodyData struct {
	Data []*DescribeDrdsSqlAuditStatusResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeDrdsSqlAuditStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyData) GetData() []*DescribeDrdsSqlAuditStatusResponseBodyDataData {
	return s.Data
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyData) SetData(v []*DescribeDrdsSqlAuditStatusResponseBodyDataData) *DescribeDrdsSqlAuditStatusResponseBodyData {
	s.Data = v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsSqlAuditStatusResponseBodyDataData struct {
	// The name of the database.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates whether the complete report of the SQL audit is supported. Valid values: true and false.
	//
	// example:
	//
	// true
	Detailed *string `json:"Detailed,omitempty" xml:"Detailed,omitempty"`
	// Indicates whether the SQL audit feature is enabled for the database. Valid values: true and false.
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The UID of the external delivery.
	//
	// > This parameter is returned only if external log delivery is enabled.
	//
	// example:
	//
	// 111
	ExtraAliUid *int64 `json:"ExtraAliUid,omitempty" xml:"ExtraAliUid,omitempty"`
	// The Log Service Logstore from which logs are delivered.
	//
	// > This parameter is returned only if external log delivery is enabled.
	//
	// example:
	//
	// test
	ExtraSlsLogStore *string `json:"ExtraSlsLogStore,omitempty" xml:"ExtraSlsLogStore,omitempty"`
	// The Log Service project from which logs are delivered.
	//
	// > This parameter is returned only if external log delivery is enabled.
	//
	// example:
	//
	// test
	ExtraSlsProject *string `json:"ExtraSlsProject,omitempty" xml:"ExtraSlsProject,omitempty"`
	// Indicates whether external log delivery is enabled. Valid values: true and false.
	//
	// example:
	//
	// false
	ExtraWriteEnabled *bool `json:"ExtraWriteEnabled,omitempty" xml:"ExtraWriteEnabled,omitempty"`
}

func (s DescribeDrdsSqlAuditStatusResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsSqlAuditStatusResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) GetDetailed() *string {
	return s.Detailed
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) GetEnabled() *string {
	return s.Enabled
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) GetExtraAliUid() *int64 {
	return s.ExtraAliUid
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) GetExtraSlsLogStore() *string {
	return s.ExtraSlsLogStore
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) GetExtraSlsProject() *string {
	return s.ExtraSlsProject
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) GetExtraWriteEnabled() *bool {
	return s.ExtraWriteEnabled
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetDbName(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetDetailed(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.Detailed = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetEnabled(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.Enabled = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraAliUid(v int64) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraAliUid = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraSlsLogStore(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraSlsLogStore = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraSlsProject(v string) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraSlsProject = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) SetExtraWriteEnabled(v bool) *DescribeDrdsSqlAuditStatusResponseBodyDataData {
	s.ExtraWriteEnabled = &v
	return s
}

func (s *DescribeDrdsSqlAuditStatusResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}
