// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTriggerResponseBody
	GetRequestId() *string
}

type UpdateTriggerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A022F78-B9A8-4ACC-BB6B-B35975******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTriggerResponseBody) SetRequestId(v string) *UpdateTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
