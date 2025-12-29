// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySceneListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *QuerySceneListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySceneListResponseBody
	GetRequestId() *string
	SetResults(v []*QuerySceneListResponseBodyResults) *QuerySceneListResponseBody
	GetResults() []*QuerySceneListResponseBodyResults
	SetStatusCode(v int32) *QuerySceneListResponseBody
	GetStatusCode() *int32
}

type QuerySceneListResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FAFCD152-4791-5F2F-B0BE-2DC06FD4F05B
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*QuerySceneListResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s QuerySceneListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySceneListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySceneListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySceneListResponseBody) GetResults() []*QuerySceneListResponseBodyResults {
	return s.Results
}

func (s *QuerySceneListResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySceneListResponseBody) SetMessage(v string) *QuerySceneListResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySceneListResponseBody) SetRequestId(v string) *QuerySceneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySceneListResponseBody) SetResults(v []*QuerySceneListResponseBodyResults) *QuerySceneListResponseBody {
	s.Results = v
	return s
}

func (s *QuerySceneListResponseBody) SetStatusCode(v int32) *QuerySceneListResponseBody {
	s.StatusCode = &v
	return s
}

func (s *QuerySceneListResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySceneListResponseBodyResults struct {
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingmoshi/shuimian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 73
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// external
	SceneSource *string `json:"SceneSource,omitempty" xml:"SceneSource,omitempty"`
	// example:
	//
	// 1
	SceneState *int32 `json:"SceneState,omitempty" xml:"SceneState,omitempty"`
	// example:
	//
	// common
	SceneType           *string                                                 `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	TemplateInfoDTOList []*QuerySceneListResponseBodyResultsTemplateInfoDTOList `json:"TemplateInfoDTOList,omitempty" xml:"TemplateInfoDTOList,omitempty" type:"Repeated"`
	UnavailableReason   *string                                                 `json:"UnavailableReason,omitempty" xml:"UnavailableReason,omitempty"`
}

func (s QuerySceneListResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneListResponseBodyResults) GoString() string {
	return s.String()
}

func (s *QuerySceneListResponseBodyResults) GetIcon() *string {
	return s.Icon
}

func (s *QuerySceneListResponseBodyResults) GetSceneId() *string {
	return s.SceneId
}

func (s *QuerySceneListResponseBodyResults) GetSceneName() *string {
	return s.SceneName
}

func (s *QuerySceneListResponseBodyResults) GetSceneSource() *string {
	return s.SceneSource
}

func (s *QuerySceneListResponseBodyResults) GetSceneState() *int32 {
	return s.SceneState
}

func (s *QuerySceneListResponseBodyResults) GetSceneType() *string {
	return s.SceneType
}

func (s *QuerySceneListResponseBodyResults) GetTemplateInfoDTOList() []*QuerySceneListResponseBodyResultsTemplateInfoDTOList {
	return s.TemplateInfoDTOList
}

func (s *QuerySceneListResponseBodyResults) GetUnavailableReason() *string {
	return s.UnavailableReason
}

func (s *QuerySceneListResponseBodyResults) SetIcon(v string) *QuerySceneListResponseBodyResults {
	s.Icon = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetSceneId(v string) *QuerySceneListResponseBodyResults {
	s.SceneId = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetSceneName(v string) *QuerySceneListResponseBodyResults {
	s.SceneName = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetSceneSource(v string) *QuerySceneListResponseBodyResults {
	s.SceneSource = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetSceneState(v int32) *QuerySceneListResponseBodyResults {
	s.SceneState = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetSceneType(v string) *QuerySceneListResponseBodyResults {
	s.SceneType = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetTemplateInfoDTOList(v []*QuerySceneListResponseBodyResultsTemplateInfoDTOList) *QuerySceneListResponseBodyResults {
	s.TemplateInfoDTOList = v
	return s
}

func (s *QuerySceneListResponseBodyResults) SetUnavailableReason(v string) *QuerySceneListResponseBodyResults {
	s.UnavailableReason = &v
	return s
}

func (s *QuerySceneListResponseBodyResults) Validate() error {
	if s.TemplateInfoDTOList != nil {
		for _, item := range s.TemplateInfoDTOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySceneListResponseBodyResultsTemplateInfoDTOList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 6962
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 101
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QuerySceneListResponseBodyResultsTemplateInfoDTOList) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneListResponseBodyResultsTemplateInfoDTOList) GoString() string {
	return s.String()
}

func (s *QuerySceneListResponseBodyResultsTemplateInfoDTOList) GetDescription() *string {
	return s.Description
}

func (s *QuerySceneListResponseBodyResultsTemplateInfoDTOList) GetId() *int64 {
	return s.Id
}

func (s *QuerySceneListResponseBodyResultsTemplateInfoDTOList) GetName() *string {
	return s.Name
}

func (s *QuerySceneListResponseBodyResultsTemplateInfoDTOList) SetDescription(v string) *QuerySceneListResponseBodyResultsTemplateInfoDTOList {
	s.Description = &v
	return s
}

func (s *QuerySceneListResponseBodyResultsTemplateInfoDTOList) SetId(v int64) *QuerySceneListResponseBodyResultsTemplateInfoDTOList {
	s.Id = &v
	return s
}

func (s *QuerySceneListResponseBodyResultsTemplateInfoDTOList) SetName(v string) *QuerySceneListResponseBodyResultsTemplateInfoDTOList {
	s.Name = &v
	return s
}

func (s *QuerySceneListResponseBodyResultsTemplateInfoDTOList) Validate() error {
	return dara.Validate(s)
}
