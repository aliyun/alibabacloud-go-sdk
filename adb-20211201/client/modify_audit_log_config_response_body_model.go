// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAuditLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAuditLogConfigResponseBody
	GetRequestId() *string
	SetUpdateSucceed(v bool) *ModifyAuditLogConfigResponseBody
	GetUpdateSucceed() *bool
}

type ModifyAuditLogConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CDC59E56-BD07-56CA-A05F-B7907DE5C862
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the status of SQL audit is updated. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	UpdateSucceed *bool `json:"UpdateSucceed,omitempty" xml:"UpdateSucceed,omitempty"`
}

func (s ModifyAuditLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAuditLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAuditLogConfigResponseBody) GetUpdateSucceed() *bool {
	return s.UpdateSucceed
}

func (s *ModifyAuditLogConfigResponseBody) SetRequestId(v string) *ModifyAuditLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAuditLogConfigResponseBody) SetUpdateSucceed(v bool) *ModifyAuditLogConfigResponseBody {
	s.UpdateSucceed = &v
	return s
}

func (s *ModifyAuditLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
