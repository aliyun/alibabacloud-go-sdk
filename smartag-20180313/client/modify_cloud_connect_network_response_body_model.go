// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudConnectNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCloudConnectNetworkResponseBody
	GetRequestId() *string
}

type ModifyCloudConnectNetworkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0BAAF2B9-88B8-4574-BDBE-102A90EE3FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCloudConnectNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudConnectNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudConnectNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCloudConnectNetworkResponseBody) SetRequestId(v string) *ModifyCloudConnectNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCloudConnectNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
