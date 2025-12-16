// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ABTestGroup) *UpdateABTestGroupRequest
	GetBody() *ABTestGroup
	SetDryRun(v bool) *UpdateABTestGroupRequest
	GetDryRun() *bool
}

type UpdateABTestGroupRequest struct {
	// The request body. For more information, see [ABTestGroup](https://help.aliyun.com/document_detail/178935.html).
	Body *ABTestGroup `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. No endpoint is created. The system checks whether your AccessKey is valid, whether Resource Access Management (RAM) users are authorized, and whether the required parameters are set.
	//
	// 	- false (default): creates an endpoint immediately.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateABTestGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateABTestGroupRequest) GetBody() *ABTestGroup {
	return s.Body
}

func (s *UpdateABTestGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateABTestGroupRequest) SetBody(v *ABTestGroup) *UpdateABTestGroupRequest {
	s.Body = v
	return s
}

func (s *UpdateABTestGroupRequest) SetDryRun(v bool) *UpdateABTestGroupRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateABTestGroupRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
