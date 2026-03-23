// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickStaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *KickStaRequest
	GetAppCode() *string
	SetAppName(v string) *KickStaRequest
	GetAppName() *string
	SetStaMac(v string) *KickStaRequest
	GetStaMac() *string
}

type KickStaRequest struct {
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	StaMac *string `json:"StaMac,omitempty" xml:"StaMac,omitempty"`
}

func (s KickStaRequest) String() string {
	return dara.Prettify(s)
}

func (s KickStaRequest) GoString() string {
	return s.String()
}

func (s *KickStaRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *KickStaRequest) GetAppName() *string {
	return s.AppName
}

func (s *KickStaRequest) GetStaMac() *string {
	return s.StaMac
}

func (s *KickStaRequest) SetAppCode(v string) *KickStaRequest {
	s.AppCode = &v
	return s
}

func (s *KickStaRequest) SetAppName(v string) *KickStaRequest {
	s.AppName = &v
	return s
}

func (s *KickStaRequest) SetStaMac(v string) *KickStaRequest {
	s.StaMac = &v
	return s
}

func (s *KickStaRequest) Validate() error {
	return dara.Validate(s)
}
