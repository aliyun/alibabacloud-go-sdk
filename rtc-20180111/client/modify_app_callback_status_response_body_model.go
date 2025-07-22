// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppCallbackStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppCallbackStatusResponseBody
	GetRequestId() *string
}

type ModifyAppCallbackStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F80AAC02-87BD-5D9C-B925-8AB40171BA1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppCallbackStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppCallbackStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppCallbackStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppCallbackStatusResponseBody) SetRequestId(v string) *ModifyAppCallbackStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppCallbackStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
