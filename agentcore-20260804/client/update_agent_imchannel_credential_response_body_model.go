// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentIMChannelCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAgentIMChannelCredentialResponseBody
	GetCode() *string
	SetData(v *UpdateAgentIMChannelCredentialResponseBodyData) *UpdateAgentIMChannelCredentialResponseBody
	GetData() *UpdateAgentIMChannelCredentialResponseBodyData
	SetHttpStatusCode(v int32) *UpdateAgentIMChannelCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateAgentIMChannelCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAgentIMChannelCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAgentIMChannelCredentialResponseBody
	GetSuccess() *bool
}

type UpdateAgentIMChannelCredentialResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The summary of the updated IM channel credential.
	Data *UpdateAgentIMChannelCredentialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The result message of the request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAgentIMChannelCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAgentIMChannelCredentialResponseBody) GetData() *UpdateAgentIMChannelCredentialResponseBodyData {
	return s.Data
}

func (s *UpdateAgentIMChannelCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateAgentIMChannelCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgentIMChannelCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentIMChannelCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAgentIMChannelCredentialResponseBody) SetCode(v string) *UpdateAgentIMChannelCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentIMChannelCredentialResponseBody) SetData(v *UpdateAgentIMChannelCredentialResponseBodyData) *UpdateAgentIMChannelCredentialResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAgentIMChannelCredentialResponseBody) SetHttpStatusCode(v int32) *UpdateAgentIMChannelCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAgentIMChannelCredentialResponseBody) SetMessage(v string) *UpdateAgentIMChannelCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentIMChannelCredentialResponseBody) SetRequestId(v string) *UpdateAgentIMChannelCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentIMChannelCredentialResponseBody) SetSuccess(v bool) *UpdateAgentIMChannelCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgentIMChannelCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentIMChannelCredentialResponseBodyData struct {
	// The list of configured secret field names. Secret values are not included.
	ConfiguredSecretFields []*string `json:"configuredSecretFields,omitempty" xml:"configuredSecretFields,omitempty" type:"Repeated"`
	// The non-sensitive credential fields and their values.
	NonSecretFields map[string]*string `json:"nonSecretFields,omitempty" xml:"nonSecretFields,omitempty"`
}

func (s UpdateAgentIMChannelCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelCredentialResponseBodyData) GetConfiguredSecretFields() []*string {
	return s.ConfiguredSecretFields
}

func (s *UpdateAgentIMChannelCredentialResponseBodyData) GetNonSecretFields() map[string]*string {
	return s.NonSecretFields
}

func (s *UpdateAgentIMChannelCredentialResponseBodyData) SetConfiguredSecretFields(v []*string) *UpdateAgentIMChannelCredentialResponseBodyData {
	s.ConfiguredSecretFields = v
	return s
}

func (s *UpdateAgentIMChannelCredentialResponseBodyData) SetNonSecretFields(v map[string]*string) *UpdateAgentIMChannelCredentialResponseBodyData {
	s.NonSecretFields = v
	return s
}

func (s *UpdateAgentIMChannelCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
