// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLayoutDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubSceneId(v string) *GetLayoutDataRequest
	GetSubSceneId() *string
}

type GetLayoutDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetLayoutDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLayoutDataRequest) GoString() string {
	return s.String()
}

func (s *GetLayoutDataRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *GetLayoutDataRequest) SetSubSceneId(v string) *GetLayoutDataRequest {
	s.SubSceneId = &v
	return s
}

func (s *GetLayoutDataRequest) Validate() error {
	return dara.Validate(s)
}
