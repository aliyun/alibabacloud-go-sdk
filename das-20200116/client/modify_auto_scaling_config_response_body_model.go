// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoScalingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyAutoScalingConfigResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyAutoScalingConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAutoScalingConfigResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyAutoScalingConfigResponseBody
	GetSuccess() *string
}

type ModifyAutoScalingConfigResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAutoScalingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyAutoScalingConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAutoScalingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAutoScalingConfigResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyAutoScalingConfigResponseBody) SetCode(v string) *ModifyAutoScalingConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) SetMessage(v string) *ModifyAutoScalingConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) SetRequestId(v string) *ModifyAutoScalingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) SetSuccess(v string) *ModifyAutoScalingConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
