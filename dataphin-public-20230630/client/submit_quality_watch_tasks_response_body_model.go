// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQualityWatchTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitQualityWatchTasksResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SubmitQualityWatchTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitQualityWatchTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitQualityWatchTasksResponseBody
	GetRequestId() *string
	SetSubmitResult(v *SubmitQualityWatchTasksResponseBodySubmitResult) *SubmitQualityWatchTasksResponseBody
	GetSubmitResult() *SubmitQualityWatchTasksResponseBodySubmitResult
	SetSuccess(v bool) *SubmitQualityWatchTasksResponseBody
	GetSuccess() *bool
}

type SubmitQualityWatchTasksResponseBody struct {
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
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubmitResult *SubmitQualityWatchTasksResponseBodySubmitResult `json:"SubmitResult,omitempty" xml:"SubmitResult,omitempty" type:"Struct"`
	Success      *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitQualityWatchTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityWatchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitQualityWatchTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitQualityWatchTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitQualityWatchTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitQualityWatchTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitQualityWatchTasksResponseBody) GetSubmitResult() *SubmitQualityWatchTasksResponseBodySubmitResult {
	return s.SubmitResult
}

func (s *SubmitQualityWatchTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitQualityWatchTasksResponseBody) SetCode(v string) *SubmitQualityWatchTasksResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitQualityWatchTasksResponseBody) SetHttpStatusCode(v int32) *SubmitQualityWatchTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitQualityWatchTasksResponseBody) SetMessage(v string) *SubmitQualityWatchTasksResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitQualityWatchTasksResponseBody) SetRequestId(v string) *SubmitQualityWatchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitQualityWatchTasksResponseBody) SetSubmitResult(v *SubmitQualityWatchTasksResponseBodySubmitResult) *SubmitQualityWatchTasksResponseBody {
	s.SubmitResult = v
	return s
}

func (s *SubmitQualityWatchTasksResponseBody) SetSuccess(v bool) *SubmitQualityWatchTasksResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitQualityWatchTasksResponseBody) Validate() error {
	if s.SubmitResult != nil {
		if err := s.SubmitResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitQualityWatchTasksResponseBodySubmitResult struct {
	WatchTaskIdList []*int64 `json:"WatchTaskIdList,omitempty" xml:"WatchTaskIdList,omitempty" type:"Repeated"`
}

func (s SubmitQualityWatchTasksResponseBodySubmitResult) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityWatchTasksResponseBodySubmitResult) GoString() string {
	return s.String()
}

func (s *SubmitQualityWatchTasksResponseBodySubmitResult) GetWatchTaskIdList() []*int64 {
	return s.WatchTaskIdList
}

func (s *SubmitQualityWatchTasksResponseBodySubmitResult) SetWatchTaskIdList(v []*int64) *SubmitQualityWatchTasksResponseBodySubmitResult {
	s.WatchTaskIdList = v
	return s
}

func (s *SubmitQualityWatchTasksResponseBodySubmitResult) Validate() error {
	return dara.Validate(s)
}
