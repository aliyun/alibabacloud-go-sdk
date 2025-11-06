// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVhostRateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetVhostRateResponseBody
	GetCode() *int32
	SetData(v []*GetVhostRateResponseBodyData) *GetVhostRateResponseBody
	GetData() []*GetVhostRateResponseBodyData
	SetMessage(v string) *GetVhostRateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVhostRateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVhostRateResponseBody
	GetSuccess() *bool
}

type GetVhostRateResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetVhostRateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVhostRateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVhostRateResponseBody) GoString() string {
	return s.String()
}

func (s *GetVhostRateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetVhostRateResponseBody) GetData() []*GetVhostRateResponseBodyData {
	return s.Data
}

func (s *GetVhostRateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVhostRateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVhostRateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVhostRateResponseBody) SetCode(v int32) *GetVhostRateResponseBody {
	s.Code = &v
	return s
}

func (s *GetVhostRateResponseBody) SetData(v []*GetVhostRateResponseBodyData) *GetVhostRateResponseBody {
	s.Data = v
	return s
}

func (s *GetVhostRateResponseBody) SetMessage(v string) *GetVhostRateResponseBody {
	s.Message = &v
	return s
}

func (s *GetVhostRateResponseBody) SetRequestId(v string) *GetVhostRateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVhostRateResponseBody) SetSuccess(v bool) *GetVhostRateResponseBody {
	s.Success = &v
	return s
}

func (s *GetVhostRateResponseBody) Validate() error {
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

type GetVhostRateResponseBodyData struct {
	ChannelNum    *int32  `json:"ChannelNum,omitempty" xml:"ChannelNum,omitempty"`
	ConnectionNum *int32  `json:"ConnectionNum,omitempty" xml:"ConnectionNum,omitempty"`
	InQps         *int64  `json:"InQps,omitempty" xml:"InQps,omitempty"`
	OutQps        *int64  `json:"OutQps,omitempty" xml:"OutQps,omitempty"`
	VhostName     *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetVhostRateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVhostRateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVhostRateResponseBodyData) GetChannelNum() *int32 {
	return s.ChannelNum
}

func (s *GetVhostRateResponseBodyData) GetConnectionNum() *int32 {
	return s.ConnectionNum
}

func (s *GetVhostRateResponseBodyData) GetInQps() *int64 {
	return s.InQps
}

func (s *GetVhostRateResponseBodyData) GetOutQps() *int64 {
	return s.OutQps
}

func (s *GetVhostRateResponseBodyData) GetVhostName() *string {
	return s.VhostName
}

func (s *GetVhostRateResponseBodyData) SetChannelNum(v int32) *GetVhostRateResponseBodyData {
	s.ChannelNum = &v
	return s
}

func (s *GetVhostRateResponseBodyData) SetConnectionNum(v int32) *GetVhostRateResponseBodyData {
	s.ConnectionNum = &v
	return s
}

func (s *GetVhostRateResponseBodyData) SetInQps(v int64) *GetVhostRateResponseBodyData {
	s.InQps = &v
	return s
}

func (s *GetVhostRateResponseBodyData) SetOutQps(v int64) *GetVhostRateResponseBodyData {
	s.OutQps = &v
	return s
}

func (s *GetVhostRateResponseBodyData) SetVhostName(v string) *GetVhostRateResponseBodyData {
	s.VhostName = &v
	return s
}

func (s *GetVhostRateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
