// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDataDomainResponseBody
	GetCode() *string
	SetCreateResult(v *CreateDataDomainResponseBodyCreateResult) *CreateDataDomainResponseBody
	GetCreateResult() *CreateDataDomainResponseBodyCreateResult
	SetHttpStatusCode(v int32) *CreateDataDomainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDataDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDataDomainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataDomainResponseBody
	GetSuccess() *bool
}

type CreateDataDomainResponseBody struct {
	// example:
	//
	// OK
	Code         *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateResult *CreateDataDomainResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
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

func (s CreateDataDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDataDomainResponseBody) GetCreateResult() *CreateDataDomainResponseBodyCreateResult {
	return s.CreateResult
}

func (s *CreateDataDomainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDataDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDataDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataDomainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataDomainResponseBody) SetCode(v string) *CreateDataDomainResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDataDomainResponseBody) SetCreateResult(v *CreateDataDomainResponseBodyCreateResult) *CreateDataDomainResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateDataDomainResponseBody) SetHttpStatusCode(v int32) *CreateDataDomainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDataDomainResponseBody) SetMessage(v string) *CreateDataDomainResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDataDomainResponseBody) SetRequestId(v string) *CreateDataDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataDomainResponseBody) SetSuccess(v bool) *CreateDataDomainResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataDomainResponseBody) Validate() error {
	if s.CreateResult != nil {
		if err := s.CreateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataDomainResponseBodyCreateResult struct {
	// example:
	//
	// 1241844456
	DataDomainId *int64 `json:"DataDomainId,omitempty" xml:"DataDomainId,omitempty"`
}

func (s CreateDataDomainResponseBodyCreateResult) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDomainResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateDataDomainResponseBodyCreateResult) GetDataDomainId() *int64 {
	return s.DataDomainId
}

func (s *CreateDataDomainResponseBodyCreateResult) SetDataDomainId(v int64) *CreateDataDomainResponseBodyCreateResult {
	s.DataDomainId = &v
	return s
}

func (s *CreateDataDomainResponseBodyCreateResult) Validate() error {
	return dara.Validate(s)
}
