// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *GetFunctionCodeRequest
	GetQualifier() *string
}

type GetFunctionCodeRequest struct {
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetFunctionCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionCodeRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionCodeRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *GetFunctionCodeRequest) SetQualifier(v string) *GetFunctionCodeRequest {
	s.Qualifier = &v
	return s
}

func (s *GetFunctionCodeRequest) Validate() error {
	return dara.Validate(s)
}
