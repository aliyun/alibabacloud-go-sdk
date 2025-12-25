// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOptimizeRightAngleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubSceneId(v string) *OptimizeRightAngleRequest
	GetSubSceneId() *string
}

type OptimizeRightAngleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s OptimizeRightAngleRequest) String() string {
	return dara.Prettify(s)
}

func (s OptimizeRightAngleRequest) GoString() string {
	return s.String()
}

func (s *OptimizeRightAngleRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *OptimizeRightAngleRequest) SetSubSceneId(v string) *OptimizeRightAngleRequest {
	s.SubSceneId = &v
	return s
}

func (s *OptimizeRightAngleRequest) Validate() error {
	return dara.Validate(s)
}
