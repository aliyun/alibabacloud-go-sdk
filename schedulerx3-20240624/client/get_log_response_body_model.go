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
	SetData(v []*string) *GetLogResponseBody
	GetData() []*string
	SetMessage(v string) *GetLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLogResponseBody
	GetSuccess() *bool
}

type GetLogResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C78E2AD2-5985-515B-BAD2-31A248AFC263
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *GetLogResponseBody) GetData() []*string {
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

func (s *GetLogResponseBody) SetData(v []*string) *GetLogResponseBody {
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
	return dara.Validate(s)
}
