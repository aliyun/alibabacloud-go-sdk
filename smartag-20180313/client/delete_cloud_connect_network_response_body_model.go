// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudConnectNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudConnectNetworkResponseBody
	GetRequestId() *string
}

type DeleteCloudConnectNetworkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0BAAF2B9-88B8-4574-BDBE-102A90EE3FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudConnectNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudConnectNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudConnectNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudConnectNetworkResponseBody) SetRequestId(v string) *DeleteCloudConnectNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudConnectNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
