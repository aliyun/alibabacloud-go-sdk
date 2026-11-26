// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDeviceBelongIdentifiersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveDeviceBelongIdentifiersResponseBody
	GetRequestId() *string
}

type RemoveDeviceBelongIdentifiersResponseBody struct {
	// The request ID, which can be used to locate the call.
	//
	// example:
	//
	// 7A92B549-EC82-5F79-8301-63711AFC9DCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveDeviceBelongIdentifiersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDeviceBelongIdentifiersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDeviceBelongIdentifiersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDeviceBelongIdentifiersResponseBody) SetRequestId(v string) *RemoveDeviceBelongIdentifiersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDeviceBelongIdentifiersResponseBody) Validate() error {
	return dara.Validate(s)
}
