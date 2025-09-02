// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgSceneAddOrUpdateSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScenes(v []*DsgSceneAddOrUpdateSceneRequestScenes) *DsgSceneAddOrUpdateSceneRequest
	GetScenes() []*DsgSceneAddOrUpdateSceneRequestScenes
}

type DsgSceneAddOrUpdateSceneRequest struct {
	// The information about the level-2 data masking scenario.
	//
	// This parameter is required.
	Scenes []*DsgSceneAddOrUpdateSceneRequestScenes `json:"scenes,omitempty" xml:"scenes,omitempty" type:"Repeated"`
}

func (s DsgSceneAddOrUpdateSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgSceneAddOrUpdateSceneRequest) GoString() string {
	return s.String()
}

func (s *DsgSceneAddOrUpdateSceneRequest) GetScenes() []*DsgSceneAddOrUpdateSceneRequestScenes {
	return s.Scenes
}

func (s *DsgSceneAddOrUpdateSceneRequest) SetScenes(v []*DsgSceneAddOrUpdateSceneRequestScenes) *DsgSceneAddOrUpdateSceneRequest {
	s.Scenes = v
	return s
}

func (s *DsgSceneAddOrUpdateSceneRequest) Validate() error {
	return dara.Validate(s)
}

type DsgSceneAddOrUpdateSceneRequestScenes struct {
	// The description.
	//
	// example:
	//
	// Test scenarios
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The ID of the level-2 data masking scenario.
	//
	// 	- If you do not configure this parameter, the current operation is to add a level-2 data masking scenario.
	//
	// 	- If you configure this parameter, the current operation is to modify a level-2 data masking scenario. You can call the [DsgSceneQuerySceneListByName](https://help.aliyun.com/document_detail/2786322.html) operation to query the ID of the level-2 data masking scenario.
	//
	// example:
	//
	// 123
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The information about the compute engine for which the data masking scenario takes effect.
	Projects []*DsgSceneAddOrUpdateSceneRequestScenesProjects `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	// The code of the level-1 data masking scenario to which the level-2 data masking scenario belongs. Valid values:
	//
	// 	- dataworks_display_desense_code: masking of displayed data in DataStudio and Data Map
	//
	// 	- maxcompute_desense_code: data masking at the MaxCompute compute engine layer
	//
	// 	- maxcompute_new_desense_code: data masking at the MaxCompute compute engine layer (new)
	//
	// 	- dataworks_analysis_desense_code: masking of displayed data in DataAnalysis
	//
	// This parameter is required.
	//
	// example:
	//
	// dataworks_display_desense_code
	SceneCode *string `json:"sceneCode,omitempty" xml:"sceneCode,omitempty"`
	// The name of the level-2 data masking scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// dev_scene
	SceneName *string `json:"sceneName,omitempty" xml:"sceneName,omitempty"`
	// The information about the user group for which the data masking scenario takes effect.
	UserGroupIds []*int64 `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
}

func (s DsgSceneAddOrUpdateSceneRequestScenes) String() string {
	return dara.Prettify(s)
}

func (s DsgSceneAddOrUpdateSceneRequestScenes) GoString() string {
	return s.String()
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) GetDesc() *string {
	return s.Desc
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) GetId() *string {
	return s.Id
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) GetProjects() []*DsgSceneAddOrUpdateSceneRequestScenesProjects {
	return s.Projects
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) GetSceneName() *string {
	return s.SceneName
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) GetUserGroupIds() []*int64 {
	return s.UserGroupIds
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) SetDesc(v string) *DsgSceneAddOrUpdateSceneRequestScenes {
	s.Desc = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) SetId(v string) *DsgSceneAddOrUpdateSceneRequestScenes {
	s.Id = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) SetProjects(v []*DsgSceneAddOrUpdateSceneRequestScenesProjects) *DsgSceneAddOrUpdateSceneRequestScenes {
	s.Projects = v
	return s
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) SetSceneCode(v string) *DsgSceneAddOrUpdateSceneRequestScenes {
	s.SceneCode = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) SetSceneName(v string) *DsgSceneAddOrUpdateSceneRequestScenes {
	s.SceneName = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) SetUserGroupIds(v []*int64) *DsgSceneAddOrUpdateSceneRequestScenes {
	s.UserGroupIds = v
	return s
}

func (s *DsgSceneAddOrUpdateSceneRequestScenes) Validate() error {
	return dara.Validate(s)
}

type DsgSceneAddOrUpdateSceneRequestScenesProjects struct {
	// If the data masking scenario takes effect for an E-MapReduce (EMR) compute engine, enter the ID of an EMR cluster. This parameter is required only when you use an EMR compute engine.
	//
	// example:
	//
	// c-1234
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The type of the compute engine for which the data masking scenario takes effect. Valid values:
	//
	// 	- ODPS: ODPS.ODPS
	//
	// 	- HOLO: HOLO.POSTGRES
	//
	// 	- EMR: EMR
	//
	// example:
	//
	// ODPS.ODPS
	DbType *string `json:"dbType,omitempty" xml:"dbType,omitempty"`
	// The name of the compute engine instance for which the data masking scenario takes effect.
	//
	// example:
	//
	// dev_project
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s DsgSceneAddOrUpdateSceneRequestScenesProjects) String() string {
	return dara.Prettify(s)
}

func (s DsgSceneAddOrUpdateSceneRequestScenesProjects) GoString() string {
	return s.String()
}

func (s *DsgSceneAddOrUpdateSceneRequestScenesProjects) GetClusterId() *string {
	return s.ClusterId
}

func (s *DsgSceneAddOrUpdateSceneRequestScenesProjects) GetDbType() *string {
	return s.DbType
}

func (s *DsgSceneAddOrUpdateSceneRequestScenesProjects) GetProjectName() *string {
	return s.ProjectName
}

func (s *DsgSceneAddOrUpdateSceneRequestScenesProjects) SetClusterId(v string) *DsgSceneAddOrUpdateSceneRequestScenesProjects {
	s.ClusterId = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneRequestScenesProjects) SetDbType(v string) *DsgSceneAddOrUpdateSceneRequestScenesProjects {
	s.DbType = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneRequestScenesProjects) SetProjectName(v string) *DsgSceneAddOrUpdateSceneRequestScenesProjects {
	s.ProjectName = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneRequestScenesProjects) Validate() error {
	return dara.Validate(s)
}
