// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutProvisionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PutProvisionConfigInput) *PutProvisionConfigRequest
	GetBody() *PutProvisionConfigInput
	SetQualifier(v string) *PutProvisionConfigRequest
	GetQualifier() *string
}

type PutProvisionConfigRequest struct {
	// The provisioned configuration information.
	//
	// This parameter is required.
	Body *PutProvisionConfigInput `json:"body,omitempty" xml:"body,omitempty"`
	// The function alias.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s PutProvisionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PutProvisionConfigRequest) GoString() string {
	return s.String()
}

func (s *PutProvisionConfigRequest) GetBody() *PutProvisionConfigInput {
	return s.Body
}

func (s *PutProvisionConfigRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *PutProvisionConfigRequest) SetBody(v *PutProvisionConfigInput) *PutProvisionConfigRequest {
	s.Body = v
	return s
}

func (s *PutProvisionConfigRequest) SetQualifier(v string) *PutProvisionConfigRequest {
	s.Qualifier = &v
	return s
}

func (s *PutProvisionConfigRequest) Validate() error {
	return dara.Validate(s)
}
