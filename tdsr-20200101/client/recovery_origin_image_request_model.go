// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoveryOriginImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubSceneId(v string) *RecoveryOriginImageRequest
	GetSubSceneId() *string
}

type RecoveryOriginImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// skjjskjk****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s RecoveryOriginImageRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoveryOriginImageRequest) GoString() string {
	return s.String()
}

func (s *RecoveryOriginImageRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *RecoveryOriginImageRequest) SetSubSceneId(v string) *RecoveryOriginImageRequest {
	s.SubSceneId = &v
	return s
}

func (s *RecoveryOriginImageRequest) Validate() error {
	return dara.Validate(s)
}
