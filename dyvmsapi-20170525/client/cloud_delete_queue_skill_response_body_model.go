// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteQueueSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudDeleteQueueSkillResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudDeleteQueueSkillResponseBody
	GetCode() *string
	SetData(v *CloudDeleteQueueSkillResponseBodyData) *CloudDeleteQueueSkillResponseBody
	GetData() *CloudDeleteQueueSkillResponseBodyData
	SetMessage(v string) *CloudDeleteQueueSkillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudDeleteQueueSkillResponseBody
	GetRequestId() *string
}

type CloudDeleteQueueSkillResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudDeleteQueueSkillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudDeleteQueueSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteQueueSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CloudDeleteQueueSkillResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudDeleteQueueSkillResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudDeleteQueueSkillResponseBody) GetData() *CloudDeleteQueueSkillResponseBodyData {
	return s.Data
}

func (s *CloudDeleteQueueSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudDeleteQueueSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudDeleteQueueSkillResponseBody) SetAccessDeniedDetail(v string) *CloudDeleteQueueSkillResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudDeleteQueueSkillResponseBody) SetCode(v string) *CloudDeleteQueueSkillResponseBody {
	s.Code = &v
	return s
}

func (s *CloudDeleteQueueSkillResponseBody) SetData(v *CloudDeleteQueueSkillResponseBodyData) *CloudDeleteQueueSkillResponseBody {
	s.Data = v
	return s
}

func (s *CloudDeleteQueueSkillResponseBody) SetMessage(v string) *CloudDeleteQueueSkillResponseBody {
	s.Message = &v
	return s
}

func (s *CloudDeleteQueueSkillResponseBody) SetRequestId(v string) *CloudDeleteQueueSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudDeleteQueueSkillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudDeleteQueueSkillResponseBodyData struct {
	// 0-成功，其它-失败
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudDeleteQueueSkillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteQueueSkillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudDeleteQueueSkillResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudDeleteQueueSkillResponseBodyData) SetResult(v int64) *CloudDeleteQueueSkillResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudDeleteQueueSkillResponseBodyData) Validate() error {
	return dara.Validate(s)
}
