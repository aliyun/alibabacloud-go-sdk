// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListTableDetailsResponseBody
	GetNextPageToken() *string
	SetTableDetails(v []*Table) *ListTableDetailsResponseBody
	GetTableDetails() []*Table
}

type ListTableDetailsResponseBody struct {
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string  `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	TableDetails  []*Table `json:"tableDetails,omitempty" xml:"tableDetails,omitempty" type:"Repeated"`
}

func (s ListTableDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTableDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableDetailsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListTableDetailsResponseBody) GetTableDetails() []*Table {
	return s.TableDetails
}

func (s *ListTableDetailsResponseBody) SetNextPageToken(v string) *ListTableDetailsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListTableDetailsResponseBody) SetTableDetails(v []*Table) *ListTableDetailsResponseBody {
	s.TableDetails = v
	return s
}

func (s *ListTableDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}
