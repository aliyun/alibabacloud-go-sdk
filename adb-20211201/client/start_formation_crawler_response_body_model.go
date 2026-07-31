// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartFormationCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *StartFormationCrawlerResponseBody
	GetData() *string
	SetRequestId(v string) *StartFormationCrawlerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartFormationCrawlerResponseBody
	GetSuccess() *bool
	SetTaskId(v int64) *StartFormationCrawlerResponseBody
	GetTaskId() *int64
}

type StartFormationCrawlerResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// 69
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1564657730
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartFormationCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartFormationCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *StartFormationCrawlerResponseBody) GetData() *string {
	return s.Data
}

func (s *StartFormationCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartFormationCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartFormationCrawlerResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *StartFormationCrawlerResponseBody) SetData(v string) *StartFormationCrawlerResponseBody {
	s.Data = &v
	return s
}

func (s *StartFormationCrawlerResponseBody) SetRequestId(v string) *StartFormationCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartFormationCrawlerResponseBody) SetSuccess(v bool) *StartFormationCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *StartFormationCrawlerResponseBody) SetTaskId(v int64) *StartFormationCrawlerResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartFormationCrawlerResponseBody) Validate() error {
	return dara.Validate(s)
}
