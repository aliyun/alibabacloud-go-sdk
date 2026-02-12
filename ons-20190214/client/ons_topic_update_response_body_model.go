// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnsTopicUpdateResponseBody
	GetRequestId() *string
}

type OnsTopicUpdateResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 81979ADA-4A78-4F64-9DEC-5700446D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsTopicUpdateResponseBody) SetRequestId(v string) *OnsTopicUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTopicUpdateResponseBody) Validate() error {
	return dara.Validate(s)
}
