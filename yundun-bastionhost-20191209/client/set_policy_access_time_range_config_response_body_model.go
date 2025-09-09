// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyAccessTimeRangeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPolicyAccessTimeRangeConfigResponseBody
	GetRequestId() *string
}

type SetPolicyAccessTimeRangeConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPolicyAccessTimeRangeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAccessTimeRangeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetPolicyAccessTimeRangeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPolicyAccessTimeRangeConfigResponseBody) SetRequestId(v string) *SetPolicyAccessTimeRangeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
