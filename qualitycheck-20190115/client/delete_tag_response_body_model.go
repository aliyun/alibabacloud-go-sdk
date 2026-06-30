// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTagResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTagResponseBody
	GetSuccess() *bool
}

type DeleteTagResponseBody struct {
	// The response code. **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message, if any.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. true: The call was successful. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTagResponseBody) SetCode(v string) *DeleteTagResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTagResponseBody) SetMessage(v string) *DeleteTagResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTagResponseBody) SetRequestId(v string) *DeleteTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTagResponseBody) SetSuccess(v bool) *DeleteTagResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTagResponseBody) Validate() error {
	return dara.Validate(s)
}
