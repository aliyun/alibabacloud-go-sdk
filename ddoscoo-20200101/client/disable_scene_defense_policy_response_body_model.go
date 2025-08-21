// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSceneDefensePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableSceneDefensePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableSceneDefensePolicyResponseBody
	GetSuccess() *bool
}

type DisableSceneDefensePolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F65DF043-E0EB-4796-9467-23DDCDF92C1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSceneDefensePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSceneDefensePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSceneDefensePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSceneDefensePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableSceneDefensePolicyResponseBody) SetRequestId(v string) *DisableSceneDefensePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSceneDefensePolicyResponseBody) SetSuccess(v bool) *DisableSceneDefensePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DisableSceneDefensePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
