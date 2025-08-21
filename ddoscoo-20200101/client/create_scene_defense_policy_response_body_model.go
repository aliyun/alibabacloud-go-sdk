// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSceneDefensePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSceneDefensePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSceneDefensePolicyResponseBody
	GetSuccess() *bool
}

type CreateSceneDefensePolicyResponseBody struct {
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

func (s CreateSceneDefensePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSceneDefensePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSceneDefensePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSceneDefensePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSceneDefensePolicyResponseBody) SetRequestId(v string) *CreateSceneDefensePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSceneDefensePolicyResponseBody) SetSuccess(v bool) *CreateSceneDefensePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSceneDefensePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
