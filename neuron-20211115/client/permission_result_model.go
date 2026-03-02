// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermissionResult interface {
	dara.Model
	String() string
	GoString() string
	SetAllow(v bool) *PermissionResult
	GetAllow() *bool
	SetRequestId(v string) *PermissionResult
	GetRequestId() *string
}

type PermissionResult struct {
	Allow     *bool   `json:"allow,omitempty" xml:"allow,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PermissionResult) String() string {
	return dara.Prettify(s)
}

func (s PermissionResult) GoString() string {
	return s.String()
}

func (s *PermissionResult) GetAllow() *bool {
	return s.Allow
}

func (s *PermissionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PermissionResult) SetAllow(v bool) *PermissionResult {
	s.Allow = &v
	return s
}

func (s *PermissionResult) SetRequestId(v string) *PermissionResult {
	s.RequestId = &v
	return s
}

func (s *PermissionResult) Validate() error {
	return dara.Validate(s)
}
