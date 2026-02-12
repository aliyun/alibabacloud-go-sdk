// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerResetOffsetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnsConsumerResetOffsetResponseBody
	GetRequestId() *string
}

type OnsConsumerResetOffsetResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// D52C68F8-EC5D-4294-BFFF-1A6A25AF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsConsumerResetOffsetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerResetOffsetResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerResetOffsetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsConsumerResetOffsetResponseBody) SetRequestId(v string) *OnsConsumerResetOffsetResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsConsumerResetOffsetResponseBody) Validate() error {
	return dara.Validate(s)
}
