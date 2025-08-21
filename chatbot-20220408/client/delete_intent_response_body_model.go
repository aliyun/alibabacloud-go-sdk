// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIntentId(v int64) *DeleteIntentResponseBody
	GetIntentId() *int64
	SetRequestId(v string) *DeleteIntentResponseBody
	GetRequestId() *string
}

type DeleteIntentResponseBody struct {
	// example:
	//
	// 12345
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// 3464dfg3qwr34tf34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntentResponseBody) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DeleteIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIntentResponseBody) SetIntentId(v int64) *DeleteIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *DeleteIntentResponseBody) SetRequestId(v string) *DeleteIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIntentResponseBody) Validate() error {
	return dara.Validate(s)
}
