// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResource4ModifyRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetResource4ModifyRecordResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetResource4ModifyRecordResponseBody
	GetCode() *string
	SetData(v []*GetResource4ModifyRecordResponseBodyData) *GetResource4ModifyRecordResponseBody
	GetData() []*GetResource4ModifyRecordResponseBodyData
	SetMessage(v string) *GetResource4ModifyRecordResponseBody
	GetMessage() *string
	SetNextToken(v string) *GetResource4ModifyRecordResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetResource4ModifyRecordResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *GetResource4ModifyRecordResponseBody
	GetTotalCount() *string
}

type GetResource4ModifyRecordResponseBody struct {
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
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetResource4ModifyRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetResource4ModifyRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResource4ModifyRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetResource4ModifyRecordResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetResource4ModifyRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResource4ModifyRecordResponseBody) GetData() []*GetResource4ModifyRecordResponseBodyData {
	return s.Data
}

func (s *GetResource4ModifyRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResource4ModifyRecordResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetResource4ModifyRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResource4ModifyRecordResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *GetResource4ModifyRecordResponseBody) SetAccessDeniedDetail(v string) *GetResource4ModifyRecordResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetCode(v string) *GetResource4ModifyRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetData(v []*GetResource4ModifyRecordResponseBodyData) *GetResource4ModifyRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetMessage(v string) *GetResource4ModifyRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetNextToken(v string) *GetResource4ModifyRecordResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetRequestId(v string) *GetResource4ModifyRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetTotalCount(v string) *GetResource4ModifyRecordResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResource4ModifyRecordResponseBodyData struct {
	// example:
	//
	// {\\"InstanceId\\": \\"\\", \\"AttributeName\\": \\"drmCommand\\", \\"Id\\": 16800, \\"Desc\\": \\"test\\"}
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	Error     *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 1726645341000
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// rm-uf6308dyal1*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// Finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// rds
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetResource4ModifyRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResource4ModifyRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResource4ModifyRecordResponseBodyData) GetAttribute() *string {
	return s.Attribute
}

func (s *GetResource4ModifyRecordResponseBodyData) GetError() *string {
	return s.Error
}

func (s *GetResource4ModifyRecordResponseBodyData) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetResource4ModifyRecordResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetResource4ModifyRecordResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetResource4ModifyRecordResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetResource4ModifyRecordResponseBodyData) SetAttribute(v string) *GetResource4ModifyRecordResponseBodyData {
	s.Attribute = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBodyData) SetError(v string) *GetResource4ModifyRecordResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBodyData) SetModifyTime(v string) *GetResource4ModifyRecordResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBodyData) SetResourceId(v string) *GetResource4ModifyRecordResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBodyData) SetStatus(v string) *GetResource4ModifyRecordResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBodyData) SetType(v string) *GetResource4ModifyRecordResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
