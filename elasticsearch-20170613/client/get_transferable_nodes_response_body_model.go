// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransferableNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTransferableNodesResponseBody
	GetRequestId() *string
	SetResult(v []*GetTransferableNodesResponseBodyResult) *GetTransferableNodesResponseBody
	GetResult() []*GetTransferableNodesResponseBodyResult
}

type GetTransferableNodesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
	Result []*GetTransferableNodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetTransferableNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTransferableNodesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTransferableNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTransferableNodesResponseBody) GetResult() []*GetTransferableNodesResponseBodyResult {
	return s.Result
}

func (s *GetTransferableNodesResponseBody) SetRequestId(v string) *GetTransferableNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTransferableNodesResponseBody) SetResult(v []*GetTransferableNodesResponseBodyResult) *GetTransferableNodesResponseBody {
	s.Result = v
	return s
}

func (s *GetTransferableNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTransferableNodesResponseBodyResult struct {
	// The IP address of the node.
	//
	// example:
	//
	// ``192.168.**.**``
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The access port of the node.
	//
	// example:
	//
	// 9200
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s GetTransferableNodesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetTransferableNodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTransferableNodesResponseBodyResult) GetHost() *string {
	return s.Host
}

func (s *GetTransferableNodesResponseBodyResult) GetPort() *int32 {
	return s.Port
}

func (s *GetTransferableNodesResponseBodyResult) SetHost(v string) *GetTransferableNodesResponseBodyResult {
	s.Host = &v
	return s
}

func (s *GetTransferableNodesResponseBodyResult) SetPort(v int32) *GetTransferableNodesResponseBodyResult {
	s.Port = &v
	return s
}

func (s *GetTransferableNodesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
