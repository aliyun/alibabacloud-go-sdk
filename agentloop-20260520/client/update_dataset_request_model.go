// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDatasetRequest
	GetDescription() *string
	SetSchema(v map[string]*IndexKey) *UpdateDatasetRequest
	GetSchema() map[string]*IndexKey
	SetClientToken(v string) *UpdateDatasetRequest
	GetClientToken() *string
}

type UpdateDatasetRequest struct {
	Description *string              `json:"description,omitempty" xml:"description,omitempty"`
	Schema      map[string]*IndexKey `json:"schema,omitempty" xml:"schema,omitempty"`
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDatasetRequest) GetSchema() map[string]*IndexKey {
	return s.Schema
}

func (s *UpdateDatasetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateDatasetRequest) SetDescription(v string) *UpdateDatasetRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetRequest) SetSchema(v map[string]*IndexKey) *UpdateDatasetRequest {
	s.Schema = v
	return s
}

func (s *UpdateDatasetRequest) SetClientToken(v string) *UpdateDatasetRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDatasetRequest) Validate() error {
	return dara.Validate(s)
}
