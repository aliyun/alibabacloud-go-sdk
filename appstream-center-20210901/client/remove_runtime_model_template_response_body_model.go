// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRuntimeModelTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*RemoveRuntimeModelTemplateResponseBodyData) *RemoveRuntimeModelTemplateResponseBody
	GetData() []*RemoveRuntimeModelTemplateResponseBodyData
	SetRequestId(v string) *RemoveRuntimeModelTemplateResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *RemoveRuntimeModelTemplateResponseBody
	GetTotalCount() *int32
}

type RemoveRuntimeModelTemplateResponseBody struct {
	Data []*RemoveRuntimeModelTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RemoveRuntimeModelTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveRuntimeModelTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveRuntimeModelTemplateResponseBody) GetData() []*RemoveRuntimeModelTemplateResponseBodyData {
	return s.Data
}

func (s *RemoveRuntimeModelTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveRuntimeModelTemplateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *RemoveRuntimeModelTemplateResponseBody) SetData(v []*RemoveRuntimeModelTemplateResponseBodyData) *RemoveRuntimeModelTemplateResponseBody {
	s.Data = v
	return s
}

func (s *RemoveRuntimeModelTemplateResponseBody) SetRequestId(v string) *RemoveRuntimeModelTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveRuntimeModelTemplateResponseBody) SetTotalCount(v int32) *RemoveRuntimeModelTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *RemoveRuntimeModelTemplateResponseBody) Validate() error {
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

type RemoveRuntimeModelTemplateResponseBodyData struct {
	// example:
	//
	// Runtime.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 404
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// The runtime is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// jvs-xxxxx
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveRuntimeModelTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveRuntimeModelTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveRuntimeModelTemplateResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *RemoveRuntimeModelTemplateResponseBodyData) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveRuntimeModelTemplateResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *RemoveRuntimeModelTemplateResponseBodyData) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *RemoveRuntimeModelTemplateResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveRuntimeModelTemplateResponseBodyData) SetCode(v string) *RemoveRuntimeModelTemplateResponseBodyData {
	s.Code = &v
	return s
}

func (s *RemoveRuntimeModelTemplateResponseBodyData) SetHttpStatusCode(v int32) *RemoveRuntimeModelTemplateResponseBodyData {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveRuntimeModelTemplateResponseBodyData) SetMessage(v string) *RemoveRuntimeModelTemplateResponseBodyData {
	s.Message = &v
	return s
}

func (s *RemoveRuntimeModelTemplateResponseBodyData) SetRuntimeId(v string) *RemoveRuntimeModelTemplateResponseBodyData {
	s.RuntimeId = &v
	return s
}

func (s *RemoveRuntimeModelTemplateResponseBodyData) SetSuccess(v bool) *RemoveRuntimeModelTemplateResponseBodyData {
	s.Success = &v
	return s
}

func (s *RemoveRuntimeModelTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
