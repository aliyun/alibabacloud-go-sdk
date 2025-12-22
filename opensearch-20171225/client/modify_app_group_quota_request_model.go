// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppGroupQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Quota) *ModifyAppGroupQuotaRequest
	GetBody() *Quota
	SetClientToken(v string) *ModifyAppGroupQuotaRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyAppGroupQuotaRequest
	GetDryRun() *bool
}

type ModifyAppGroupQuotaRequest struct {
	// The request body.
	Body        *Quota  `json:"body,omitempty" xml:"body,omitempty"`
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
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
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyAppGroupQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppGroupQuotaRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaRequest) GetBody() *Quota {
	return s.Body
}

func (s *ModifyAppGroupQuotaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppGroupQuotaRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyAppGroupQuotaRequest) SetBody(v *Quota) *ModifyAppGroupQuotaRequest {
	s.Body = v
	return s
}

func (s *ModifyAppGroupQuotaRequest) SetClientToken(v string) *ModifyAppGroupQuotaRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppGroupQuotaRequest) SetDryRun(v bool) *ModifyAppGroupQuotaRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyAppGroupQuotaRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
