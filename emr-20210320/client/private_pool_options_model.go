// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrivatePoolOptions interface {
	dara.Model
	String() string
	GoString() string
	SetMatchCriteria(v string) *PrivatePoolOptions
	GetMatchCriteria() *string
	SetPrivatePoolIds(v []*string) *PrivatePoolOptions
	GetPrivatePoolIds() []*string
}

type PrivatePoolOptions struct {
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
	// example:
	//
	// eap-bp67acfmxazb4****
	PrivatePoolIds []*string `json:"PrivatePoolIds,omitempty" xml:"PrivatePoolIds,omitempty" type:"Repeated"`
}

func (s PrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s PrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *PrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *PrivatePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *PrivatePoolOptions) SetMatchCriteria(v string) *PrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *PrivatePoolOptions) SetPrivatePoolIds(v []*string) *PrivatePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *PrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
