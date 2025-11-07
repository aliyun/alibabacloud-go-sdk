// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteNetworkPathResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteNetworkPathResponseBody
	GetRequestId() *string
}

type DeleteNetworkPathResponseBody struct {
	// Result of operation.
	//
	// - **true**: Delete Success.
	//
	// - **false**: Delete Fail.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C4331873-C534-590F-A905-F66C53B88A47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkPathResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPathResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteNetworkPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkPathResponseBody) SetData(v bool) *DeleteNetworkPathResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNetworkPathResponseBody) SetRequestId(v string) *DeleteNetworkPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkPathResponseBody) Validate() error {
	return dara.Validate(s)
}
