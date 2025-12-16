// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABTestGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ABTestGroup) *CreateABTestGroupRequest
	GetBody() *ABTestGroup
	SetDryRun(v bool) *CreateABTestGroupRequest
	GetDryRun() *bool
}

type CreateABTestGroupRequest struct {
	// The request body. For more information, see [ABTestGroup](https://help.aliyun.com/document_detail/178935.html).
	Body *ABTestGroup `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to check the validity of input parameters. Default value: false.
	//
	// Valid values:
	//
	// 	- **true**: checks only the validity of input parameters.
	//
	// 	- **false**: checks the validity of input parameters and creates an attribution configuration.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateABTestGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateABTestGroupRequest) GetBody() *ABTestGroup {
	return s.Body
}

func (s *CreateABTestGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateABTestGroupRequest) SetBody(v *ABTestGroup) *CreateABTestGroupRequest {
	s.Body = v
	return s
}

func (s *CreateABTestGroupRequest) SetDryRun(v bool) *CreateABTestGroupRequest {
	s.DryRun = &v
	return s
}

func (s *CreateABTestGroupRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
