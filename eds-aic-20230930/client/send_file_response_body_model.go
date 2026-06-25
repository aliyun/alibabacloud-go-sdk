// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SendFileResponseBodyData) *SendFileResponseBody
	GetData() []*SendFileResponseBodyData
	SetRequestId(v string) *SendFileResponseBody
	GetRequestId() *string
	SetTaskId(v string) *SendFileResponseBody
	GetTaskId() *string
}

type SendFileResponseBody struct {
	// The details of the created tasks.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	Data []*SendFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A unique identifier for the request. If you encounter an issue, provide this request ID to technical support for troubleshooting.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The batch task ID.
	//
	// example:
	//
	// t-ehs0yoedj0xe9****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SendFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendFileResponseBody) GoString() string {
	return s.String()
}

func (s *SendFileResponseBody) GetData() []*SendFileResponseBodyData {
	return s.Data
}

func (s *SendFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendFileResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *SendFileResponseBody) SetData(v []*SendFileResponseBodyData) *SendFileResponseBody {
	s.Data = v
	return s
}

func (s *SendFileResponseBody) SetRequestId(v string) *SendFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendFileResponseBody) SetTaskId(v string) *SendFileResponseBody {
	s.TaskId = &v
	return s
}

func (s *SendFileResponseBody) Validate() error {
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

type SendFileResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The ID of the individual task for a specific cloud phone.
	//
	// example:
	//
	// t-4ks224ujixw****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SendFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendFileResponseBodyData) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *SendFileResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SendFileResponseBodyData) SetAndroidInstanceId(v string) *SendFileResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *SendFileResponseBodyData) SetTaskId(v string) *SendFileResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SendFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
