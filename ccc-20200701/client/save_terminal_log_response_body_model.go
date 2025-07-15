// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTerminalLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveTerminalLogResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *SaveTerminalLogResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *SaveTerminalLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveTerminalLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveTerminalLogResponseBody
	GetSuccess() *bool
	SetTimeStamp(v int64) *SaveTerminalLogResponseBody
	GetTimeStamp() *int64
}

type SaveTerminalLogResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1647309061000
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s SaveTerminalLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTerminalLogResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTerminalLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveTerminalLogResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *SaveTerminalLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveTerminalLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTerminalLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveTerminalLogResponseBody) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *SaveTerminalLogResponseBody) SetCode(v string) *SaveTerminalLogResponseBody {
	s.Code = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetHttpStatusCode(v int64) *SaveTerminalLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetMessage(v string) *SaveTerminalLogResponseBody {
	s.Message = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetRequestId(v string) *SaveTerminalLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetSuccess(v bool) *SaveTerminalLogResponseBody {
	s.Success = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetTimeStamp(v int64) *SaveTerminalLogResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *SaveTerminalLogResponseBody) Validate() error {
	return dara.Validate(s)
}
