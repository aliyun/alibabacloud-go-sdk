// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradeApplicationApmServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DowngradeApplicationApmServiceRequest
	GetAppId() *string
}

type DowngradeApplicationApmServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DowngradeApplicationApmServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DowngradeApplicationApmServiceRequest) GoString() string {
	return s.String()
}

func (s *DowngradeApplicationApmServiceRequest) GetAppId() *string {
	return s.AppId
}

func (s *DowngradeApplicationApmServiceRequest) SetAppId(v string) *DowngradeApplicationApmServiceRequest {
	s.AppId = &v
	return s
}

func (s *DowngradeApplicationApmServiceRequest) Validate() error {
	return dara.Validate(s)
}
