// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *CreateConfigResponseBody
	GetCreatedAt() *string
	SetKey(v string) *CreateConfigResponseBody
	GetKey() *string
	SetUpdatedAt(v string) *CreateConfigResponseBody
	GetUpdatedAt() *string
	SetValue(v string) *CreateConfigResponseBody
	GetValue() *string
}

type CreateConfigResponseBody struct {
	// example:
	//
	// 2024-01-01T00:00:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// model-config
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// example:
	//
	// {"model": "gpt-4"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateConfigResponseBody) GetKey() *string {
	return s.Key
}

func (s *CreateConfigResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateConfigResponseBody) GetValue() *string {
	return s.Value
}

func (s *CreateConfigResponseBody) SetCreatedAt(v string) *CreateConfigResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateConfigResponseBody) SetKey(v string) *CreateConfigResponseBody {
	s.Key = &v
	return s
}

func (s *CreateConfigResponseBody) SetUpdatedAt(v string) *CreateConfigResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *CreateConfigResponseBody) SetValue(v string) *CreateConfigResponseBody {
	s.Value = &v
	return s
}

func (s *CreateConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
