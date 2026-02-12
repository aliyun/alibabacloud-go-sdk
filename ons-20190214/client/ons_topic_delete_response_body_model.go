// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnsTopicDeleteResponseBody
	GetRequestId() *string
}

type OnsTopicDeleteResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 4189D4A6-231A-4028-8D89-F66A76C1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsTopicDeleteResponseBody) SetRequestId(v string) *OnsTopicDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTopicDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
