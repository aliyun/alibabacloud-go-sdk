// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportDefinitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNbid(v string) *ListReportDefinitionsRequest
	GetNbid() *string
}

type ListReportDefinitionsRequest struct {
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s ListReportDefinitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReportDefinitionsRequest) GoString() string {
	return s.String()
}

func (s *ListReportDefinitionsRequest) GetNbid() *string {
	return s.Nbid
}

func (s *ListReportDefinitionsRequest) SetNbid(v string) *ListReportDefinitionsRequest {
	s.Nbid = &v
	return s
}

func (s *ListReportDefinitionsRequest) Validate() error {
	return dara.Validate(s)
}
