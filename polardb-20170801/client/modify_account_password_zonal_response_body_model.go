// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPasswordZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccountPasswordZonalResponseBody
	GetRequestId() *string
}

type ModifyAccountPasswordZonalResponseBody struct {
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountPasswordZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPasswordZonalResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountPasswordZonalResponseBody) SetRequestId(v string) *ModifyAccountPasswordZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountPasswordZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
