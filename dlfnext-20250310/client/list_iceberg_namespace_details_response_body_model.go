// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIcebergNamespaceDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceDetails(v []*Namespace) *ListIcebergNamespaceDetailsResponseBody
	GetNamespaceDetails() []*Namespace
	SetNextPageToken(v string) *ListIcebergNamespaceDetailsResponseBody
	GetNextPageToken() *string
}

type ListIcebergNamespaceDetailsResponseBody struct {
	NamespaceDetails []*Namespace `json:"namespaceDetails,omitempty" xml:"namespaceDetails,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
}

func (s ListIcebergNamespaceDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIcebergNamespaceDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIcebergNamespaceDetailsResponseBody) GetNamespaceDetails() []*Namespace {
	return s.NamespaceDetails
}

func (s *ListIcebergNamespaceDetailsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListIcebergNamespaceDetailsResponseBody) SetNamespaceDetails(v []*Namespace) *ListIcebergNamespaceDetailsResponseBody {
	s.NamespaceDetails = v
	return s
}

func (s *ListIcebergNamespaceDetailsResponseBody) SetNextPageToken(v string) *ListIcebergNamespaceDetailsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListIcebergNamespaceDetailsResponseBody) Validate() error {
	if s.NamespaceDetails != nil {
		for _, item := range s.NamespaceDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
