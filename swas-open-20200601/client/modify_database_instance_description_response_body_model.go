// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseInstanceDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDatabaseInstanceDescriptionResponseBody
	GetRequestId() *string
}

type ModifyDatabaseInstanceDescriptionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatabaseInstanceDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDatabaseInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyDatabaseInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
