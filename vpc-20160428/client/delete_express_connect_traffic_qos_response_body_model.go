// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectTrafficQosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExpressConnectTrafficQosResponseBody
	GetRequestId() *string
}

type DeleteExpressConnectTrafficQosResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 606998F0-B94D-48FE-8316-ACA81BB230DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExpressConnectTrafficQosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectTrafficQosResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectTrafficQosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExpressConnectTrafficQosResponseBody) SetRequestId(v string) *DeleteExpressConnectTrafficQosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosResponseBody) Validate() error {
	return dara.Validate(s)
}
