// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountOneKeyDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AccountOneKeyDeleteRequest
	GetAppName() *string
	SetPk(v string) *AccountOneKeyDeleteRequest
	GetPk() *string
}

type AccountOneKeyDeleteRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s AccountOneKeyDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s AccountOneKeyDeleteRequest) GoString() string {
	return s.String()
}

func (s *AccountOneKeyDeleteRequest) GetAppName() *string {
	return s.AppName
}

func (s *AccountOneKeyDeleteRequest) GetPk() *string {
	return s.Pk
}

func (s *AccountOneKeyDeleteRequest) SetAppName(v string) *AccountOneKeyDeleteRequest {
	s.AppName = &v
	return s
}

func (s *AccountOneKeyDeleteRequest) SetPk(v string) *AccountOneKeyDeleteRequest {
	s.Pk = &v
	return s
}

func (s *AccountOneKeyDeleteRequest) Validate() error {
	return dara.Validate(s)
}
