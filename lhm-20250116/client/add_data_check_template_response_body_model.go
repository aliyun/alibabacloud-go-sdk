// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataCheckTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *AddDataCheckTemplateResponseBody
	GetData() *string
	SetErrCode(v string) *AddDataCheckTemplateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AddDataCheckTemplateResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *AddDataCheckTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDataCheckTemplateResponseBody
	GetSuccess() *bool
}

type AddDataCheckTemplateResponseBody struct {
	// The business data returned by the operation in string format. The specific content varies by operation.
	//
	// example:
	//
	// demo
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues of this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values: true and false. If the call fails, check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddDataCheckTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataCheckTemplateResponseBody) GetData() *string {
	return s.Data
}

func (s *AddDataCheckTemplateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AddDataCheckTemplateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AddDataCheckTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataCheckTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDataCheckTemplateResponseBody) SetData(v string) *AddDataCheckTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *AddDataCheckTemplateResponseBody) SetErrCode(v string) *AddDataCheckTemplateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddDataCheckTemplateResponseBody) SetErrMessage(v string) *AddDataCheckTemplateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AddDataCheckTemplateResponseBody) SetRequestId(v string) *AddDataCheckTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataCheckTemplateResponseBody) SetSuccess(v bool) *AddDataCheckTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *AddDataCheckTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
