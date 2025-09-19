// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProcessWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyProcessWhiteListResponseBody
	GetRequestId() *string
}

type ModifyProcessWhiteListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 11C96623-E106-59C9-866D-A6C82911****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProcessWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyProcessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProcessWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyProcessWhiteListResponseBody) SetRequestId(v string) *ModifyProcessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyProcessWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
