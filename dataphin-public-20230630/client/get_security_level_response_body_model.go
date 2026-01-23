// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecurityLevelResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetSecurityLevelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSecurityLevelResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecurityLevelResponseBody
	GetRequestId() *string
	SetSecurityLevelInfo(v *GetSecurityLevelResponseBodySecurityLevelInfo) *GetSecurityLevelResponseBody
	GetSecurityLevelInfo() *GetSecurityLevelResponseBodySecurityLevelInfo
	SetSuccess(v bool) *GetSecurityLevelResponseBody
	GetSuccess() *bool
}

type GetSecurityLevelResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityLevelInfo *GetSecurityLevelResponseBodySecurityLevelInfo `json:"SecurityLevelInfo,omitempty" xml:"SecurityLevelInfo,omitempty" type:"Struct"`
	Success           *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSecurityLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityLevelResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityLevelResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecurityLevelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSecurityLevelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecurityLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecurityLevelResponseBody) GetSecurityLevelInfo() *GetSecurityLevelResponseBodySecurityLevelInfo {
	return s.SecurityLevelInfo
}

func (s *GetSecurityLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSecurityLevelResponseBody) SetCode(v string) *GetSecurityLevelResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecurityLevelResponseBody) SetHttpStatusCode(v int32) *GetSecurityLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSecurityLevelResponseBody) SetMessage(v string) *GetSecurityLevelResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecurityLevelResponseBody) SetRequestId(v string) *GetSecurityLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityLevelResponseBody) SetSecurityLevelInfo(v *GetSecurityLevelResponseBodySecurityLevelInfo) *GetSecurityLevelResponseBody {
	s.SecurityLevelInfo = v
	return s
}

func (s *GetSecurityLevelResponseBody) SetSuccess(v bool) *GetSecurityLevelResponseBody {
	s.Success = &v
	return s
}

func (s *GetSecurityLevelResponseBody) Validate() error {
	if s.SecurityLevelInfo != nil {
		if err := s.SecurityLevelInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityLevelResponseBodySecurityLevelInfo struct {
	// example:
	//
	// test
	Abbreviation *string `json:"Abbreviation,omitempty" xml:"Abbreviation,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// test
	Name                  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	RelatedClassifyIdList []*int64 `json:"RelatedClassifyIdList,omitempty" xml:"RelatedClassifyIdList,omitempty" type:"Repeated"`
}

func (s GetSecurityLevelResponseBodySecurityLevelInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityLevelResponseBodySecurityLevelInfo) GoString() string {
	return s.String()
}

func (s *GetSecurityLevelResponseBodySecurityLevelInfo) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *GetSecurityLevelResponseBodySecurityLevelInfo) GetDescription() *string {
	return s.Description
}

func (s *GetSecurityLevelResponseBodySecurityLevelInfo) GetIndex() *int64 {
	return s.Index
}

func (s *GetSecurityLevelResponseBodySecurityLevelInfo) GetName() *string {
	return s.Name
}

func (s *GetSecurityLevelResponseBodySecurityLevelInfo) GetRelatedClassifyIdList() []*int64 {
	return s.RelatedClassifyIdList
}

func (s *GetSecurityLevelResponseBodySecurityLevelInfo) SetAbbreviation(v string) *GetSecurityLevelResponseBodySecurityLevelInfo {
	s.Abbreviation = &v
	return s
}

func (s *GetSecurityLevelResponseBodySecurityLevelInfo) SetDescription(v string) *GetSecurityLevelResponseBodySecurityLevelInfo {
	s.Description = &v
	return s
}

func (s *GetSecurityLevelResponseBodySecurityLevelInfo) SetIndex(v int64) *GetSecurityLevelResponseBodySecurityLevelInfo {
	s.Index = &v
	return s
}

func (s *GetSecurityLevelResponseBodySecurityLevelInfo) SetName(v string) *GetSecurityLevelResponseBodySecurityLevelInfo {
	s.Name = &v
	return s
}

func (s *GetSecurityLevelResponseBodySecurityLevelInfo) SetRelatedClassifyIdList(v []*int64) *GetSecurityLevelResponseBodySecurityLevelInfo {
	s.RelatedClassifyIdList = v
	return s
}

func (s *GetSecurityLevelResponseBodySecurityLevelInfo) Validate() error {
	return dara.Validate(s)
}
