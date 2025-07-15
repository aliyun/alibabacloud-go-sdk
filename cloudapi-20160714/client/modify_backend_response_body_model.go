// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBackendResponseBody
	GetRequestId() *string
}

type ModifyBackendResponseBody struct {
	// example:
	//
	// 06DACA61-9359-5EC6-AEDA-C73E620E49A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackendResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackendResponseBody) SetRequestId(v string) *ModifyBackendResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackendResponseBody) Validate() error {
	return dara.Validate(s)
}
