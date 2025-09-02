// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPublicModelEngineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *QueryPublicModelEngineRequest
	GetProjectId() *string
	SetText(v string) *QueryPublicModelEngineRequest
	GetText() *string
}

type QueryPublicModelEngineRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FML statement that is used to query information about objects that are created in Data Modeling. For more information, see [Use FML statements to configure and manage data tables](https://help.aliyun.com/document_detail/298128.html). Only SHOW statements are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// show dim tables
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s QueryPublicModelEngineRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPublicModelEngineRequest) GoString() string {
	return s.String()
}

func (s *QueryPublicModelEngineRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryPublicModelEngineRequest) GetText() *string {
	return s.Text
}

func (s *QueryPublicModelEngineRequest) SetProjectId(v string) *QueryPublicModelEngineRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryPublicModelEngineRequest) SetText(v string) *QueryPublicModelEngineRequest {
	s.Text = &v
	return s
}

func (s *QueryPublicModelEngineRequest) Validate() error {
	return dara.Validate(s)
}
