// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuntimeChannelQrCodeDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetRuntimeChannelQrCodeDetailResponseBodyData) *GetRuntimeChannelQrCodeDetailResponseBody
	GetData() *GetRuntimeChannelQrCodeDetailResponseBodyData
	SetRequestId(v string) *GetRuntimeChannelQrCodeDetailResponseBody
	GetRequestId() *string
}

type GetRuntimeChannelQrCodeDetailResponseBody struct {
	// The QR code configuration details of the channel. If no QR code configuration process exists for the channel under the specified runtime, this field is empty.
	Data *GetRuntimeChannelQrCodeDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, which uniquely identifies the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRuntimeChannelQrCodeDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeChannelQrCodeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuntimeChannelQrCodeDetailResponseBody) GetData() *GetRuntimeChannelQrCodeDetailResponseBodyData {
	return s.Data
}

func (s *GetRuntimeChannelQrCodeDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRuntimeChannelQrCodeDetailResponseBody) SetData(v *GetRuntimeChannelQrCodeDetailResponseBodyData) *GetRuntimeChannelQrCodeDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailResponseBody) SetRequestId(v string) *GetRuntimeChannelQrCodeDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuntimeChannelQrCodeDetailResponseBodyData struct {
	// The URL or content of the QR code. If the QR code is not yet generated or the content is empty, this field is empty.
	//
	// example:
	//
	// https://qr.example.com/abc
	QrCode *string `json:"QrCode,omitempty" xml:"QrCode,omitempty"`
	// The access URL of the QR code configuration channel page. This URL is generated and sent to the user when the QR code configuration is initiated, and carries the taskId, channel code, and region parameters of the current process. If the URL is not recorded, this field is empty.
	//
	// example:
	//
	// https://eds.console.aliyun.com/public/scan-pair?taskId=8f14e45f-ceea-4672-9f6e-3c5b2a1d****&code=dingtalk-connector&region=cn-shanghai
	QrCodeNotifyUrl *string `json:"QrCodeNotifyUrl,omitempty" xml:"QrCodeNotifyUrl,omitempty"`
	// The scan status. Valid values: Generating (generating), Generated (generated), GenerateFailed (generation failed), Scanned (scanned), Configuring (configuring), Success (success), Failed (failed), and Expired (expired).
	//
	// example:
	//
	// Generated
	QrCodeStatus *string `json:"QrCodeStatus,omitempty" xml:"QrCodeStatus,omitempty"`
	// The status of the QR code configuration process. The valid values are the same as those of QrCodeStatus. Valid values: Generating (generating), Generated (generated), GenerateFailed (generation failed), Scanned (scanned), Configuring (configuring), Success (success), Failed (failed), and Expired (expired).
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRuntimeChannelQrCodeDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeChannelQrCodeDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuntimeChannelQrCodeDetailResponseBodyData) GetQrCode() *string {
	return s.QrCode
}

func (s *GetRuntimeChannelQrCodeDetailResponseBodyData) GetQrCodeNotifyUrl() *string {
	return s.QrCodeNotifyUrl
}

func (s *GetRuntimeChannelQrCodeDetailResponseBodyData) GetQrCodeStatus() *string {
	return s.QrCodeStatus
}

func (s *GetRuntimeChannelQrCodeDetailResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetRuntimeChannelQrCodeDetailResponseBodyData) SetQrCode(v string) *GetRuntimeChannelQrCodeDetailResponseBodyData {
	s.QrCode = &v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailResponseBodyData) SetQrCodeNotifyUrl(v string) *GetRuntimeChannelQrCodeDetailResponseBodyData {
	s.QrCodeNotifyUrl = &v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailResponseBodyData) SetQrCodeStatus(v string) *GetRuntimeChannelQrCodeDetailResponseBodyData {
	s.QrCodeStatus = &v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailResponseBodyData) SetStatus(v string) *GetRuntimeChannelQrCodeDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRuntimeChannelQrCodeDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
