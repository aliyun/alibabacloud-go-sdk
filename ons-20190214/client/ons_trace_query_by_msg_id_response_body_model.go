// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTraceQueryByMsgIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *OnsTraceQueryByMsgIdResponseBody
	GetQueryId() *string
	SetRequestId(v string) *OnsTraceQueryByMsgIdResponseBody
	GetRequestId() *string
}

type OnsTraceQueryByMsgIdResponseBody struct {
	// The ID of the query task. You can call the [OnsTraceGetResult](https://help.aliyun.com/document_detail/59832.html) operation to query the details of the message trace based on the task ID.
	//
	// example:
	//
	// 272967562652883649157096685****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// B93332A3-160D-404F-880F-1F8736D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTraceQueryByMsgIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceQueryByMsgIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgIdResponseBody) GetQueryId() *string {
	return s.QueryId
}

func (s *OnsTraceQueryByMsgIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsTraceQueryByMsgIdResponseBody) SetQueryId(v string) *OnsTraceQueryByMsgIdResponseBody {
	s.QueryId = &v
	return s
}

func (s *OnsTraceQueryByMsgIdResponseBody) SetRequestId(v string) *OnsTraceQueryByMsgIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTraceQueryByMsgIdResponseBody) Validate() error {
	return dara.Validate(s)
}
