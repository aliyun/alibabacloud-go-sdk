// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDNADBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDNADB(v *GetDNADBResponseBodyDNADB) *GetDNADBResponseBody
	GetDNADB() *GetDNADBResponseBodyDNADB
	SetRequestId(v string) *GetDNADBResponseBody
	GetRequestId() *string
}

type GetDNADBResponseBody struct {
	DNADB     *GetDNADBResponseBodyDNADB `json:"DNADB,omitempty" xml:"DNADB,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDNADBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDNADBResponseBody) GoString() string {
	return s.String()
}

func (s *GetDNADBResponseBody) GetDNADB() *GetDNADBResponseBodyDNADB {
	return s.DNADB
}

func (s *GetDNADBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDNADBResponseBody) SetDNADB(v *GetDNADBResponseBodyDNADB) *GetDNADBResponseBody {
	s.DNADB = v
	return s
}

func (s *GetDNADBResponseBody) SetRequestId(v string) *GetDNADBResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDNADBResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDNADBResponseBodyDNADB struct {
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	DBId          *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	DBName        *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DBRegion      *string `json:"DBRegion,omitempty" xml:"DBRegion,omitempty"`
	DBType        *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDNADBResponseBodyDNADB) String() string {
	return dara.Prettify(s)
}

func (s GetDNADBResponseBodyDNADB) GoString() string {
	return s.String()
}

func (s *GetDNADBResponseBodyDNADB) GetDBDescription() *string {
	return s.DBDescription
}

func (s *GetDNADBResponseBodyDNADB) GetDBId() *string {
	return s.DBId
}

func (s *GetDNADBResponseBodyDNADB) GetDBName() *string {
	return s.DBName
}

func (s *GetDNADBResponseBodyDNADB) GetDBRegion() *string {
	return s.DBRegion
}

func (s *GetDNADBResponseBodyDNADB) GetDBType() *string {
	return s.DBType
}

func (s *GetDNADBResponseBodyDNADB) GetStatus() *string {
	return s.Status
}

func (s *GetDNADBResponseBodyDNADB) SetDBDescription(v string) *GetDNADBResponseBodyDNADB {
	s.DBDescription = &v
	return s
}

func (s *GetDNADBResponseBodyDNADB) SetDBId(v string) *GetDNADBResponseBodyDNADB {
	s.DBId = &v
	return s
}

func (s *GetDNADBResponseBodyDNADB) SetDBName(v string) *GetDNADBResponseBodyDNADB {
	s.DBName = &v
	return s
}

func (s *GetDNADBResponseBodyDNADB) SetDBRegion(v string) *GetDNADBResponseBodyDNADB {
	s.DBRegion = &v
	return s
}

func (s *GetDNADBResponseBodyDNADB) SetDBType(v string) *GetDNADBResponseBodyDNADB {
	s.DBType = &v
	return s
}

func (s *GetDNADBResponseBodyDNADB) SetStatus(v string) *GetDNADBResponseBodyDNADB {
	s.Status = &v
	return s
}

func (s *GetDNADBResponseBodyDNADB) Validate() error {
	return dara.Validate(s)
}
