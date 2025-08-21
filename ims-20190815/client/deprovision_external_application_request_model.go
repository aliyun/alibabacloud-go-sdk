// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprovisionExternalApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeprovisionExternalApplicationRequest
	GetAppId() *string
}

type DeprovisionExternalApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 417273362044613****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeprovisionExternalApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeprovisionExternalApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeprovisionExternalApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeprovisionExternalApplicationRequest) SetAppId(v string) *DeprovisionExternalApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeprovisionExternalApplicationRequest) Validate() error {
	return dara.Validate(s)
}
