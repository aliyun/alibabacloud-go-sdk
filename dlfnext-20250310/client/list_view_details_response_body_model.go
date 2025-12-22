// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListViewDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListViewDetailsResponseBody
	GetNextPageToken() *string
	SetViewDetails(v []*View) *ListViewDetailsResponseBody
	GetViewDetails() []*View
}

type ListViewDetailsResponseBody struct {
	// example:
	//
	// “”
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	ViewDetails   []*View `json:"viewDetails,omitempty" xml:"viewDetails,omitempty" type:"Repeated"`
}

func (s ListViewDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListViewDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListViewDetailsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListViewDetailsResponseBody) GetViewDetails() []*View {
	return s.ViewDetails
}

func (s *ListViewDetailsResponseBody) SetNextPageToken(v string) *ListViewDetailsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListViewDetailsResponseBody) SetViewDetails(v []*View) *ListViewDetailsResponseBody {
	s.ViewDetails = v
	return s
}

func (s *ListViewDetailsResponseBody) Validate() error {
	if s.ViewDetails != nil {
		for _, item := range s.ViewDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
