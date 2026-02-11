// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRetcodeAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteRetcodeAppRequest
	GetAppId() *string
	SetRegionId(v string) *DeleteRetcodeAppRequest
	GetRegionId() *string
}

type DeleteRetcodeAppRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRetcodeAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRetcodeAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteRetcodeAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRetcodeAppRequest) SetAppId(v string) *DeleteRetcodeAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRetcodeAppRequest) SetRegionId(v string) *DeleteRetcodeAppRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRetcodeAppRequest) Validate() error {
	return dara.Validate(s)
}
