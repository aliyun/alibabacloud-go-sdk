// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *DeleteSessionRequest
	GetQualifier() *string
}

type DeleteSessionRequest struct {
	// The function alias or version associated with the session to be deleted.
	//
	// example:
	//
	// aliasName1
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s DeleteSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSessionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSessionRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *DeleteSessionRequest) SetQualifier(v string) *DeleteSessionRequest {
	s.Qualifier = &v
	return s
}

func (s *DeleteSessionRequest) Validate() error {
	return dara.Validate(s)
}
