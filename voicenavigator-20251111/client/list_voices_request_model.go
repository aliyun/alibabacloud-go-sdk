// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListVoicesRequest
	GetInstanceId() *string
	SetModel(v string) *ListVoicesRequest
	GetModel() *string
	SetNlsAccessType(v string) *ListVoicesRequest
	GetNlsAccessType() *string
	SetNlsEngine(v string) *ListVoicesRequest
	GetNlsEngine() *string
	SetOmniModel(v bool) *ListVoicesRequest
	GetOmniModel() *bool
	SetPageNumber(v int32) *ListVoicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVoicesRequest
	GetPageSize() *int32
}

type ListVoicesRequest struct {
	// The instance ID.
	//
	// example:
	//
	// e48e45dd-e47a-4744-a063-f08cbebb1c5a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The model.
	//
	// example:
	//
	// qwen-3.8-omni-flash
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The TTS invocation method.
	//
	// example:
	//
	// MANAGED
	NlsAccessType *string `json:"NlsAccessType,omitempty" xml:"NlsAccessType,omitempty"`
	// The TTS engine.
	//
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// Specifies whether the model is an Omni model.
	//
	// example:
	//
	// false
	OmniModel *bool `json:"OmniModel,omitempty" xml:"OmniModel,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVoicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVoicesRequest) GoString() string {
	return s.String()
}

func (s *ListVoicesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVoicesRequest) GetModel() *string {
	return s.Model
}

func (s *ListVoicesRequest) GetNlsAccessType() *string {
	return s.NlsAccessType
}

func (s *ListVoicesRequest) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *ListVoicesRequest) GetOmniModel() *bool {
	return s.OmniModel
}

func (s *ListVoicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoicesRequest) SetInstanceId(v string) *ListVoicesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVoicesRequest) SetModel(v string) *ListVoicesRequest {
	s.Model = &v
	return s
}

func (s *ListVoicesRequest) SetNlsAccessType(v string) *ListVoicesRequest {
	s.NlsAccessType = &v
	return s
}

func (s *ListVoicesRequest) SetNlsEngine(v string) *ListVoicesRequest {
	s.NlsEngine = &v
	return s
}

func (s *ListVoicesRequest) SetOmniModel(v bool) *ListVoicesRequest {
	s.OmniModel = &v
	return s
}

func (s *ListVoicesRequest) SetPageNumber(v int32) *ListVoicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVoicesRequest) SetPageSize(v int32) *ListVoicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVoicesRequest) Validate() error {
	return dara.Validate(s)
}
