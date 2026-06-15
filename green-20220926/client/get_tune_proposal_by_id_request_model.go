// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTuneProposalByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetTuneProposalByIdRequest
	GetId() *string
}

type GetTuneProposalByIdRequest struct {
	// example:
	//
	// prop-xxxxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTuneProposalByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTuneProposalByIdRequest) GoString() string {
	return s.String()
}

func (s *GetTuneProposalByIdRequest) GetId() *string {
	return s.Id
}

func (s *GetTuneProposalByIdRequest) SetId(v string) *GetTuneProposalByIdRequest {
	s.Id = &v
	return s
}

func (s *GetTuneProposalByIdRequest) Validate() error {
	return dara.Validate(s)
}
