// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceSSLResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceSSLResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6D806B11-078F-4154-BF9F-844F56D08964
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceSSLResponseBody) SetRequestId(v string) *ModifyDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
