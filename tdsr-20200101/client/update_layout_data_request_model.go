// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLayoutDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLayoutData(v string) *UpdateLayoutDataRequest
	GetLayoutData() *string
	SetSubSceneId(v string) *UpdateLayoutDataRequest
	GetSubSceneId() *string
}

type UpdateLayoutDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {}
	LayoutData *string `json:"LayoutData,omitempty" xml:"LayoutData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s UpdateLayoutDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLayoutDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateLayoutDataRequest) GetLayoutData() *string {
	return s.LayoutData
}

func (s *UpdateLayoutDataRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *UpdateLayoutDataRequest) SetLayoutData(v string) *UpdateLayoutDataRequest {
	s.LayoutData = &v
	return s
}

func (s *UpdateLayoutDataRequest) SetSubSceneId(v string) *UpdateLayoutDataRequest {
	s.SubSceneId = &v
	return s
}

func (s *UpdateLayoutDataRequest) Validate() error {
	return dara.Validate(s)
}
