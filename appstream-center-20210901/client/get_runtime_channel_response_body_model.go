// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuntimeChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetRuntimeChannelResponseBodyData) *GetRuntimeChannelResponseBody
	GetData() []*GetRuntimeChannelResponseBodyData
	SetRequestId(v string) *GetRuntimeChannelResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetRuntimeChannelResponseBody
	GetTotalCount() *int32
}

type GetRuntimeChannelResponseBody struct {
	Data []*GetRuntimeChannelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetRuntimeChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuntimeChannelResponseBody) GetData() []*GetRuntimeChannelResponseBodyData {
	return s.Data
}

func (s *GetRuntimeChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRuntimeChannelResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetRuntimeChannelResponseBody) SetData(v []*GetRuntimeChannelResponseBodyData) *GetRuntimeChannelResponseBody {
	s.Data = v
	return s
}

func (s *GetRuntimeChannelResponseBody) SetRequestId(v string) *GetRuntimeChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuntimeChannelResponseBody) SetTotalCount(v int32) *GetRuntimeChannelResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetRuntimeChannelResponseBody) Validate() error {
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

type GetRuntimeChannelResponseBodyData struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// example:
	//
	// dingtalk-connector
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {
	//
	//     "appKey": "abc",
	//
	//     "appSecret": "efg"
	//
	// }
	Config          *string `json:"Config,omitempty" xml:"Config,omitempty"`
	ConfigMode      *string `json:"ConfigMode,omitempty" xml:"ConfigMode,omitempty"`
	QrCodeNotifyUrl *string `json:"QrCodeNotifyUrl,omitempty" xml:"QrCodeNotifyUrl,omitempty"`
	QrCodeStatus    *string `json:"QrCodeStatus,omitempty" xml:"QrCodeStatus,omitempty"`
	// example:
	//
	// Normal
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	// example:
	//
	// Configured
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRuntimeChannelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeChannelResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuntimeChannelResponseBodyData) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetRuntimeChannelResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetRuntimeChannelResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *GetRuntimeChannelResponseBodyData) GetConfigMode() *string {
	return s.ConfigMode
}

func (s *GetRuntimeChannelResponseBodyData) GetQrCodeNotifyUrl() *string {
	return s.QrCodeNotifyUrl
}

func (s *GetRuntimeChannelResponseBodyData) GetQrCodeStatus() *string {
	return s.QrCodeStatus
}

func (s *GetRuntimeChannelResponseBodyData) GetRiskType() *string {
	return s.RiskType
}

func (s *GetRuntimeChannelResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetRuntimeChannelResponseBodyData) SetAvatarUrl(v string) *GetRuntimeChannelResponseBodyData {
	s.AvatarUrl = &v
	return s
}

func (s *GetRuntimeChannelResponseBodyData) SetCode(v string) *GetRuntimeChannelResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetRuntimeChannelResponseBodyData) SetConfig(v string) *GetRuntimeChannelResponseBodyData {
	s.Config = &v
	return s
}

func (s *GetRuntimeChannelResponseBodyData) SetConfigMode(v string) *GetRuntimeChannelResponseBodyData {
	s.ConfigMode = &v
	return s
}

func (s *GetRuntimeChannelResponseBodyData) SetQrCodeNotifyUrl(v string) *GetRuntimeChannelResponseBodyData {
	s.QrCodeNotifyUrl = &v
	return s
}

func (s *GetRuntimeChannelResponseBodyData) SetQrCodeStatus(v string) *GetRuntimeChannelResponseBodyData {
	s.QrCodeStatus = &v
	return s
}

func (s *GetRuntimeChannelResponseBodyData) SetRiskType(v string) *GetRuntimeChannelResponseBodyData {
	s.RiskType = &v
	return s
}

func (s *GetRuntimeChannelResponseBodyData) SetStatus(v string) *GetRuntimeChannelResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRuntimeChannelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
