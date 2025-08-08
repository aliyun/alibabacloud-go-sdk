// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInfoFromMdpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *QueryInfoFromMdpResponseBody
	GetData() *string
	SetRequestId(v string) *QueryInfoFromMdpResponseBody
	GetRequestId() *string
	SetResultCode(v int32) *QueryInfoFromMdpResponseBody
	GetResultCode() *int32
	SetResultMessage(v string) *QueryInfoFromMdpResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *QueryInfoFromMdpResponseBody
	GetSuccess() *bool
}

type QueryInfoFromMdpResponseBody struct {
	// example:
	//
	// {\\"db0797452ccafce84d7c151eb81596099bda3f097693d1e18b588804e6742ced\\":{\\"riskScore\\":\\"3\\"}}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode *int32 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInfoFromMdpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInfoFromMdpResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInfoFromMdpResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryInfoFromMdpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInfoFromMdpResponseBody) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *QueryInfoFromMdpResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryInfoFromMdpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryInfoFromMdpResponseBody) SetData(v string) *QueryInfoFromMdpResponseBody {
	s.Data = &v
	return s
}

func (s *QueryInfoFromMdpResponseBody) SetRequestId(v string) *QueryInfoFromMdpResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInfoFromMdpResponseBody) SetResultCode(v int32) *QueryInfoFromMdpResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryInfoFromMdpResponseBody) SetResultMessage(v string) *QueryInfoFromMdpResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryInfoFromMdpResponseBody) SetSuccess(v bool) *QueryInfoFromMdpResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInfoFromMdpResponseBody) Validate() error {
	return dara.Validate(s)
}
