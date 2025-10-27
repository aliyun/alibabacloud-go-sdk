// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetLogResponseBody
	GetCode() *int32
	SetData(v *GetLogResponseBodyData) *GetLogResponseBody
	GetData() *GetLogResponseBodyData
	SetMessage(v string) *GetLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLogResponseBody
	GetSuccess() *bool
}

type GetLogResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// jobid=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
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

func (s GetLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetLogResponseBody) GetData() *GetLogResponseBodyData {
	return s.Data
}

func (s *GetLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLogResponseBody) SetCode(v int32) *GetLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetLogResponseBody) SetData(v *GetLogResponseBodyData) *GetLogResponseBody {
	s.Data = v
	return s
}

func (s *GetLogResponseBody) SetMessage(v string) *GetLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetLogResponseBody) SetRequestId(v string) *GetLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogResponseBody) SetSuccess(v bool) *GetLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLogResponseBodyData struct {
	// The logs. The value is an array of strings.
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
}

func (s GetLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLogResponseBodyData) GetLogs() []*string {
	return s.Logs
}

func (s *GetLogResponseBodyData) SetLogs(v []*string) *GetLogResponseBodyData {
	s.Logs = v
	return s
}

func (s *GetLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
