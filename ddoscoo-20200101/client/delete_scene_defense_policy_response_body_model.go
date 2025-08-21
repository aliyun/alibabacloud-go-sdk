// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSceneDefensePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSceneDefensePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSceneDefensePolicyResponseBody
	GetSuccess() *bool
}

type DeleteSceneDefensePolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F65DF043-E0EB-4796-9467-23DDCDF92C1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSceneDefensePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSceneDefensePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSceneDefensePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSceneDefensePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSceneDefensePolicyResponseBody) SetRequestId(v string) *DeleteSceneDefensePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSceneDefensePolicyResponseBody) SetSuccess(v bool) *DeleteSceneDefensePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSceneDefensePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
