// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLumaTableResponseBody
	GetCode() *string
	SetData(v *LumaTable) *GetLumaTableResponseBody
	GetData() *LumaTable
	SetMessage(v string) *GetLumaTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLumaTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLumaTableResponseBody
	GetSuccess() *bool
}

type GetLumaTableResponseBody struct {
	// The response code. A value of Success indicates that the call succeeds. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the event table bound to the Agent, including column definitions.
	Data *LumaTable `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned by the operation. The value is Operation success if the call succeeds, or a specific error description if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. You can use this ID for troubleshooting and when you submit a ticket.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call succeeds.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLumaTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLumaTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetLumaTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLumaTableResponseBody) GetData() *LumaTable {
	return s.Data
}

func (s *GetLumaTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLumaTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLumaTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLumaTableResponseBody) SetCode(v string) *GetLumaTableResponseBody {
	s.Code = &v
	return s
}

func (s *GetLumaTableResponseBody) SetData(v *LumaTable) *GetLumaTableResponseBody {
	s.Data = v
	return s
}

func (s *GetLumaTableResponseBody) SetMessage(v string) *GetLumaTableResponseBody {
	s.Message = &v
	return s
}

func (s *GetLumaTableResponseBody) SetRequestId(v string) *GetLumaTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLumaTableResponseBody) SetSuccess(v bool) *GetLumaTableResponseBody {
	s.Success = &v
	return s
}

func (s *GetLumaTableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
