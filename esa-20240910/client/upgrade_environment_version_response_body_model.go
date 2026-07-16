// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeEnvironmentVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeEnvironmentVersionResponseBody
	GetRequestId() *string
}

type UpgradeEnvironmentVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B49E6DDA-F413-422B-B58E-2FA23F286726
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeEnvironmentVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEnvironmentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeEnvironmentVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeEnvironmentVersionResponseBody) SetRequestId(v string) *UpgradeEnvironmentVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeEnvironmentVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
