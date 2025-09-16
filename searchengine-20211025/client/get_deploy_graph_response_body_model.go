// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeployGraphResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDeployGraphResponseBody
	GetRequestId() *string
	SetResult(v *GetDeployGraphResponseBodyResult) *GetDeployGraphResponseBody
	GetResult() *GetDeployGraphResponseBodyResult
}

type GetDeployGraphResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The results returned.
	//
	// example:
	//
	// {}
	Result *GetDeployGraphResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetDeployGraphResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeployGraphResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeployGraphResponseBody) GetResult() *GetDeployGraphResponseBodyResult {
	return s.Result
}

func (s *GetDeployGraphResponseBody) SetRequestId(v string) *GetDeployGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeployGraphResponseBody) SetResult(v *GetDeployGraphResponseBodyResult) *GetDeployGraphResponseBody {
	s.Result = v
	return s
}

func (s *GetDeployGraphResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeployGraphResponseBodyResult struct {
	// The deployment information.
	Graph *GetDeployGraphResponseBodyResultGraph `json:"graph,omitempty" xml:"graph,omitempty" type:"Struct"`
}

func (s GetDeployGraphResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDeployGraphResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResult) GetGraph() *GetDeployGraphResponseBodyResultGraph {
	return s.Graph
}

func (s *GetDeployGraphResponseBodyResult) SetGraph(v *GetDeployGraphResponseBodyResultGraph) *GetDeployGraphResponseBodyResult {
	s.Graph = v
	return s
}

func (s *GetDeployGraphResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetDeployGraphResponseBodyResultGraph struct {
	// The index metadata.
	IndexMetas []*GetDeployGraphResponseBodyResultGraphIndexMetas `json:"indexMetas,omitempty" xml:"indexMetas,omitempty" type:"Repeated"`
	// The metadata of online clusters.
	OnlineMaster []*GetDeployGraphResponseBodyResultGraphOnlineMaster `json:"onlineMaster,omitempty" xml:"onlineMaster,omitempty" type:"Repeated"`
	// The association relationships between data sources and indexes.
	TableIndexRelation map[string][]*string `json:"tableIndexRelation,omitempty" xml:"tableIndexRelation,omitempty"`
	// The metadata of data sources.
	TableMetas []*GetDeployGraphResponseBodyResultGraphTableMetas `json:"tableMetas,omitempty" xml:"tableMetas,omitempty" type:"Repeated"`
	// The association relationships between zones and indexes.
	ZoneIndexRelation map[string][]*string `json:"zoneIndexRelation,omitempty" xml:"zoneIndexRelation,omitempty"`
	// The zone metadata.
	ZoneMetas []*GetDeployGraphResponseBodyResultGraphZoneMetas `json:"zoneMetas,omitempty" xml:"zoneMetas,omitempty" type:"Repeated"`
}

func (s GetDeployGraphResponseBodyResultGraph) String() string {
	return dara.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraph) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraph) GetIndexMetas() []*GetDeployGraphResponseBodyResultGraphIndexMetas {
	return s.IndexMetas
}

func (s *GetDeployGraphResponseBodyResultGraph) GetOnlineMaster() []*GetDeployGraphResponseBodyResultGraphOnlineMaster {
	return s.OnlineMaster
}

func (s *GetDeployGraphResponseBodyResultGraph) GetTableIndexRelation() map[string][]*string {
	return s.TableIndexRelation
}

func (s *GetDeployGraphResponseBodyResultGraph) GetTableMetas() []*GetDeployGraphResponseBodyResultGraphTableMetas {
	return s.TableMetas
}

func (s *GetDeployGraphResponseBodyResultGraph) GetZoneIndexRelation() map[string][]*string {
	return s.ZoneIndexRelation
}

func (s *GetDeployGraphResponseBodyResultGraph) GetZoneMetas() []*GetDeployGraphResponseBodyResultGraphZoneMetas {
	return s.ZoneMetas
}

func (s *GetDeployGraphResponseBodyResultGraph) SetIndexMetas(v []*GetDeployGraphResponseBodyResultGraphIndexMetas) *GetDeployGraphResponseBodyResultGraph {
	s.IndexMetas = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetOnlineMaster(v []*GetDeployGraphResponseBodyResultGraphOnlineMaster) *GetDeployGraphResponseBodyResultGraph {
	s.OnlineMaster = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetTableIndexRelation(v map[string][]*string) *GetDeployGraphResponseBodyResultGraph {
	s.TableIndexRelation = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetTableMetas(v []*GetDeployGraphResponseBodyResultGraphTableMetas) *GetDeployGraphResponseBodyResultGraph {
	s.TableMetas = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetZoneIndexRelation(v map[string][]*string) *GetDeployGraphResponseBodyResultGraph {
	s.ZoneIndexRelation = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) SetZoneMetas(v []*GetDeployGraphResponseBodyResultGraphZoneMetas) *GetDeployGraphResponseBodyResultGraph {
	s.ZoneMetas = v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraph) Validate() error {
	return dara.Validate(s)
}

type GetDeployGraphResponseBodyResultGraphIndexMetas struct {
	// The name of the data center.
	//
	// example:
	//
	// hz_pre_vpc_domain_1
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The index name.
	//
	// example:
	//
	// test_api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The deployment ID of the table.
	//
	// example:
	//
	// 2409
	TableDeployId *int64 `json:"tableDeployId,omitempty" xml:"tableDeployId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// ha-cn-pl32rf0****_test_api
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// The tag.
	//
	// example:
	//
	// test_api_hz_pre_vpc_domain_1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The name of the QRS worker.
	//
	// example:
	//
	// ha-cn-pl32rf0****_qrs
	ZoneName *string `json:"zoneName,omitempty" xml:"zoneName,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphIndexMetas) String() string {
	return dara.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphIndexMetas) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) GetDomainName() *string {
	return s.DomainName
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) GetName() *string {
	return s.Name
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) GetTableDeployId() *int64 {
	return s.TableDeployId
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) GetTableName() *string {
	return s.TableName
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) GetTag() *string {
	return s.Tag
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) GetZoneName() *string {
	return s.ZoneName
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetDomainName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.DomainName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.Name = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetTableDeployId(v int64) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.TableDeployId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetTableName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.TableName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetTag(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.Tag = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) SetZoneName(v string) *GetDeployGraphResponseBodyResultGraphIndexMetas {
	s.ZoneName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphIndexMetas) Validate() error {
	return dara.Validate(s)
}

type GetDeployGraphResponseBodyResultGraphOnlineMaster struct {
	// The name of the data center.
	//
	// example:
	//
	// hz_pre_vpc_domain_1
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// 5377
	HippoId *string `json:"hippoId,omitempty" xml:"hippoId,omitempty"`
	// The ID of the data center.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the online cluster.
	//
	// example:
	//
	// ha-cn-pl32rf0****_hz_pre_vpc_domain_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphOnlineMaster) String() string {
	return dara.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphOnlineMaster) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) GetDomainName() *string {
	return s.DomainName
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) GetHippoId() *string {
	return s.HippoId
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) GetId() *int64 {
	return s.Id
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) GetName() *string {
	return s.Name
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetDomainName(v string) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.DomainName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetHippoId(v string) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.HippoId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetId(v int64) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.Id = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) SetName(v string) *GetDeployGraphResponseBodyResultGraphOnlineMaster {
	s.Name = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphOnlineMaster) Validate() error {
	return dara.Validate(s)
}

type GetDeployGraphResponseBodyResultGraphTableMetas struct {
	// The ID of the offline deployment.
	//
	// example:
	//
	// 1
	BuildDeployId *int64 `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	// The name of the data center.
	//
	// example:
	//
	// hz_pre_vpc_domain_1
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// ha-cn-pl32rf0****_test_api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The deployment ID of the table.
	//
	// example:
	//
	// 2177
	TableDeployId *int64 `json:"tableDeployId,omitempty" xml:"tableDeployId,omitempty"`
	// The tag.
	//
	// example:
	//
	// ha-cn-pl32rf0****_test_api_hz_pre_vpc_domain_1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphTableMetas) String() string {
	return dara.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphTableMetas) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) GetBuildDeployId() *int64 {
	return s.BuildDeployId
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) GetDomainName() *string {
	return s.DomainName
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) GetName() *string {
	return s.Name
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) GetTableDeployId() *int64 {
	return s.TableDeployId
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) GetTag() *string {
	return s.Tag
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) GetType() *string {
	return s.Type
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetBuildDeployId(v int64) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.BuildDeployId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetDomainName(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.DomainName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetName(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.Name = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetTableDeployId(v int64) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.TableDeployId = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetTag(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.Tag = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) SetType(v string) *GetDeployGraphResponseBodyResultGraphTableMetas {
	s.Type = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphTableMetas) Validate() error {
	return dara.Validate(s)
}

type GetDeployGraphResponseBodyResultGraphZoneMetas struct {
	// The name of the data center.
	//
	// example:
	//
	// hz_pre_vpc_domain_1
	DomainInfo *string `json:"domainInfo,omitempty" xml:"domainInfo,omitempty"`
	// The name of the Query Result Searcher (QRS) worker.
	//
	// example:
	//
	// ha-cn-pl32rf0****_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the service that is used to manage the relationships between online clusters and indexes.
	//
	// example:
	//
	// ha-cn-pl32rf0****_hz_pre_vpc_domain_1
	SuezAdminName *string `json:"suezAdminName,omitempty" xml:"suezAdminName,omitempty"`
	// The tag.
	//
	// example:
	//
	// ha-cn-pl32rf0****_qrs_hz_pre_vpc_domain_1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The node type.
	//
	// example:
	//
	// qrs
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDeployGraphResponseBodyResultGraphZoneMetas) String() string {
	return dara.Prettify(s)
}

func (s GetDeployGraphResponseBodyResultGraphZoneMetas) GoString() string {
	return s.String()
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) GetDomainInfo() *string {
	return s.DomainInfo
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) GetName() *string {
	return s.Name
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) GetSuezAdminName() *string {
	return s.SuezAdminName
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) GetTag() *string {
	return s.Tag
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) GetType() *string {
	return s.Type
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetDomainInfo(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.DomainInfo = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetName(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.Name = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetSuezAdminName(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.SuezAdminName = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetTag(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.Tag = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) SetType(v string) *GetDeployGraphResponseBodyResultGraphZoneMetas {
	s.Type = &v
	return s
}

func (s *GetDeployGraphResponseBodyResultGraphZoneMetas) Validate() error {
	return dara.Validate(s)
}
