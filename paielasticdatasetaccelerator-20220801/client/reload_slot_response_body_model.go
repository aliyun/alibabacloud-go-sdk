// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReloadSlotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReloadSlotResponseBody
	GetRequestId() *string
}

type ReloadSlotResponseBody struct {
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReloadSlotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReloadSlotResponseBody) GoString() string {
	return s.String()
}

func (s *ReloadSlotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReloadSlotResponseBody) SetRequestId(v string) *ReloadSlotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReloadSlotResponseBody) Validate() error {
	return dara.Validate(s)
}
