// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorksId(v string) *QueryWorksRequest
	GetWorksId() *string
}

type QueryWorksRequest struct {
	// Report ID
	//
	// This parameter is required.
	//
	// example:
	//
	// abcd****
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s QueryWorksRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksRequest) GoString() string {
	return s.String()
}

func (s *QueryWorksRequest) GetWorksId() *string {
	return s.WorksId
}

func (s *QueryWorksRequest) SetWorksId(v string) *QueryWorksRequest {
	s.WorksId = &v
	return s
}

func (s *QueryWorksRequest) Validate() error {
	return dara.Validate(s)
}
