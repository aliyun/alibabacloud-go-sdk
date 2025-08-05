// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdmEcsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteUdmEcsInstanceResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteUdmEcsInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteUdmEcsInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteUdmEcsInstanceResponseBody
	GetSuccess() *bool
}

type DeleteUdmEcsInstanceResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0497C0D3-82B5-56B2-8D64-D62E61B90E95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUdmEcsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdmEcsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUdmEcsInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteUdmEcsInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteUdmEcsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUdmEcsInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUdmEcsInstanceResponseBody) SetCode(v string) *DeleteUdmEcsInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUdmEcsInstanceResponseBody) SetMessage(v string) *DeleteUdmEcsInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUdmEcsInstanceResponseBody) SetRequestId(v string) *DeleteUdmEcsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUdmEcsInstanceResponseBody) SetSuccess(v bool) *DeleteUdmEcsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUdmEcsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
