// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFCTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddFCTriggerResponseBody
	GetRequestId() *string
}

type AddFCTriggerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC046C5D-8CB4-4B6B-B7F8-B335E51EF90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFCTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFCTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *AddFCTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFCTriggerResponseBody) SetRequestId(v string) *AddFCTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFCTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
