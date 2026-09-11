// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteHttpApiRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteHttpApiRequest
	GetDryRun() *bool
}

type DeleteHttpApiRequest struct {
	// The idempotence token. If you call this operation multiple times for the same HTTP API by using the same clientToken value and the same delete request parameters, the result of the first successful call is returned.
	//
	// example:
	//
	// delete-http-api-client-token-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Specifies whether to perform only a dry run. If set to true, all synchronous validations identical to an actual deletion are executed (including admission checks such as whether a published API can be deleted), but the API is not deleted, no associated configurations are cleaned up, and no side effects are produced. If this parameter is not specified or is set to false, the behavior is the same as the existing version.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s DeleteHttpApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteHttpApiRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteHttpApiRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteHttpApiRequest) SetClientToken(v string) *DeleteHttpApiRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteHttpApiRequest) SetDryRun(v bool) *DeleteHttpApiRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteHttpApiRequest) Validate() error {
	return dara.Validate(s)
}
