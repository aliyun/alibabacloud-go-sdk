// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddConnectableClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddConnectableClusterResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddConnectableClusterResponseBody
	GetResult() *bool
}

type AddConnectableClusterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5A5D8E74-565C-43DC-B031-29289FA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The following information is returned:
	//
	// 	- true: The configuration is successful.
	//
	// 	- false: The configuration failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AddConnectableClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddConnectableClusterResponseBody) GoString() string {
	return s.String()
}

func (s *AddConnectableClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddConnectableClusterResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddConnectableClusterResponseBody) SetRequestId(v string) *AddConnectableClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddConnectableClusterResponseBody) SetResult(v bool) *AddConnectableClusterResponseBody {
	s.Result = &v
	return s
}

func (s *AddConnectableClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
