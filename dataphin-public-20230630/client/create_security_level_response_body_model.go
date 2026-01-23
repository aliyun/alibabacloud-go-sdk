// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSecurityLevelResponseBody
	GetCode() *string
	SetData(v int64) *CreateSecurityLevelResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateSecurityLevelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateSecurityLevelResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSecurityLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSecurityLevelResponseBody
	GetSuccess() *bool
}

type CreateSecurityLevelResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSecurityLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityLevelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityLevelResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSecurityLevelResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateSecurityLevelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateSecurityLevelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSecurityLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecurityLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSecurityLevelResponseBody) SetCode(v string) *CreateSecurityLevelResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSecurityLevelResponseBody) SetData(v int64) *CreateSecurityLevelResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSecurityLevelResponseBody) SetHttpStatusCode(v int32) *CreateSecurityLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSecurityLevelResponseBody) SetMessage(v string) *CreateSecurityLevelResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSecurityLevelResponseBody) SetRequestId(v string) *CreateSecurityLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityLevelResponseBody) SetSuccess(v bool) *CreateSecurityLevelResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSecurityLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
