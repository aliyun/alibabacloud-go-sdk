// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreatePipelineGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreatePipelineGroupResponseBody
	GetErrorMessage() *string
	SetPipelineGroup(v *CreatePipelineGroupResponseBodyPipelineGroup) *CreatePipelineGroupResponseBody
	GetPipelineGroup() *CreatePipelineGroupResponseBodyPipelineGroup
	SetRequestId(v string) *CreatePipelineGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePipelineGroupResponseBody
	GetSuccess() *bool
}

type CreatePipelineGroupResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage  *string                                       `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	PipelineGroup *CreatePipelineGroupResponseBodyPipelineGroup `json:"pipelineGroup,omitempty" xml:"pipelineGroup,omitempty" type:"Struct"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreatePipelineGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreatePipelineGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreatePipelineGroupResponseBody) GetPipelineGroup() *CreatePipelineGroupResponseBodyPipelineGroup {
	return s.PipelineGroup
}

func (s *CreatePipelineGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePipelineGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePipelineGroupResponseBody) SetErrorCode(v string) *CreatePipelineGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreatePipelineGroupResponseBody) SetErrorMessage(v string) *CreatePipelineGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePipelineGroupResponseBody) SetPipelineGroup(v *CreatePipelineGroupResponseBodyPipelineGroup) *CreatePipelineGroupResponseBody {
	s.PipelineGroup = v
	return s
}

func (s *CreatePipelineGroupResponseBody) SetRequestId(v string) *CreatePipelineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelineGroupResponseBody) SetSuccess(v bool) *CreatePipelineGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePipelineGroupResponseBody) Validate() error {
	if s.PipelineGroup != nil {
		if err := s.PipelineGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePipelineGroupResponseBodyPipelineGroup struct {
	// example:
	//
	// 111
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreatePipelineGroupResponseBodyPipelineGroup) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineGroupResponseBodyPipelineGroup) GoString() string {
	return s.String()
}

func (s *CreatePipelineGroupResponseBodyPipelineGroup) GetId() *int64 {
	return s.Id
}

func (s *CreatePipelineGroupResponseBodyPipelineGroup) GetName() *string {
	return s.Name
}

func (s *CreatePipelineGroupResponseBodyPipelineGroup) SetId(v int64) *CreatePipelineGroupResponseBodyPipelineGroup {
	s.Id = &v
	return s
}

func (s *CreatePipelineGroupResponseBodyPipelineGroup) SetName(v string) *CreatePipelineGroupResponseBodyPipelineGroup {
	s.Name = &v
	return s
}

func (s *CreatePipelineGroupResponseBodyPipelineGroup) Validate() error {
	return dara.Validate(s)
}
