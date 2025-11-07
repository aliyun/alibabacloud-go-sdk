// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPathId(v string) *CreateNetworkPathResponseBody
	GetNetworkPathId() *string
	SetRequestId(v string) *CreateNetworkPathResponseBody
	GetRequestId() *string
}

type CreateNetworkPathResponseBody struct {
	// The ID of the network path.
	//
	// example:
	//
	// np-4cbf598673d14d27****
	NetworkPathId *string `json:"NetworkPathId,omitempty" xml:"NetworkPathId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 92DD9FFB-06FB-56F7-83EF-5CEF98F5562A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkPathResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkPathResponseBody) GetNetworkPathId() *string {
	return s.NetworkPathId
}

func (s *CreateNetworkPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkPathResponseBody) SetNetworkPathId(v string) *CreateNetworkPathResponseBody {
	s.NetworkPathId = &v
	return s
}

func (s *CreateNetworkPathResponseBody) SetRequestId(v string) *CreateNetworkPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkPathResponseBody) Validate() error {
	return dara.Validate(s)
}
