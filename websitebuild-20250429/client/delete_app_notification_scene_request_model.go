// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppNotificationSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DeleteAppNotificationSceneRequest
	GetBizId() *string
	SetSceneId(v string) *DeleteAppNotificationSceneRequest
	GetSceneId() *string
}

type DeleteAppNotificationSceneRequest struct {
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 1000006921
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DeleteAppNotificationSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppNotificationSceneRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppNotificationSceneRequest) GetBizId() *string {
	return s.BizId
}

func (s *DeleteAppNotificationSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DeleteAppNotificationSceneRequest) SetBizId(v string) *DeleteAppNotificationSceneRequest {
	s.BizId = &v
	return s
}

func (s *DeleteAppNotificationSceneRequest) SetSceneId(v string) *DeleteAppNotificationSceneRequest {
	s.SceneId = &v
	return s
}

func (s *DeleteAppNotificationSceneRequest) Validate() error {
	return dara.Validate(s)
}
