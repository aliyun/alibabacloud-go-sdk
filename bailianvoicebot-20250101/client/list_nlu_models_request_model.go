// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNluModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *ListNluModelsRequest
	GetBusinessUnitId() *string
}

type ListNluModelsRequest struct {
	// example:
	//
	// llm-3pptowd2olrctsvc
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
}

func (s ListNluModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNluModelsRequest) GoString() string {
	return s.String()
}

func (s *ListNluModelsRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *ListNluModelsRequest) SetBusinessUnitId(v string) *ListNluModelsRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *ListNluModelsRequest) Validate() error {
	return dara.Validate(s)
}
