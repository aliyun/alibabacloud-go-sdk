// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewInstanceRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ViewInstanceRecordsResponseBody
	GetCode() *string
	SetMessage(v string) *ViewInstanceRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ViewInstanceRecordsResponseBody
	GetRequestId() *string
	SetRoot(v bool) *ViewInstanceRecordsResponseBody
	GetRoot() *bool
	SetSuccess(v bool) *ViewInstanceRecordsResponseBody
	GetSuccess() *bool
}

type ViewInstanceRecordsResponseBody struct {
	// example:
	//
	// InvalidParamter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Root *bool `json:"Root,omitempty" xml:"Root,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ViewInstanceRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ViewInstanceRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ViewInstanceRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ViewInstanceRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ViewInstanceRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ViewInstanceRecordsResponseBody) GetRoot() *bool {
	return s.Root
}

func (s *ViewInstanceRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ViewInstanceRecordsResponseBody) SetCode(v string) *ViewInstanceRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ViewInstanceRecordsResponseBody) SetMessage(v string) *ViewInstanceRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ViewInstanceRecordsResponseBody) SetRequestId(v string) *ViewInstanceRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ViewInstanceRecordsResponseBody) SetRoot(v bool) *ViewInstanceRecordsResponseBody {
	s.Root = &v
	return s
}

func (s *ViewInstanceRecordsResponseBody) SetSuccess(v bool) *ViewInstanceRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *ViewInstanceRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}
