// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectTrafficQosQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyExpressConnectTrafficQosQueueResponseBody
	GetRequestId() *string
}

type ModifyExpressConnectTrafficQosQueueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 079874CD-AEC1-43E6-AC03-ADD96B6E4907
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyExpressConnectTrafficQosQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectTrafficQosQueueResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectTrafficQosQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExpressConnectTrafficQosQueueResponseBody) SetRequestId(v string) *ModifyExpressConnectTrafficQosQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
