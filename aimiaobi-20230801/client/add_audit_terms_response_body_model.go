// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuditTermsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddAuditTermsResponseBody
	GetCode() *string
	SetData(v bool) *AddAuditTermsResponseBody
	GetData() *bool
	SetDataV1(v *AddAuditTermsResponseBodyDataV1) *AddAuditTermsResponseBody
	GetDataV1() *AddAuditTermsResponseBodyDataV1
	SetHttpStatusCode(v int32) *AddAuditTermsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddAuditTermsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddAuditTermsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAuditTermsResponseBody
	GetSuccess() *bool
}

type AddAuditTermsResponseBody struct {
	// Status code
	//
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Business data (whether the update succeeded). This field is deprecated. Use DataV1 to get the primary key ID instead.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// ID of the added dictionary term
	//
	// example:
	//
	// 返回添加的实体信息
	DataV1 *AddAuditTermsResponseBodyDataV1 `json:"DataV1,omitempty" xml:"DataV1,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAuditTermsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAuditTermsResponseBody) GoString() string {
	return s.String()
}

func (s *AddAuditTermsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddAuditTermsResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddAuditTermsResponseBody) GetDataV1() *AddAuditTermsResponseBodyDataV1 {
	return s.DataV1
}

func (s *AddAuditTermsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddAuditTermsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAuditTermsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAuditTermsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAuditTermsResponseBody) SetCode(v string) *AddAuditTermsResponseBody {
	s.Code = &v
	return s
}

func (s *AddAuditTermsResponseBody) SetData(v bool) *AddAuditTermsResponseBody {
	s.Data = &v
	return s
}

func (s *AddAuditTermsResponseBody) SetDataV1(v *AddAuditTermsResponseBodyDataV1) *AddAuditTermsResponseBody {
	s.DataV1 = v
	return s
}

func (s *AddAuditTermsResponseBody) SetHttpStatusCode(v int32) *AddAuditTermsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddAuditTermsResponseBody) SetMessage(v string) *AddAuditTermsResponseBody {
	s.Message = &v
	return s
}

func (s *AddAuditTermsResponseBody) SetRequestId(v string) *AddAuditTermsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAuditTermsResponseBody) SetSuccess(v bool) *AddAuditTermsResponseBody {
	s.Success = &v
	return s
}

func (s *AddAuditTermsResponseBody) Validate() error {
	if s.DataV1 != nil {
		if err := s.DataV1.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAuditTermsResponseBodyDataV1 struct {
	// ID
	//
	// example:
	//
	// 562fe4163a59d7bcb44bfdde4e3d5046
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddAuditTermsResponseBodyDataV1) String() string {
	return dara.Prettify(s)
}

func (s AddAuditTermsResponseBodyDataV1) GoString() string {
	return s.String()
}

func (s *AddAuditTermsResponseBodyDataV1) GetId() *int64 {
	return s.Id
}

func (s *AddAuditTermsResponseBodyDataV1) SetId(v int64) *AddAuditTermsResponseBodyDataV1 {
	s.Id = &v
	return s
}

func (s *AddAuditTermsResponseBodyDataV1) Validate() error {
	return dara.Validate(s)
}
