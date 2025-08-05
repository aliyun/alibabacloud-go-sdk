// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdmDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteUdmDiskResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteUdmDiskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteUdmDiskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteUdmDiskResponseBody
	GetSuccess() *bool
}

type DeleteUdmDiskResponseBody struct {
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
	// Id of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
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

func (s DeleteUdmDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdmDiskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUdmDiskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteUdmDiskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteUdmDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUdmDiskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUdmDiskResponseBody) SetCode(v string) *DeleteUdmDiskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUdmDiskResponseBody) SetMessage(v string) *DeleteUdmDiskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUdmDiskResponseBody) SetRequestId(v string) *DeleteUdmDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUdmDiskResponseBody) SetSuccess(v bool) *DeleteUdmDiskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUdmDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
