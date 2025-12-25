// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSingleConnDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubSceneId(v string) *GetSingleConnDataRequest
	GetSubSceneId() *string
}

type GetSingleConnDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetSingleConnDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSingleConnDataRequest) GoString() string {
	return s.String()
}

func (s *GetSingleConnDataRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *GetSingleConnDataRequest) SetSubSceneId(v string) *GetSingleConnDataRequest {
	s.SubSceneId = &v
	return s
}

func (s *GetSingleConnDataRequest) Validate() error {
	return dara.Validate(s)
}
