// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateAgentSpecRequestBody) *UpdateAgentSpecRequest
	GetBody() *UpdateAgentSpecRequestBody
}

type UpdateAgentSpecRequest struct {
	// The request body.
	Body *UpdateAgentSpecRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateAgentSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentSpecRequest) GetBody() *UpdateAgentSpecRequestBody {
	return s.Body
}

func (s *UpdateAgentSpecRequest) SetBody(v *UpdateAgentSpecRequestBody) *UpdateAgentSpecRequest {
	s.Body = v
	return s
}

func (s *UpdateAgentSpecRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentSpecRequestBody struct {
	// The business tags as a JSON-formatted string.
	//
	// example:
	//
	// ["ai","agent"]
	BizTags *string `json:"bizTags,omitempty" xml:"bizTags,omitempty"`
	// The label mapping as a JSON-formatted string.
	//
	// example:
	//
	// {"latest":"0.0.1"}
	Labels *string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The visibility scope. Valid values:
	//
	// - PUBLIC
	//
	// - PRIVATE
	//
	// example:
	//
	// PUBLIC
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s UpdateAgentSpecRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSpecRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentSpecRequestBody) GetBizTags() *string {
	return s.BizTags
}

func (s *UpdateAgentSpecRequestBody) GetLabels() *string {
	return s.Labels
}

func (s *UpdateAgentSpecRequestBody) GetScope() *string {
	return s.Scope
}

func (s *UpdateAgentSpecRequestBody) SetBizTags(v string) *UpdateAgentSpecRequestBody {
	s.BizTags = &v
	return s
}

func (s *UpdateAgentSpecRequestBody) SetLabels(v string) *UpdateAgentSpecRequestBody {
	s.Labels = &v
	return s
}

func (s *UpdateAgentSpecRequestBody) SetScope(v string) *UpdateAgentSpecRequestBody {
	s.Scope = &v
	return s
}

func (s *UpdateAgentSpecRequestBody) Validate() error {
	return dara.Validate(s)
}
