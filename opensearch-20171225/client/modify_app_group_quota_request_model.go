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
	Body *Quota `json:"body,omitempty" xml:"body,omitempty"`
	// A client token that is used to ensure the idempotence of the request. The client generates this value to make sure that it is unique among different requests. The value can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 74db41d8cd3c784209093aa76afbe89e
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Specifies whether to perform a dry run. Default value: false.
	//
	// Valid values:
	//
	// - **true**: Validates the request parameters without creating the attribution configuration.
	//
	// - **false**: Validates the request parameters and creates the attribution configuration.
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
