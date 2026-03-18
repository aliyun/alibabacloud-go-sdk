// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiKeyUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ApiKeyUpdateCmd
	GetName() *string
	SetStatus(v int32) *ApiKeyUpdateCmd
	GetStatus() *int32
}

type ApiKeyUpdateCmd struct {
	// example:
	//
	// MyApiKey
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ApiKeyUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s ApiKeyUpdateCmd) GoString() string {
	return s.String()
}

func (s *ApiKeyUpdateCmd) GetName() *string {
	return s.Name
}

func (s *ApiKeyUpdateCmd) GetStatus() *int32 {
	return s.Status
}

func (s *ApiKeyUpdateCmd) SetName(v string) *ApiKeyUpdateCmd {
	s.Name = &v
	return s
}

func (s *ApiKeyUpdateCmd) SetStatus(v int32) *ApiKeyUpdateCmd {
	s.Status = &v
	return s
}

func (s *ApiKeyUpdateCmd) Validate() error {
	return dara.Validate(s)
}
