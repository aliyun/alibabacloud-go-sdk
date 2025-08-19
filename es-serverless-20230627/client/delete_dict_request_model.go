// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteDictRequest
	GetName() *string
	SetType(v string) *DeleteDictRequest
	GetType() *string
}

type DeleteDictRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DeleteDictRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDictRequest) GoString() string {
	return s.String()
}

func (s *DeleteDictRequest) GetName() *string {
	return s.Name
}

func (s *DeleteDictRequest) GetType() *string {
	return s.Type
}

func (s *DeleteDictRequest) SetName(v string) *DeleteDictRequest {
	s.Name = &v
	return s
}

func (s *DeleteDictRequest) SetType(v string) *DeleteDictRequest {
	s.Type = &v
	return s
}

func (s *DeleteDictRequest) Validate() error {
	return dara.Validate(s)
}
