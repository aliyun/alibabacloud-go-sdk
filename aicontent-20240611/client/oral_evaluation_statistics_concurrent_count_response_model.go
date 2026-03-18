// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOralEvaluationStatisticsConcurrentCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetProjectData(v *OralEvaluationStatisticsConcurrentCountResponseProjectData) *OralEvaluationStatisticsConcurrentCountResponse
	GetProjectData() *OralEvaluationStatisticsConcurrentCountResponseProjectData
	SetProjectId(v string) *OralEvaluationStatisticsConcurrentCountResponse
	GetProjectId() *string
}

type OralEvaluationStatisticsConcurrentCountResponse struct {
	// This parameter is required.
	ProjectData *OralEvaluationStatisticsConcurrentCountResponseProjectData `json:"projectData,omitempty" xml:"projectData,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s OralEvaluationStatisticsConcurrentCountResponse) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsConcurrentCountResponse) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsConcurrentCountResponse) GetProjectData() *OralEvaluationStatisticsConcurrentCountResponseProjectData {
	return s.ProjectData
}

func (s *OralEvaluationStatisticsConcurrentCountResponse) GetProjectId() *string {
	return s.ProjectId
}

func (s *OralEvaluationStatisticsConcurrentCountResponse) SetProjectData(v *OralEvaluationStatisticsConcurrentCountResponseProjectData) *OralEvaluationStatisticsConcurrentCountResponse {
	s.ProjectData = v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponse) SetProjectId(v string) *OralEvaluationStatisticsConcurrentCountResponse {
	s.ProjectId = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponse) Validate() error {
	if s.ProjectData != nil {
		if err := s.ProjectData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OralEvaluationStatisticsConcurrentCountResponseProjectData struct {
	ApplicationData []*OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData `json:"ApplicationData,omitempty" xml:"ApplicationData,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	ApplicationInternalId *string `json:"applicationInternalId,omitempty" xml:"applicationInternalId,omitempty"`
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectData) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectData) GetApplicationData() []*OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData {
	return s.ApplicationData
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectData) GetApplicationInternalId() *string {
	return s.ApplicationInternalId
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectData) SetApplicationData(v []*OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) *OralEvaluationStatisticsConcurrentCountResponseProjectData {
	s.ApplicationData = v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectData) SetApplicationInternalId(v string) *OralEvaluationStatisticsConcurrentCountResponseProjectData {
	s.ApplicationInternalId = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectData) Validate() error {
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

type OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData struct {
	Data []*OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// appkey
	//
	// This parameter is required.
	//
	// example:
	//
	// a0007g7
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) GetData() []*OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData {
	return s.Data
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) SetData(v []*OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData {
	s.Data = v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) SetApplicationAccessId(v string) *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData {
	s.ApplicationAccessId = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) Validate() error {
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

type OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData struct {
	// This parameter is required.
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) GetCount() *int32 {
	return s.Count
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) GetName() *string {
	return s.Name
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) SetCount(v int32) *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData {
	s.Count = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) SetName(v string) *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData {
	s.Name = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) Validate() error {
	return dara.Validate(s)
}
