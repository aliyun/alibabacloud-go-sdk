// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnKvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailKeys(v []*string) *BatchDeleteDcdnKvResponseBody
	GetFailKeys() []*string
	SetRequestId(v string) *BatchDeleteDcdnKvResponseBody
	GetRequestId() *string
	SetSuccessKeys(v []*string) *BatchDeleteDcdnKvResponseBody
	GetSuccessKeys() []*string
}

type BatchDeleteDcdnKvResponseBody struct {
	FailKeys []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchDeleteDcdnKvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnKvResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnKvResponseBody) GetFailKeys() []*string {
	return s.FailKeys
}

func (s *BatchDeleteDcdnKvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteDcdnKvResponseBody) GetSuccessKeys() []*string {
	return s.SuccessKeys
}

func (s *BatchDeleteDcdnKvResponseBody) SetFailKeys(v []*string) *BatchDeleteDcdnKvResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchDeleteDcdnKvResponseBody) SetRequestId(v string) *BatchDeleteDcdnKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteDcdnKvResponseBody) SetSuccessKeys(v []*string) *BatchDeleteDcdnKvResponseBody {
	s.SuccessKeys = v
	return s
}

func (s *BatchDeleteDcdnKvResponseBody) Validate() error {
	return dara.Validate(s)
}
