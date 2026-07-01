// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySystemEventAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySystemEventAttributeResponseBody
	GetRequestId() *string
}

type ModifySystemEventAttributeResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySystemEventAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySystemEventAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySystemEventAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySystemEventAttributeResponseBody) SetRequestId(v string) *ModifySystemEventAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySystemEventAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
