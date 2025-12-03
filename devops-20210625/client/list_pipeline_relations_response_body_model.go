// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListPipelineRelationsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListPipelineRelationsResponseBody
	GetErrorMessage() *string
	SetPipelineRelations(v []*ListPipelineRelationsResponseBodyPipelineRelations) *ListPipelineRelationsResponseBody
	GetPipelineRelations() []*ListPipelineRelationsResponseBodyPipelineRelations
	SetRequestId(v string) *ListPipelineRelationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPipelineRelationsResponseBody
	GetSuccess() *bool
}

type ListPipelineRelationsResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage      *string                                               `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	PipelineRelations []*ListPipelineRelationsResponseBodyPipelineRelations `json:"pipelineRelations,omitempty" xml:"pipelineRelations,omitempty" type:"Repeated"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListPipelineRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineRelationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPipelineRelationsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListPipelineRelationsResponseBody) GetPipelineRelations() []*ListPipelineRelationsResponseBodyPipelineRelations {
	return s.PipelineRelations
}

func (s *ListPipelineRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelineRelationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPipelineRelationsResponseBody) SetErrorCode(v string) *ListPipelineRelationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineRelationsResponseBody) SetErrorMessage(v string) *ListPipelineRelationsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineRelationsResponseBody) SetPipelineRelations(v []*ListPipelineRelationsResponseBodyPipelineRelations) *ListPipelineRelationsResponseBody {
	s.PipelineRelations = v
	return s
}

func (s *ListPipelineRelationsResponseBody) SetRequestId(v string) *ListPipelineRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineRelationsResponseBody) SetSuccess(v bool) *ListPipelineRelationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelineRelationsResponseBody) Validate() error {
	if s.PipelineRelations != nil {
		for _, item := range s.PipelineRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPipelineRelationsResponseBodyPipelineRelations struct {
	// example:
	//
	// 12234
	RefObjectId *int64 `json:"refObjectId,omitempty" xml:"refObjectId,omitempty"`
}

func (s ListPipelineRelationsResponseBodyPipelineRelations) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRelationsResponseBodyPipelineRelations) GoString() string {
	return s.String()
}

func (s *ListPipelineRelationsResponseBodyPipelineRelations) GetRefObjectId() *int64 {
	return s.RefObjectId
}

func (s *ListPipelineRelationsResponseBodyPipelineRelations) SetRefObjectId(v int64) *ListPipelineRelationsResponseBodyPipelineRelations {
	s.RefObjectId = &v
	return s
}

func (s *ListPipelineRelationsResponseBodyPipelineRelations) Validate() error {
	return dara.Validate(s)
}
