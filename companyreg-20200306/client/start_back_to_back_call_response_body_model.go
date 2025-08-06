// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBackToBackCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *StartBackToBackCallResponseBody
	GetData() *bool
	SetErrorCode(v string) *StartBackToBackCallResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *StartBackToBackCallResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *StartBackToBackCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartBackToBackCallResponseBody
	GetSuccess() *bool
}

type StartBackToBackCallResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 2174AA97-56FB-50FA-B243-0460B9E4CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartBackToBackCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartBackToBackCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartBackToBackCallResponseBody) GetData() *bool {
	return s.Data
}

func (s *StartBackToBackCallResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartBackToBackCallResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *StartBackToBackCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartBackToBackCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartBackToBackCallResponseBody) SetData(v bool) *StartBackToBackCallResponseBody {
	s.Data = &v
	return s
}

func (s *StartBackToBackCallResponseBody) SetErrorCode(v string) *StartBackToBackCallResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartBackToBackCallResponseBody) SetErrorMsg(v string) *StartBackToBackCallResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *StartBackToBackCallResponseBody) SetRequestId(v string) *StartBackToBackCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartBackToBackCallResponseBody) SetSuccess(v bool) *StartBackToBackCallResponseBody {
	s.Success = &v
	return s
}

func (s *StartBackToBackCallResponseBody) Validate() error {
	return dara.Validate(s)
}
