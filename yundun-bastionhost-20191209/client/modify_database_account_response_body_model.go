// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDatabaseAccountResponseBody
	GetRequestId() *string
}

type ModifyDatabaseAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 850FA4B4-5BD2-5269-903E-3B7E07E6C975
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatabaseAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDatabaseAccountResponseBody) SetRequestId(v string) *ModifyDatabaseAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabaseAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
