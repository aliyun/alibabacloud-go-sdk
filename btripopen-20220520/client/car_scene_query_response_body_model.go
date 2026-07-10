// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarSceneQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CarSceneQueryResponseBody
	GetCode() *string
	SetMessage(v string) *CarSceneQueryResponseBody
	GetMessage() *string
	SetModule(v []*CarSceneQueryResponseBodyModule) *CarSceneQueryResponseBody
	GetModule() []*CarSceneQueryResponseBodyModule
	SetRequestId(v string) *CarSceneQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CarSceneQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CarSceneQueryResponseBody
	GetTraceId() *string
}

type CarSceneQueryResponseBody struct {
	Code      *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                            `json:"message,omitempty" xml:"message,omitempty"`
	Module    []*CarSceneQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                              `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                            `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CarSceneQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CarSceneQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CarSceneQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CarSceneQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CarSceneQueryResponseBody) GetModule() []*CarSceneQueryResponseBodyModule {
	return s.Module
}

func (s *CarSceneQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CarSceneQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CarSceneQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CarSceneQueryResponseBody) SetCode(v string) *CarSceneQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CarSceneQueryResponseBody) SetMessage(v string) *CarSceneQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CarSceneQueryResponseBody) SetModule(v []*CarSceneQueryResponseBodyModule) *CarSceneQueryResponseBody {
	s.Module = v
	return s
}

func (s *CarSceneQueryResponseBody) SetRequestId(v string) *CarSceneQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarSceneQueryResponseBody) SetSuccess(v bool) *CarSceneQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CarSceneQueryResponseBody) SetTraceId(v string) *CarSceneQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *CarSceneQueryResponseBody) Validate() error {
	if s.Module != nil {
		for _, item := range s.Module {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CarSceneQueryResponseBodyModule struct {
	ScenarioTemplateId   *string `json:"scenarioTemplateId,omitempty" xml:"scenarioTemplateId,omitempty"`
	ScenarioTemplateName *string `json:"scenarioTemplateName,omitempty" xml:"scenarioTemplateName,omitempty"`
	State                *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s CarSceneQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CarSceneQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CarSceneQueryResponseBodyModule) GetScenarioTemplateId() *string {
	return s.ScenarioTemplateId
}

func (s *CarSceneQueryResponseBodyModule) GetScenarioTemplateName() *string {
	return s.ScenarioTemplateName
}

func (s *CarSceneQueryResponseBodyModule) GetState() *string {
	return s.State
}

func (s *CarSceneQueryResponseBodyModule) SetScenarioTemplateId(v string) *CarSceneQueryResponseBodyModule {
	s.ScenarioTemplateId = &v
	return s
}

func (s *CarSceneQueryResponseBodyModule) SetScenarioTemplateName(v string) *CarSceneQueryResponseBodyModule {
	s.ScenarioTemplateName = &v
	return s
}

func (s *CarSceneQueryResponseBodyModule) SetState(v string) *CarSceneQueryResponseBodyModule {
	s.State = &v
	return s
}

func (s *CarSceneQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
