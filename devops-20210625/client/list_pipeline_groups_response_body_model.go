// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListPipelineGroupsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListPipelineGroupsResponseBody
	GetErrorMessage() *string
	SetNextToken(v string) *ListPipelineGroupsResponseBody
	GetNextToken() *string
	SetPipelineGroups(v []*ListPipelineGroupsResponseBodyPipelineGroups) *ListPipelineGroupsResponseBody
	GetPipelineGroups() []*ListPipelineGroupsResponseBodyPipelineGroups
	SetRequestId(v string) *ListPipelineGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPipelineGroupsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListPipelineGroupsResponseBody
	GetTotalCount() *int32
}

type ListPipelineGroupsResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ssaassasass
	NextToken      *string                                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PipelineGroups []*ListPipelineGroupsResponseBodyPipelineGroups `json:"pipelineGroups,omitempty" xml:"pipelineGroups,omitempty" type:"Repeated"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelineGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPipelineGroupsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListPipelineGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelineGroupsResponseBody) GetPipelineGroups() []*ListPipelineGroupsResponseBodyPipelineGroups {
	return s.PipelineGroups
}

func (s *ListPipelineGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelineGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPipelineGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPipelineGroupsResponseBody) SetErrorCode(v string) *ListPipelineGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetErrorMessage(v string) *ListPipelineGroupsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetNextToken(v string) *ListPipelineGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetPipelineGroups(v []*ListPipelineGroupsResponseBodyPipelineGroups) *ListPipelineGroupsResponseBody {
	s.PipelineGroups = v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetRequestId(v string) *ListPipelineGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetSuccess(v bool) *ListPipelineGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelineGroupsResponseBody) SetTotalCount(v int32) *ListPipelineGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPipelineGroupsResponseBody) Validate() error {
	if s.PipelineGroups != nil {
		for _, item := range s.PipelineGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPipelineGroupsResponseBodyPipelineGroups struct {
	// example:
	//
	// 1586863220000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 111
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPipelineGroupsResponseBodyPipelineGroups) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineGroupsResponseBodyPipelineGroups) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupsResponseBodyPipelineGroups) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListPipelineGroupsResponseBodyPipelineGroups) GetId() *int64 {
	return s.Id
}

func (s *ListPipelineGroupsResponseBodyPipelineGroups) GetName() *string {
	return s.Name
}

func (s *ListPipelineGroupsResponseBodyPipelineGroups) SetCreateTime(v int64) *ListPipelineGroupsResponseBodyPipelineGroups {
	s.CreateTime = &v
	return s
}

func (s *ListPipelineGroupsResponseBodyPipelineGroups) SetId(v int64) *ListPipelineGroupsResponseBodyPipelineGroups {
	s.Id = &v
	return s
}

func (s *ListPipelineGroupsResponseBodyPipelineGroups) SetName(v string) *ListPipelineGroupsResponseBodyPipelineGroups {
	s.Name = &v
	return s
}

func (s *ListPipelineGroupsResponseBodyPipelineGroups) Validate() error {
	return dara.Validate(s)
}
