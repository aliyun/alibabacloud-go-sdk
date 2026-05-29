// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableTtsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAvailableTtsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListAvailableTtsResponseBody
	GetCode() *string
	SetData(v []*ListAvailableTtsResponseBodyData) *ListAvailableTtsResponseBody
	GetData() []*ListAvailableTtsResponseBodyData
	SetMessage(v string) *ListAvailableTtsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAvailableTtsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAvailableTtsResponseBody
	GetSuccess() *bool
}

type ListAvailableTtsResponseBody struct {
	// example:
	//
	// Access Denied
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListAvailableTtsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5EBAEBE6-0E77-5E1F-99E4-7B20512F22222
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAvailableTtsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableTtsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableTtsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAvailableTtsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAvailableTtsResponseBody) GetData() []*ListAvailableTtsResponseBodyData {
	return s.Data
}

func (s *ListAvailableTtsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAvailableTtsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvailableTtsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAvailableTtsResponseBody) SetAccessDeniedDetail(v string) *ListAvailableTtsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAvailableTtsResponseBody) SetCode(v string) *ListAvailableTtsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAvailableTtsResponseBody) SetData(v []*ListAvailableTtsResponseBodyData) *ListAvailableTtsResponseBody {
	s.Data = v
	return s
}

func (s *ListAvailableTtsResponseBody) SetMessage(v string) *ListAvailableTtsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAvailableTtsResponseBody) SetRequestId(v string) *ListAvailableTtsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableTtsResponseBody) SetSuccess(v bool) *ListAvailableTtsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAvailableTtsResponseBody) Validate() error {
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

type ListAvailableTtsResponseBodyData struct {
	// 音色试听文件
	//
	// example:
	//
	// 示例值
	TtsAuditionFileUrl *string `json:"TtsAuditionFileUrl,omitempty" xml:"TtsAuditionFileUrl,omitempty"`
	// 音色平台CosyVoice/Volcano
	//
	// example:
	//
	// CosyVoice/Volcano
	TtsEngine *string `json:"TtsEngine,omitempty" xml:"TtsEngine,omitempty"`
	// 音色cosy类型 cosyvoice-v2-XXXX-XXXX
	//
	// example:
	//
	// 5EBAEBE6-0E77-5E1F-99E4-7B20512FCB3C
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// 音色编码
	//
	// example:
	//
	// V123456789
	TtsVoiceCode *string `json:"TtsVoiceCode,omitempty" xml:"TtsVoiceCode,omitempty"`
	// 音色名称
	//
	// example:
	//
	// 示例值示例值
	TtsVoiceName *string `json:"TtsVoiceName,omitempty" xml:"TtsVoiceName,omitempty"`
	// example:
	//
	// 示例值示例值
	VoiceType *string `json:"VoiceType,omitempty" xml:"VoiceType,omitempty"`
}

func (s ListAvailableTtsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableTtsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAvailableTtsResponseBodyData) GetTtsAuditionFileUrl() *string {
	return s.TtsAuditionFileUrl
}

func (s *ListAvailableTtsResponseBodyData) GetTtsEngine() *string {
	return s.TtsEngine
}

func (s *ListAvailableTtsResponseBodyData) GetTtsStyle() *string {
	return s.TtsStyle
}

func (s *ListAvailableTtsResponseBodyData) GetTtsVoiceCode() *string {
	return s.TtsVoiceCode
}

func (s *ListAvailableTtsResponseBodyData) GetTtsVoiceName() *string {
	return s.TtsVoiceName
}

func (s *ListAvailableTtsResponseBodyData) GetVoiceType() *string {
	return s.VoiceType
}

func (s *ListAvailableTtsResponseBodyData) SetTtsAuditionFileUrl(v string) *ListAvailableTtsResponseBodyData {
	s.TtsAuditionFileUrl = &v
	return s
}

func (s *ListAvailableTtsResponseBodyData) SetTtsEngine(v string) *ListAvailableTtsResponseBodyData {
	s.TtsEngine = &v
	return s
}

func (s *ListAvailableTtsResponseBodyData) SetTtsStyle(v string) *ListAvailableTtsResponseBodyData {
	s.TtsStyle = &v
	return s
}

func (s *ListAvailableTtsResponseBodyData) SetTtsVoiceCode(v string) *ListAvailableTtsResponseBodyData {
	s.TtsVoiceCode = &v
	return s
}

func (s *ListAvailableTtsResponseBodyData) SetTtsVoiceName(v string) *ListAvailableTtsResponseBodyData {
	s.TtsVoiceName = &v
	return s
}

func (s *ListAvailableTtsResponseBodyData) SetVoiceType(v string) *ListAvailableTtsResponseBodyData {
	s.VoiceType = &v
	return s
}

func (s *ListAvailableTtsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
