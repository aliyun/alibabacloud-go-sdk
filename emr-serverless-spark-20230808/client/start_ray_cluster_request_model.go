// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRayClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartRayClusterRequest
	GetClientToken() *string
}

type StartRayClusterRequest struct {
	// example:
	//
	// c533e141-bf99-4236-8b6b-30e133db113c
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s StartRayClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRayClusterRequest) GoString() string {
	return s.String()
}

func (s *StartRayClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartRayClusterRequest) SetClientToken(v string) *StartRayClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *StartRayClusterRequest) Validate() error {
	return dara.Validate(s)
}
