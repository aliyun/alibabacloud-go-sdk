// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteKvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailKeys(v []*string) *BatchDeleteKvResponseBody
	GetFailKeys() []*string
	SetRequestId(v string) *BatchDeleteKvResponseBody
	GetRequestId() *string
	SetSuccessKeys(v []*string) *BatchDeleteKvResponseBody
	GetSuccessKeys() []*string
}

type BatchDeleteKvResponseBody struct {
	FailKeys    []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchDeleteKvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteKvResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvResponseBody) GetFailKeys() []*string {
	return s.FailKeys
}

func (s *BatchDeleteKvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteKvResponseBody) GetSuccessKeys() []*string {
	return s.SuccessKeys
}

func (s *BatchDeleteKvResponseBody) SetFailKeys(v []*string) *BatchDeleteKvResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchDeleteKvResponseBody) SetRequestId(v string) *BatchDeleteKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteKvResponseBody) SetSuccessKeys(v []*string) *BatchDeleteKvResponseBody {
	s.SuccessKeys = v
	return s
}

func (s *BatchDeleteKvResponseBody) Validate() error {
	return dara.Validate(s)
}
