// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppServiceGroup interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AppServiceGroup
	GetName() *string
	SetQrCode(v string) *AppServiceGroup
	GetQrCode() *string
	SetType(v string) *AppServiceGroup
	GetType() *string
	SetUrl(v string) *AppServiceGroup
	GetUrl() *string
}

type AppServiceGroup struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	QrCode *string `json:"QrCode,omitempty" xml:"QrCode,omitempty"`
	// 例如：dingtalk、wx 等
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s AppServiceGroup) String() string {
	return dara.Prettify(s)
}

func (s AppServiceGroup) GoString() string {
	return s.String()
}

func (s *AppServiceGroup) GetName() *string {
	return s.Name
}

func (s *AppServiceGroup) GetQrCode() *string {
	return s.QrCode
}

func (s *AppServiceGroup) GetType() *string {
	return s.Type
}

func (s *AppServiceGroup) GetUrl() *string {
	return s.Url
}

func (s *AppServiceGroup) SetName(v string) *AppServiceGroup {
	s.Name = &v
	return s
}

func (s *AppServiceGroup) SetQrCode(v string) *AppServiceGroup {
	s.QrCode = &v
	return s
}

func (s *AppServiceGroup) SetType(v string) *AppServiceGroup {
	s.Type = &v
	return s
}

func (s *AppServiceGroup) SetUrl(v string) *AppServiceGroup {
	s.Url = &v
	return s
}

func (s *AppServiceGroup) Validate() error {
	return dara.Validate(s)
}
