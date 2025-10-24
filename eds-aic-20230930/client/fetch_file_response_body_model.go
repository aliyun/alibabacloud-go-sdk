// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*FetchFileResponseBodyData) *FetchFileResponseBody
	GetData() []*FetchFileResponseBodyData
	SetRequestId(v string) *FetchFileResponseBody
	GetRequestId() *string
	SetTaskId(v string) *FetchFileResponseBody
	GetTaskId() *string
}

type FetchFileResponseBody struct {
	// The objects that are returned.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	Data []*FetchFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request. If the request fails, share this ID with technical support to help diagnose the issue.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the batch task.
	//
	// example:
	//
	// t-ehs0yoedj0xe9****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FetchFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchFileResponseBody) GoString() string {
	return s.String()
}

func (s *FetchFileResponseBody) GetData() []*FetchFileResponseBodyData {
	return s.Data
}

func (s *FetchFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchFileResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *FetchFileResponseBody) SetData(v []*FetchFileResponseBodyData) *FetchFileResponseBody {
	s.Data = v
	return s
}

func (s *FetchFileResponseBody) SetRequestId(v string) *FetchFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchFileResponseBody) SetTaskId(v string) *FetchFileResponseBody {
	s.TaskId = &v
	return s
}

func (s *FetchFileResponseBody) Validate() error {
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

type FetchFileResponseBodyData struct {
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	TaskId            *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FetchFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FetchFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchFileResponseBodyData) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *FetchFileResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *FetchFileResponseBodyData) SetAndroidInstanceId(v string) *FetchFileResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *FetchFileResponseBodyData) SetTaskId(v string) *FetchFileResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *FetchFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
