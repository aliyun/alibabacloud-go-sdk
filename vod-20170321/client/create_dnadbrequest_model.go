// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDNADBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDNADBRequest
	GetClientToken() *string
	SetDBDescription(v string) *CreateDNADBRequest
	GetDBDescription() *string
	SetDBName(v string) *CreateDNADBRequest
	GetDBName() *string
	SetDBRegion(v string) *CreateDNADBRequest
	GetDBRegion() *string
	SetDBType(v string) *CreateDNADBRequest
	GetDBType() *string
}

type CreateDNADBRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// This parameter is required.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// This parameter is required.
	DBRegion *string `json:"DBRegion,omitempty" xml:"DBRegion,omitempty"`
	DBType   *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
}

func (s CreateDNADBRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDNADBRequest) GoString() string {
	return s.String()
}

func (s *CreateDNADBRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDNADBRequest) GetDBDescription() *string {
	return s.DBDescription
}

func (s *CreateDNADBRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateDNADBRequest) GetDBRegion() *string {
	return s.DBRegion
}

func (s *CreateDNADBRequest) GetDBType() *string {
	return s.DBType
}

func (s *CreateDNADBRequest) SetClientToken(v string) *CreateDNADBRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDNADBRequest) SetDBDescription(v string) *CreateDNADBRequest {
	s.DBDescription = &v
	return s
}

func (s *CreateDNADBRequest) SetDBName(v string) *CreateDNADBRequest {
	s.DBName = &v
	return s
}

func (s *CreateDNADBRequest) SetDBRegion(v string) *CreateDNADBRequest {
	s.DBRegion = &v
	return s
}

func (s *CreateDNADBRequest) SetDBType(v string) *CreateDNADBRequest {
	s.DBType = &v
	return s
}

func (s *CreateDNADBRequest) Validate() error {
	return dara.Validate(s)
}
