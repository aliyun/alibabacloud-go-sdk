// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UninstallPluginResponseBody
	GetCode() *string
	SetMessage(v string) *UninstallPluginResponseBody
	GetMessage() *string
	SetRequestId(v string) *UninstallPluginResponseBody
	GetRequestId() *string
}

type UninstallPluginResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F61D96E8-4E6D-5896-86E7-F1202AC31280
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UninstallPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallPluginResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallPluginResponseBody) GetCode() *string {
	return s.Code
}

func (s *UninstallPluginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallPluginResponseBody) SetCode(v string) *UninstallPluginResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallPluginResponseBody) SetMessage(v string) *UninstallPluginResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallPluginResponseBody) SetRequestId(v string) *UninstallPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
