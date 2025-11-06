// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAckInfoOfMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAckInfoOfMessageResponseBody
	GetCode() *int32
	SetData(v []*GetAckInfoOfMessageResponseBodyData) *GetAckInfoOfMessageResponseBody
	GetData() []*GetAckInfoOfMessageResponseBodyData
	SetMessage(v string) *GetAckInfoOfMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAckInfoOfMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAckInfoOfMessageResponseBody
	GetSuccess() *bool
}

type GetAckInfoOfMessageResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetAckInfoOfMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAckInfoOfMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAckInfoOfMessageResponseBody) GoString() string {
	return s.String()
}

func (s *GetAckInfoOfMessageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAckInfoOfMessageResponseBody) GetData() []*GetAckInfoOfMessageResponseBodyData {
	return s.Data
}

func (s *GetAckInfoOfMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAckInfoOfMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAckInfoOfMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAckInfoOfMessageResponseBody) SetCode(v int32) *GetAckInfoOfMessageResponseBody {
	s.Code = &v
	return s
}

func (s *GetAckInfoOfMessageResponseBody) SetData(v []*GetAckInfoOfMessageResponseBodyData) *GetAckInfoOfMessageResponseBody {
	s.Data = v
	return s
}

func (s *GetAckInfoOfMessageResponseBody) SetMessage(v string) *GetAckInfoOfMessageResponseBody {
	s.Message = &v
	return s
}

func (s *GetAckInfoOfMessageResponseBody) SetRequestId(v string) *GetAckInfoOfMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAckInfoOfMessageResponseBody) SetSuccess(v bool) *GetAckInfoOfMessageResponseBody {
	s.Success = &v
	return s
}

func (s *GetAckInfoOfMessageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAckInfoOfMessageResponseBodyData struct {
	AckErrorInfo *string                `json:"AckErrorInfo,omitempty" xml:"AckErrorInfo,omitempty"`
	AckResult    *string                `json:"AckResult,omitempty" xml:"AckResult,omitempty"`
	Action       *string                `json:"Action,omitempty" xml:"Action,omitempty"`
	Code         *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Property     map[string]interface{} `json:"Property,omitempty" xml:"Property,omitempty"`
	TimeStamp    *string                `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s GetAckInfoOfMessageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAckInfoOfMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAckInfoOfMessageResponseBodyData) GetAckErrorInfo() *string {
	return s.AckErrorInfo
}

func (s *GetAckInfoOfMessageResponseBodyData) GetAckResult() *string {
	return s.AckResult
}

func (s *GetAckInfoOfMessageResponseBodyData) GetAction() *string {
	return s.Action
}

func (s *GetAckInfoOfMessageResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetAckInfoOfMessageResponseBodyData) GetProperty() map[string]interface{} {
	return s.Property
}

func (s *GetAckInfoOfMessageResponseBodyData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *GetAckInfoOfMessageResponseBodyData) SetAckErrorInfo(v string) *GetAckInfoOfMessageResponseBodyData {
	s.AckErrorInfo = &v
	return s
}

func (s *GetAckInfoOfMessageResponseBodyData) SetAckResult(v string) *GetAckInfoOfMessageResponseBodyData {
	s.AckResult = &v
	return s
}

func (s *GetAckInfoOfMessageResponseBodyData) SetAction(v string) *GetAckInfoOfMessageResponseBodyData {
	s.Action = &v
	return s
}

func (s *GetAckInfoOfMessageResponseBodyData) SetCode(v string) *GetAckInfoOfMessageResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetAckInfoOfMessageResponseBodyData) SetProperty(v map[string]interface{}) *GetAckInfoOfMessageResponseBodyData {
	s.Property = v
	return s
}

func (s *GetAckInfoOfMessageResponseBodyData) SetTimeStamp(v string) *GetAckInfoOfMessageResponseBodyData {
	s.TimeStamp = &v
	return s
}

func (s *GetAckInfoOfMessageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
