// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprovisionApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeprovisionApplicationRequest
	GetAppId() *string
}

type DeprovisionApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 452392483381546****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeprovisionApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeprovisionApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeprovisionApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeprovisionApplicationRequest) SetAppId(v string) *DeprovisionApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeprovisionApplicationRequest) Validate() error {
	return dara.Validate(s)
}
