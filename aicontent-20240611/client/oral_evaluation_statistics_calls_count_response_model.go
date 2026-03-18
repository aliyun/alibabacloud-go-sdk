// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOralEvaluationStatisticsCallsCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetProjectData(v *OralEvaluationStatisticsCallsCountResponseProjectData) *OralEvaluationStatisticsCallsCountResponse
	GetProjectData() *OralEvaluationStatisticsCallsCountResponseProjectData
	SetProjectId(v string) *OralEvaluationStatisticsCallsCountResponse
	GetProjectId() *string
}

type OralEvaluationStatisticsCallsCountResponse struct {
	// This parameter is required.
	ProjectData *OralEvaluationStatisticsCallsCountResponseProjectData `json:"projectData,omitempty" xml:"projectData,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s OralEvaluationStatisticsCallsCountResponse) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsCallsCountResponse) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsCallsCountResponse) GetProjectData() *OralEvaluationStatisticsCallsCountResponseProjectData {
	return s.ProjectData
}

func (s *OralEvaluationStatisticsCallsCountResponse) GetProjectId() *string {
	return s.ProjectId
}

func (s *OralEvaluationStatisticsCallsCountResponse) SetProjectData(v *OralEvaluationStatisticsCallsCountResponseProjectData) *OralEvaluationStatisticsCallsCountResponse {
	s.ProjectData = v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponse) SetProjectId(v string) *OralEvaluationStatisticsCallsCountResponse {
	s.ProjectId = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponse) Validate() error {
	if s.ProjectData != nil {
		if err := s.ProjectData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OralEvaluationStatisticsCallsCountResponseProjectData struct {
	ApplicationData []*OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData `json:"ApplicationData,omitempty" xml:"ApplicationData,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	ApplicationInternalId *string `json:"applicationInternalId,omitempty" xml:"applicationInternalId,omitempty"`
}

func (s OralEvaluationStatisticsCallsCountResponseProjectData) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsCallsCountResponseProjectData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectData) GetApplicationData() []*OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData {
	return s.ApplicationData
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectData) GetApplicationInternalId() *string {
	return s.ApplicationInternalId
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectData) SetApplicationData(v []*OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) *OralEvaluationStatisticsCallsCountResponseProjectData {
	s.ApplicationData = v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectData) SetApplicationInternalId(v string) *OralEvaluationStatisticsCallsCountResponseProjectData {
	s.ApplicationInternalId = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectData) Validate() error {
	if s.ApplicationData != nil {
		for _, item := range s.ApplicationData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData struct {
	Data []*OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// appkey
	//
	// This parameter is required.
	//
	// example:
	//
	// a0007g7
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
}

func (s OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) GetData() []*OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData {
	return s.Data
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) SetData(v []*OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData {
	s.Data = v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) SetApplicationAccessId(v string) *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData {
	s.ApplicationAccessId = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) Validate() error {
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

type OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) GetCount() *int32 {
	return s.Count
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) GetName() *string {
	return s.Name
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) SetCount(v int32) *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData {
	s.Count = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) SetName(v string) *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData {
	s.Name = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) Validate() error {
	return dara.Validate(s)
}
