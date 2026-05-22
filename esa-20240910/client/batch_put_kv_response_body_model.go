// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutKvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailKeys(v []*string) *BatchPutKvResponseBody
	GetFailKeys() []*string
	SetRequestId(v string) *BatchPutKvResponseBody
	GetRequestId() *string
	SetSuccessKeys(v []*string) *BatchPutKvResponseBody
	GetSuccessKeys() []*string
}

type BatchPutKvResponseBody struct {
	FailKeys    []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchPutKvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchPutKvResponseBody) GoString() string {
	return s.String()
}

func (s *BatchPutKvResponseBody) GetFailKeys() []*string {
	return s.FailKeys
}

func (s *BatchPutKvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchPutKvResponseBody) GetSuccessKeys() []*string {
	return s.SuccessKeys
}

func (s *BatchPutKvResponseBody) SetFailKeys(v []*string) *BatchPutKvResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchPutKvResponseBody) SetRequestId(v string) *BatchPutKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchPutKvResponseBody) SetSuccessKeys(v []*string) *BatchPutKvResponseBody {
	s.SuccessKeys = v
	return s
}

func (s *BatchPutKvResponseBody) Validate() error {
	return dara.Validate(s)
}
