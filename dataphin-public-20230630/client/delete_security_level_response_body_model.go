// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSecurityLevelResponseBody
	GetCode() *string
	SetDeleteResult(v *DeleteSecurityLevelResponseBodyDeleteResult) *DeleteSecurityLevelResponseBody
	GetDeleteResult() *DeleteSecurityLevelResponseBodyDeleteResult
	SetHttpStatusCode(v int32) *DeleteSecurityLevelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteSecurityLevelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSecurityLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSecurityLevelResponseBody
	GetSuccess() *bool
}

type DeleteSecurityLevelResponseBody struct {
	// example:
	//
	// OK
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	DeleteResult *DeleteSecurityLevelResponseBodyDeleteResult `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty" type:"Struct"`
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

func (s DeleteSecurityLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityLevelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityLevelResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSecurityLevelResponseBody) GetDeleteResult() *DeleteSecurityLevelResponseBodyDeleteResult {
	return s.DeleteResult
}

func (s *DeleteSecurityLevelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteSecurityLevelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSecurityLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSecurityLevelResponseBody) SetCode(v string) *DeleteSecurityLevelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecurityLevelResponseBody) SetDeleteResult(v *DeleteSecurityLevelResponseBodyDeleteResult) *DeleteSecurityLevelResponseBody {
	s.DeleteResult = v
	return s
}

func (s *DeleteSecurityLevelResponseBody) SetHttpStatusCode(v int32) *DeleteSecurityLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSecurityLevelResponseBody) SetMessage(v string) *DeleteSecurityLevelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecurityLevelResponseBody) SetRequestId(v string) *DeleteSecurityLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityLevelResponseBody) SetSuccess(v bool) *DeleteSecurityLevelResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSecurityLevelResponseBody) Validate() error {
	if s.DeleteResult != nil {
		if err := s.DeleteResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSecurityLevelResponseBodyDeleteResult struct {
	// example:
	//
	// E10012011
	ErrorCode             *string  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RelatedClassifyIdList []*int64 `json:"RelatedClassifyIdList,omitempty" xml:"RelatedClassifyIdList,omitempty" type:"Repeated"`
	Success               *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSecurityLevelResponseBodyDeleteResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityLevelResponseBodyDeleteResult) GoString() string {
	return s.String()
}

func (s *DeleteSecurityLevelResponseBodyDeleteResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteSecurityLevelResponseBodyDeleteResult) GetRelatedClassifyIdList() []*int64 {
	return s.RelatedClassifyIdList
}

func (s *DeleteSecurityLevelResponseBodyDeleteResult) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSecurityLevelResponseBodyDeleteResult) SetErrorCode(v string) *DeleteSecurityLevelResponseBodyDeleteResult {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSecurityLevelResponseBodyDeleteResult) SetRelatedClassifyIdList(v []*int64) *DeleteSecurityLevelResponseBodyDeleteResult {
	s.RelatedClassifyIdList = v
	return s
}

func (s *DeleteSecurityLevelResponseBodyDeleteResult) SetSuccess(v bool) *DeleteSecurityLevelResponseBodyDeleteResult {
	s.Success = &v
	return s
}

func (s *DeleteSecurityLevelResponseBodyDeleteResult) Validate() error {
	return dara.Validate(s)
}
