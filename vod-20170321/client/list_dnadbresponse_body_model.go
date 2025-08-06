// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDNADBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDNADBList(v []*ListDNADBResponseBodyDNADBList) *ListDNADBResponseBody
	GetDNADBList() []*ListDNADBResponseBodyDNADBList
	SetRequestId(v string) *ListDNADBResponseBody
	GetRequestId() *string
}

type ListDNADBResponseBody struct {
	DNADBList []*ListDNADBResponseBodyDNADBList `json:"DNADBList,omitempty" xml:"DNADBList,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDNADBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDNADBResponseBody) GoString() string {
	return s.String()
}

func (s *ListDNADBResponseBody) GetDNADBList() []*ListDNADBResponseBodyDNADBList {
	return s.DNADBList
}

func (s *ListDNADBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDNADBResponseBody) SetDNADBList(v []*ListDNADBResponseBodyDNADBList) *ListDNADBResponseBody {
	s.DNADBList = v
	return s
}

func (s *ListDNADBResponseBody) SetRequestId(v string) *ListDNADBResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDNADBResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDNADBResponseBodyDNADBList struct {
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	DBId          *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	DBName        *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DBRegion      *string `json:"DBRegion,omitempty" xml:"DBRegion,omitempty"`
	DBType        *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDNADBResponseBodyDNADBList) String() string {
	return dara.Prettify(s)
}

func (s ListDNADBResponseBodyDNADBList) GoString() string {
	return s.String()
}

func (s *ListDNADBResponseBodyDNADBList) GetDBDescription() *string {
	return s.DBDescription
}

func (s *ListDNADBResponseBodyDNADBList) GetDBId() *string {
	return s.DBId
}

func (s *ListDNADBResponseBodyDNADBList) GetDBName() *string {
	return s.DBName
}

func (s *ListDNADBResponseBodyDNADBList) GetDBRegion() *string {
	return s.DBRegion
}

func (s *ListDNADBResponseBodyDNADBList) GetDBType() *string {
	return s.DBType
}

func (s *ListDNADBResponseBodyDNADBList) GetStatus() *string {
	return s.Status
}

func (s *ListDNADBResponseBodyDNADBList) SetDBDescription(v string) *ListDNADBResponseBodyDNADBList {
	s.DBDescription = &v
	return s
}

func (s *ListDNADBResponseBodyDNADBList) SetDBId(v string) *ListDNADBResponseBodyDNADBList {
	s.DBId = &v
	return s
}

func (s *ListDNADBResponseBodyDNADBList) SetDBName(v string) *ListDNADBResponseBodyDNADBList {
	s.DBName = &v
	return s
}

func (s *ListDNADBResponseBodyDNADBList) SetDBRegion(v string) *ListDNADBResponseBodyDNADBList {
	s.DBRegion = &v
	return s
}

func (s *ListDNADBResponseBodyDNADBList) SetDBType(v string) *ListDNADBResponseBodyDNADBList {
	s.DBType = &v
	return s
}

func (s *ListDNADBResponseBodyDNADBList) SetStatus(v string) *ListDNADBResponseBodyDNADBList {
	s.Status = &v
	return s
}

func (s *ListDNADBResponseBodyDNADBList) Validate() error {
	return dara.Validate(s)
}
