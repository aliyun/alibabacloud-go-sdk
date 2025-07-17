// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateDatabaseRequest
	GetName() *string
	SetOptions(v map[string]*string) *CreateDatabaseRequest
	GetOptions() map[string]*string
}

type CreateDatabaseRequest struct {
	// example:
	//
	// database_name
	Name    *string            `json:"name,omitempty" xml:"name,omitempty"`
	Options map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
}

func (s CreateDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseRequest) GetName() *string {
	return s.Name
}

func (s *CreateDatabaseRequest) GetOptions() map[string]*string {
	return s.Options
}

func (s *CreateDatabaseRequest) SetName(v string) *CreateDatabaseRequest {
	s.Name = &v
	return s
}

func (s *CreateDatabaseRequest) SetOptions(v map[string]*string) *CreateDatabaseRequest {
	s.Options = v
	return s
}

func (s *CreateDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
