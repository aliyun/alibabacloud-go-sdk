// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateServiceRequest
	GetDryRun() *bool
	SetMemberToUpdate(v string) *UpdateServiceRequest
	GetMemberToUpdate() *string
	SetUpdateType(v string) *UpdateServiceRequest
	GetUpdateType() *string
	SetBody(v string) *UpdateServiceRequest
	GetBody() *string
}

type UpdateServiceRequest struct {
	// The client token used to ensure the idempotence of the request.
	//
	// example:
	//
	// 5db780a0-9d3e-4f35-a84a-xxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request.
	//
	// - true: sends a precheck request without performing the actual operation. If the precheck fails, the corresponding error is returned. If the precheck succeeds, the error code DryRunOperation is returned.
	//
	// - false: sends a normal request. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The type of the child service to update. Note: This parameter takes effect only for aggregated services.
	//
	// example:
	//
	// LLMPrefill
	MemberToUpdate *string `json:"MemberToUpdate,omitempty" xml:"MemberToUpdate,omitempty"`
	// The update type of the service. Both merge and replace modes are supported. If this parameter is not specified, the merge mode is used by default.
	//
	// - merge: the merge mode. If the original service configuration JSON is
	//
	// `{"a":"b"}` and the updated configuration is `{"c":"d"}`, the final configuration in merge mode is `{"a":"b","c":"d"}`.
	//
	// - replace: the replace mode. If the original service configuration JSON is
	//
	// `{"a":"b"}` and the updated configuration is `{"c":"d"}`, the final configuration in replace mode is `{"c":"d"}`.
	//
	// example:
	//
	// merge
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
	// The updated configuration. For more information about parameter settings, see [CreateService](https://help.aliyun.com/document_detail/412086.html).
	//
	// example:
	//
	// {
	//
	//       "name": "foo",
	//
	//       "model_path": "http://path/to/model.tar.gz",
	//
	//       "processor": "tensorflow_cpu",
	//
	//       "metadata": {
	//
	//             "instance": 2,
	//
	//             "memory": 7000,
	//
	//             "cpu": 4
	//
	//       }
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateServiceRequest) GetMemberToUpdate() *string {
	return s.MemberToUpdate
}

func (s *UpdateServiceRequest) GetUpdateType() *string {
	return s.UpdateType
}

func (s *UpdateServiceRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateServiceRequest) SetClientToken(v string) *UpdateServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceRequest) SetDryRun(v bool) *UpdateServiceRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceRequest) SetMemberToUpdate(v string) *UpdateServiceRequest {
	s.MemberToUpdate = &v
	return s
}

func (s *UpdateServiceRequest) SetUpdateType(v string) *UpdateServiceRequest {
	s.UpdateType = &v
	return s
}

func (s *UpdateServiceRequest) SetBody(v string) *UpdateServiceRequest {
	s.Body = &v
	return s
}

func (s *UpdateServiceRequest) Validate() error {
	return dara.Validate(s)
}
