// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatStockOssCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreatStockOssCheckTaskResponseBody
	GetData() *bool
	SetRequestId(v string) *CreatStockOssCheckTaskResponseBody
	GetRequestId() *string
}

type CreatStockOssCheckTaskResponseBody struct {
	// Returned data
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

func (s CreatStockOssCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatStockOssCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatStockOssCheckTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreatStockOssCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatStockOssCheckTaskResponseBody) SetData(v bool) *CreatStockOssCheckTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreatStockOssCheckTaskResponseBody) SetRequestId(v string) *CreatStockOssCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatStockOssCheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
