// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupClientResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteBackupClientResourceResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteBackupClientResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBackupClientResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBackupClientResourceResponseBody
	GetSuccess() *bool
}

type DeleteBackupClientResourceResponseBody struct {
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request is successful, \\`successful\\` is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBackupClientResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupClientResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBackupClientResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBackupClientResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupClientResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBackupClientResourceResponseBody) SetCode(v string) *DeleteBackupClientResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBackupClientResourceResponseBody) SetMessage(v string) *DeleteBackupClientResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBackupClientResourceResponseBody) SetRequestId(v string) *DeleteBackupClientResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupClientResourceResponseBody) SetSuccess(v bool) *DeleteBackupClientResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBackupClientResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
