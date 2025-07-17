// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*string) *ListDatabasesResponseBody
	GetDatabases() []*string
	SetNextPageToken(v string) *ListDatabasesResponseBody
	GetNextPageToken() *string
}

type ListDatabasesResponseBody struct {
	Databases []*string `json:"databases,omitempty" xml:"databases,omitempty" type:"Repeated"`
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
}

func (s ListDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBody) GetDatabases() []*string {
	return s.Databases
}

func (s *ListDatabasesResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListDatabasesResponseBody) SetDatabases(v []*string) *ListDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *ListDatabasesResponseBody) SetNextPageToken(v string) *ListDatabasesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListDatabasesResponseBody) Validate() error {
	return dara.Validate(s)
}
