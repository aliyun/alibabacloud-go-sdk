// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProhibitedPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProhibitedPoliciesResponseBody
	GetRequestId() *string
}

type DeleteProhibitedPoliciesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0325E55E-BD76-5856-894F-65AEEF01E84B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProhibitedPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProhibitedPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProhibitedPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProhibitedPoliciesResponseBody) SetRequestId(v string) *DeleteProhibitedPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProhibitedPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}
