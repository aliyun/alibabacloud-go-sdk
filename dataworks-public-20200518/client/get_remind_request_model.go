// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRemindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRemindId(v int64) *GetRemindRequest
	GetRemindId() *int64
}

type GetRemindRequest struct {
	// The custom alert rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	RemindId *int64 `json:"RemindId,omitempty" xml:"RemindId,omitempty"`
}

func (s GetRemindRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRemindRequest) GoString() string {
	return s.String()
}

func (s *GetRemindRequest) GetRemindId() *int64 {
	return s.RemindId
}

func (s *GetRemindRequest) SetRemindId(v int64) *GetRemindRequest {
	s.RemindId = &v
	return s
}

func (s *GetRemindRequest) Validate() error {
	return dara.Validate(s)
}
