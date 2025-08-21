// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIntentId(v int64) *CreateIntentResponseBody
	GetIntentId() *int64
	SetRequestId(v string) *CreateIntentResponseBody
	GetRequestId() *string
}

type CreateIntentResponseBody struct {
	// example:
	//
	// 43546474
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// df56gjh5et34g3g3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIntentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntentResponseBody) GetIntentId() *int64 {
	return s.IntentId
}

func (s *CreateIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIntentResponseBody) SetIntentId(v int64) *CreateIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *CreateIntentResponseBody) SetRequestId(v string) *CreateIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIntentResponseBody) Validate() error {
	return dara.Validate(s)
}
