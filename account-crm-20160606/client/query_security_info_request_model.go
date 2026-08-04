// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecurityInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *QuerySecurityInfoRequest
	GetAppName() *string
	SetPk(v string) *QuerySecurityInfoRequest
	GetPk() *string
}

type QuerySecurityInfoRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s QuerySecurityInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySecurityInfoRequest) GoString() string {
	return s.String()
}

func (s *QuerySecurityInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *QuerySecurityInfoRequest) GetPk() *string {
	return s.Pk
}

func (s *QuerySecurityInfoRequest) SetAppName(v string) *QuerySecurityInfoRequest {
	s.AppName = &v
	return s
}

func (s *QuerySecurityInfoRequest) SetPk(v string) *QuerySecurityInfoRequest {
	s.Pk = &v
	return s
}

func (s *QuerySecurityInfoRequest) Validate() error {
	return dara.Validate(s)
}
