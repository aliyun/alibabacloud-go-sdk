// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeApplicationApmServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpgradeApplicationApmServiceRequest
	GetAppId() *string
}

type UpgradeApplicationApmServiceRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s UpgradeApplicationApmServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeApplicationApmServiceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationApmServiceRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpgradeApplicationApmServiceRequest) SetAppId(v string) *UpgradeApplicationApmServiceRequest {
	s.AppId = &v
	return s
}

func (s *UpgradeApplicationApmServiceRequest) Validate() error {
	return dara.Validate(s)
}
