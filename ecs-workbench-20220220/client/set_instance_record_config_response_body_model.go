// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceRecordConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetInstanceRecordConfigResponseBody
	GetCode() *string
	SetMessage(v string) *SetInstanceRecordConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetInstanceRecordConfigResponseBody
	GetRequestId() *string
	SetRoot(v bool) *SetInstanceRecordConfigResponseBody
	GetRoot() *bool
	SetSuccess(v bool) *SetInstanceRecordConfigResponseBody
	GetSuccess() *bool
}

type SetInstanceRecordConfigResponseBody struct {
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

func (s SetInstanceRecordConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceRecordConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstanceRecordConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetInstanceRecordConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetInstanceRecordConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetInstanceRecordConfigResponseBody) GetRoot() *bool {
	return s.Root
}

func (s *SetInstanceRecordConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetInstanceRecordConfigResponseBody) SetCode(v string) *SetInstanceRecordConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SetInstanceRecordConfigResponseBody) SetMessage(v string) *SetInstanceRecordConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SetInstanceRecordConfigResponseBody) SetRequestId(v string) *SetInstanceRecordConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetInstanceRecordConfigResponseBody) SetRoot(v bool) *SetInstanceRecordConfigResponseBody {
	s.Root = &v
	return s
}

func (s *SetInstanceRecordConfigResponseBody) SetSuccess(v bool) *SetInstanceRecordConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SetInstanceRecordConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
