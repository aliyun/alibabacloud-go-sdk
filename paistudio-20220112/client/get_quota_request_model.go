// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v bool) *GetQuotaRequest
	GetVerbose() *bool
}

type GetQuotaRequest struct {
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetQuotaRequest) SetVerbose(v bool) *GetQuotaRequest {
	s.Verbose = &v
	return s
}

func (s *GetQuotaRequest) Validate() error {
	return dara.Validate(s)
}
