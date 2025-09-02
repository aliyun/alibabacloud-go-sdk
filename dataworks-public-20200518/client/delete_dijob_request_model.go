// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDIJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *DeleteDIJobRequest
	GetDIJobId() *int64
}

type DeleteDIJobRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11126
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
}

func (s DeleteDIJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDIJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDIJobRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *DeleteDIJobRequest) SetDIJobId(v int64) *DeleteDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *DeleteDIJobRequest) Validate() error {
	return dara.Validate(s)
}
