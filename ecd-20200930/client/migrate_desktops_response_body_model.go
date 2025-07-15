// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MigrateDesktopsResponseBody
	GetRequestId() *string
}

type MigrateDesktopsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E46DECEC-AC72-570E-958B-B52A4B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateDesktopsResponseBody) SetRequestId(v string) *MigrateDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateDesktopsResponseBody) Validate() error {
	return dara.Validate(s)
}
