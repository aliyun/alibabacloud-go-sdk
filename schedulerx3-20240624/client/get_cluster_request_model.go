// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetClusterRequest
	GetClusterId() *string
}

type GetClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-d6a5243b6fa
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRequest) GoString() string {
	return s.String()
}

func (s *GetClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterRequest) SetClusterId(v string) *GetClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterRequest) Validate() error {
	return dara.Validate(s)
}
