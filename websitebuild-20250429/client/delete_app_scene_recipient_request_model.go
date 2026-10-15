// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppSceneRecipientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DeleteAppSceneRecipientRequest
	GetBizId() *string
	SetRecipientId(v string) *DeleteAppSceneRecipientRequest
	GetRecipientId() *string
	SetSceneId(v string) *DeleteAppSceneRecipientRequest
	GetSceneId() *string
}

type DeleteAppSceneRecipientRequest struct {
	// The business ID.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The recipient ID.
	//
	// example:
	//
	// 1
	RecipientId *string `json:"RecipientId,omitempty" xml:"RecipientId,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// 1000006921
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DeleteAppSceneRecipientRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppSceneRecipientRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppSceneRecipientRequest) GetBizId() *string {
	return s.BizId
}

func (s *DeleteAppSceneRecipientRequest) GetRecipientId() *string {
	return s.RecipientId
}

func (s *DeleteAppSceneRecipientRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DeleteAppSceneRecipientRequest) SetBizId(v string) *DeleteAppSceneRecipientRequest {
	s.BizId = &v
	return s
}

func (s *DeleteAppSceneRecipientRequest) SetRecipientId(v string) *DeleteAppSceneRecipientRequest {
	s.RecipientId = &v
	return s
}

func (s *DeleteAppSceneRecipientRequest) SetSceneId(v string) *DeleteAppSceneRecipientRequest {
	s.SceneId = &v
	return s
}

func (s *DeleteAppSceneRecipientRequest) Validate() error {
	return dara.Validate(s)
}
