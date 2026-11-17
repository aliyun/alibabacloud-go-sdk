// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiModelSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *UpdateAiModelSourceRequest
	GetDisplayName() *string
	SetExpectedVersion(v int64) *UpdateAiModelSourceRequest
	GetExpectedVersion() *int64
	SetServiceId(v string) *UpdateAiModelSourceRequest
	GetServiceId() *string
}

type UpdateAiModelSourceRequest struct {
	// The display name of the model source or model provider.
	//
	// example:
	//
	// Model Studio model source
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The expected object version. Use the version returned by the most recent query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExpectedVersion *int64 `json:"expectedVersion,omitempty" xml:"expectedVersion,omitempty"`
	// The ID of the actual AI service bound to the model source.
	//
	// example:
	//
	// svc-d00000000000000000000
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s UpdateAiModelSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiModelSourceRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateAiModelSourceRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *UpdateAiModelSourceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateAiModelSourceRequest) SetDisplayName(v string) *UpdateAiModelSourceRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateAiModelSourceRequest) SetExpectedVersion(v int64) *UpdateAiModelSourceRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateAiModelSourceRequest) SetServiceId(v string) *UpdateAiModelSourceRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateAiModelSourceRequest) Validate() error {
	return dara.Validate(s)
}
