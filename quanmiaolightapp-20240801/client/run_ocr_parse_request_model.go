// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunOcrParseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *RunOcrParseRequest
	GetFileKey() *string
	SetModelId(v string) *RunOcrParseRequest
	GetModelId() *string
	SetUrl(v string) *RunOcrParseRequest
	GetUrl() *string
}

type RunOcrParseRequest struct {
	// example:
	//
	// oss://default/aimiaobi-service-prod/aimiaobi/temp/public/government_service_experience_feedback_summary.txt
	FileKey *string `json:"fileKey,omitempty" xml:"fileKey,omitempty"`
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// http://www.example.com/xxxx.txt
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RunOcrParseRequest) String() string {
	return dara.Prettify(s)
}

func (s RunOcrParseRequest) GoString() string {
	return s.String()
}

func (s *RunOcrParseRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *RunOcrParseRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunOcrParseRequest) GetUrl() *string {
	return s.Url
}

func (s *RunOcrParseRequest) SetFileKey(v string) *RunOcrParseRequest {
	s.FileKey = &v
	return s
}

func (s *RunOcrParseRequest) SetModelId(v string) *RunOcrParseRequest {
	s.ModelId = &v
	return s
}

func (s *RunOcrParseRequest) SetUrl(v string) *RunOcrParseRequest {
	s.Url = &v
	return s
}

func (s *RunOcrParseRequest) Validate() error {
	return dara.Validate(s)
}
