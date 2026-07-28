// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeApplicationVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpgradeApplicationVersionRequest
	GetApplicationId() *string
	SetUpgradePolicy(v string) *UpgradeApplicationVersionRequest
	GetUpgradePolicy() *string
}

type UpgradeApplicationVersionRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The upgrade policy.
	//
	// example:
	//
	// HOT
	UpgradePolicy *string `json:"UpgradePolicy,omitempty" xml:"UpgradePolicy,omitempty"`
}

func (s UpgradeApplicationVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeApplicationVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationVersionRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpgradeApplicationVersionRequest) GetUpgradePolicy() *string {
	return s.UpgradePolicy
}

func (s *UpgradeApplicationVersionRequest) SetApplicationId(v string) *UpgradeApplicationVersionRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpgradeApplicationVersionRequest) SetUpgradePolicy(v string) *UpgradeApplicationVersionRequest {
	s.UpgradePolicy = &v
	return s
}

func (s *UpgradeApplicationVersionRequest) Validate() error {
	return dara.Validate(s)
}
