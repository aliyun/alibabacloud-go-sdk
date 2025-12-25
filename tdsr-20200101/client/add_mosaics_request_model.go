// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMosaicsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarkPosition(v string) *AddMosaicsRequest
	GetMarkPosition() *string
	SetSubSceneId(v string) *AddMosaicsRequest
	GetSubSceneId() *string
}

type AddMosaicsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"x": 504,"y": 450,"w": 256,"h": 153}]
	MarkPosition *string `json:"MarkPosition,omitempty" xml:"MarkPosition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skjjskjk****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s AddMosaicsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMosaicsRequest) GoString() string {
	return s.String()
}

func (s *AddMosaicsRequest) GetMarkPosition() *string {
	return s.MarkPosition
}

func (s *AddMosaicsRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *AddMosaicsRequest) SetMarkPosition(v string) *AddMosaicsRequest {
	s.MarkPosition = &v
	return s
}

func (s *AddMosaicsRequest) SetSubSceneId(v string) *AddMosaicsRequest {
	s.SubSceneId = &v
	return s
}

func (s *AddMosaicsRequest) Validate() error {
	return dara.Validate(s)
}
