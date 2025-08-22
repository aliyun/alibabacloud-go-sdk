// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutDcdnKvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailKeys(v []*string) *BatchPutDcdnKvResponseBody
	GetFailKeys() []*string
	SetRequestId(v string) *BatchPutDcdnKvResponseBody
	GetRequestId() *string
	SetSuccessKeys(v []*string) *BatchPutDcdnKvResponseBody
	GetSuccessKeys() []*string
}

type BatchPutDcdnKvResponseBody struct {
	// The keys that failed to be written.
	FailKeys []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 156A6B-677B1A-4297B7-9187B7-2B44792
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The keys that were written.
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchPutDcdnKvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchPutDcdnKvResponseBody) GoString() string {
	return s.String()
}

func (s *BatchPutDcdnKvResponseBody) GetFailKeys() []*string {
	return s.FailKeys
}

func (s *BatchPutDcdnKvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchPutDcdnKvResponseBody) GetSuccessKeys() []*string {
	return s.SuccessKeys
}

func (s *BatchPutDcdnKvResponseBody) SetFailKeys(v []*string) *BatchPutDcdnKvResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchPutDcdnKvResponseBody) SetRequestId(v string) *BatchPutDcdnKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchPutDcdnKvResponseBody) SetSuccessKeys(v []*string) *BatchPutDcdnKvResponseBody {
	s.SuccessKeys = v
	return s
}

func (s *BatchPutDcdnKvResponseBody) Validate() error {
	return dara.Validate(s)
}
