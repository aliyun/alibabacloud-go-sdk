// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFactoryId(v string) *GetDeviceListRequest
	GetFactoryId() *string
}

type GetDeviceListRequest struct {
	// The ID of the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// pn_95
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
}

func (s GetDeviceListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceListRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceListRequest) GetFactoryId() *string {
	return s.FactoryId
}

func (s *GetDeviceListRequest) SetFactoryId(v string) *GetDeviceListRequest {
	s.FactoryId = &v
	return s
}

func (s *GetDeviceListRequest) Validate() error {
	return dara.Validate(s)
}
