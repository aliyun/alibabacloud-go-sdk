// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySceneConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *QuerySceneConfigsRequest
	GetType() *string
}

type QuerySceneConfigsRequest struct {
	// Scene type.
	//
	// This parameter is required.
	//
	// example:
	//
	// VOLUNTARY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QuerySceneConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneConfigsRequest) GoString() string {
	return s.String()
}

func (s *QuerySceneConfigsRequest) GetType() *string {
	return s.Type
}

func (s *QuerySceneConfigsRequest) SetType(v string) *QuerySceneConfigsRequest {
	s.Type = &v
	return s
}

func (s *QuerySceneConfigsRequest) Validate() error {
	return dara.Validate(s)
}
