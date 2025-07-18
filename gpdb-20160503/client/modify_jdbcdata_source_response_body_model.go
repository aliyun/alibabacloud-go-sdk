// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJDBCDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyJDBCDataSourceResponseBody
	GetRequestId() *string
}

type ModifyJDBCDataSourceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 2C125605-266F-41CA-8AC5-3A643D4F42C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyJDBCDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyJDBCDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyJDBCDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyJDBCDataSourceResponseBody) SetRequestId(v string) *ModifyJDBCDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyJDBCDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
