// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySourceServerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySourceServerAttributeResponseBody
	GetRequestId() *string
}

type ModifySourceServerAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySourceServerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySourceServerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySourceServerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySourceServerAttributeResponseBody) SetRequestId(v string) *ModifySourceServerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySourceServerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
