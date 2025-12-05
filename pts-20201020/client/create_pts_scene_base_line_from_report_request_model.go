// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePtsSceneBaseLineFromReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *CreatePtsSceneBaseLineFromReportRequest
	GetReportId() *string
	SetSceneId(v string) *CreatePtsSceneBaseLineFromReportRequest
	GetSceneId() *string
}

type CreatePtsSceneBaseLineFromReportRequest struct {
	// The ID of the report. For more information, see the [table](https://help.aliyun.com/document_detail/201321.html) provided in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// HNB78HB
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the scene. For more information, see the [table](https://help.aliyun.com/document_detail/201321.html) provided in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// VCB78HB
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s CreatePtsSceneBaseLineFromReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePtsSceneBaseLineFromReportRequest) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneBaseLineFromReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *CreatePtsSceneBaseLineFromReportRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CreatePtsSceneBaseLineFromReportRequest) SetReportId(v string) *CreatePtsSceneBaseLineFromReportRequest {
	s.ReportId = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportRequest) SetSceneId(v string) *CreatePtsSceneBaseLineFromReportRequest {
	s.SceneId = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportRequest) Validate() error {
	return dara.Validate(s)
}
