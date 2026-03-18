// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOralEvaluationStatisticsErrorCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetProjectData(v *OralEvaluationStatisticsErrorCountResponseProjectData) *OralEvaluationStatisticsErrorCountResponse
	GetProjectData() *OralEvaluationStatisticsErrorCountResponseProjectData
	SetProjectId(v string) *OralEvaluationStatisticsErrorCountResponse
	GetProjectId() *string
}

type OralEvaluationStatisticsErrorCountResponse struct {
	ProjectData *OralEvaluationStatisticsErrorCountResponseProjectData `json:"ProjectData,omitempty" xml:"ProjectData,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountResponse) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountResponse) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountResponse) GetProjectData() *OralEvaluationStatisticsErrorCountResponseProjectData {
	return s.ProjectData
}

func (s *OralEvaluationStatisticsErrorCountResponse) GetProjectId() *string {
	return s.ProjectId
}

func (s *OralEvaluationStatisticsErrorCountResponse) SetProjectData(v *OralEvaluationStatisticsErrorCountResponseProjectData) *OralEvaluationStatisticsErrorCountResponse {
	s.ProjectData = v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponse) SetProjectId(v string) *OralEvaluationStatisticsErrorCountResponse {
	s.ProjectId = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponse) Validate() error {
	if s.ProjectData != nil {
		if err := s.ProjectData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OralEvaluationStatisticsErrorCountResponseProjectData struct {
	ApplicationData []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData `json:"ApplicationData,omitempty" xml:"ApplicationData,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ApplicationInternalId *string `json:"applicationInternalId,omitempty" xml:"applicationInternalId,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountResponseProjectData) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountResponseProjectData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectData) GetApplicationData() []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData {
	return s.ApplicationData
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectData) GetApplicationInternalId() *string {
	return s.ApplicationInternalId
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectData) SetApplicationData(v []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) *OralEvaluationStatisticsErrorCountResponseProjectData {
	s.ApplicationData = v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectData) SetApplicationInternalId(v string) *OralEvaluationStatisticsErrorCountResponseProjectData {
	s.ApplicationInternalId = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectData) Validate() error {
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

type OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData struct {
	Data []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// appId,appkey
	//
	// This parameter is required.
	//
	// example:
	//
	// t000797
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) GetData() []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData {
	return s.Data
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) SetData(v []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData {
	s.Data = v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) SetApplicationAccessId(v string) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData {
	s.ApplicationAccessId = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) Validate() error {
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

type OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData struct {
	Data []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 51000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// start the core unsuccessfull.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) GetData() []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData {
	return s.Data
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) SetData(v []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData {
	s.Data = v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) SetErrorCode(v string) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData {
	s.ErrorCode = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) SetErrorMessage(v string) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData {
	s.ErrorMessage = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) Validate() error {
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

type OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData struct {
	// This parameter is required.
	//
	// example:
	//
	// 230
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-02
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) GetCount() *int32 {
	return s.Count
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) GetName() *string {
	return s.Name
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) SetCount(v int32) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData {
	s.Count = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) SetName(v string) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData {
	s.Name = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) Validate() error {
	return dara.Validate(s)
}
