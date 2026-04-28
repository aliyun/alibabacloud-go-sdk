// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkInfo interface {
	dara.Model
	String() string
	GoString() string
	SetExtra(v string) *LinkInfo
	GetExtra() *string
	SetIdentity(v string) *LinkInfo
	GetIdentity() *string
	SetType(v string) *LinkInfo
	GetType() *string
}

type LinkInfo struct {
	Extra    *string `json:"extra,omitempty" xml:"extra,omitempty"`
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s LinkInfo) String() string {
	return dara.Prettify(s)
}

func (s LinkInfo) GoString() string {
	return s.String()
}

func (s *LinkInfo) GetExtra() *string {
	return s.Extra
}

func (s *LinkInfo) GetIdentity() *string {
	return s.Identity
}

func (s *LinkInfo) GetType() *string {
	return s.Type
}

func (s *LinkInfo) SetExtra(v string) *LinkInfo {
	s.Extra = &v
	return s
}

func (s *LinkInfo) SetIdentity(v string) *LinkInfo {
	s.Identity = &v
	return s
}

func (s *LinkInfo) SetType(v string) *LinkInfo {
	s.Type = &v
	return s
}

func (s *LinkInfo) Validate() error {
	return dara.Validate(s)
}
