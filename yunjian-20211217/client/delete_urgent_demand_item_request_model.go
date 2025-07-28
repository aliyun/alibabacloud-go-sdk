// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUrgentDemandItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteUrgentDemandItemRequest
	GetId() *int64
	SetModifier(v string) *DeleteUrgentDemandItemRequest
	GetModifier() *string
}

type DeleteUrgentDemandItemRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111222
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
}

func (s DeleteUrgentDemandItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUrgentDemandItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandItemRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteUrgentDemandItemRequest) GetModifier() *string {
	return s.Modifier
}

func (s *DeleteUrgentDemandItemRequest) SetId(v int64) *DeleteUrgentDemandItemRequest {
	s.Id = &v
	return s
}

func (s *DeleteUrgentDemandItemRequest) SetModifier(v string) *DeleteUrgentDemandItemRequest {
	s.Modifier = &v
	return s
}

func (s *DeleteUrgentDemandItemRequest) Validate() error {
	return dara.Validate(s)
}
