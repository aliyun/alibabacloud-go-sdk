// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDBNameZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBName(v string) *CheckDBNameZonalResponseBody
	GetDBName() *string
	SetRequestId(v string) *CheckDBNameZonalResponseBody
	GetRequestId() *string
}

type CheckDBNameZonalResponseBody struct {
	// example:
	//
	// test_db
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDBNameZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDBNameZonalResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDBNameZonalResponseBody) GetDBName() *string {
	return s.DBName
}

func (s *CheckDBNameZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDBNameZonalResponseBody) SetDBName(v string) *CheckDBNameZonalResponseBody {
	s.DBName = &v
	return s
}

func (s *CheckDBNameZonalResponseBody) SetRequestId(v string) *CheckDBNameZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDBNameZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
