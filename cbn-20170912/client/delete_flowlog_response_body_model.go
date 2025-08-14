// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowlogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFlowlogResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteFlowlogResponseBody
	GetSuccess() *string
}

type DeleteFlowlogResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFlowlogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowlogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowlogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowlogResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteFlowlogResponseBody) SetRequestId(v string) *DeleteFlowlogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowlogResponseBody) SetSuccess(v string) *DeleteFlowlogResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFlowlogResponseBody) Validate() error {
	return dara.Validate(s)
}
