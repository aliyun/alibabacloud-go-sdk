// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelStockOssCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CancelStockOssCheckTaskResponseBody
	GetData() *bool
	SetRequestId(v string) *CancelStockOssCheckTaskResponseBody
	GetRequestId() *string
}

type CancelStockOssCheckTaskResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelStockOssCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelStockOssCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelStockOssCheckTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *CancelStockOssCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelStockOssCheckTaskResponseBody) SetData(v bool) *CancelStockOssCheckTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelStockOssCheckTaskResponseBody) SetRequestId(v string) *CancelStockOssCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelStockOssCheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
