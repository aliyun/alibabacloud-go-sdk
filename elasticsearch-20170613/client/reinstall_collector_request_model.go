// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstallCollectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ReinstallCollectorRequest
	GetClientToken() *string
	SetBody(v string) *ReinstallCollectorRequest
	GetBody() *string
}

type ReinstallCollectorRequest struct {
	// A unique token used to ensure the idempotence of the request. The client generates this value. The value must be unique among different requests and cannot exceed 64 ASCII characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The request body parameters. For more information, see the RequestBody section below.
	//
	// example:
	//
	// {
	//
	//   "restartType": "nodeEcsId",
	//
	//   "nodes":["i-bp1gyhphjaj73jsr****","i-bp10piq1mkfnyw9t****"]
	//
	// }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReinstallCollectorRequest) String() string {
	return dara.Prettify(s)
}

func (s ReinstallCollectorRequest) GoString() string {
	return s.String()
}

func (s *ReinstallCollectorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReinstallCollectorRequest) GetBody() *string {
	return s.Body
}

func (s *ReinstallCollectorRequest) SetClientToken(v string) *ReinstallCollectorRequest {
	s.ClientToken = &v
	return s
}

func (s *ReinstallCollectorRequest) SetBody(v string) *ReinstallCollectorRequest {
	s.Body = &v
	return s
}

func (s *ReinstallCollectorRequest) Validate() error {
	return dara.Validate(s)
}
