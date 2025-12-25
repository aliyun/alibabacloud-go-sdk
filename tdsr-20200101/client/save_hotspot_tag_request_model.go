// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveHotspotTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamTag(v string) *SaveHotspotTagRequest
	GetParamTag() *string
	SetSubSceneUuid(v string) *SaveHotspotTagRequest
	GetSubSceneUuid() *string
}

type SaveHotspotTagRequest struct {
	ParamTag *string `json:"ParamTag,omitempty" xml:"ParamTag,omitempty"`
	// example:
	//
	// 2345****
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
}

func (s SaveHotspotTagRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotTagRequest) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagRequest) GetParamTag() *string {
	return s.ParamTag
}

func (s *SaveHotspotTagRequest) GetSubSceneUuid() *string {
	return s.SubSceneUuid
}

func (s *SaveHotspotTagRequest) SetParamTag(v string) *SaveHotspotTagRequest {
	s.ParamTag = &v
	return s
}

func (s *SaveHotspotTagRequest) SetSubSceneUuid(v string) *SaveHotspotTagRequest {
	s.SubSceneUuid = &v
	return s
}

func (s *SaveHotspotTagRequest) Validate() error {
	return dara.Validate(s)
}
