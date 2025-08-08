// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnection interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *Connection
	GetCreatedTime() *string
	SetDescription(v string) *Connection
	GetDescription() *string
	SetKind(v string) *Connection
	GetKind() *string
	SetLabels(v map[string]*string) *Connection
	GetLabels() map[string]*string
	SetName(v string) *Connection
	GetName() *string
	SetSpec(v *ConnectionSpec) *Connection
	GetSpec() *ConnectionSpec
	SetStatus(v *ConnectionStatus) *Connection
	GetStatus() *ConnectionStatus
	SetUid(v string) *Connection
	GetUid() *string
}

type Connection struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// test-description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// Connection
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-connection
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Spec   *ConnectionSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status *ConnectionStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Connection) String() string {
	return dara.Prettify(s)
}

func (s Connection) GoString() string {
	return s.String()
}

func (s *Connection) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Connection) GetDescription() *string {
	return s.Description
}

func (s *Connection) GetKind() *string {
	return s.Kind
}

func (s *Connection) GetLabels() map[string]*string {
	return s.Labels
}

func (s *Connection) GetName() *string {
	return s.Name
}

func (s *Connection) GetSpec() *ConnectionSpec {
	return s.Spec
}

func (s *Connection) GetStatus() *ConnectionStatus {
	return s.Status
}

func (s *Connection) GetUid() *string {
	return s.Uid
}

func (s *Connection) SetCreatedTime(v string) *Connection {
	s.CreatedTime = &v
	return s
}

func (s *Connection) SetDescription(v string) *Connection {
	s.Description = &v
	return s
}

func (s *Connection) SetKind(v string) *Connection {
	s.Kind = &v
	return s
}

func (s *Connection) SetLabels(v map[string]*string) *Connection {
	s.Labels = v
	return s
}

func (s *Connection) SetName(v string) *Connection {
	s.Name = &v
	return s
}

func (s *Connection) SetSpec(v *ConnectionSpec) *Connection {
	s.Spec = v
	return s
}

func (s *Connection) SetStatus(v *ConnectionStatus) *Connection {
	s.Status = v
	return s
}

func (s *Connection) SetUid(v string) *Connection {
	s.Uid = &v
	return s
}

func (s *Connection) Validate() error {
	return dara.Validate(s)
}
