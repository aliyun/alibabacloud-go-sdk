// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHibernateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *HibernateClusterRequest
	GetClientToken() *string
}

type HibernateClusterRequest struct {
	// example:
	//
	// c82e6987e2961451182edacd74faf
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s HibernateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s HibernateClusterRequest) GoString() string {
	return s.String()
}

func (s *HibernateClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *HibernateClusterRequest) SetClientToken(v string) *HibernateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *HibernateClusterRequest) Validate() error {
	return dara.Validate(s)
}
