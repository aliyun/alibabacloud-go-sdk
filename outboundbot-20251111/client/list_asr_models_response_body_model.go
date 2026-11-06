// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsrModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAsrModelsResponseBody
	GetCode() *string
	SetData(v []*ListAsrModelsResponseBodyData) *ListAsrModelsResponseBody
	GetData() []*ListAsrModelsResponseBodyData
	SetHttpStatusCode(v int32) *ListAsrModelsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAsrModelsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListAsrModelsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListAsrModelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAsrModelsResponseBody
	GetSuccess() *bool
}

type ListAsrModelsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of ASR models.
	Data []*ListAsrModelsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAsrModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAsrModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsrModelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAsrModelsResponseBody) GetData() []*ListAsrModelsResponseBodyData {
	return s.Data
}

func (s *ListAsrModelsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAsrModelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAsrModelsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListAsrModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAsrModelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAsrModelsResponseBody) SetCode(v string) *ListAsrModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAsrModelsResponseBody) SetData(v []*ListAsrModelsResponseBodyData) *ListAsrModelsResponseBody {
	s.Data = v
	return s
}

func (s *ListAsrModelsResponseBody) SetHttpStatusCode(v int32) *ListAsrModelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAsrModelsResponseBody) SetMessage(v string) *ListAsrModelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAsrModelsResponseBody) SetParams(v []*string) *ListAsrModelsResponseBody {
	s.Params = v
	return s
}

func (s *ListAsrModelsResponseBody) SetRequestId(v string) *ListAsrModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsrModelsResponseBody) SetSuccess(v bool) *ListAsrModelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAsrModelsResponseBody) Validate() error {
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

type ListAsrModelsResponseBodyData struct {
	// The ASR model identifier.
	//
	// example:
	//
	// paraformer-realtime-v2
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The ASR model name.
	//
	// example:
	//
	// Standard model
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The NLS engine.
	//
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// The list of parameters supported by the model.
	SupportedParams []*string `json:"SupportedParams,omitempty" xml:"SupportedParams,omitempty" type:"Repeated"`
}

func (s ListAsrModelsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAsrModelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAsrModelsResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *ListAsrModelsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAsrModelsResponseBodyData) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *ListAsrModelsResponseBodyData) GetSupportedParams() []*string {
	return s.SupportedParams
}

func (s *ListAsrModelsResponseBodyData) SetModel(v string) *ListAsrModelsResponseBodyData {
	s.Model = &v
	return s
}

func (s *ListAsrModelsResponseBodyData) SetName(v string) *ListAsrModelsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAsrModelsResponseBodyData) SetNlsEngine(v string) *ListAsrModelsResponseBodyData {
	s.NlsEngine = &v
	return s
}

func (s *ListAsrModelsResponseBodyData) SetSupportedParams(v []*string) *ListAsrModelsResponseBodyData {
	s.SupportedParams = v
	return s
}

func (s *ListAsrModelsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
