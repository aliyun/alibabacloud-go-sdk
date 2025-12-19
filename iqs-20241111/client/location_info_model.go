// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLocationInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *LocationInfo
	GetCity() *string
	SetIp(v string) *LocationInfo
	GetIp() *string
}

type LocationInfo struct {
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	Ip   *string `json:"ip,omitempty" xml:"ip,omitempty"`
}

func (s LocationInfo) String() string {
	return dara.Prettify(s)
}

func (s LocationInfo) GoString() string {
	return s.String()
}

func (s *LocationInfo) GetCity() *string {
	return s.City
}

func (s *LocationInfo) GetIp() *string {
	return s.Ip
}

func (s *LocationInfo) SetCity(v string) *LocationInfo {
	s.City = &v
	return s
}

func (s *LocationInfo) SetIp(v string) *LocationInfo {
	s.Ip = &v
	return s
}

func (s *LocationInfo) Validate() error {
	return dara.Validate(s)
}
