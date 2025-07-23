// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigAuditLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigAuditLogResponseBody
	GetRequestId() *string
}

type ConfigAuditLogResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 42B118FB-16A6-56FB-B877-D58637EEC6AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigAuditLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigAuditLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigAuditLogResponseBody) SetRequestId(v string) *ConfigAuditLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigAuditLogResponseBody) Validate() error {
	return dara.Validate(s)
}
