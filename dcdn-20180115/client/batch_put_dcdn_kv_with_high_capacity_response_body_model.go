// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutDcdnKvWithHighCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailKeys(v []*string) *BatchPutDcdnKvWithHighCapacityResponseBody
	GetFailKeys() []*string
	SetRequestId(v string) *BatchPutDcdnKvWithHighCapacityResponseBody
	GetRequestId() *string
	SetSuccessKeys(v []*string) *BatchPutDcdnKvWithHighCapacityResponseBody
	GetSuccessKeys() []*string
}

type BatchPutDcdnKvWithHighCapacityResponseBody struct {
	FailKeys []*string `json:"FailKeys,omitempty" xml:"FailKeys,omitempty" type:"Repeated"`
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessKeys []*string `json:"SuccessKeys,omitempty" xml:"SuccessKeys,omitempty" type:"Repeated"`
}

func (s BatchPutDcdnKvWithHighCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchPutDcdnKvWithHighCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *BatchPutDcdnKvWithHighCapacityResponseBody) GetFailKeys() []*string {
	return s.FailKeys
}

func (s *BatchPutDcdnKvWithHighCapacityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchPutDcdnKvWithHighCapacityResponseBody) GetSuccessKeys() []*string {
	return s.SuccessKeys
}

func (s *BatchPutDcdnKvWithHighCapacityResponseBody) SetFailKeys(v []*string) *BatchPutDcdnKvWithHighCapacityResponseBody {
	s.FailKeys = v
	return s
}

func (s *BatchPutDcdnKvWithHighCapacityResponseBody) SetRequestId(v string) *BatchPutDcdnKvWithHighCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchPutDcdnKvWithHighCapacityResponseBody) SetSuccessKeys(v []*string) *BatchPutDcdnKvWithHighCapacityResponseBody {
	s.SuccessKeys = v
	return s
}

func (s *BatchPutDcdnKvWithHighCapacityResponseBody) Validate() error {
	return dara.Validate(s)
}
