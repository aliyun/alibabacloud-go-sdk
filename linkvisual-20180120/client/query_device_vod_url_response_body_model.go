// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceVodUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryDeviceVodUrlResponseBody
	GetCode() *string
	SetData(v *QueryDeviceVodUrlResponseBodyData) *QueryDeviceVodUrlResponseBody
	GetData() *QueryDeviceVodUrlResponseBodyData
	SetErrorMessage(v string) *QueryDeviceVodUrlResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryDeviceVodUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDeviceVodUrlResponseBody
	GetSuccess() *bool
}

type QueryDeviceVodUrlResponseBody struct {
	// The response code. Valid values:
	//
	// - **200**: The call was successful.
	//
	// - Other values: The call failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data *QueryDeviceVodUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// product is not active.
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

func (s QueryDeviceVodUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceVodUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryDeviceVodUrlResponseBody) GetData() *QueryDeviceVodUrlResponseBodyData {
	return s.Data
}

func (s *QueryDeviceVodUrlResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryDeviceVodUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDeviceVodUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDeviceVodUrlResponseBody) SetCode(v string) *QueryDeviceVodUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBody) SetData(v *QueryDeviceVodUrlResponseBodyData) *QueryDeviceVodUrlResponseBody {
	s.Data = v
	return s
}

func (s *QueryDeviceVodUrlResponseBody) SetErrorMessage(v string) *QueryDeviceVodUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBody) SetRequestId(v string) *QueryDeviceVodUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBody) SetSuccess(v bool) *QueryDeviceVodUrlResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDeviceVodUrlResponseBodyData struct {
	// The decryption key.
	//
	// A decryption key is generated when the **ShouldEncrypt*	- parameter is set to **true**.
	//
	// example:
	//
	// {"iv": "MDEy****OTo7PD0+Pw==","key": "eA2r****fHjOmnyBfQ=="}
	DecryptKey *string `json:"DecryptKey,omitempty" xml:"DecryptKey,omitempty"`
	StunInfo   *string `json:"StunInfo,omitempty" xml:"StunInfo,omitempty"`
	// The URL of the local recording.
	VodUrl *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s QueryDeviceVodUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceVodUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDeviceVodUrlResponseBodyData) GetDecryptKey() *string {
	return s.DecryptKey
}

func (s *QueryDeviceVodUrlResponseBodyData) GetStunInfo() *string {
	return s.StunInfo
}

func (s *QueryDeviceVodUrlResponseBodyData) GetVodUrl() *string {
	return s.VodUrl
}

func (s *QueryDeviceVodUrlResponseBodyData) SetDecryptKey(v string) *QueryDeviceVodUrlResponseBodyData {
	s.DecryptKey = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBodyData) SetStunInfo(v string) *QueryDeviceVodUrlResponseBodyData {
	s.StunInfo = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBodyData) SetVodUrl(v string) *QueryDeviceVodUrlResponseBodyData {
	s.VodUrl = &v
	return s
}

func (s *QueryDeviceVodUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
