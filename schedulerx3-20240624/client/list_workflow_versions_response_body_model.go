// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListWorkflowVersionsResponseBody
	GetCode() *int32
	SetData(v []*ListWorkflowVersionsResponseBodyData) *ListWorkflowVersionsResponseBody
	GetData() []*ListWorkflowVersionsResponseBodyData
	SetMaxResults(v int32) *ListWorkflowVersionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListWorkflowVersionsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListWorkflowVersionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkflowVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkflowVersionsResponseBody
	GetSuccess() *bool
}

type ListWorkflowVersionsResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListWorkflowVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkflowVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowVersionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListWorkflowVersionsResponseBody) GetData() []*ListWorkflowVersionsResponseBodyData {
	return s.Data
}

func (s *ListWorkflowVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkflowVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkflowVersionsResponseBody) SetCode(v int32) *ListWorkflowVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkflowVersionsResponseBody) SetData(v []*ListWorkflowVersionsResponseBodyData) *ListWorkflowVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkflowVersionsResponseBody) SetMaxResults(v int32) *ListWorkflowVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowVersionsResponseBody) SetMessage(v string) *ListWorkflowVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkflowVersionsResponseBody) SetNextToken(v string) *ListWorkflowVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowVersionsResponseBody) SetRequestId(v string) *ListWorkflowVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowVersionsResponseBody) SetSuccess(v bool) *ListWorkflowVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkflowVersionsResponseBody) Validate() error {
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

type ListWorkflowVersionsResponseBodyData struct {
	// example:
	//
	// 2024-10-29 15:56:36
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	Current *bool `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// v1
	DagVersion *string `json:"DagVersion,omitempty" xml:"DagVersion,omitempty"`
}

func (s ListWorkflowVersionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowVersionsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkflowVersionsResponseBodyData) GetCurrent() *bool {
	return s.Current
}

func (s *ListWorkflowVersionsResponseBodyData) GetDagVersion() *string {
	return s.DagVersion
}

func (s *ListWorkflowVersionsResponseBodyData) SetCreateTime(v string) *ListWorkflowVersionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListWorkflowVersionsResponseBodyData) SetCurrent(v bool) *ListWorkflowVersionsResponseBodyData {
	s.Current = &v
	return s
}

func (s *ListWorkflowVersionsResponseBodyData) SetDagVersion(v string) *ListWorkflowVersionsResponseBodyData {
	s.DagVersion = &v
	return s
}

func (s *ListWorkflowVersionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
