// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubSceneId(v string) *GetOssPolicyRequest
	GetSubSceneId() *string
}

type GetOssPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s GetOssPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetOssPolicyRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *GetOssPolicyRequest) SetSubSceneId(v string) *GetOssPolicyRequest {
	s.SubSceneId = &v
	return s
}

func (s *GetOssPolicyRequest) Validate() error {
	return dara.Validate(s)
}
