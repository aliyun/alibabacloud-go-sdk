// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSparkLogAnalyzeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SubmitSparkLogAnalyzeTaskRequest
	GetAppId() *string
}

type SubmitSparkLogAnalyzeTaskRequest struct {
	// The ID of the Spark application.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202301121553hzd9c6f7xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s SubmitSparkLogAnalyzeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSparkLogAnalyzeTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSparkLogAnalyzeTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *SubmitSparkLogAnalyzeTaskRequest) SetAppId(v string) *SubmitSparkLogAnalyzeTaskRequest {
	s.AppId = &v
	return s
}

func (s *SubmitSparkLogAnalyzeTaskRequest) Validate() error {
	return dara.Validate(s)
}
