// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWakeUpClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *WakeUpClusterRequest
	GetClientToken() *string
}

type WakeUpClusterRequest struct {
	// example:
	//
	// c82e6987e2961451182edacd74faf
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s WakeUpClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s WakeUpClusterRequest) GoString() string {
	return s.String()
}

func (s *WakeUpClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *WakeUpClusterRequest) SetClientToken(v string) *WakeUpClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *WakeUpClusterRequest) Validate() error {
	return dara.Validate(s)
}
