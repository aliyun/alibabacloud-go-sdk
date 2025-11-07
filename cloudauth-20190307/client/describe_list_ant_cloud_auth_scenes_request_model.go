// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListAntCloudAuthScenesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v int64) *DescribeListAntCloudAuthScenesRequest
	GetSceneId() *int64
}

type DescribeListAntCloudAuthScenesRequest struct {
	// Scenario ID.
	//
	// example:
	//
	// 100000xxxx
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeListAntCloudAuthScenesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeListAntCloudAuthScenesRequest) GoString() string {
	return s.String()
}

func (s *DescribeListAntCloudAuthScenesRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeListAntCloudAuthScenesRequest) SetSceneId(v int64) *DescribeListAntCloudAuthScenesRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeListAntCloudAuthScenesRequest) Validate() error {
	return dara.Validate(s)
}
