// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteModelApiRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteModelApiRequest
	GetDryRun() *bool
}

type DeleteModelApiRequest struct {
	// The idempotency token. For the same tenant, model API, and parameters, you can use the original clientToken to retry the deletion if a background deletion fails. A conflict is returned if an active operation is still in progress or if the configuration has been replaced by a newer operation. A replay of a fully completed deletion returns an empty success response. Using the same token for a different target is rejected.
	//
	// example:
	//
	// delete-model-api-client-token-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Specifies whether to perform only a dry run. If you set this parameter to true, the system performs the same synchronous admission checks as an actual deletion, but does not create a deletion run or idempotent record, stop the service, or delete the API or associated configurations. If you set this parameter to false or omit it, the actual deletion is processed.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s DeleteModelApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelApiRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteModelApiRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteModelApiRequest) SetClientToken(v string) *DeleteModelApiRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteModelApiRequest) SetDryRun(v bool) *DeleteModelApiRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteModelApiRequest) Validate() error {
	return dara.Validate(s)
}
