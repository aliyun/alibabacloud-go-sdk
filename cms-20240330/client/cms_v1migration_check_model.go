// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCmsV1MigrationCheck interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CmsV1MigrationCheck
	GetCode() *string
	SetMessage(v string) *CmsV1MigrationCheck
	GetMessage() *string
}

type CmsV1MigrationCheck struct {
	// The check result code, where NAME_ indicates a naming issue and UNSUPPORTED_ indicates an unsupported resource.
	//
	// example:
	//
	// NAME_INVALID
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The check result description.
	//
	// example:
	//
	// Target name must contain 1 to 256 characters
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CmsV1MigrationCheck) String() string {
	return dara.Prettify(s)
}

func (s CmsV1MigrationCheck) GoString() string {
	return s.String()
}

func (s *CmsV1MigrationCheck) GetCode() *string {
	return s.Code
}

func (s *CmsV1MigrationCheck) GetMessage() *string {
	return s.Message
}

func (s *CmsV1MigrationCheck) SetCode(v string) *CmsV1MigrationCheck {
	s.Code = &v
	return s
}

func (s *CmsV1MigrationCheck) SetMessage(v string) *CmsV1MigrationCheck {
	s.Message = &v
	return s
}

func (s *CmsV1MigrationCheck) Validate() error {
	return dara.Validate(s)
}
