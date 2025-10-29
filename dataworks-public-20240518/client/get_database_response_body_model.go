// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v *Database) *GetDatabaseResponseBody
	GetDatabase() *Database
	SetRequestId(v string) *GetDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDatabaseResponseBody
	GetSuccess() *bool
}

type GetDatabaseResponseBody struct {
	// The database details.
	Database *Database `json:"Database,omitempty" xml:"Database,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBody) GetDatabase() *Database {
	return s.Database
}

func (s *GetDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDatabaseResponseBody) SetDatabase(v *Database) *GetDatabaseResponseBody {
	s.Database = v
	return s
}

func (s *GetDatabaseResponseBody) SetRequestId(v string) *GetDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabaseResponseBody) SetSuccess(v bool) *GetDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *GetDatabaseResponseBody) Validate() error {
	if s.Database != nil {
		if err := s.Database.Validate(); err != nil {
			return err
		}
	}
	return nil
}
