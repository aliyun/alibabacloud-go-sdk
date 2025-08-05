// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteBackupClientResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteBackupClientResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBackupClientResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBackupClientResponseBody
	GetSuccess() *bool
}

type DeleteBackupClientResponseBody struct {
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
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

func (s DeleteBackupClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupClientResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBackupClientResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBackupClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupClientResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBackupClientResponseBody) SetCode(v string) *DeleteBackupClientResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBackupClientResponseBody) SetMessage(v string) *DeleteBackupClientResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBackupClientResponseBody) SetRequestId(v string) *DeleteBackupClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupClientResponseBody) SetSuccess(v bool) *DeleteBackupClientResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBackupClientResponseBody) Validate() error {
	return dara.Validate(s)
}
