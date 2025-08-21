// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIntentId(v int64) *UpdateIntentResponseBody
	GetIntentId() *int64
	SetRequestId(v string) *UpdateIntentResponseBody
	GetRequestId() *string
}

type UpdateIntentResponseBody struct {
	// example:
	//
	// 234234234534
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// dfaf23dfas234234234534
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIntentResponseBody) GetIntentId() *int64 {
	return s.IntentId
}

func (s *UpdateIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIntentResponseBody) SetIntentId(v int64) *UpdateIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *UpdateIntentResponseBody) SetRequestId(v string) *UpdateIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIntentResponseBody) Validate() error {
	return dara.Validate(s)
}
