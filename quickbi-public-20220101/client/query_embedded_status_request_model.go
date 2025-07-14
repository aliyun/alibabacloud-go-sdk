// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmbeddedStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorksId(v string) *QueryEmbeddedStatusRequest
	GetWorksId() *string
}

type QueryEmbeddedStatusRequest struct {
	// The work ID of the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s QueryEmbeddedStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEmbeddedStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryEmbeddedStatusRequest) GetWorksId() *string {
	return s.WorksId
}

func (s *QueryEmbeddedStatusRequest) SetWorksId(v string) *QueryEmbeddedStatusRequest {
	s.WorksId = &v
	return s
}

func (s *QueryEmbeddedStatusRequest) Validate() error {
	return dara.Validate(s)
}
