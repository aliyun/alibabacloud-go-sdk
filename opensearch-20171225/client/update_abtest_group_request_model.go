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
	// Specifies whether to perform a dry run. Valid values:
	//
	// - true: Performs a dry run. The system checks if the AccessKey is valid, if the RAM user is authorized, and if all required parameters are specified. The test group is not updated.
	//
	// - false (default): Sends the request to update the test group.
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
