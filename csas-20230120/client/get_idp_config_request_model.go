// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdpConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetIdpConfigRequest
	GetId() *string
}

type GetIdpConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1465
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetIdpConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdpConfigRequest) GoString() string {
	return s.String()
}

func (s *GetIdpConfigRequest) GetId() *string {
	return s.Id
}

func (s *GetIdpConfigRequest) SetId(v string) *GetIdpConfigRequest {
	s.Id = &v
	return s
}

func (s *GetIdpConfigRequest) Validate() error {
	return dara.Validate(s)
}
