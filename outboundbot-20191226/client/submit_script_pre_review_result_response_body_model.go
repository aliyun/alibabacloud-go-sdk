// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitScriptPreReviewResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitScriptPreReviewResultResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SubmitScriptPreReviewResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitScriptPreReviewResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitScriptPreReviewResultResponseBody
	GetRequestId() *string
	SetScript(v *SubmitScriptPreReviewResultResponseBodyScript) *SubmitScriptPreReviewResultResponseBody
	GetScript() *SubmitScriptPreReviewResultResponseBodyScript
	SetSuccess(v bool) *SubmitScriptPreReviewResultResponseBody
	GetSuccess() *bool
}

type SubmitScriptPreReviewResultResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API operation.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The script information for which the pre-review result is submitted.
	Script *SubmitScriptPreReviewResultResponseBodyScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitScriptPreReviewResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitScriptPreReviewResultResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitScriptPreReviewResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitScriptPreReviewResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitScriptPreReviewResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitScriptPreReviewResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitScriptPreReviewResultResponseBody) GetScript() *SubmitScriptPreReviewResultResponseBodyScript {
	return s.Script
}

func (s *SubmitScriptPreReviewResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitScriptPreReviewResultResponseBody) SetCode(v string) *SubmitScriptPreReviewResultResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitScriptPreReviewResultResponseBody) SetHttpStatusCode(v int32) *SubmitScriptPreReviewResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitScriptPreReviewResultResponseBody) SetMessage(v string) *SubmitScriptPreReviewResultResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitScriptPreReviewResultResponseBody) SetRequestId(v string) *SubmitScriptPreReviewResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitScriptPreReviewResultResponseBody) SetScript(v *SubmitScriptPreReviewResultResponseBodyScript) *SubmitScriptPreReviewResultResponseBody {
	s.Script = v
	return s
}

func (s *SubmitScriptPreReviewResultResponseBody) SetSuccess(v bool) *SubmitScriptPreReviewResultResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitScriptPreReviewResultResponseBody) Validate() error {
	if s.Script != nil {
		if err := s.Script.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitScriptPreReviewResultResponseBodyScript struct {
	// The script description.
	//
	// example:
	//
	// The script used for outbound calls to notify members of campaign information
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// The script ID.
	//
	// example:
	//
	// 0f6b8f2c-3e14-4a89-9d5e-1f2a3b4c5d6e
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The script name.
	//
	// example:
	//
	// Member notification script
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
}

func (s SubmitScriptPreReviewResultResponseBodyScript) String() string {
	return dara.Prettify(s)
}

func (s SubmitScriptPreReviewResultResponseBodyScript) GoString() string {
	return s.String()
}

func (s *SubmitScriptPreReviewResultResponseBodyScript) GetScriptDescription() *string {
	return s.ScriptDescription
}

func (s *SubmitScriptPreReviewResultResponseBodyScript) GetScriptId() *string {
	return s.ScriptId
}

func (s *SubmitScriptPreReviewResultResponseBodyScript) GetScriptName() *string {
	return s.ScriptName
}

func (s *SubmitScriptPreReviewResultResponseBodyScript) SetScriptDescription(v string) *SubmitScriptPreReviewResultResponseBodyScript {
	s.ScriptDescription = &v
	return s
}

func (s *SubmitScriptPreReviewResultResponseBodyScript) SetScriptId(v string) *SubmitScriptPreReviewResultResponseBodyScript {
	s.ScriptId = &v
	return s
}

func (s *SubmitScriptPreReviewResultResponseBodyScript) SetScriptName(v string) *SubmitScriptPreReviewResultResponseBodyScript {
	s.ScriptName = &v
	return s
}

func (s *SubmitScriptPreReviewResultResponseBodyScript) Validate() error {
	return dara.Validate(s)
}
