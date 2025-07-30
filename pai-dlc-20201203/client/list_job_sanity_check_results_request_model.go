// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobSanityCheckResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *ListJobSanityCheckResultsRequest
	GetOrder() *string
}

type ListJobSanityCheckResultsRequest struct {
	// The sorting order:
	//
	// 	- desc: descending order
	//
	// 	- asc: ascending order
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListJobSanityCheckResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobSanityCheckResultsRequest) GoString() string {
	return s.String()
}

func (s *ListJobSanityCheckResultsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListJobSanityCheckResultsRequest) SetOrder(v string) *ListJobSanityCheckResultsRequest {
	s.Order = &v
	return s
}

func (s *ListJobSanityCheckResultsRequest) Validate() error {
	return dara.Validate(s)
}
