// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceCLSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceCLSResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceCLSResponseBody struct {
	// example:
	//
	// 2144F5CC-10C5-3B72-8C74-E5***********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceCLSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceCLSResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceCLSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceCLSResponseBody) SetRequestId(v string) *ModifyDBInstanceCLSResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceCLSResponseBody) Validate() error {
	return dara.Validate(s)
}
