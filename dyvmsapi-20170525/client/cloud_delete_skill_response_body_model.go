// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudDeleteSkillResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudDeleteSkillResponseBody
	GetCode() *string
	SetData(v *CloudDeleteSkillResponseBodyData) *CloudDeleteSkillResponseBody
	GetData() *CloudDeleteSkillResponseBodyData
	SetMessage(v string) *CloudDeleteSkillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudDeleteSkillResponseBody
	GetRequestId() *string
}

type CloudDeleteSkillResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudDeleteSkillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ED18A5AE-9BBC-5851-1111-8E5B8C23CEDE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudDeleteSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CloudDeleteSkillResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudDeleteSkillResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudDeleteSkillResponseBody) GetData() *CloudDeleteSkillResponseBodyData {
	return s.Data
}

func (s *CloudDeleteSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudDeleteSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudDeleteSkillResponseBody) SetAccessDeniedDetail(v string) *CloudDeleteSkillResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudDeleteSkillResponseBody) SetCode(v string) *CloudDeleteSkillResponseBody {
	s.Code = &v
	return s
}

func (s *CloudDeleteSkillResponseBody) SetData(v *CloudDeleteSkillResponseBodyData) *CloudDeleteSkillResponseBody {
	s.Data = v
	return s
}

func (s *CloudDeleteSkillResponseBody) SetMessage(v string) *CloudDeleteSkillResponseBody {
	s.Message = &v
	return s
}

func (s *CloudDeleteSkillResponseBody) SetRequestId(v string) *CloudDeleteSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudDeleteSkillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudDeleteSkillResponseBodyData struct {
	// 0-成功，-1 失败
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudDeleteSkillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteSkillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudDeleteSkillResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudDeleteSkillResponseBodyData) SetResult(v int64) *CloudDeleteSkillResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudDeleteSkillResponseBodyData) Validate() error {
	return dara.Validate(s)
}
