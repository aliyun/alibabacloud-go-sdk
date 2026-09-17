// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialRecognitionIntlV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CredentialRecognitionIntlV2ResponseBody
	GetCode() *string
	SetMessage(v string) *CredentialRecognitionIntlV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *CredentialRecognitionIntlV2ResponseBody
	GetRequestId() *string
	SetResult(v *CredentialRecognitionIntlV2ResponseBodyResult) *CredentialRecognitionIntlV2ResponseBody
	GetResult() *CredentialRecognitionIntlV2ResponseBodyResult
}

type CredentialRecognitionIntlV2ResponseBody struct {
	// The return code. A value of 200 indicates a successful request. Other values indicate failures.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7F971622-38C0-5F56-B2EC-315367979B4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response result.
	Result *CredentialRecognitionIntlV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CredentialRecognitionIntlV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CredentialRecognitionIntlV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CredentialRecognitionIntlV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *CredentialRecognitionIntlV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CredentialRecognitionIntlV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CredentialRecognitionIntlV2ResponseBody) GetResult() *CredentialRecognitionIntlV2ResponseBodyResult {
	return s.Result
}

func (s *CredentialRecognitionIntlV2ResponseBody) SetCode(v string) *CredentialRecognitionIntlV2ResponseBody {
	s.Code = &v
	return s
}

func (s *CredentialRecognitionIntlV2ResponseBody) SetMessage(v string) *CredentialRecognitionIntlV2ResponseBody {
	s.Message = &v
	return s
}

func (s *CredentialRecognitionIntlV2ResponseBody) SetRequestId(v string) *CredentialRecognitionIntlV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CredentialRecognitionIntlV2ResponseBody) SetResult(v *CredentialRecognitionIntlV2ResponseBodyResult) *CredentialRecognitionIntlV2ResponseBody {
	s.Result = v
	return s
}

func (s *CredentialRecognitionIntlV2ResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CredentialRecognitionIntlV2ResponseBodyResult struct {
	// The recognized key information, in JSON format.
	//
	// example:
	//
	// {
	//
	//   "address": "",
	//
	//   "name":""
	//
	// }
	ExtIdInfo *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	// The result code. Valid values:
	//
	// - 200: OCR extraction succeeded and all rule checks passed.
	//
	// - 204: Validation result is inconsistent. OCR extraction succeeded, but some fields in CheckRuleConfig did not pass (N).
	//
	// - 211: Quality does not meet requirements. Quality detection did not pass when idQuality is set to Y (not yet supported in the current version).
	//
	// - 212: Anti-forgery check did not pass. fraudCheck was triggered and anti-forgery verification failed.
	//
	// - 213: No text was extracted, or the credential type check did not pass.
	//
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// The extraction result. Valid values:
	//
	// - S: Succeeded.
	//
	// - F: Failed.
	//
	// example:
	//
	// S
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CredentialRecognitionIntlV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CredentialRecognitionIntlV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CredentialRecognitionIntlV2ResponseBodyResult) GetExtIdInfo() *string {
	return s.ExtIdInfo
}

func (s *CredentialRecognitionIntlV2ResponseBodyResult) GetSubCode() *string {
	return s.SubCode
}

func (s *CredentialRecognitionIntlV2ResponseBodyResult) GetSuccess() *string {
	return s.Success
}

func (s *CredentialRecognitionIntlV2ResponseBodyResult) SetExtIdInfo(v string) *CredentialRecognitionIntlV2ResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *CredentialRecognitionIntlV2ResponseBodyResult) SetSubCode(v string) *CredentialRecognitionIntlV2ResponseBodyResult {
	s.SubCode = &v
	return s
}

func (s *CredentialRecognitionIntlV2ResponseBodyResult) SetSuccess(v string) *CredentialRecognitionIntlV2ResponseBodyResult {
	s.Success = &v
	return s
}

func (s *CredentialRecognitionIntlV2ResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
