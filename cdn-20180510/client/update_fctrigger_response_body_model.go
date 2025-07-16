// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFCTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFCTriggerResponseBody
	GetRequestId() *string
}

type UpdateFCTriggerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC046C5D-8CB4-4B6B-B7F8-B335E51EF90E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFCTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFCTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFCTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFCTriggerResponseBody) SetRequestId(v string) *UpdateFCTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFCTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
