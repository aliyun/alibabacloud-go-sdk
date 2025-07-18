// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v *ListDatabasesResponseBodyDatabases) *ListDatabasesResponseBody
	GetDatabases() *ListDatabasesResponseBodyDatabases
	SetMessage(v string) *ListDatabasesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListDatabasesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDatabasesResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListDatabasesResponseBody
	GetStatus() *string
}

type ListDatabasesResponseBody struct {
	// The queried databases.
	Databases *ListDatabasesResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBody) GetDatabases() *ListDatabasesResponseBodyDatabases {
	return s.Databases
}

func (s *ListDatabasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDatabasesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabasesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListDatabasesResponseBody) SetDatabases(v *ListDatabasesResponseBodyDatabases) *ListDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *ListDatabasesResponseBody) SetMessage(v string) *ListDatabasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDatabasesResponseBody) SetNextToken(v string) *ListDatabasesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDatabasesResponseBody) SetRequestId(v string) *ListDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabasesResponseBody) SetStatus(v string) *ListDatabasesResponseBody {
	s.Status = &v
	return s
}

func (s *ListDatabasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDatabasesResponseBodyDatabases struct {
	Databases []*string `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
}

func (s ListDatabasesResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyDatabases) GetDatabases() []*string {
	return s.Databases
}

func (s *ListDatabasesResponseBodyDatabases) SetDatabases(v []*string) *ListDatabasesResponseBodyDatabases {
	s.Databases = v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) Validate() error {
	return dara.Validate(s)
}
