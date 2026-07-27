// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgSceneQuerySceneListByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DsgSceneQuerySceneListByNameResponseBodyData) *DsgSceneQuerySceneListByNameResponseBody
	GetData() []*DsgSceneQuerySceneListByNameResponseBodyData
	SetErrorCode(v string) *DsgSceneQuerySceneListByNameResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgSceneQuerySceneListByNameResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgSceneQuerySceneListByNameResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgSceneQuerySceneListByNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgSceneQuerySceneListByNameResponseBody
	GetSuccess() *bool
}

type DsgSceneQuerySceneListByNameResponseBody struct {
	// The list of data masking scenarios.
	Data []*DsgSceneQuerySceneListByNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 1029030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 102400001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgSceneQuerySceneListByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgSceneQuerySceneListByNameResponseBody) GoString() string {
	return s.String()
}

func (s *DsgSceneQuerySceneListByNameResponseBody) GetData() []*DsgSceneQuerySceneListByNameResponseBodyData {
	return s.Data
}

func (s *DsgSceneQuerySceneListByNameResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgSceneQuerySceneListByNameResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgSceneQuerySceneListByNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgSceneQuerySceneListByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgSceneQuerySceneListByNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgSceneQuerySceneListByNameResponseBody) SetData(v []*DsgSceneQuerySceneListByNameResponseBodyData) *DsgSceneQuerySceneListByNameResponseBody {
	s.Data = v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBody) SetErrorCode(v string) *DsgSceneQuerySceneListByNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBody) SetErrorMessage(v string) *DsgSceneQuerySceneListByNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBody) SetHttpStatusCode(v int32) *DsgSceneQuerySceneListByNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBody) SetRequestId(v string) *DsgSceneQuerySceneListByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBody) SetSuccess(v bool) *DsgSceneQuerySceneListByNameResponseBody {
	s.Success = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBody) Validate() error {
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

type DsgSceneQuerySceneListByNameResponseBodyData struct {
	// The nested data masking scenarios.
	Children []interface{} `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The description of the data masking scenario.
	//
	// example:
	//
	// Test scenarios
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The ID of the data masking scenario.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The engine instances to which the data masking scenario applies.
	Projects []*DsgSceneQuerySceneListByNameResponseBodyDataProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// The code for the level-1 scenario. Valid values:
	//
	// - Data masking in Data Map and DataStudio: `dataworks_display_desense_code`
	//
	// - Data masking at the MaxCompute engine layer: `maxcompute_desense_code`
	//
	// - Data masking at the MaxCompute engine layer (new): `maxcompute_new_desense_code`
	//
	// - Data masking at the Hologres engine layer: `hologres_display_desense_code`
	//
	// - Static data masking in Data Integration: `dataworks_data_integration_desense_code`
	//
	// - Data masking in Data Analysis: `dataworks_analysis_desense_code`
	//
	// example:
	//
	// dataworks_display_desense_code
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The level of the data masking scenario. Valid values:
	//
	// - `0`: level-1 scenario
	//
	// - `1`: level-2 scenario
	//
	// example:
	//
	// 1
	SceneLevel *int32 `json:"SceneLevel,omitempty" xml:"SceneLevel,omitempty"`
	// The name of the data masking scenario.
	//
	// example:
	//
	// test_scene
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The user groups to which the data masking scenario applies. Multiple user group names are separated by a comma (,).
	//
	// example:
	//
	// user1,user2
	UserGroups   *string `json:"UserGroups,omitempty" xml:"UserGroups,omitempty"`
	ScenceDbType *string `json:"scenceDbType,omitempty" xml:"scenceDbType,omitempty"`
}

func (s DsgSceneQuerySceneListByNameResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DsgSceneQuerySceneListByNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) GetChildren() []interface{} {
	return s.Children
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) GetDesc() *string {
	return s.Desc
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) GetProjects() []*DsgSceneQuerySceneListByNameResponseBodyDataProjects {
	return s.Projects
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) GetSceneLevel() *int32 {
	return s.SceneLevel
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) GetSceneName() *string {
	return s.SceneName
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) GetUserGroups() *string {
	return s.UserGroups
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) GetScenceDbType() *string {
	return s.ScenceDbType
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) SetChildren(v []interface{}) *DsgSceneQuerySceneListByNameResponseBodyData {
	s.Children = v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) SetDesc(v string) *DsgSceneQuerySceneListByNameResponseBodyData {
	s.Desc = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) SetId(v int64) *DsgSceneQuerySceneListByNameResponseBodyData {
	s.Id = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) SetProjects(v []*DsgSceneQuerySceneListByNameResponseBodyDataProjects) *DsgSceneQuerySceneListByNameResponseBodyData {
	s.Projects = v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) SetSceneCode(v string) *DsgSceneQuerySceneListByNameResponseBodyData {
	s.SceneCode = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) SetSceneLevel(v int32) *DsgSceneQuerySceneListByNameResponseBodyData {
	s.SceneLevel = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) SetSceneName(v string) *DsgSceneQuerySceneListByNameResponseBodyData {
	s.SceneName = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) SetUserGroups(v string) *DsgSceneQuerySceneListByNameResponseBodyData {
	s.UserGroups = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) SetScenceDbType(v string) *DsgSceneQuerySceneListByNameResponseBodyData {
	s.ScenceDbType = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyData) Validate() error {
	if s.Projects != nil {
		for _, item := range s.Projects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DsgSceneQuerySceneListByNameResponseBodyDataProjects struct {
	// The ID of the E-MapReduce (EMR) cluster. This parameter is returned only if the `DbType` is `EMR`.
	//
	// example:
	//
	// c-123456
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The engine type. Valid values:
	//
	// - MaxCompute: `ODPS.ODPS`
	//
	// - Hologres: `HOLO.POSTGRES`
	//
	// - E-MapReduce (EMR): `EMR`
	//
	// example:
	//
	// ODPS.ODPS
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The name of the engine instance.
	//
	// example:
	//
	// dev_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DsgSceneQuerySceneListByNameResponseBodyDataProjects) String() string {
	return dara.Prettify(s)
}

func (s DsgSceneQuerySceneListByNameResponseBodyDataProjects) GoString() string {
	return s.String()
}

func (s *DsgSceneQuerySceneListByNameResponseBodyDataProjects) GetClusterId() *string {
	return s.ClusterId
}

func (s *DsgSceneQuerySceneListByNameResponseBodyDataProjects) GetDbType() *string {
	return s.DbType
}

func (s *DsgSceneQuerySceneListByNameResponseBodyDataProjects) GetProjectName() *string {
	return s.ProjectName
}

func (s *DsgSceneQuerySceneListByNameResponseBodyDataProjects) SetClusterId(v string) *DsgSceneQuerySceneListByNameResponseBodyDataProjects {
	s.ClusterId = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyDataProjects) SetDbType(v string) *DsgSceneQuerySceneListByNameResponseBodyDataProjects {
	s.DbType = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyDataProjects) SetProjectName(v string) *DsgSceneQuerySceneListByNameResponseBodyDataProjects {
	s.ProjectName = &v
	return s
}

func (s *DsgSceneQuerySceneListByNameResponseBodyDataProjects) Validate() error {
	return dara.Validate(s)
}
