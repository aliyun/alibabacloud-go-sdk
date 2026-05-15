// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineMessageLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotlineMessageLogResponseBody
	GetCode() *string
	SetData(v []*GetHotlineMessageLogResponseBodyData) *GetHotlineMessageLogResponseBody
	GetData() []*GetHotlineMessageLogResponseBodyData
	SetMessage(v string) *GetHotlineMessageLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotlineMessageLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotlineMessageLogResponseBody
	GetSuccess() *bool
}

type GetHotlineMessageLogResponseBody struct {
	// example:
	//
	// Success
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetHotlineMessageLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineMessageLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineMessageLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineMessageLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotlineMessageLogResponseBody) GetData() []*GetHotlineMessageLogResponseBodyData {
	return s.Data
}

func (s *GetHotlineMessageLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotlineMessageLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotlineMessageLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotlineMessageLogResponseBody) SetCode(v string) *GetHotlineMessageLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineMessageLogResponseBody) SetData(v []*GetHotlineMessageLogResponseBodyData) *GetHotlineMessageLogResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineMessageLogResponseBody) SetMessage(v string) *GetHotlineMessageLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineMessageLogResponseBody) SetRequestId(v string) *GetHotlineMessageLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineMessageLogResponseBody) SetSuccess(v bool) *GetHotlineMessageLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineMessageLogResponseBody) Validate() error {
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

type GetHotlineMessageLogResponseBodyData struct {
	// example:
	//
	// 100****2077
	Acid    *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1623738027480
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 11deca999****
	Mid *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	// example:
	//
	// 1
	SenderType *int32 `json:"SenderType,omitempty" xml:"SenderType,omitempty"`
	// example:
	//
	// 1623738026460
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetHotlineMessageLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineMessageLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineMessageLogResponseBodyData) GetAcid() *string {
	return s.Acid
}

func (s *GetHotlineMessageLogResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetHotlineMessageLogResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetHotlineMessageLogResponseBodyData) GetMid() *string {
	return s.Mid
}

func (s *GetHotlineMessageLogResponseBodyData) GetSenderType() *int32 {
	return s.SenderType
}

func (s *GetHotlineMessageLogResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetHotlineMessageLogResponseBodyData) SetAcid(v string) *GetHotlineMessageLogResponseBodyData {
	s.Acid = &v
	return s
}

func (s *GetHotlineMessageLogResponseBodyData) SetContent(v string) *GetHotlineMessageLogResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetHotlineMessageLogResponseBodyData) SetEndTime(v int64) *GetHotlineMessageLogResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetHotlineMessageLogResponseBodyData) SetMid(v string) *GetHotlineMessageLogResponseBodyData {
	s.Mid = &v
	return s
}

func (s *GetHotlineMessageLogResponseBodyData) SetSenderType(v int32) *GetHotlineMessageLogResponseBodyData {
	s.SenderType = &v
	return s
}

func (s *GetHotlineMessageLogResponseBodyData) SetStartTime(v int64) *GetHotlineMessageLogResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetHotlineMessageLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
