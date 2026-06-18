// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartSupabaseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartSupabaseProjectResponseBody
	GetRequestId() *string
}

type RestartSupabaseProjectResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartSupabaseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartSupabaseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *RestartSupabaseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartSupabaseProjectResponseBody) SetRequestId(v string) *RestartSupabaseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartSupabaseProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
