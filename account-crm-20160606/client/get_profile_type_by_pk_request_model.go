// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProfileTypeByPkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetProfileTypeByPkRequest
	GetAppName() *string
	SetPk(v string) *GetProfileTypeByPkRequest
	GetPk() *string
}

type GetProfileTypeByPkRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s GetProfileTypeByPkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProfileTypeByPkRequest) GoString() string {
	return s.String()
}

func (s *GetProfileTypeByPkRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetProfileTypeByPkRequest) GetPk() *string {
	return s.Pk
}

func (s *GetProfileTypeByPkRequest) SetAppName(v string) *GetProfileTypeByPkRequest {
	s.AppName = &v
	return s
}

func (s *GetProfileTypeByPkRequest) SetPk(v string) *GetProfileTypeByPkRequest {
	s.Pk = &v
	return s
}

func (s *GetProfileTypeByPkRequest) Validate() error {
	return dara.Validate(s)
}
