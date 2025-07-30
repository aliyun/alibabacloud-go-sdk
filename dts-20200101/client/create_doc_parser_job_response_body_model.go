// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocParserJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *CreateDocParserJobResponseBody
	GetDtsJobId() *string
	SetDynamicCode(v string) *CreateDocParserJobResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateDocParserJobResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *CreateDocParserJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateDocParserJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateDocParserJobResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateDocParserJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDocParserJobResponseBody
	GetSuccess() *bool
}

type CreateDocParserJobResponseBody struct {
	// example:
	//
	// dts-20250729-boss6pn1w******
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// present environment is not support, so skip
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// FDC111B1-ACBF-457D-9656-247FDEE9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDocParserJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDocParserJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocParserJobResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *CreateDocParserJobResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateDocParserJobResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateDocParserJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateDocParserJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateDocParserJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDocParserJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDocParserJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDocParserJobResponseBody) SetDtsJobId(v string) *CreateDocParserJobResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetDynamicCode(v string) *CreateDocParserJobResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetDynamicMessage(v string) *CreateDocParserJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetErrCode(v string) *CreateDocParserJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetErrMessage(v string) *CreateDocParserJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetHttpStatusCode(v int32) *CreateDocParserJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetRequestId(v string) *CreateDocParserJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDocParserJobResponseBody) SetSuccess(v bool) *CreateDocParserJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDocParserJobResponseBody) Validate() error {
	return dara.Validate(s)
}
