// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileDeepAnalyzeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateFileDeepAnalyzeResponseBody
	GetCode() *string
	SetData(v string) *CreateFileDeepAnalyzeResponseBody
	GetData() *string
	SetMessage(v string) *CreateFileDeepAnalyzeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFileDeepAnalyzeResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateFileDeepAnalyzeResponseBody
	GetStatus() *string
}

type CreateFileDeepAnalyzeResponseBody struct {
	// The status code.
	//
	// example:
	//
	// FileNotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The analysis result.
	//
	// example:
	//
	// {
	//
	//   "result": 0,
	//
	//   "file_analysis": {
	//
	//     "summary": {
	//
	//       "ai_interpretation": "The file is a script for the JetBrains PhpStorm stub generation tool, which is a legitimate developer tool. The code uses the PhpParser library to parse PHP source code and generate the PhpStormStubsMap.php mapping file. The engine analysis did not detect any attempts to invoke dangerous functions, external input processing, code obfuscation, or malicious behavior characteristics, and the file is determined to be normal."
	//
	//     }
	//
	//   }
	//
	// }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The stage description.
	//
	// example:
	//
	// File hash parsed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7F9C0A1-2B3D-4E5F-8A90-123456XXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stage status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateFileDeepAnalyzeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDeepAnalyzeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileDeepAnalyzeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFileDeepAnalyzeResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateFileDeepAnalyzeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFileDeepAnalyzeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileDeepAnalyzeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateFileDeepAnalyzeResponseBody) SetCode(v string) *CreateFileDeepAnalyzeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFileDeepAnalyzeResponseBody) SetData(v string) *CreateFileDeepAnalyzeResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFileDeepAnalyzeResponseBody) SetMessage(v string) *CreateFileDeepAnalyzeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFileDeepAnalyzeResponseBody) SetRequestId(v string) *CreateFileDeepAnalyzeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileDeepAnalyzeResponseBody) SetStatus(v string) *CreateFileDeepAnalyzeResponseBody {
	s.Status = &v
	return s
}

func (s *CreateFileDeepAnalyzeResponseBody) Validate() error {
	return dara.Validate(s)
}
