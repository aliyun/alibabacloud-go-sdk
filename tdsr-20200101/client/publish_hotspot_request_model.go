// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishHotspotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamTag(v string) *PublishHotspotRequest
	GetParamTag() *string
	SetSubSceneUuid(v string) *PublishHotspotRequest
	GetSubSceneUuid() *string
}

type PublishHotspotRequest struct {
	ParamTag *string `json:"ParamTag,omitempty" xml:"ParamTag,omitempty"`
	// example:
	//
	// 2345****
	SubSceneUuid *string `json:"SubSceneUuid,omitempty" xml:"SubSceneUuid,omitempty"`
}

func (s PublishHotspotRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishHotspotRequest) GoString() string {
	return s.String()
}

func (s *PublishHotspotRequest) GetParamTag() *string {
	return s.ParamTag
}

func (s *PublishHotspotRequest) GetSubSceneUuid() *string {
	return s.SubSceneUuid
}

func (s *PublishHotspotRequest) SetParamTag(v string) *PublishHotspotRequest {
	s.ParamTag = &v
	return s
}

func (s *PublishHotspotRequest) SetSubSceneUuid(v string) *PublishHotspotRequest {
	s.SubSceneUuid = &v
	return s
}

func (s *PublishHotspotRequest) Validate() error {
	return dara.Validate(s)
}
