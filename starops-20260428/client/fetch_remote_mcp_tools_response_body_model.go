// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchRemoteMcpToolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FetchRemoteMcpToolsResponseBody
	GetRequestId() *string
	SetTools(v string) *FetchRemoteMcpToolsResponseBody
	GetTools() *string
}

type FetchRemoteMcpToolsResponseBody struct {
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// {"jsonrpc":"2.0","id":1,"result":{"tools":[]}}
	Tools *string `json:"tools,omitempty" xml:"tools,omitempty"`
}

func (s FetchRemoteMcpToolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchRemoteMcpToolsResponseBody) GoString() string {
	return s.String()
}

func (s *FetchRemoteMcpToolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchRemoteMcpToolsResponseBody) GetTools() *string {
	return s.Tools
}

func (s *FetchRemoteMcpToolsResponseBody) SetRequestId(v string) *FetchRemoteMcpToolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchRemoteMcpToolsResponseBody) SetTools(v string) *FetchRemoteMcpToolsResponseBody {
	s.Tools = &v
	return s
}

func (s *FetchRemoteMcpToolsResponseBody) Validate() error {
	return dara.Validate(s)
}
