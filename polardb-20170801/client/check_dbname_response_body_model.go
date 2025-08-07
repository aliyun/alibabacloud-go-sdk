// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDBNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBName(v string) *CheckDBNameResponseBody
	GetDBName() *string
	SetRequestId(v string) *CheckDBNameResponseBody
	GetRequestId() *string
}

type CheckDBNameResponseBody struct {
	// The name of the database.
	//
	// example:
	//
	// test_db
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDBNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDBNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDBNameResponseBody) GetDBName() *string {
	return s.DBName
}

func (s *CheckDBNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDBNameResponseBody) SetDBName(v string) *CheckDBNameResponseBody {
	s.DBName = &v
	return s
}

func (s *CheckDBNameResponseBody) SetRequestId(v string) *CheckDBNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDBNameResponseBody) Validate() error {
	return dara.Validate(s)
}
