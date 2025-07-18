// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebalanceDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebalanceDBInstanceResponseBody
	GetRequestId() *string
}

type RebalanceDBInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5414A4E5-4C36-4461-95FC-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebalanceDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebalanceDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebalanceDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebalanceDBInstanceResponseBody) SetRequestId(v string) *RebalanceDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebalanceDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
