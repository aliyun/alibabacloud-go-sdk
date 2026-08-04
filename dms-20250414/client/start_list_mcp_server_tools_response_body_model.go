// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartListMcpServerToolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StartListMcpServerToolsResponseBodyData) *StartListMcpServerToolsResponseBody
	GetData() *StartListMcpServerToolsResponseBodyData
	SetErrorCode(v string) *StartListMcpServerToolsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StartListMcpServerToolsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *StartListMcpServerToolsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartListMcpServerToolsResponseBody
	GetSuccess() *bool
}

type StartListMcpServerToolsResponseBody struct {
	// The result of the asynchronous detection startup. Only StartTimestamp may be returned if the resources are still being provisioned.
	Data *StartListMcpServerToolsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The return code. The value success is returned if the request was successful. An error code is returned if the request failed.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if a system-level request failure occurs.
	//
	// example:
	//
	// agent status=wait_resource_running
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate this call.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartListMcpServerToolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartListMcpServerToolsResponseBody) GoString() string {
	return s.String()
}

func (s *StartListMcpServerToolsResponseBody) GetData() *StartListMcpServerToolsResponseBodyData {
	return s.Data
}

func (s *StartListMcpServerToolsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartListMcpServerToolsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StartListMcpServerToolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartListMcpServerToolsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartListMcpServerToolsResponseBody) SetData(v *StartListMcpServerToolsResponseBodyData) *StartListMcpServerToolsResponseBody {
	s.Data = v
	return s
}

func (s *StartListMcpServerToolsResponseBody) SetErrorCode(v string) *StartListMcpServerToolsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartListMcpServerToolsResponseBody) SetErrorMessage(v string) *StartListMcpServerToolsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartListMcpServerToolsResponseBody) SetRequestId(v string) *StartListMcpServerToolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartListMcpServerToolsResponseBody) SetSuccess(v bool) *StartListMcpServerToolsResponseBody {
	s.Success = &v
	return s
}

func (s *StartListMcpServerToolsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartListMcpServerToolsResponseBodyData struct {
	// The temporary session ID for this detection. After the startup succeeds, use this value to call GetListMcpServerToolsResult to poll for the result.
	//
	// example:
	//
	// 1vw***6wr
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The UNIX timestamp at which the server started the detection, in milliseconds.
	//
	// example:
	//
	// 1785819600000
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s StartListMcpServerToolsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartListMcpServerToolsResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartListMcpServerToolsResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *StartListMcpServerToolsResponseBodyData) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *StartListMcpServerToolsResponseBodyData) SetSessionId(v string) *StartListMcpServerToolsResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *StartListMcpServerToolsResponseBodyData) SetStartTimestamp(v string) *StartListMcpServerToolsResponseBodyData {
	s.StartTimestamp = &v
	return s
}

func (s *StartListMcpServerToolsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
