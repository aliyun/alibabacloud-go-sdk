// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagSupabaseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagSupabaseProjectResponseBody
	GetRequestId() *string
}

type TagSupabaseProjectResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagSupabaseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagSupabaseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *TagSupabaseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagSupabaseProjectResponseBody) SetRequestId(v string) *TagSupabaseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagSupabaseProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
