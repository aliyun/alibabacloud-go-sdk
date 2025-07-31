// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamBatchJobMappingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *CreateStreamBatchJobMappingShrinkRequest
	GetOpTenantId() *int64
	SetStreamBatchJobMappingCreateCommandShrink(v string) *CreateStreamBatchJobMappingShrinkRequest
	GetStreamBatchJobMappingCreateCommandShrink() *string
}

type CreateStreamBatchJobMappingShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	StreamBatchJobMappingCreateCommandShrink *string `json:"StreamBatchJobMappingCreateCommand,omitempty" xml:"StreamBatchJobMappingCreateCommand,omitempty"`
}

func (s CreateStreamBatchJobMappingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamBatchJobMappingShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamBatchJobMappingShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStreamBatchJobMappingShrinkRequest) GetStreamBatchJobMappingCreateCommandShrink() *string {
	return s.StreamBatchJobMappingCreateCommandShrink
}

func (s *CreateStreamBatchJobMappingShrinkRequest) SetOpTenantId(v int64) *CreateStreamBatchJobMappingShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStreamBatchJobMappingShrinkRequest) SetStreamBatchJobMappingCreateCommandShrink(v string) *CreateStreamBatchJobMappingShrinkRequest {
	s.StreamBatchJobMappingCreateCommandShrink = &v
	return s
}

func (s *CreateStreamBatchJobMappingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
