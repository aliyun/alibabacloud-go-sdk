// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MoveResourceGroupResponseBody
	GetCode() *string
	SetMessage(v string) *MoveResourceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *MoveResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MoveResourceGroupResponseBody
	GetSuccess() *bool
}

type MoveResourceGroupResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FABF516-FED3-5697-BDA2-B18C5D9A****
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

func (s MoveResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *MoveResourceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MoveResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveResourceGroupResponseBody) SetCode(v string) *MoveResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *MoveResourceGroupResponseBody) SetMessage(v string) *MoveResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceGroupResponseBody) SetSuccess(v bool) *MoveResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *MoveResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
