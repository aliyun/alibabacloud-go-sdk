// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFlowLogResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteFlowLogResponseBody
	GetSuccess() *string
}

type DeleteFlowLogResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
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

func (s DeleteFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowLogResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteFlowLogResponseBody) SetRequestId(v string) *DeleteFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowLogResponseBody) SetSuccess(v string) *DeleteFlowLogResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
