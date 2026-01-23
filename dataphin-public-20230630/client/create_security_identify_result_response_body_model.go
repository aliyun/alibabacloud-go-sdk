// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityIdentifyResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSecurityIdentifyResultResponseBody
	GetCode() *string
	SetData(v int64) *CreateSecurityIdentifyResultResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateSecurityIdentifyResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateSecurityIdentifyResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSecurityIdentifyResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSecurityIdentifyResultResponseBody
	GetSuccess() *bool
}

type CreateSecurityIdentifyResultResponseBody struct {
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSecurityIdentifyResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityIdentifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityIdentifyResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSecurityIdentifyResultResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateSecurityIdentifyResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateSecurityIdentifyResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSecurityIdentifyResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecurityIdentifyResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSecurityIdentifyResultResponseBody) SetCode(v string) *CreateSecurityIdentifyResultResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSecurityIdentifyResultResponseBody) SetData(v int64) *CreateSecurityIdentifyResultResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSecurityIdentifyResultResponseBody) SetHttpStatusCode(v int32) *CreateSecurityIdentifyResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSecurityIdentifyResultResponseBody) SetMessage(v string) *CreateSecurityIdentifyResultResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSecurityIdentifyResultResponseBody) SetRequestId(v string) *CreateSecurityIdentifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityIdentifyResultResponseBody) SetSuccess(v bool) *CreateSecurityIdentifyResultResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSecurityIdentifyResultResponseBody) Validate() error {
	return dara.Validate(s)
}
