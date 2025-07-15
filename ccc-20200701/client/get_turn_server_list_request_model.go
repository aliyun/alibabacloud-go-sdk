// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTurnServerListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetTurnServerListRequest
	GetInstanceId() *string
}

type GetTurnServerListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTurnServerListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTurnServerListRequest) GoString() string {
	return s.String()
}

func (s *GetTurnServerListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTurnServerListRequest) SetInstanceId(v string) *GetTurnServerListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTurnServerListRequest) Validate() error {
	return dara.Validate(s)
}
