// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTotalPublicUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTotalPublicUrlResponseBody
	GetCode() *string
	SetData(v *GetTotalPublicUrlResponseBodyData) *GetTotalPublicUrlResponseBody
	GetData() *GetTotalPublicUrlResponseBodyData
	SetMessage(v string) *GetTotalPublicUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTotalPublicUrlResponseBody
	GetRequestId() *string
}

type GetTotalPublicUrlResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The download URLs of the recording files.
	Data *GetTotalPublicUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AB3CEF7-DCBE-488C-9C33-D180982CE031
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTotalPublicUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTotalPublicUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTotalPublicUrlResponseBody) GetData() *GetTotalPublicUrlResponseBodyData {
	return s.Data
}

func (s *GetTotalPublicUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTotalPublicUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTotalPublicUrlResponseBody) SetCode(v string) *GetTotalPublicUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetTotalPublicUrlResponseBody) SetData(v *GetTotalPublicUrlResponseBodyData) *GetTotalPublicUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetTotalPublicUrlResponseBody) SetMessage(v string) *GetTotalPublicUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetTotalPublicUrlResponseBody) SetRequestId(v string) *GetTotalPublicUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTotalPublicUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTotalPublicUrlResponseBodyData struct {
	// The download URL of the recorded call.
	//
	// >  The download URL of the recorded call is valid for 30 days.
	//
	// example:
	//
	// http://secret-axb-reco****cn-shanghai.aliyuncs.com/1000000820257625_66647243838006067251551752068865.mp3?Expires=1551****07&OSSAccessKeyId=LTAIP00vvvv****v&Signature=tK6Yq9KusU4n%2BZQWX****4/WmEA%3D
	PhonePublicUrl *string `json:"PhonePublicUrl,omitempty" xml:"PhonePublicUrl,omitempty"`
	// The download URL of the recorded ringing tone.
	//
	// >  The download URL of the recorded ringing tone is valid for 30 days.
	//
	// example:
	//
	// http://secret-ab-reco****cn-shanghai.aliyuncs.com/1000000820257625_66647243838006067251551752068865.mp3?Expires=155175****&OSSAccessKeyId=LTAIP00vvv****vv&Signature=tK6Yq9KusU4n%2BZQW****g4/WmEA%3D
	RingPublicUrl *string `json:"RingPublicUrl,omitempty" xml:"RingPublicUrl,omitempty"`
}

func (s GetTotalPublicUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTotalPublicUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTotalPublicUrlResponseBodyData) GetPhonePublicUrl() *string {
	return s.PhonePublicUrl
}

func (s *GetTotalPublicUrlResponseBodyData) GetRingPublicUrl() *string {
	return s.RingPublicUrl
}

func (s *GetTotalPublicUrlResponseBodyData) SetPhonePublicUrl(v string) *GetTotalPublicUrlResponseBodyData {
	s.PhonePublicUrl = &v
	return s
}

func (s *GetTotalPublicUrlResponseBodyData) SetRingPublicUrl(v string) *GetTotalPublicUrlResponseBodyData {
	s.RingPublicUrl = &v
	return s
}

func (s *GetTotalPublicUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
