// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityClassifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecurityClassifyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetSecurityClassifyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSecurityClassifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecurityClassifyResponseBody
	GetRequestId() *string
	SetSecurityClassifyInfo(v *GetSecurityClassifyResponseBodySecurityClassifyInfo) *GetSecurityClassifyResponseBody
	GetSecurityClassifyInfo() *GetSecurityClassifyResponseBodySecurityClassifyInfo
	SetSuccess(v bool) *GetSecurityClassifyResponseBody
	GetSuccess() *bool
}

type GetSecurityClassifyResponseBody struct {
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
	RequestId            *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityClassifyInfo *GetSecurityClassifyResponseBodySecurityClassifyInfo `json:"SecurityClassifyInfo,omitempty" xml:"SecurityClassifyInfo,omitempty" type:"Struct"`
	Success              *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSecurityClassifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityClassifyResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityClassifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecurityClassifyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSecurityClassifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecurityClassifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecurityClassifyResponseBody) GetSecurityClassifyInfo() *GetSecurityClassifyResponseBodySecurityClassifyInfo {
	return s.SecurityClassifyInfo
}

func (s *GetSecurityClassifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSecurityClassifyResponseBody) SetCode(v string) *GetSecurityClassifyResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecurityClassifyResponseBody) SetHttpStatusCode(v int32) *GetSecurityClassifyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSecurityClassifyResponseBody) SetMessage(v string) *GetSecurityClassifyResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecurityClassifyResponseBody) SetRequestId(v string) *GetSecurityClassifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityClassifyResponseBody) SetSecurityClassifyInfo(v *GetSecurityClassifyResponseBodySecurityClassifyInfo) *GetSecurityClassifyResponseBody {
	s.SecurityClassifyInfo = v
	return s
}

func (s *GetSecurityClassifyResponseBody) SetSuccess(v bool) *GetSecurityClassifyResponseBody {
	s.Success = &v
	return s
}

func (s *GetSecurityClassifyResponseBody) Validate() error {
	if s.SecurityClassifyInfo != nil {
		if err := s.SecurityClassifyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityClassifyResponseBodySecurityClassifyInfo struct {
	// example:
	//
	// test
	Abbreviation *string `json:"Abbreviation,omitempty" xml:"Abbreviation,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// lv_test
	LevelAbbreviation *string `json:"LevelAbbreviation,omitempty" xml:"LevelAbbreviation,omitempty"`
	// example:
	//
	// 1
	LevelIndex *int64 `json:"LevelIndex,omitempty" xml:"LevelIndex,omitempty"`
	// example:
	//
	// lv_teat
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// /a/b/c
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s GetSecurityClassifyResponseBodySecurityClassifyInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityClassifyResponseBodySecurityClassifyInfo) GoString() string {
	return s.String()
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) GetId() *int64 {
	return s.Id
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) GetLevelAbbreviation() *string {
	return s.LevelAbbreviation
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) GetLevelIndex() *int64 {
	return s.LevelIndex
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) GetLevelName() *string {
	return s.LevelName
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) GetName() *string {
	return s.Name
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) GetPath() *string {
	return s.Path
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) SetAbbreviation(v string) *GetSecurityClassifyResponseBodySecurityClassifyInfo {
	s.Abbreviation = &v
	return s
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) SetId(v int64) *GetSecurityClassifyResponseBodySecurityClassifyInfo {
	s.Id = &v
	return s
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) SetLevelAbbreviation(v string) *GetSecurityClassifyResponseBodySecurityClassifyInfo {
	s.LevelAbbreviation = &v
	return s
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) SetLevelIndex(v int64) *GetSecurityClassifyResponseBodySecurityClassifyInfo {
	s.LevelIndex = &v
	return s
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) SetLevelName(v string) *GetSecurityClassifyResponseBodySecurityClassifyInfo {
	s.LevelName = &v
	return s
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) SetName(v string) *GetSecurityClassifyResponseBodySecurityClassifyInfo {
	s.Name = &v
	return s
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) SetPath(v string) *GetSecurityClassifyResponseBodySecurityClassifyInfo {
	s.Path = &v
	return s
}

func (s *GetSecurityClassifyResponseBodySecurityClassifyInfo) Validate() error {
	return dara.Validate(s)
}
