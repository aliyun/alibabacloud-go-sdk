// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRoomGenieScenesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ImportRoomGenieScenesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportRoomGenieScenesResponseBody
	GetRequestId() *string
	SetResult(v bool) *ImportRoomGenieScenesResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *ImportRoomGenieScenesResponseBody
	GetStatusCode() *int32
}

type ImportRoomGenieScenesResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ImportRoomGenieScenesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportRoomGenieScenesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportRoomGenieScenesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ImportRoomGenieScenesResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportRoomGenieScenesResponseBody) SetMessage(v string) *ImportRoomGenieScenesResponseBody {
	s.Message = &v
	return s
}

func (s *ImportRoomGenieScenesResponseBody) SetRequestId(v string) *ImportRoomGenieScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportRoomGenieScenesResponseBody) SetResult(v bool) *ImportRoomGenieScenesResponseBody {
	s.Result = &v
	return s
}

func (s *ImportRoomGenieScenesResponseBody) SetStatusCode(v int32) *ImportRoomGenieScenesResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ImportRoomGenieScenesResponseBody) Validate() error {
	return dara.Validate(s)
}
