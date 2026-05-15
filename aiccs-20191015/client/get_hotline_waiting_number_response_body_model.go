// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineWaitingNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotlineWaitingNumberResponseBody
	GetCode() *string
	SetData(v int64) *GetHotlineWaitingNumberResponseBody
	GetData() *int64
	SetMessage(v string) *GetHotlineWaitingNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotlineWaitingNumberResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetHotlineWaitingNumberResponseBody
	GetSuccess() *string
}

type GetHotlineWaitingNumberResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineWaitingNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineWaitingNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotlineWaitingNumberResponseBody) GetData() *int64 {
	return s.Data
}

func (s *GetHotlineWaitingNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotlineWaitingNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotlineWaitingNumberResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetHotlineWaitingNumberResponseBody) SetCode(v string) *GetHotlineWaitingNumberResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetData(v int64) *GetHotlineWaitingNumberResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetMessage(v string) *GetHotlineWaitingNumberResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetRequestId(v string) *GetHotlineWaitingNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetSuccess(v string) *GetHotlineWaitingNumberResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
