// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByDeviceFileIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendByDeviceFileIdResponseBody
	GetCode() *string
	SetData(v *SendByDeviceFileIdResponseBodyData) *SendByDeviceFileIdResponseBody
	GetData() *SendByDeviceFileIdResponseBodyData
	SetHttpStatusCode(v int32) *SendByDeviceFileIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SendByDeviceFileIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendByDeviceFileIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendByDeviceFileIdResponseBody
	GetSuccess() *bool
}

type SendByDeviceFileIdResponseBody struct {
	// example:
	//
	// 0
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendByDeviceFileIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendByDeviceFileIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendByDeviceFileIdResponseBody) GoString() string {
	return s.String()
}

func (s *SendByDeviceFileIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendByDeviceFileIdResponseBody) GetData() *SendByDeviceFileIdResponseBodyData {
	return s.Data
}

func (s *SendByDeviceFileIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SendByDeviceFileIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendByDeviceFileIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendByDeviceFileIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendByDeviceFileIdResponseBody) SetCode(v string) *SendByDeviceFileIdResponseBody {
	s.Code = &v
	return s
}

func (s *SendByDeviceFileIdResponseBody) SetData(v *SendByDeviceFileIdResponseBodyData) *SendByDeviceFileIdResponseBody {
	s.Data = v
	return s
}

func (s *SendByDeviceFileIdResponseBody) SetHttpStatusCode(v int32) *SendByDeviceFileIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByDeviceFileIdResponseBody) SetMessage(v string) *SendByDeviceFileIdResponseBody {
	s.Message = &v
	return s
}

func (s *SendByDeviceFileIdResponseBody) SetRequestId(v string) *SendByDeviceFileIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByDeviceFileIdResponseBody) SetSuccess(v bool) *SendByDeviceFileIdResponseBody {
	s.Success = &v
	return s
}

func (s *SendByDeviceFileIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendByDeviceFileIdResponseBodyData struct {
	// example:
	//
	// ufe29y2167046828041801
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByDeviceFileIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendByDeviceFileIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByDeviceFileIdResponseBodyData) GetMsgId() *string {
	return s.MsgId
}

func (s *SendByDeviceFileIdResponseBodyData) SetMsgId(v string) *SendByDeviceFileIdResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *SendByDeviceFileIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
