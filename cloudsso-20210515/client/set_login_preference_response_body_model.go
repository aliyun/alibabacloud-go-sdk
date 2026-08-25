// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoginPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLoginPreferenceResponseBody
	GetRequestId() *string
}

type SetLoginPreferenceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B13E4EE-3853-5852-9165-597C32AD8FB7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoginPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLoginPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoginPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLoginPreferenceResponseBody) SetRequestId(v string) *SetLoginPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLoginPreferenceResponseBody) Validate() error {
	return dara.Validate(s)
}
