// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySceneDefensePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySceneDefensePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifySceneDefensePolicyResponseBody
	GetSuccess() *bool
}

type ModifySceneDefensePolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F65DF043-E0EB-4796-9467-23DDCDF92C1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
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

func (s ModifySceneDefensePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySceneDefensePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySceneDefensePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySceneDefensePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifySceneDefensePolicyResponseBody) SetRequestId(v string) *ModifySceneDefensePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySceneDefensePolicyResponseBody) SetSuccess(v bool) *ModifySceneDefensePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySceneDefensePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
