// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTpsByTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTpsByTimeResponseBody
	GetCode() *int32
	SetData(v *GetTpsByTimeResponseBodyData) *GetTpsByTimeResponseBody
	GetData() *GetTpsByTimeResponseBodyData
	SetMessage(v string) *GetTpsByTimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTpsByTimeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTpsByTimeResponseBody
	GetSuccess() *bool
}

type GetTpsByTimeResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetTpsByTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTpsByTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTpsByTimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetTpsByTimeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTpsByTimeResponseBody) GetData() *GetTpsByTimeResponseBodyData {
	return s.Data
}

func (s *GetTpsByTimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTpsByTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTpsByTimeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTpsByTimeResponseBody) SetCode(v int32) *GetTpsByTimeResponseBody {
	s.Code = &v
	return s
}

func (s *GetTpsByTimeResponseBody) SetData(v *GetTpsByTimeResponseBodyData) *GetTpsByTimeResponseBody {
	s.Data = v
	return s
}

func (s *GetTpsByTimeResponseBody) SetMessage(v string) *GetTpsByTimeResponseBody {
	s.Message = &v
	return s
}

func (s *GetTpsByTimeResponseBody) SetRequestId(v string) *GetTpsByTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTpsByTimeResponseBody) SetSuccess(v bool) *GetTpsByTimeResponseBody {
	s.Success = &v
	return s
}

func (s *GetTpsByTimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTpsByTimeResponseBodyData struct {
	EndTime   *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxTps    *int32   `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	StartTime *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TpsList   []*int32 `json:"tpsList,omitempty" xml:"tpsList,omitempty" type:"Repeated"`
}

func (s GetTpsByTimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTpsByTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTpsByTimeResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTpsByTimeResponseBodyData) GetMaxTps() *int32 {
	return s.MaxTps
}

func (s *GetTpsByTimeResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetTpsByTimeResponseBodyData) GetTpsList() []*int32 {
	return s.TpsList
}

func (s *GetTpsByTimeResponseBodyData) SetEndTime(v string) *GetTpsByTimeResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetTpsByTimeResponseBodyData) SetMaxTps(v int32) *GetTpsByTimeResponseBodyData {
	s.MaxTps = &v
	return s
}

func (s *GetTpsByTimeResponseBodyData) SetStartTime(v int64) *GetTpsByTimeResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetTpsByTimeResponseBodyData) SetTpsList(v []*int32) *GetTpsByTimeResponseBodyData {
	s.TpsList = v
	return s
}

func (s *GetTpsByTimeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
