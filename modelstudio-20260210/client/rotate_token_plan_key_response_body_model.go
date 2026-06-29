// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRotateTokenPlanKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RotateTokenPlanKeyResponseBody
	GetCode() *string
	SetData(v *RotateTokenPlanKeyResponseBodyData) *RotateTokenPlanKeyResponseBody
	GetData() *RotateTokenPlanKeyResponseBodyData
	SetMessage(v string) *RotateTokenPlanKeyResponseBody
	GetMessage() *string
	SetSuccess(v bool) *RotateTokenPlanKeyResponseBody
	GetSuccess() *bool
}

type RotateTokenPlanKeyResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *RotateTokenPlanKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RotateTokenPlanKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RotateTokenPlanKeyResponseBody) GoString() string {
	return s.String()
}

func (s *RotateTokenPlanKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *RotateTokenPlanKeyResponseBody) GetData() *RotateTokenPlanKeyResponseBodyData {
	return s.Data
}

func (s *RotateTokenPlanKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RotateTokenPlanKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RotateTokenPlanKeyResponseBody) SetCode(v string) *RotateTokenPlanKeyResponseBody {
	s.Code = &v
	return s
}

func (s *RotateTokenPlanKeyResponseBody) SetData(v *RotateTokenPlanKeyResponseBodyData) *RotateTokenPlanKeyResponseBody {
	s.Data = v
	return s
}

func (s *RotateTokenPlanKeyResponseBody) SetMessage(v string) *RotateTokenPlanKeyResponseBody {
	s.Message = &v
	return s
}

func (s *RotateTokenPlanKeyResponseBody) SetSuccess(v bool) *RotateTokenPlanKeyResponseBody {
	s.Success = &v
	return s
}

func (s *RotateTokenPlanKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RotateTokenPlanKeyResponseBodyData struct {
	// The API Key ID, which is system-generated.
	//
	// example:
	//
	// ak_123456
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// The new masked_api_key returned by BaiLian, such as sk_***cdef.
	//
	// example:
	//
	// sk_***6
	MaskedApiKey *string `json:"MaskedApiKey,omitempty" xml:"MaskedApiKey,omitempty"`
	// The new plaintext API Key returned by BaiLian. This value is returned only once during the reset operation.
	//
	// example:
	//
	// sk-ws.abc123456
	PlainApiKey *string `json:"PlainApiKey,omitempty" xml:"PlainApiKey,omitempty"`
	// The time when the API key was reset.
	//
	// example:
	//
	// 2025-07-18T03:19:17Z
	ResetAt *string `json:"ResetAt,omitempty" xml:"ResetAt,omitempty"`
	// The source_id returned by BaiLian.
	//
	// example:
	//
	// 123456
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s RotateTokenPlanKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RotateTokenPlanKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *RotateTokenPlanKeyResponseBodyData) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *RotateTokenPlanKeyResponseBodyData) GetMaskedApiKey() *string {
	return s.MaskedApiKey
}

func (s *RotateTokenPlanKeyResponseBodyData) GetPlainApiKey() *string {
	return s.PlainApiKey
}

func (s *RotateTokenPlanKeyResponseBodyData) GetResetAt() *string {
	return s.ResetAt
}

func (s *RotateTokenPlanKeyResponseBodyData) GetSourceId() *string {
	return s.SourceId
}

func (s *RotateTokenPlanKeyResponseBodyData) SetApiKeyId(v string) *RotateTokenPlanKeyResponseBodyData {
	s.ApiKeyId = &v
	return s
}

func (s *RotateTokenPlanKeyResponseBodyData) SetMaskedApiKey(v string) *RotateTokenPlanKeyResponseBodyData {
	s.MaskedApiKey = &v
	return s
}

func (s *RotateTokenPlanKeyResponseBodyData) SetPlainApiKey(v string) *RotateTokenPlanKeyResponseBodyData {
	s.PlainApiKey = &v
	return s
}

func (s *RotateTokenPlanKeyResponseBodyData) SetResetAt(v string) *RotateTokenPlanKeyResponseBodyData {
	s.ResetAt = &v
	return s
}

func (s *RotateTokenPlanKeyResponseBodyData) SetSourceId(v string) *RotateTokenPlanKeyResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *RotateTokenPlanKeyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
