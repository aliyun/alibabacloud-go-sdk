// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceAttributeResponseBody
	GetRequestId() *string
}

type ModifyInstanceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceAttributeResponseBody) SetRequestId(v string) *ModifyInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
