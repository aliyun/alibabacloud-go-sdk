// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAccessWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterAccessWhitelistResponseBody
	GetRequestId() *string
}

type ModifyDBClusterAccessWhitelistResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterAccessWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAccessWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterAccessWhitelistResponseBody) SetRequestId(v string) *ModifyDBClusterAccessWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
