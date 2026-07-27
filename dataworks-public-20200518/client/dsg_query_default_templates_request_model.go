// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryDefaultTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v int64) *DsgQueryDefaultTemplatesRequest
	GetSceneId() *int64
}

type DsgQueryDefaultTemplatesRequest struct {
	// The ID of the Level-2 Desensitization Scene. Call the [DsgSceneQuerySceneListByName](https://help.aliyun.com/document_detail/2786322.html) API to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DsgQueryDefaultTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDefaultTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DsgQueryDefaultTemplatesRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DsgQueryDefaultTemplatesRequest) SetSceneId(v int64) *DsgQueryDefaultTemplatesRequest {
	s.SceneId = &v
	return s
}

func (s *DsgQueryDefaultTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
