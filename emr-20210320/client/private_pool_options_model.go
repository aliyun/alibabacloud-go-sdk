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
	// The type of private pool that you want to use to start ECS instances. A private pool is generated when an elasticity assurance or a capacity reservation takes effect. You can select a private pool for starting ECS instances. Valid values: Open, Target, and None. If you set the parameter to Open, the system selects an open private pool to start instances. If no matching open private pools exist, the resources in the public pool are used. In this case, you do not need to specify the ID of the private pool. If you set the parameter to Target, the resources in the specified private pool are used to start ECS instances. If the specified private pool does not exist, ECS instances cannot be started. You must specify the ID of the private pool when you set the parameter to Target. If you set the parameter to None, the resources in private pools are not used to start instances. Default value: None.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
	// The IDs of the private pools.
	//
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
