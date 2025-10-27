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
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the status of SQL audit is updated. Valid values:
	//
	// 	- **true**: The status of SQL audit is updated.
	//
	// 	- **false**: The status of SQL audit is not updated.
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
