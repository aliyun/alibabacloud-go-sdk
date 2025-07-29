// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunClusterInspectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RunClusterInspectRequest
	GetClientToken() *string
}

type RunClusterInspectRequest struct {
	// The idempotency token that ensures an API request completes no more than one time.
	//
	// example:
	//
	// c82e6987e2961451182edacd74faf
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s RunClusterInspectRequest) String() string {
	return dara.Prettify(s)
}

func (s RunClusterInspectRequest) GoString() string {
	return s.String()
}

func (s *RunClusterInspectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunClusterInspectRequest) SetClientToken(v string) *RunClusterInspectRequest {
	s.ClientToken = &v
	return s
}

func (s *RunClusterInspectRequest) Validate() error {
	return dara.Validate(s)
}
