// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *MigrateApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *MigrateApplicationResponseBody
	GetMessage() *string
	SetData(v *MigrateApplicationResponseBodyData) *MigrateApplicationResponseBody
	GetData() *MigrateApplicationResponseBodyData
}

type MigrateApplicationResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The API information.
	Data *MigrateApplicationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s MigrateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *MigrateApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MigrateApplicationResponseBody) GetData() *MigrateApplicationResponseBodyData {
	return s.Data
}

func (s *MigrateApplicationResponseBody) SetCode(v int32) *MigrateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *MigrateApplicationResponseBody) SetMessage(v string) *MigrateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateApplicationResponseBody) SetData(v *MigrateApplicationResponseBodyData) *MigrateApplicationResponseBody {
	s.Data = v
	return s
}

func (s *MigrateApplicationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MigrateApplicationResponseBodyData struct {
	// The migration ID.
	//
	// example:
	//
	// a3de82d7-83a4-4cca-8d1e-63f87651ce78
	MigrationId *string `json:"migrationId,omitempty" xml:"migrationId,omitempty"`
}

func (s MigrateApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MigrateApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *MigrateApplicationResponseBodyData) GetMigrationId() *string {
	return s.MigrationId
}

func (s *MigrateApplicationResponseBodyData) SetMigrationId(v string) *MigrateApplicationResponseBodyData {
	s.MigrationId = &v
	return s
}

func (s *MigrateApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
