// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginLayoutDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubSceneId(v string) *GetOriginLayoutDataRequest
	GetSubSceneId() *string
}

type GetOriginLayoutDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetOriginLayoutDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOriginLayoutDataRequest) GoString() string {
	return s.String()
}

func (s *GetOriginLayoutDataRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *GetOriginLayoutDataRequest) SetSubSceneId(v string) *GetOriginLayoutDataRequest {
	s.SubSceneId = &v
	return s
}

func (s *GetOriginLayoutDataRequest) Validate() error {
	return dara.Validate(s)
}
