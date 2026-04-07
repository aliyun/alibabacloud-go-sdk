// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *DeleteVariableRequest
	GetBusinessUnitId() *string
	SetVariableId(v string) *DeleteVariableRequest
	GetVariableId() *string
}

type DeleteVariableRequest struct {
	// example:
	//
	// llm-rj6aqmctjcit4acy
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// example:
	//
	// 82ea16d1-425c-4c03-9be5-cc91de9779ed
	VariableId *string `json:"VariableId,omitempty" xml:"VariableId,omitempty"`
}

func (s DeleteVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVariableRequest) GoString() string {
	return s.String()
}

func (s *DeleteVariableRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *DeleteVariableRequest) GetVariableId() *string {
	return s.VariableId
}

func (s *DeleteVariableRequest) SetBusinessUnitId(v string) *DeleteVariableRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *DeleteVariableRequest) SetVariableId(v string) *DeleteVariableRequest {
	s.VariableId = &v
	return s
}

func (s *DeleteVariableRequest) Validate() error {
	return dara.Validate(s)
}
