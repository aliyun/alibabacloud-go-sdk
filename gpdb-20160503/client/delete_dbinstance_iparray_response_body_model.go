// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceIPArrayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBInstanceIPArrayResponseBody
	GetRequestId() *string
}

type DeleteDBInstanceIPArrayResponseBody struct {
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceIPArrayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceIPArrayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceIPArrayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBInstanceIPArrayResponseBody) SetRequestId(v string) *DeleteDBInstanceIPArrayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBInstanceIPArrayResponseBody) Validate() error {
	return dara.Validate(s)
}
