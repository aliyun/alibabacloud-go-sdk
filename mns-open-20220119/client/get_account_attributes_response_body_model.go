// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetAccountAttributesResponseBody
	GetCode() *int64
	SetData(v *GetAccountAttributesResponseBodyData) *GetAccountAttributesResponseBody
	GetData() *GetAccountAttributesResponseBodyData
	SetMessage(v string) *GetAccountAttributesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAccountAttributesResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetAccountAttributesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetAccountAttributesResponseBody
	GetSuccess() *bool
}

type GetAccountAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetAccountAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAccountAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountAttributesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetAccountAttributesResponseBody) GetData() *GetAccountAttributesResponseBodyData {
	return s.Data
}

func (s *GetAccountAttributesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAccountAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountAttributesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAccountAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAccountAttributesResponseBody) SetCode(v int64) *GetAccountAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *GetAccountAttributesResponseBody) SetData(v *GetAccountAttributesResponseBodyData) *GetAccountAttributesResponseBody {
	s.Data = v
	return s
}

func (s *GetAccountAttributesResponseBody) SetMessage(v string) *GetAccountAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *GetAccountAttributesResponseBody) SetRequestId(v string) *GetAccountAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountAttributesResponseBody) SetStatus(v string) *GetAccountAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *GetAccountAttributesResponseBody) SetSuccess(v bool) *GetAccountAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *GetAccountAttributesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccountAttributesResponseBodyData struct {
	// Indicates whether the Log Management feature is enabled.
	//
	// example:
	//
	// true
	LogEnabled *bool `json:"LogEnabled,omitempty" xml:"LogEnabled,omitempty"`
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// test-logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// Indicates whether the message trace feature is enabled.
	//
	// example:
	//
	// true
	MessageTraceEnabled *bool `json:"MessageTraceEnabled,omitempty" xml:"MessageTraceEnabled,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetAccountAttributesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAccountAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAccountAttributesResponseBodyData) GetLogEnabled() *bool {
	return s.LogEnabled
}

func (s *GetAccountAttributesResponseBodyData) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *GetAccountAttributesResponseBodyData) GetMessageTraceEnabled() *bool {
	return s.MessageTraceEnabled
}

func (s *GetAccountAttributesResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetAccountAttributesResponseBodyData) SetLogEnabled(v bool) *GetAccountAttributesResponseBodyData {
	s.LogEnabled = &v
	return s
}

func (s *GetAccountAttributesResponseBodyData) SetLogStoreName(v string) *GetAccountAttributesResponseBodyData {
	s.LogStoreName = &v
	return s
}

func (s *GetAccountAttributesResponseBodyData) SetMessageTraceEnabled(v bool) *GetAccountAttributesResponseBodyData {
	s.MessageTraceEnabled = &v
	return s
}

func (s *GetAccountAttributesResponseBodyData) SetProjectName(v string) *GetAccountAttributesResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetAccountAttributesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
