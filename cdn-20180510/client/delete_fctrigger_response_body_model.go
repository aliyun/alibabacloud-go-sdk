// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFCTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFCTriggerResponseBody
	GetRequestId() *string
}

type DeleteFCTriggerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC046C5D-8CB4-4B6B-B7F8-B335E51EF90E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFCTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFCTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFCTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFCTriggerResponseBody) SetRequestId(v string) *DeleteFCTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFCTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
