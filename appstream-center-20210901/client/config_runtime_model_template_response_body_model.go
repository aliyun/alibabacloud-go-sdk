// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigRuntimeModelTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ConfigRuntimeModelTemplateResponseBodyData) *ConfigRuntimeModelTemplateResponseBody
	GetData() []*ConfigRuntimeModelTemplateResponseBodyData
	SetRequestId(v string) *ConfigRuntimeModelTemplateResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ConfigRuntimeModelTemplateResponseBody
	GetTotalCount() *int32
}

type ConfigRuntimeModelTemplateResponseBody struct {
	Data []*ConfigRuntimeModelTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ConfigRuntimeModelTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigRuntimeModelTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigRuntimeModelTemplateResponseBody) GetData() []*ConfigRuntimeModelTemplateResponseBodyData {
	return s.Data
}

func (s *ConfigRuntimeModelTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigRuntimeModelTemplateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ConfigRuntimeModelTemplateResponseBody) SetData(v []*ConfigRuntimeModelTemplateResponseBodyData) *ConfigRuntimeModelTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ConfigRuntimeModelTemplateResponseBody) SetRequestId(v string) *ConfigRuntimeModelTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigRuntimeModelTemplateResponseBody) SetTotalCount(v int32) *ConfigRuntimeModelTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ConfigRuntimeModelTemplateResponseBody) Validate() error {
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

type ConfigRuntimeModelTemplateResponseBodyData struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// The parameter callerUid may not be null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// jvs-xxxxxxxx
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigRuntimeModelTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConfigRuntimeModelTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfigRuntimeModelTemplateResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *ConfigRuntimeModelTemplateResponseBodyData) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ConfigRuntimeModelTemplateResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ConfigRuntimeModelTemplateResponseBodyData) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *ConfigRuntimeModelTemplateResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ConfigRuntimeModelTemplateResponseBodyData) SetCode(v string) *ConfigRuntimeModelTemplateResponseBodyData {
	s.Code = &v
	return s
}

func (s *ConfigRuntimeModelTemplateResponseBodyData) SetHttpStatusCode(v int32) *ConfigRuntimeModelTemplateResponseBodyData {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfigRuntimeModelTemplateResponseBodyData) SetMessage(v string) *ConfigRuntimeModelTemplateResponseBodyData {
	s.Message = &v
	return s
}

func (s *ConfigRuntimeModelTemplateResponseBodyData) SetRuntimeId(v string) *ConfigRuntimeModelTemplateResponseBodyData {
	s.RuntimeId = &v
	return s
}

func (s *ConfigRuntimeModelTemplateResponseBodyData) SetSuccess(v bool) *ConfigRuntimeModelTemplateResponseBodyData {
	s.Success = &v
	return s
}

func (s *ConfigRuntimeModelTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
