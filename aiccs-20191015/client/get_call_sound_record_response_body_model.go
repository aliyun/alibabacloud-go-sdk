// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallSoundRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCallSoundRecordResponseBody
	GetCode() *string
	SetData(v string) *GetCallSoundRecordResponseBody
	GetData() *string
	SetMessage(v string) *GetCallSoundRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCallSoundRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCallSoundRecordResponseBody
	GetSuccess() *bool
}

type GetCallSoundRecordResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// http://aliccrec-shvpc.oss-cn-shanghai.aliyuncs.com/accrec_tmp/1001067****.wav
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCallSoundRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCallSoundRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallSoundRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCallSoundRecordResponseBody) GetData() *string {
	return s.Data
}

func (s *GetCallSoundRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCallSoundRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCallSoundRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCallSoundRecordResponseBody) SetCode(v string) *GetCallSoundRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallSoundRecordResponseBody) SetData(v string) *GetCallSoundRecordResponseBody {
	s.Data = &v
	return s
}

func (s *GetCallSoundRecordResponseBody) SetMessage(v string) *GetCallSoundRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallSoundRecordResponseBody) SetRequestId(v string) *GetCallSoundRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCallSoundRecordResponseBody) SetSuccess(v bool) *GetCallSoundRecordResponseBody {
	s.Success = &v
	return s
}

func (s *GetCallSoundRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
