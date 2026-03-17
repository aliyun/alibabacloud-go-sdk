// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseSupabaseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PauseSupabaseProjectResponseBody
	GetRequestId() *string
}

type PauseSupabaseProjectResponseBody struct {
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PauseSupabaseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseSupabaseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PauseSupabaseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseSupabaseProjectResponseBody) SetRequestId(v string) *PauseSupabaseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseSupabaseProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
