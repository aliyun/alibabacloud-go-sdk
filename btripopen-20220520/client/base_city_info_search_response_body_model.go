// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseCityInfoSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BaseCityInfoSearchResponseBody
	GetCode() *string
	SetMessage(v string) *BaseCityInfoSearchResponseBody
	GetMessage() *string
	SetModule(v []*BaseCityInfoSearchResponseBodyModule) *BaseCityInfoSearchResponseBody
	GetModule() []*BaseCityInfoSearchResponseBodyModule
	SetRequestId(v string) *BaseCityInfoSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BaseCityInfoSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *BaseCityInfoSearchResponseBody
	GetTraceId() *string
}

type BaseCityInfoSearchResponseBody struct {
	Code      *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Module    []*BaseCityInfoSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                 `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s BaseCityInfoSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BaseCityInfoSearchResponseBody) GoString() string {
	return s.String()
}

func (s *BaseCityInfoSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *BaseCityInfoSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BaseCityInfoSearchResponseBody) GetModule() []*BaseCityInfoSearchResponseBodyModule {
	return s.Module
}

func (s *BaseCityInfoSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BaseCityInfoSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BaseCityInfoSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *BaseCityInfoSearchResponseBody) SetCode(v string) *BaseCityInfoSearchResponseBody {
	s.Code = &v
	return s
}

func (s *BaseCityInfoSearchResponseBody) SetMessage(v string) *BaseCityInfoSearchResponseBody {
	s.Message = &v
	return s
}

func (s *BaseCityInfoSearchResponseBody) SetModule(v []*BaseCityInfoSearchResponseBodyModule) *BaseCityInfoSearchResponseBody {
	s.Module = v
	return s
}

func (s *BaseCityInfoSearchResponseBody) SetRequestId(v string) *BaseCityInfoSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *BaseCityInfoSearchResponseBody) SetSuccess(v bool) *BaseCityInfoSearchResponseBody {
	s.Success = &v
	return s
}

func (s *BaseCityInfoSearchResponseBody) SetTraceId(v string) *BaseCityInfoSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *BaseCityInfoSearchResponseBody) Validate() error {
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

type BaseCityInfoSearchResponseBodyModule struct {
	Code     *string `json:"code,omitempty" xml:"code,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	NameTree *string `json:"nameTree,omitempty" xml:"nameTree,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s BaseCityInfoSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s BaseCityInfoSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *BaseCityInfoSearchResponseBodyModule) GetCode() *string {
	return s.Code
}

func (s *BaseCityInfoSearchResponseBodyModule) GetName() *string {
	return s.Name
}

func (s *BaseCityInfoSearchResponseBodyModule) GetNameTree() *string {
	return s.NameTree
}

func (s *BaseCityInfoSearchResponseBodyModule) GetRegion() *string {
	return s.Region
}

func (s *BaseCityInfoSearchResponseBodyModule) SetCode(v string) *BaseCityInfoSearchResponseBodyModule {
	s.Code = &v
	return s
}

func (s *BaseCityInfoSearchResponseBodyModule) SetName(v string) *BaseCityInfoSearchResponseBodyModule {
	s.Name = &v
	return s
}

func (s *BaseCityInfoSearchResponseBodyModule) SetNameTree(v string) *BaseCityInfoSearchResponseBodyModule {
	s.NameTree = &v
	return s
}

func (s *BaseCityInfoSearchResponseBodyModule) SetRegion(v string) *BaseCityInfoSearchResponseBodyModule {
	s.Region = &v
	return s
}

func (s *BaseCityInfoSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
