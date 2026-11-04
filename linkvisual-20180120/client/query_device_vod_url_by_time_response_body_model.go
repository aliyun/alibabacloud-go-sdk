// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceVodUrlByTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryDeviceVodUrlByTimeResponseBody
	GetCode() *string
	SetData(v *QueryDeviceVodUrlByTimeResponseBodyData) *QueryDeviceVodUrlByTimeResponseBody
	GetData() *QueryDeviceVodUrlByTimeResponseBodyData
	SetErrorMessage(v string) *QueryDeviceVodUrlByTimeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryDeviceVodUrlByTimeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDeviceVodUrlByTimeResponseBody
	GetSuccess() *bool
}

type QueryDeviceVodUrlByTimeResponseBody struct {
	// The response code. Valid values:
	//
	// - **200**: The call was successful.
	//
	// - Other values: An error occurred. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryDeviceVodUrlByTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// tenant auth error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06DC77A0-4622-42DB-9EE0-25FIOHS82JK1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDeviceVodUrlByTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceVodUrlByTimeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlByTimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryDeviceVodUrlByTimeResponseBody) GetData() *QueryDeviceVodUrlByTimeResponseBodyData {
	return s.Data
}

func (s *QueryDeviceVodUrlByTimeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryDeviceVodUrlByTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDeviceVodUrlByTimeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDeviceVodUrlByTimeResponseBody) SetCode(v string) *QueryDeviceVodUrlByTimeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBody) SetData(v *QueryDeviceVodUrlByTimeResponseBodyData) *QueryDeviceVodUrlByTimeResponseBody {
	s.Data = v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBody) SetErrorMessage(v string) *QueryDeviceVodUrlByTimeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBody) SetRequestId(v string) *QueryDeviceVodUrlByTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBody) SetSuccess(v bool) *QueryDeviceVodUrlByTimeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDeviceVodUrlByTimeResponseBodyData struct {
	// The decryption key.
	//
	// A decryption key is generated when the **ShouldEncrypt*	- parameter is set to **true**.
	//
	// example:
	//
	// {"iv": "MDEy****OTo7PD0+Pw==","key": "eA2r****fHjOmnyBfQ=="}
	DecryptKey *string `json:"DecryptKey,omitempty" xml:"DecryptKey,omitempty"`
	StunInfo   *string `json:"StunInfo,omitempty" xml:"StunInfo,omitempty"`
	// The playback URL of the local recording.
	//
	// example:
	//
	// https://iotx-vision-vod-vpc-hz-pre.aliyun-inc.com/vod/device/localrecord/flv/play/1****668.flv?token=e9****a9e14482c56be52&session=470e****3d015
	VodUrl *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s QueryDeviceVodUrlByTimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceVodUrlByTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlByTimeResponseBodyData) GetDecryptKey() *string {
	return s.DecryptKey
}

func (s *QueryDeviceVodUrlByTimeResponseBodyData) GetStunInfo() *string {
	return s.StunInfo
}

func (s *QueryDeviceVodUrlByTimeResponseBodyData) GetVodUrl() *string {
	return s.VodUrl
}

func (s *QueryDeviceVodUrlByTimeResponseBodyData) SetDecryptKey(v string) *QueryDeviceVodUrlByTimeResponseBodyData {
	s.DecryptKey = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBodyData) SetStunInfo(v string) *QueryDeviceVodUrlByTimeResponseBodyData {
	s.StunInfo = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBodyData) SetVodUrl(v string) *QueryDeviceVodUrlByTimeResponseBodyData {
	s.VodUrl = &v
	return s
}

func (s *QueryDeviceVodUrlByTimeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
