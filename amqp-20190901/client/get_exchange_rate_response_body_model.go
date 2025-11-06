// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExchangeRateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetExchangeRateResponseBody
	GetCode() *int32
	SetData(v *GetExchangeRateResponseBodyData) *GetExchangeRateResponseBody
	GetData() *GetExchangeRateResponseBodyData
	SetMessage(v string) *GetExchangeRateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetExchangeRateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetExchangeRateResponseBody
	GetSuccess() *bool
}

type GetExchangeRateResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetExchangeRateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetExchangeRateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeRateResponseBody) GoString() string {
	return s.String()
}

func (s *GetExchangeRateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetExchangeRateResponseBody) GetData() *GetExchangeRateResponseBodyData {
	return s.Data
}

func (s *GetExchangeRateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetExchangeRateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExchangeRateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetExchangeRateResponseBody) SetCode(v int32) *GetExchangeRateResponseBody {
	s.Code = &v
	return s
}

func (s *GetExchangeRateResponseBody) SetData(v *GetExchangeRateResponseBodyData) *GetExchangeRateResponseBody {
	s.Data = v
	return s
}

func (s *GetExchangeRateResponseBody) SetMessage(v string) *GetExchangeRateResponseBody {
	s.Message = &v
	return s
}

func (s *GetExchangeRateResponseBody) SetRequestId(v string) *GetExchangeRateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExchangeRateResponseBody) SetSuccess(v bool) *GetExchangeRateResponseBody {
	s.Success = &v
	return s
}

func (s *GetExchangeRateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExchangeRateResponseBodyData struct {
	ExchangeQuotaVO []*GetExchangeRateResponseBodyDataExchangeQuotaVO `json:"ExchangeQuotaVO,omitempty" xml:"ExchangeQuotaVO,omitempty" type:"Repeated"`
}

func (s GetExchangeRateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeRateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExchangeRateResponseBodyData) GetExchangeQuotaVO() []*GetExchangeRateResponseBodyDataExchangeQuotaVO {
	return s.ExchangeQuotaVO
}

func (s *GetExchangeRateResponseBodyData) SetExchangeQuotaVO(v []*GetExchangeRateResponseBodyDataExchangeQuotaVO) *GetExchangeRateResponseBodyData {
	s.ExchangeQuotaVO = v
	return s
}

func (s *GetExchangeRateResponseBodyData) Validate() error {
	if s.ExchangeQuotaVO != nil {
		for _, item := range s.ExchangeQuotaVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetExchangeRateResponseBodyDataExchangeQuotaVO struct {
	ExchangeName *string `json:"ExchangeName,omitempty" xml:"ExchangeName,omitempty"`
	InQps        *int64  `json:"InQps,omitempty" xml:"InQps,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutQps       *int64  `json:"OutQps,omitempty" xml:"OutQps,omitempty"`
	VhostName    *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetExchangeRateResponseBodyDataExchangeQuotaVO) String() string {
	return dara.Prettify(s)
}

func (s GetExchangeRateResponseBodyDataExchangeQuotaVO) GoString() string {
	return s.String()
}

func (s *GetExchangeRateResponseBodyDataExchangeQuotaVO) GetExchangeName() *string {
	return s.ExchangeName
}

func (s *GetExchangeRateResponseBodyDataExchangeQuotaVO) GetInQps() *int64 {
	return s.InQps
}

func (s *GetExchangeRateResponseBodyDataExchangeQuotaVO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetExchangeRateResponseBodyDataExchangeQuotaVO) GetOutQps() *int64 {
	return s.OutQps
}

func (s *GetExchangeRateResponseBodyDataExchangeQuotaVO) GetVhostName() *string {
	return s.VhostName
}

func (s *GetExchangeRateResponseBodyDataExchangeQuotaVO) SetExchangeName(v string) *GetExchangeRateResponseBodyDataExchangeQuotaVO {
	s.ExchangeName = &v
	return s
}

func (s *GetExchangeRateResponseBodyDataExchangeQuotaVO) SetInQps(v int64) *GetExchangeRateResponseBodyDataExchangeQuotaVO {
	s.InQps = &v
	return s
}

func (s *GetExchangeRateResponseBodyDataExchangeQuotaVO) SetInstanceId(v string) *GetExchangeRateResponseBodyDataExchangeQuotaVO {
	s.InstanceId = &v
	return s
}

func (s *GetExchangeRateResponseBodyDataExchangeQuotaVO) SetOutQps(v int64) *GetExchangeRateResponseBodyDataExchangeQuotaVO {
	s.OutQps = &v
	return s
}

func (s *GetExchangeRateResponseBodyDataExchangeQuotaVO) SetVhostName(v string) *GetExchangeRateResponseBodyDataExchangeQuotaVO {
	s.VhostName = &v
	return s
}

func (s *GetExchangeRateResponseBodyDataExchangeQuotaVO) Validate() error {
	return dara.Validate(s)
}
