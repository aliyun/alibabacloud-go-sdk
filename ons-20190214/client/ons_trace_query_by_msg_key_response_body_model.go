// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTraceQueryByMsgKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *OnsTraceQueryByMsgKeyResponseBody
	GetQueryId() *string
	SetRequestId(v string) *OnsTraceQueryByMsgKeyResponseBody
	GetRequestId() *string
}

type OnsTraceQueryByMsgKeyResponseBody struct {
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
	// F8654231-122A-4DBD-801F-38E35538****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTraceQueryByMsgKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceQueryByMsgKeyResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgKeyResponseBody) GetQueryId() *string {
	return s.QueryId
}

func (s *OnsTraceQueryByMsgKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsTraceQueryByMsgKeyResponseBody) SetQueryId(v string) *OnsTraceQueryByMsgKeyResponseBody {
	s.QueryId = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyResponseBody) SetRequestId(v string) *OnsTraceQueryByMsgKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
