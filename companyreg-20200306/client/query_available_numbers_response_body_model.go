// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *QueryAvailableNumbersResponseBody
	GetData() []*string
	SetErrorCode(v string) *QueryAvailableNumbersResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryAvailableNumbersResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryAvailableNumbersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAvailableNumbersResponseBody
	GetSuccess() *bool
}

type QueryAvailableNumbersResponseBody struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// PARTNER.CONFIG.NOT.FOUND
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 6A603AA0-73BA-52B3-AC7D-0F846ECF7A9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAvailableNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvailableNumbersResponseBody) GetData() []*string {
	return s.Data
}

func (s *QueryAvailableNumbersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryAvailableNumbersResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryAvailableNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAvailableNumbersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAvailableNumbersResponseBody) SetData(v []*string) *QueryAvailableNumbersResponseBody {
	s.Data = v
	return s
}

func (s *QueryAvailableNumbersResponseBody) SetErrorCode(v string) *QueryAvailableNumbersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryAvailableNumbersResponseBody) SetErrorMsg(v string) *QueryAvailableNumbersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryAvailableNumbersResponseBody) SetRequestId(v string) *QueryAvailableNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvailableNumbersResponseBody) SetSuccess(v bool) *QueryAvailableNumbersResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAvailableNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}
