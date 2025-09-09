// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDatabaseResponseBody
	GetRequestId() *string
}

type ModifyDatabaseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 162088A7-7D47-56A3-9D04-93DE7B6DBE1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDatabaseResponseBody) SetRequestId(v string) *ModifyDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
