// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableProcessDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableProcessDefinitionRequest
	GetClientToken() *string
	SetId(v string) *DisableProcessDefinitionRequest
	GetId() *string
}

type DisableProcessDefinitionRequest struct {
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 11792
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DisableProcessDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableProcessDefinitionRequest) GoString() string {
	return s.String()
}

func (s *DisableProcessDefinitionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableProcessDefinitionRequest) GetId() *string {
	return s.Id
}

func (s *DisableProcessDefinitionRequest) SetClientToken(v string) *DisableProcessDefinitionRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableProcessDefinitionRequest) SetId(v string) *DisableProcessDefinitionRequest {
	s.Id = &v
	return s
}

func (s *DisableProcessDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
