// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceIPArrayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDBInstanceIPArrayResponseBody
	GetRequestId() *string
}

type CreateDBInstanceIPArrayResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// API-d971c90f-c801-41bd-b8e5-5b8bd79326c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBInstanceIPArrayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceIPArrayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceIPArrayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceIPArrayResponseBody) SetRequestId(v string) *CreateDBInstanceIPArrayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceIPArrayResponseBody) Validate() error {
	return dara.Validate(s)
}
