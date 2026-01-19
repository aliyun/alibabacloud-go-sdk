// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCorpNumberGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateCorpNumberGroupRequest
	GetDescription() *string
	SetName(v string) *CreateCorpNumberGroupRequest
	GetName() *string
}

type CreateCorpNumberGroupRequest struct {
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateCorpNumberGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCorpNumberGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateCorpNumberGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCorpNumberGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateCorpNumberGroupRequest) SetDescription(v string) *CreateCorpNumberGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateCorpNumberGroupRequest) SetName(v string) *CreateCorpNumberGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateCorpNumberGroupRequest) Validate() error {
	return dara.Validate(s)
}
