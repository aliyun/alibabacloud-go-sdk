// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *DeleteDataAgentRequest
	GetDMSUnit() *string
}

type DeleteDataAgentRequest struct {
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
}

func (s DeleteDataAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *DeleteDataAgentRequest) SetDMSUnit(v string) *DeleteDataAgentRequest {
	s.DMSUnit = &v
	return s
}

func (s *DeleteDataAgentRequest) Validate() error {
	return dara.Validate(s)
}
