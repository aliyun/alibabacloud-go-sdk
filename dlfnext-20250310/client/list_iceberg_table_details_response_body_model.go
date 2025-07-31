// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIcebergTableDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListIcebergTableDetailsResponseBody
	GetNextPageToken() *string
	SetTableDetails(v []*IcebergTable) *ListIcebergTableDetailsResponseBody
	GetTableDetails() []*IcebergTable
}

type ListIcebergTableDetailsResponseBody struct {
	// example:
	//
	// ""
	NextPageToken *string         `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	TableDetails  []*IcebergTable `json:"tableDetails,omitempty" xml:"tableDetails,omitempty" type:"Repeated"`
}

func (s ListIcebergTableDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIcebergTableDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIcebergTableDetailsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListIcebergTableDetailsResponseBody) GetTableDetails() []*IcebergTable {
	return s.TableDetails
}

func (s *ListIcebergTableDetailsResponseBody) SetNextPageToken(v string) *ListIcebergTableDetailsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListIcebergTableDetailsResponseBody) SetTableDetails(v []*IcebergTable) *ListIcebergTableDetailsResponseBody {
	s.TableDetails = v
	return s
}

func (s *ListIcebergTableDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}
