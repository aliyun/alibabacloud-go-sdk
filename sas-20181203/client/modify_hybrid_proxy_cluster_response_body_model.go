// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridProxyClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHybridProxyClusterResponseBody
	GetRequestId() *string
}

type ModifyHybridProxyClusterResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8B4B6E6D-B0B0-5F05-A14E-82917D9648EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHybridProxyClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridProxyClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridProxyClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridProxyClusterResponseBody) SetRequestId(v string) *ModifyHybridProxyClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridProxyClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
