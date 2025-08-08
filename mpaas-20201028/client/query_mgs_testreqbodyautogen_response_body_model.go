// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMgsTestreqbodyautogenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryMgsTestreqbodyautogenResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMgsTestreqbodyautogenResponseBody
	GetResultCode() *string
	SetResultContent(v string) *QueryMgsTestreqbodyautogenResponseBody
	GetResultContent() *string
	SetResultMessage(v string) *QueryMgsTestreqbodyautogenResponseBody
	GetResultMessage() *string
}

type QueryMgsTestreqbodyautogenResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *string `json:"ResultContent,omitempty" xml:"ResultContent,omitempty"`
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMgsTestreqbodyautogenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMgsTestreqbodyautogenResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMgsTestreqbodyautogenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMgsTestreqbodyautogenResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMgsTestreqbodyautogenResponseBody) GetResultContent() *string {
	return s.ResultContent
}

func (s *QueryMgsTestreqbodyautogenResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMgsTestreqbodyautogenResponseBody) SetRequestId(v string) *QueryMgsTestreqbodyautogenResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMgsTestreqbodyautogenResponseBody) SetResultCode(v string) *QueryMgsTestreqbodyautogenResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMgsTestreqbodyautogenResponseBody) SetResultContent(v string) *QueryMgsTestreqbodyautogenResponseBody {
	s.ResultContent = &v
	return s
}

func (s *QueryMgsTestreqbodyautogenResponseBody) SetResultMessage(v string) *QueryMgsTestreqbodyautogenResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMgsTestreqbodyautogenResponseBody) Validate() error {
	return dara.Validate(s)
}
