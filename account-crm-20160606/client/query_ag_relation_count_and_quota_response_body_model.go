// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgRelationCountAndQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAgRelationCountAndQuotaResponseBody
	GetCode() *string
	SetData(v *QueryAgRelationCountAndQuotaResponseBodyData) *QueryAgRelationCountAndQuotaResponseBody
	GetData() *QueryAgRelationCountAndQuotaResponseBodyData
	SetHttpCode(v string) *QueryAgRelationCountAndQuotaResponseBody
	GetHttpCode() *string
	SetMessage(v string) *QueryAgRelationCountAndQuotaResponseBody
	GetMessage() *string
	SetNullObject(v bool) *QueryAgRelationCountAndQuotaResponseBody
	GetNullObject() *bool
	SetRequestId(v string) *QueryAgRelationCountAndQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAgRelationCountAndQuotaResponseBody
	GetSuccess() *bool
}

type QueryAgRelationCountAndQuotaResponseBody struct {
	Code       *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *QueryAgRelationCountAndQuotaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpCode   *string                                       `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message    *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	NullObject *bool                                         `json:"NullObject,omitempty" xml:"NullObject,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAgRelationCountAndQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAgRelationCountAndQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAgRelationCountAndQuotaResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAgRelationCountAndQuotaResponseBody) GetData() *QueryAgRelationCountAndQuotaResponseBodyData {
	return s.Data
}

func (s *QueryAgRelationCountAndQuotaResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *QueryAgRelationCountAndQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAgRelationCountAndQuotaResponseBody) GetNullObject() *bool {
	return s.NullObject
}

func (s *QueryAgRelationCountAndQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAgRelationCountAndQuotaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAgRelationCountAndQuotaResponseBody) SetCode(v string) *QueryAgRelationCountAndQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponseBody) SetData(v *QueryAgRelationCountAndQuotaResponseBodyData) *QueryAgRelationCountAndQuotaResponseBody {
	s.Data = v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponseBody) SetHttpCode(v string) *QueryAgRelationCountAndQuotaResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponseBody) SetMessage(v string) *QueryAgRelationCountAndQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponseBody) SetNullObject(v bool) *QueryAgRelationCountAndQuotaResponseBody {
	s.NullObject = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponseBody) SetRequestId(v string) *QueryAgRelationCountAndQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponseBody) SetSuccess(v bool) *QueryAgRelationCountAndQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAgRelationCountAndQuotaResponseBodyData struct {
	AccountCount *int64  `json:"AccountCount,omitempty" xml:"AccountCount,omitempty"`
	Mpk          *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	NullObject   *bool   `json:"NullObject,omitempty" xml:"NullObject,omitempty"`
	Quota        *int32  `json:"Quota,omitempty" xml:"Quota,omitempty"`
}

func (s QueryAgRelationCountAndQuotaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAgRelationCountAndQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAgRelationCountAndQuotaResponseBodyData) GetAccountCount() *int64 {
	return s.AccountCount
}

func (s *QueryAgRelationCountAndQuotaResponseBodyData) GetMpk() *string {
	return s.Mpk
}

func (s *QueryAgRelationCountAndQuotaResponseBodyData) GetNullObject() *bool {
	return s.NullObject
}

func (s *QueryAgRelationCountAndQuotaResponseBodyData) GetQuota() *int32 {
	return s.Quota
}

func (s *QueryAgRelationCountAndQuotaResponseBodyData) SetAccountCount(v int64) *QueryAgRelationCountAndQuotaResponseBodyData {
	s.AccountCount = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponseBodyData) SetMpk(v string) *QueryAgRelationCountAndQuotaResponseBodyData {
	s.Mpk = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponseBodyData) SetNullObject(v bool) *QueryAgRelationCountAndQuotaResponseBodyData {
	s.NullObject = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponseBodyData) SetQuota(v int32) *QueryAgRelationCountAndQuotaResponseBodyData {
	s.Quota = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
