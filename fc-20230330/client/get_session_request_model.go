// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *GetSessionRequest
	GetQualifier() *string
}

type GetSessionRequest struct {
	// example:
	//
	// aliasName1
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSessionRequest) GoString() string {
	return s.String()
}

func (s *GetSessionRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *GetSessionRequest) SetQualifier(v string) *GetSessionRequest {
	s.Qualifier = &v
	return s
}

func (s *GetSessionRequest) Validate() error {
	return dara.Validate(s)
}
