// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConsoleProxyResponseBody
	GetRequestId() *string
	SetResultJson(v string) *ConsoleProxyResponseBody
	GetResultJson() *string
}

type ConsoleProxyResponseBody struct {
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultJson *string `json:"resultJson,omitempty" xml:"resultJson,omitempty"`
}

func (s ConsoleProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConsoleProxyResponseBody) GoString() string {
	return s.String()
}

func (s *ConsoleProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConsoleProxyResponseBody) GetResultJson() *string {
	return s.ResultJson
}

func (s *ConsoleProxyResponseBody) SetRequestId(v string) *ConsoleProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConsoleProxyResponseBody) SetResultJson(v string) *ConsoleProxyResponseBody {
	s.ResultJson = &v
	return s
}

func (s *ConsoleProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
