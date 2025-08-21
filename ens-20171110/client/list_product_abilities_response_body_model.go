// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductAbilitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductAbilities(v []*string) *ListProductAbilitiesResponseBody
	GetProductAbilities() []*string
	SetRequestId(v string) *ListProductAbilitiesResponseBody
	GetRequestId() *string
}

type ListProductAbilitiesResponseBody struct {
	// Products supported by the edge node.
	ProductAbilities []*string `json:"ProductAbilities,omitempty" xml:"ProductAbilities,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// xxxxx-75ED-422E-A022-7121FA18C968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProductAbilitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductAbilitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductAbilitiesResponseBody) GetProductAbilities() []*string {
	return s.ProductAbilities
}

func (s *ListProductAbilitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductAbilitiesResponseBody) SetProductAbilities(v []*string) *ListProductAbilitiesResponseBody {
	s.ProductAbilities = v
	return s
}

func (s *ListProductAbilitiesResponseBody) SetRequestId(v string) *ListProductAbilitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductAbilitiesResponseBody) Validate() error {
	return dara.Validate(s)
}
