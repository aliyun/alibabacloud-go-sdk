// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTargetId(v string) *GetAttackTargetRequest
	GetTargetId() *string
}

type GetAttackTargetRequest struct {
	// The unique identifier of the scan target. If the target does not exist or belongs to another tenant, a 400 error is returned to avoid exposing whether the resource exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// target-abc123def4567
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s GetAttackTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAttackTargetRequest) GoString() string {
	return s.String()
}

func (s *GetAttackTargetRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *GetAttackTargetRequest) SetTargetId(v string) *GetAttackTargetRequest {
	s.TargetId = &v
	return s
}

func (s *GetAttackTargetRequest) Validate() error {
	return dara.Validate(s)
}
