// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBuildBreakpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryBuildBreakpointResponseBody
	GetCode() *string
	SetData(v *QueryBuildBreakpointResponseBodyData) *QueryBuildBreakpointResponseBody
	GetData() *QueryBuildBreakpointResponseBodyData
	SetErrorName(v string) *QueryBuildBreakpointResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *QueryBuildBreakpointResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *QueryBuildBreakpointResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryBuildBreakpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryBuildBreakpointResponseBody
	GetSuccess() *bool
}

type QueryBuildBreakpointResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryBuildBreakpointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorName *string                               `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32                                `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBuildBreakpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBuildBreakpointResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBuildBreakpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryBuildBreakpointResponseBody) GetData() *QueryBuildBreakpointResponseBodyData {
	return s.Data
}

func (s *QueryBuildBreakpointResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *QueryBuildBreakpointResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *QueryBuildBreakpointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryBuildBreakpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBuildBreakpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryBuildBreakpointResponseBody) SetCode(v string) *QueryBuildBreakpointResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBuildBreakpointResponseBody) SetData(v *QueryBuildBreakpointResponseBodyData) *QueryBuildBreakpointResponseBody {
	s.Data = v
	return s
}

func (s *QueryBuildBreakpointResponseBody) SetErrorName(v string) *QueryBuildBreakpointResponseBody {
	s.ErrorName = &v
	return s
}

func (s *QueryBuildBreakpointResponseBody) SetHttpCode(v int32) *QueryBuildBreakpointResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryBuildBreakpointResponseBody) SetMessage(v string) *QueryBuildBreakpointResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBuildBreakpointResponseBody) SetRequestId(v string) *QueryBuildBreakpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBuildBreakpointResponseBody) SetSuccess(v bool) *QueryBuildBreakpointResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBuildBreakpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryBuildBreakpointResponseBodyData struct {
	Breakpoints []*QueryBuildBreakpointResponseBodyDataBreakpoints `json:"Breakpoints,omitempty" xml:"Breakpoints,omitempty" type:"Repeated"`
	ProjectId   *string                                            `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s QueryBuildBreakpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryBuildBreakpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBuildBreakpointResponseBodyData) GetBreakpoints() []*QueryBuildBreakpointResponseBodyDataBreakpoints {
	return s.Breakpoints
}

func (s *QueryBuildBreakpointResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryBuildBreakpointResponseBodyData) SetBreakpoints(v []*QueryBuildBreakpointResponseBodyDataBreakpoints) *QueryBuildBreakpointResponseBodyData {
	s.Breakpoints = v
	return s
}

func (s *QueryBuildBreakpointResponseBodyData) SetProjectId(v string) *QueryBuildBreakpointResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *QueryBuildBreakpointResponseBodyData) Validate() error {
	if s.Breakpoints != nil {
		for _, item := range s.Breakpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryBuildBreakpointResponseBodyDataBreakpoints struct {
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s QueryBuildBreakpointResponseBodyDataBreakpoints) String() string {
	return dara.Prettify(s)
}

func (s QueryBuildBreakpointResponseBodyDataBreakpoints) GoString() string {
	return s.String()
}

func (s *QueryBuildBreakpointResponseBodyDataBreakpoints) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *QueryBuildBreakpointResponseBodyDataBreakpoints) GetJobId() *string {
	return s.JobId
}

func (s *QueryBuildBreakpointResponseBodyDataBreakpoints) SetAlgorithm(v string) *QueryBuildBreakpointResponseBodyDataBreakpoints {
	s.Algorithm = &v
	return s
}

func (s *QueryBuildBreakpointResponseBodyDataBreakpoints) SetJobId(v string) *QueryBuildBreakpointResponseBodyDataBreakpoints {
	s.JobId = &v
	return s
}

func (s *QueryBuildBreakpointResponseBodyDataBreakpoints) Validate() error {
	return dara.Validate(s)
}
