// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPromClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UninstallPromClusterResponseBody
	GetCode() *int32
	SetData(v string) *UninstallPromClusterResponseBody
	GetData() *string
	SetMessage(v string) *UninstallPromClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *UninstallPromClusterResponseBody
	GetRequestId() *string
}

type UninstallPromClusterResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. You can use the ID to find logs and troubleshoot issues.
	//
	// example:
	//
	// 53980F48-DE82-53A1-9ADE-D2629226DD9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallPromClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallPromClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallPromClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UninstallPromClusterResponseBody) GetData() *string {
	return s.Data
}

func (s *UninstallPromClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallPromClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallPromClusterResponseBody) SetCode(v int32) *UninstallPromClusterResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallPromClusterResponseBody) SetData(v string) *UninstallPromClusterResponseBody {
	s.Data = &v
	return s
}

func (s *UninstallPromClusterResponseBody) SetMessage(v string) *UninstallPromClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallPromClusterResponseBody) SetRequestId(v string) *UninstallPromClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallPromClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
