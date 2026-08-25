// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTerminalCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTerminalCountResponseBody
	GetCode() *string
	SetData(v *GetTerminalCountResponseBodyData) *GetTerminalCountResponseBody
	GetData() *GetTerminalCountResponseBodyData
	SetHttpStatusCode(v int32) *GetTerminalCountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTerminalCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTerminalCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTerminalCountResponseBody
	GetSuccess() *bool
}

type GetTerminalCountResponseBody struct {
	// The status code. 200 is returned if the call is successful. An error code is returned if the call fails.
	//
	// example:
	//
	// PARAM_ERROR
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The terminal count statistics information.
	Data *GetTerminalCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. This parameter is empty if the call is successful.
	//
	// example:
	//
	// parameter error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C5DCE54A-B266-522E-A6ED-468AF45F5AAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTerminalCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTerminalCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetTerminalCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTerminalCountResponseBody) GetData() *GetTerminalCountResponseBodyData {
	return s.Data
}

func (s *GetTerminalCountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTerminalCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTerminalCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTerminalCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTerminalCountResponseBody) SetCode(v string) *GetTerminalCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetTerminalCountResponseBody) SetData(v *GetTerminalCountResponseBodyData) *GetTerminalCountResponseBody {
	s.Data = v
	return s
}

func (s *GetTerminalCountResponseBody) SetHttpStatusCode(v int32) *GetTerminalCountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTerminalCountResponseBody) SetMessage(v string) *GetTerminalCountResponseBody {
	s.Message = &v
	return s
}

func (s *GetTerminalCountResponseBody) SetRequestId(v string) *GetTerminalCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTerminalCountResponseBody) SetSuccess(v bool) *GetTerminalCountResponseBody {
	s.Success = &v
	return s
}

func (s *GetTerminalCountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTerminalCountResponseBodyData struct {
	// The number of hardware terminals that are bound to users. This parameter is returned only when ClientType is set to 1.
	//
	// example:
	//
	// 60
	BindUserCount *int64 `json:"BindUserCount,omitempty" xml:"BindUserCount,omitempty"`
	// The number of managed terminals.
	//
	// example:
	//
	// 80
	InManageCount *int64 `json:"InManageCount,omitempty" xml:"InManageCount,omitempty"`
	// The number of unmanaged terminals.
	//
	// example:
	//
	// 20
	NotInManageCount *int64 `json:"NotInManageCount,omitempty" xml:"NotInManageCount,omitempty"`
	// The total number of terminals.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTerminalCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTerminalCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTerminalCountResponseBodyData) GetBindUserCount() *int64 {
	return s.BindUserCount
}

func (s *GetTerminalCountResponseBodyData) GetInManageCount() *int64 {
	return s.InManageCount
}

func (s *GetTerminalCountResponseBodyData) GetNotInManageCount() *int64 {
	return s.NotInManageCount
}

func (s *GetTerminalCountResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTerminalCountResponseBodyData) SetBindUserCount(v int64) *GetTerminalCountResponseBodyData {
	s.BindUserCount = &v
	return s
}

func (s *GetTerminalCountResponseBodyData) SetInManageCount(v int64) *GetTerminalCountResponseBodyData {
	s.InManageCount = &v
	return s
}

func (s *GetTerminalCountResponseBodyData) SetNotInManageCount(v int64) *GetTerminalCountResponseBodyData {
	s.NotInManageCount = &v
	return s
}

func (s *GetTerminalCountResponseBodyData) SetTotalCount(v int64) *GetTerminalCountResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetTerminalCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
