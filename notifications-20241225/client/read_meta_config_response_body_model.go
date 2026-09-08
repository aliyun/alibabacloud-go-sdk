// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMetaConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadMetaConfigResponseBody
	GetCode() *string
	SetData(v interface{}) *ReadMetaConfigResponseBody
	GetData() interface{}
	SetMessage(v string) *ReadMetaConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadMetaConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadMetaConfigResponseBody
	GetSuccess() *bool
}

type ReadMetaConfigResponseBody struct {
	// The error code returned if the call fails. For more information, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result.
	//
	// example:
	//
	// {
	//
	//     "isSubNextUser": "YES",
	//
	//     "IS_INNER_USER": "NO",
	//
	//     "isXsWhite": "NO",
	//
	//     "isXs": "NO",
	//
	//     "SITEID": "cn",
	//
	//     "webhookSignatures": [
	//
	//         "Alibaba",
	//
	//         "阿里云"
	//
	//     ]
	//
	// }
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned if the call fails.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is successful. A value of false indicates that the call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMetaConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadMetaConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMetaConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadMetaConfigResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ReadMetaConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadMetaConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadMetaConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadMetaConfigResponseBody) SetCode(v string) *ReadMetaConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMetaConfigResponseBody) SetData(v interface{}) *ReadMetaConfigResponseBody {
	s.Data = v
	return s
}

func (s *ReadMetaConfigResponseBody) SetMessage(v string) *ReadMetaConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMetaConfigResponseBody) SetRequestId(v string) *ReadMetaConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMetaConfigResponseBody) SetSuccess(v bool) *ReadMetaConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ReadMetaConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
