// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePushReviewOnOffRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationId(v string) *UpdatePushReviewOnOffRequest
	GetOrganizationId() *string
	SetTrunkMode(v bool) *UpdatePushReviewOnOffRequest
	GetTrunkMode() *bool
}

type UpdatePushReviewOnOffRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	TrunkMode *bool `json:"trunkMode,omitempty" xml:"trunkMode,omitempty"`
}

func (s UpdatePushReviewOnOffRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePushReviewOnOffRequest) GoString() string {
	return s.String()
}

func (s *UpdatePushReviewOnOffRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdatePushReviewOnOffRequest) GetTrunkMode() *bool {
	return s.TrunkMode
}

func (s *UpdatePushReviewOnOffRequest) SetOrganizationId(v string) *UpdatePushReviewOnOffRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdatePushReviewOnOffRequest) SetTrunkMode(v bool) *UpdatePushReviewOnOffRequest {
	s.TrunkMode = &v
	return s
}

func (s *UpdatePushReviewOnOffRequest) Validate() error {
	return dara.Validate(s)
}
