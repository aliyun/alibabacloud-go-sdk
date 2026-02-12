// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessagePushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnsMessagePushResponseBody
	GetRequestId() *string
}

type OnsMessagePushResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// B8EDC90D-F726-4B9E-8BEF-F0DD25EC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsMessagePushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsMessagePushResponseBody) GoString() string {
	return s.String()
}

func (s *OnsMessagePushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsMessagePushResponseBody) SetRequestId(v string) *OnsMessagePushResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsMessagePushResponseBody) Validate() error {
	return dara.Validate(s)
}
