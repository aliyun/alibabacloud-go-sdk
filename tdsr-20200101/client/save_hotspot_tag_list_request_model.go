// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveHotspotTagListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotspotListJson(v string) *SaveHotspotTagListRequest
	GetHotspotListJson() *string
	SetSceneId(v string) *SaveHotspotTagListRequest
	GetSceneId() *string
}

type SaveHotspotTagListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{}]
	HotspotListJson *string `json:"HotspotListJson,omitempty" xml:"HotspotListJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tqwiuwetwet****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s SaveHotspotTagListRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotTagListRequest) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagListRequest) GetHotspotListJson() *string {
	return s.HotspotListJson
}

func (s *SaveHotspotTagListRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *SaveHotspotTagListRequest) SetHotspotListJson(v string) *SaveHotspotTagListRequest {
	s.HotspotListJson = &v
	return s
}

func (s *SaveHotspotTagListRequest) SetSceneId(v string) *SaveHotspotTagListRequest {
	s.SceneId = &v
	return s
}

func (s *SaveHotspotTagListRequest) Validate() error {
	return dara.Validate(s)
}
