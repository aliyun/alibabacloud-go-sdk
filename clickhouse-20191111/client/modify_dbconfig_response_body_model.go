// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBConfigResponseBody
	GetRequestId() *string
}

type ModifyDBConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BF3844B6-1B12-57A0-A259-476D2079EE83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBConfigResponseBody) SetRequestId(v string) *ModifyDBConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
