// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleApiProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConsoleApiProxyResponseBody
	GetRequestId() *string
	SetResultJson(v string) *ConsoleApiProxyResponseBody
	GetResultJson() *string
}

type ConsoleApiProxyResponseBody struct {
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultJson *string `json:"resultJson,omitempty" xml:"resultJson,omitempty"`
}

func (s ConsoleApiProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConsoleApiProxyResponseBody) GoString() string {
	return s.String()
}

func (s *ConsoleApiProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConsoleApiProxyResponseBody) GetResultJson() *string {
	return s.ResultJson
}

func (s *ConsoleApiProxyResponseBody) SetRequestId(v string) *ConsoleApiProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConsoleApiProxyResponseBody) SetResultJson(v string) *ConsoleApiProxyResponseBody {
	s.ResultJson = &v
	return s
}

func (s *ConsoleApiProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
