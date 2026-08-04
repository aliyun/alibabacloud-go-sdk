// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOneKeyDeleteBlockListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *QueryOneKeyDeleteBlockListRequest
	GetAppName() *string
	SetPk(v string) *QueryOneKeyDeleteBlockListRequest
	GetPk() *string
}

type QueryOneKeyDeleteBlockListRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s QueryOneKeyDeleteBlockListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOneKeyDeleteBlockListRequest) GoString() string {
	return s.String()
}

func (s *QueryOneKeyDeleteBlockListRequest) GetAppName() *string {
	return s.AppName
}

func (s *QueryOneKeyDeleteBlockListRequest) GetPk() *string {
	return s.Pk
}

func (s *QueryOneKeyDeleteBlockListRequest) SetAppName(v string) *QueryOneKeyDeleteBlockListRequest {
	s.AppName = &v
	return s
}

func (s *QueryOneKeyDeleteBlockListRequest) SetPk(v string) *QueryOneKeyDeleteBlockListRequest {
	s.Pk = &v
	return s
}

func (s *QueryOneKeyDeleteBlockListRequest) Validate() error {
	return dara.Validate(s)
}
