// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterSceneRunningDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *GetJMeterSceneRunningDataRequest
	GetSceneId() *string
}

type GetJMeterSceneRunningDataRequest struct {
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetJMeterSceneRunningDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterSceneRunningDataRequest) GoString() string {
	return s.String()
}

func (s *GetJMeterSceneRunningDataRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *GetJMeterSceneRunningDataRequest) SetSceneId(v string) *GetJMeterSceneRunningDataRequest {
	s.SceneId = &v
	return s
}

func (s *GetJMeterSceneRunningDataRequest) Validate() error {
	return dara.Validate(s)
}
