// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v *DescribeDatabasesResponseBodyDatabases) *DescribeDatabasesResponseBody
	GetDatabases() *DescribeDatabasesResponseBodyDatabases
	SetRequestId(v string) *DescribeDatabasesResponseBody
	GetRequestId() *string
}

type DescribeDatabasesResponseBody struct {
	Databases *DescribeDatabasesResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 009D2572-6839-43C0-8D89-44DB46CBA2EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBody) GetDatabases() *DescribeDatabasesResponseBodyDatabases {
	return s.Databases
}

func (s *DescribeDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatabasesResponseBody) SetDatabases(v *DescribeDatabasesResponseBodyDatabases) *DescribeDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDatabasesResponseBody) SetRequestId(v string) *DescribeDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabasesResponseBody) Validate() error {
	if s.Databases != nil {
		if err := s.Databases.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDatabasesResponseBodyDatabases struct {
	Database []*string `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabases) GetDatabase() []*string {
	return s.Database
}

func (s *DescribeDatabasesResponseBodyDatabases) SetDatabase(v []*string) *DescribeDatabasesResponseBodyDatabases {
	s.Database = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) Validate() error {
	return dara.Validate(s)
}
