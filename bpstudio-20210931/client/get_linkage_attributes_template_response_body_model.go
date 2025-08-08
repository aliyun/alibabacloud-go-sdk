// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLinkageAttributesTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetLinkageAttributesTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetLinkageAttributesTemplateResponseBody
	GetCode() *string
	SetData(v *GetLinkageAttributesTemplateResponseBodyData) *GetLinkageAttributesTemplateResponseBody
	GetData() *GetLinkageAttributesTemplateResponseBodyData
	SetMessage(v string) *GetLinkageAttributesTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLinkageAttributesTemplateResponseBody
	GetRequestId() *string
}

type GetLinkageAttributesTemplateResponseBody struct {
	// example:
	//
	// {
	//
	//     "PolicyType": "",
	//
	//     "AuthPrincipalOwnerId": "",
	//
	//     "EncodedDiagnosticMessage": "",
	//
	//     "AuthPrincipalType": "",
	//
	//     "AuthPrincipalDisplayName": "",
	//
	//     "NoPermissionType": "",
	//
	//     "AuthAction": ""
	//
	//   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetLinkageAttributesTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLinkageAttributesTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLinkageAttributesTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetLinkageAttributesTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetLinkageAttributesTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLinkageAttributesTemplateResponseBody) GetData() *GetLinkageAttributesTemplateResponseBodyData {
	return s.Data
}

func (s *GetLinkageAttributesTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLinkageAttributesTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLinkageAttributesTemplateResponseBody) SetAccessDeniedDetail(v string) *GetLinkageAttributesTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLinkageAttributesTemplateResponseBody) SetCode(v string) *GetLinkageAttributesTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetLinkageAttributesTemplateResponseBody) SetData(v *GetLinkageAttributesTemplateResponseBodyData) *GetLinkageAttributesTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetLinkageAttributesTemplateResponseBody) SetMessage(v string) *GetLinkageAttributesTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetLinkageAttributesTemplateResponseBody) SetRequestId(v string) *GetLinkageAttributesTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLinkageAttributesTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLinkageAttributesTemplateResponseBodyData struct {
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// SyntaxError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// [
	//
	//       {
	//
	//         "label": "ecs.c6.2xlarge (8c 16g)",
	//
	//         "value": "ecs.c6.2xlarge"
	//
	//       },
	//
	//       {
	//
	//         "label": "ecs.c6.3xlarge (12c 24g)",
	//
	//         "value": "ecs.c6.3xlarge"
	//
	//       }
	//
	// ]
	OptionalValues *string `json:"OptionalValues,omitempty" xml:"OptionalValues,omitempty"`
	// example:
	//
	// 19
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetLinkageAttributesTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLinkageAttributesTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLinkageAttributesTemplateResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetLinkageAttributesTemplateResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetLinkageAttributesTemplateResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *GetLinkageAttributesTemplateResponseBodyData) GetOptionalValues() *string {
	return s.OptionalValues
}

func (s *GetLinkageAttributesTemplateResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetLinkageAttributesTemplateResponseBodyData) SetErrorCode(v string) *GetLinkageAttributesTemplateResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetLinkageAttributesTemplateResponseBodyData) SetErrorMessage(v string) *GetLinkageAttributesTemplateResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetLinkageAttributesTemplateResponseBodyData) SetNextToken(v string) *GetLinkageAttributesTemplateResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *GetLinkageAttributesTemplateResponseBodyData) SetOptionalValues(v string) *GetLinkageAttributesTemplateResponseBodyData {
	s.OptionalValues = &v
	return s
}

func (s *GetLinkageAttributesTemplateResponseBodyData) SetTotalCount(v int64) *GetLinkageAttributesTemplateResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetLinkageAttributesTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
