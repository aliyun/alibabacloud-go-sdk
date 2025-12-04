// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAccessWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterAccessWhiteListResponseBody
	GetRequestId() *string
}

type ModifyDBClusterAccessWhiteListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterAccessWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *ModifyDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterAccessWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
