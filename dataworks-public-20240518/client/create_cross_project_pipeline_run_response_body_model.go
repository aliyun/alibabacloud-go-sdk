// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrossProjectPipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateCrossProjectPipelineRunResponseBodyData) *CreateCrossProjectPipelineRunResponseBody
	GetData() *CreateCrossProjectPipelineRunResponseBodyData
	SetRequestId(v string) *CreateCrossProjectPipelineRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCrossProjectPipelineRunResponseBody
	GetSuccess() *bool
}

type CreateCrossProjectPipelineRunResponseBody struct {
	// The business response.
	//
	// example:
	//
	// {"RequestId":"735894D1-D5E5-50B8-8A6D-041C90A98B23","PipelineRunId":"fcfd4160-e2ff-4603-9719-09128fe733df"}
	Data *CreateCrossProjectPipelineRunResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, which is used to locate and troubleshoot this API call.
	//
	// example:
	//
	// 735894D1-D5E5-50B8-8A6D-041C90A98B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCrossProjectPipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCrossProjectPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCrossProjectPipelineRunResponseBody) GetData() *CreateCrossProjectPipelineRunResponseBodyData {
	return s.Data
}

func (s *CreateCrossProjectPipelineRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCrossProjectPipelineRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCrossProjectPipelineRunResponseBody) SetData(v *CreateCrossProjectPipelineRunResponseBodyData) *CreateCrossProjectPipelineRunResponseBody {
	s.Data = v
	return s
}

func (s *CreateCrossProjectPipelineRunResponseBody) SetRequestId(v string) *CreateCrossProjectPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCrossProjectPipelineRunResponseBody) SetSuccess(v bool) *CreateCrossProjectPipelineRunResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCrossProjectPipelineRunResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCrossProjectPipelineRunResponseBodyData struct {
	// The cross-workspace deployment flow ID.
	//
	// example:
	//
	// fcfd4160-e2ff-4603-9719-09128fe733df
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 735894D1-D5E5-50B8-8A6D-041C90A98B23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCrossProjectPipelineRunResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCrossProjectPipelineRunResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCrossProjectPipelineRunResponseBodyData) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *CreateCrossProjectPipelineRunResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCrossProjectPipelineRunResponseBodyData) SetPipelineRunId(v string) *CreateCrossProjectPipelineRunResponseBodyData {
	s.PipelineRunId = &v
	return s
}

func (s *CreateCrossProjectPipelineRunResponseBodyData) SetRequestId(v string) *CreateCrossProjectPipelineRunResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CreateCrossProjectPipelineRunResponseBodyData) Validate() error {
	return dara.Validate(s)
}
