// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeOnBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v int64) *GetNodeOnBaselineRequest
	GetBaselineId() *int64
}

type GetNodeOnBaselineRequest struct {
	// The baseline ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
}

func (s GetNodeOnBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeOnBaselineRequest) GoString() string {
	return s.String()
}

func (s *GetNodeOnBaselineRequest) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetNodeOnBaselineRequest) SetBaselineId(v int64) *GetNodeOnBaselineRequest {
	s.BaselineId = &v
	return s
}

func (s *GetNodeOnBaselineRequest) Validate() error {
	return dara.Validate(s)
}
