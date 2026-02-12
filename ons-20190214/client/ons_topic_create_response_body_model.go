// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnsTopicCreateResponseBody
	GetRequestId() *string
}

type OnsTopicCreateResponseBody struct {
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// B6949B58-223E-4B75-B4FE-7797C15E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsTopicCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicCreateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsTopicCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsTopicCreateResponseBody) SetRequestId(v string) *OnsTopicCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsTopicCreateResponseBody) Validate() error {
	return dara.Validate(s)
}
