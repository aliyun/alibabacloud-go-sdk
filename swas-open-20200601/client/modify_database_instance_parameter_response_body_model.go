// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseInstanceParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDatabaseInstanceParameterResponseBody
	GetRequestId() *string
}

type ModifyDatabaseInstanceParameterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatabaseInstanceParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseInstanceParameterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDatabaseInstanceParameterResponseBody) SetRequestId(v string) *ModifyDatabaseInstanceParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
