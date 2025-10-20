// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseDetails(v []*Database) *ListDatabaseDetailsResponseBody
	GetDatabaseDetails() []*Database
	SetNextPageToken(v string) *ListDatabaseDetailsResponseBody
	GetNextPageToken() *string
}

type ListDatabaseDetailsResponseBody struct {
	DatabaseDetails []*Database `json:"databaseDetails,omitempty" xml:"databaseDetails,omitempty" type:"Repeated"`
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
}

func (s ListDatabaseDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabaseDetailsResponseBody) GetDatabaseDetails() []*Database {
	return s.DatabaseDetails
}

func (s *ListDatabaseDetailsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListDatabaseDetailsResponseBody) SetDatabaseDetails(v []*Database) *ListDatabaseDetailsResponseBody {
	s.DatabaseDetails = v
	return s
}

func (s *ListDatabaseDetailsResponseBody) SetNextPageToken(v string) *ListDatabaseDetailsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListDatabaseDetailsResponseBody) Validate() error {
	if s.DatabaseDetails != nil {
		for _, item := range s.DatabaseDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
