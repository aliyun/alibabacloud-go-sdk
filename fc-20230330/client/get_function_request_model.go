// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *GetFunctionRequest
	GetQualifier() *string
}

type GetFunctionRequest struct {
	// 2023-03-10T10:10:10Z
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *GetFunctionRequest) SetQualifier(v string) *GetFunctionRequest {
	s.Qualifier = &v
	return s
}

func (s *GetFunctionRequest) Validate() error {
	return dara.Validate(s)
}
