// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRemindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRemindId(v int64) *DeleteRemindRequest
	GetRemindId() *int64
}

type DeleteRemindRequest struct {
	// The ID of the custom alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	RemindId *int64 `json:"RemindId,omitempty" xml:"RemindId,omitempty"`
}

func (s DeleteRemindRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRemindRequest) GoString() string {
	return s.String()
}

func (s *DeleteRemindRequest) GetRemindId() *int64 {
	return s.RemindId
}

func (s *DeleteRemindRequest) SetRemindId(v int64) *DeleteRemindRequest {
	s.RemindId = &v
	return s
}

func (s *DeleteRemindRequest) Validate() error {
	return dara.Validate(s)
}
