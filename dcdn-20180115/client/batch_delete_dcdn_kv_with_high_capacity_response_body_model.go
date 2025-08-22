// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnKvWithHighCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailKeys(v []*string) *BatchDeleteDcdnKvWithHighCapacityResponseBody
	GetFailKeys() []*string
	SetRequestId(v string) *BatchDeleteDcdnKvWithHighCapacityResponseBody
	GetRequestId() *string
	SetSuccessKeys(v []*string) *BatchDeleteDcdnKvWithHighCapacityResponseBody
	GetSuccessKeys() []*string
}

type BatchDeleteDcdnKvWithHighCapacityResponseBody struct {
	FailKeys []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchDeleteDcdnKvWithHighCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnKvWithHighCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponseBody) GetFailKeys() []*string {
	return s.FailKeys
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponseBody) GetSuccessKeys() []*string {
	return s.SuccessKeys
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponseBody) SetFailKeys(v []*string) *BatchDeleteDcdnKvWithHighCapacityResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponseBody) SetRequestId(v string) *BatchDeleteDcdnKvWithHighCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponseBody) SetSuccessKeys(v []*string) *BatchDeleteDcdnKvWithHighCapacityResponseBody {
	s.SuccessKeys = v
	return s
}

func (s *BatchDeleteDcdnKvWithHighCapacityResponseBody) Validate() error {
	return dara.Validate(s)
}
