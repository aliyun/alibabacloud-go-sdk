// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ActiveFlowLogResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ActiveFlowLogResponseBody
	GetSuccess() *string
}

type ActiveFlowLogResponseBody struct {
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

func (s ActiveFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActiveFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActiveFlowLogResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ActiveFlowLogResponseBody) SetRequestId(v string) *ActiveFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveFlowLogResponseBody) SetSuccess(v string) *ActiveFlowLogResponseBody {
	s.Success = &v
	return s
}

func (s *ActiveFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
