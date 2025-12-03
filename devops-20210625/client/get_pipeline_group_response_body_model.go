// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetPipelineGroupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPipelineGroupResponseBody
	GetErrorMessage() *string
	SetPipelineGroup(v *GetPipelineGroupResponseBodyPipelineGroup) *GetPipelineGroupResponseBody
	GetPipelineGroup() *GetPipelineGroupResponseBodyPipelineGroup
	SetRequestId(v string) *GetPipelineGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPipelineGroupResponseBody
	GetSuccess() *bool
}

type GetPipelineGroupResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage  *string                                    `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	PipelineGroup *GetPipelineGroupResponseBodyPipelineGroup `json:"pipelineGroup,omitempty" xml:"pipelineGroup,omitempty" type:"Struct"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPipelineGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPipelineGroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPipelineGroupResponseBody) GetPipelineGroup() *GetPipelineGroupResponseBodyPipelineGroup {
	return s.PipelineGroup
}

func (s *GetPipelineGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPipelineGroupResponseBody) SetErrorCode(v string) *GetPipelineGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineGroupResponseBody) SetErrorMessage(v string) *GetPipelineGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineGroupResponseBody) SetPipelineGroup(v *GetPipelineGroupResponseBodyPipelineGroup) *GetPipelineGroupResponseBody {
	s.PipelineGroup = v
	return s
}

func (s *GetPipelineGroupResponseBody) SetRequestId(v string) *GetPipelineGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineGroupResponseBody) SetSuccess(v bool) *GetPipelineGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetPipelineGroupResponseBody) Validate() error {
	if s.PipelineGroup != nil {
		if err := s.PipelineGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineGroupResponseBodyPipelineGroup struct {
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

func (s GetPipelineGroupResponseBodyPipelineGroup) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineGroupResponseBodyPipelineGroup) GoString() string {
	return s.String()
}

func (s *GetPipelineGroupResponseBodyPipelineGroup) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetPipelineGroupResponseBodyPipelineGroup) GetId() *int64 {
	return s.Id
}

func (s *GetPipelineGroupResponseBodyPipelineGroup) GetName() *string {
	return s.Name
}

func (s *GetPipelineGroupResponseBodyPipelineGroup) SetCreateTime(v int64) *GetPipelineGroupResponseBodyPipelineGroup {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineGroupResponseBodyPipelineGroup) SetId(v int64) *GetPipelineGroupResponseBodyPipelineGroup {
	s.Id = &v
	return s
}

func (s *GetPipelineGroupResponseBodyPipelineGroup) SetName(v string) *GetPipelineGroupResponseBodyPipelineGroup {
	s.Name = &v
	return s
}

func (s *GetPipelineGroupResponseBodyPipelineGroup) Validate() error {
	return dara.Validate(s)
}
