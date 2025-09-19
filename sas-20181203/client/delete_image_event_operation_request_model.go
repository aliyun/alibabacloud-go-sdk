// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageEventOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteImageEventOperationRequest
	GetId() *int64
}

type DeleteImageEventOperationRequest struct {
	// The primary key of the alert handling rule.
	//
	// example:
	//
	// 1404656
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteImageEventOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageEventOperationRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageEventOperationRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteImageEventOperationRequest) SetId(v int64) *DeleteImageEventOperationRequest {
	s.Id = &v
	return s
}

func (s *DeleteImageEventOperationRequest) Validate() error {
	return dara.Validate(s)
}
