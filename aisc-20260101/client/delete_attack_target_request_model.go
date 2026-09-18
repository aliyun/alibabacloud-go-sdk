// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttackTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTargetId(v string) *DeleteAttackTargetRequest
	GetTargetId() *string
}

type DeleteAttackTargetRequest struct {
	// The unique identifier of the scan target. This is the TargetId returned by CreateAttackTarget or ListAttackTargets. If the target does not exist or belongs to another tenant, a 400 error is returned. This prevents exposing whether the resource exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// target-abc123def4567
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s DeleteAttackTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttackTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteAttackTargetRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *DeleteAttackTargetRequest) SetTargetId(v string) *DeleteAttackTargetRequest {
	s.TargetId = &v
	return s
}

func (s *DeleteAttackTargetRequest) Validate() error {
	return dara.Validate(s)
}
