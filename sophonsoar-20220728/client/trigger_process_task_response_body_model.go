// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerProcessTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TriggerProcessTaskResponseBody
	GetRequestId() *string
}

type TriggerProcessTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 58A518BC-E4A8-5BD7-AFEA-366046ED9073
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerProcessTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerProcessTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerProcessTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerProcessTaskResponseBody) SetRequestId(v string) *TriggerProcessTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerProcessTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
