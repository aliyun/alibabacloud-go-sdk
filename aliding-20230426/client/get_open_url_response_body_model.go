// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOpenUrlResponseBody
	GetRequestId() *string
	SetResult(v string) *GetOpenUrlResponseBody
	GetResult() *string
}

type GetOpenUrlResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// https/www.aliwork.com
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetOpenUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpenUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpenUrlResponseBody) GetResult() *string {
	return s.Result
}

func (s *GetOpenUrlResponseBody) SetRequestId(v string) *GetOpenUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpenUrlResponseBody) SetResult(v string) *GetOpenUrlResponseBody {
	s.Result = &v
	return s
}

func (s *GetOpenUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
