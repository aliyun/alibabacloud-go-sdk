// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCronExecTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *GetCronExecTimeResponseBody
	GetData() []*string
	SetErrCode(v string) *GetCronExecTimeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetCronExecTimeResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetCronExecTimeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCronExecTimeResponseBody
	GetSuccess() *bool
}

type GetCronExecTimeResponseBody struct {
	// The data list returned by the operation. For the structure of each element, see the child field descriptions.
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID. Use this ID to locate and troubleshoot issues related to this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values: true and false. If false is returned, use errCode and errMessage to troubleshoot.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetCronExecTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCronExecTimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetCronExecTimeResponseBody) GetData() []*string {
	return s.Data
}

func (s *GetCronExecTimeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetCronExecTimeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetCronExecTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCronExecTimeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCronExecTimeResponseBody) SetData(v []*string) *GetCronExecTimeResponseBody {
	s.Data = v
	return s
}

func (s *GetCronExecTimeResponseBody) SetErrCode(v string) *GetCronExecTimeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetCronExecTimeResponseBody) SetErrMessage(v string) *GetCronExecTimeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetCronExecTimeResponseBody) SetRequestId(v string) *GetCronExecTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCronExecTimeResponseBody) SetSuccess(v bool) *GetCronExecTimeResponseBody {
	s.Success = &v
	return s
}

func (s *GetCronExecTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
